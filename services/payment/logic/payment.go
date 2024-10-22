package logic

import (
	"microservice/entity"
	"microservice/services/payment/model"
	"microservice/shared"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type PaymentService struct {
	db *gorm.DB
	js nats.JetStreamContext
}

func NewPaymentService(db *gorm.DB, nc *nats.Conn) (*PaymentService, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return &PaymentService{db: db, js: js}, nil
}

// GetAllPayment godoc
//
// @Summary		Get all payment
// @Description	Get all payment
// @Tags			payment
// @Accept			json
// @Produce		json
//
// @Router			/payment/ [get]
func (ps *PaymentService) GetAllPayment(ctx *fiber.Ctx) error {
	var payments []entity.Payment
	result := ps.db.Find(&payments)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	var paymentModel []model.PaymentInfo
	for _, payment := range payments {
		paymentModel = append(paymentModel, model.PaymentInfo{
			ID:              payment.ID,
			SourceAccountID: payment.SourceAccountID,
			ReferenceCode:   payment.ReferenceCode,
			Amount:          payment.Amount,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payments": paymentModel,
	})
}

// GetPaymentByID godoc
//
// @Summary		Get payment by ID
// @Description	Get payment by ID
// @Tags			payment
// @Accept			json
// @Produce		json
//
// @Param			id path string true "Payment ID"
// @Router			/payment/{id} [get]
func (ps *PaymentService) GetPaymentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payment entity.Payment
	result := ps.db.Where("id = ?", id).First(&payment)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payment": payment,
	})
}

// CreatePayment godoc
//
// @Summary		Create payment
// @Description	Create a new payment between accounts
// @Tags			payment
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			payment body model.CreatePayment true "Payment information"
// @Router			/payment/ [post]
func (ps *PaymentService) CreatePayment(ctx *fiber.Ctx) error {
	var createPayment model.CreatePayment
	if err := ctx.BodyParser(&createPayment); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Get token from header for authentication
	tokenHeader := ctx.Get("Authorization")
	if tokenHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Validate payment amount
	if createPayment.Amount <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Payment amount must be greater than zero.",
		})
	}

	// Fetch source and destination accounts
	var sourceAccount entity.Account
	if err := ps.db.Where("id = ?", createPayment.SourceAccountID).First(&sourceAccount).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Source account not found.",
		})
	}

	// Check if source account has enough balance
	if sourceAccount.Balance < createPayment.Amount {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Insufficient balance.",
		})
	}

	// Process payment within a database payment
	err = ps.db.Transaction(func(tx *gorm.DB) error {
		sourceAccount.Balance -= createPayment.Amount

		if err := tx.Save(&sourceAccount).Error; err != nil {
			return err
		}

		// Create payment record
		payment := entity.Payment{
			SourceAccountID: createPayment.SourceAccountID,
			ReferenceCode:   createPayment.ReferenceCode,
			Amount:          createPayment.Amount,
		}
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		// Publish payment event
		event := map[string]interface{}{
			"ID":              payment.ID,
			"sourceAccountID": payment.SourceAccountID,
			"referenceCode":   payment.ReferenceCode,
			"amount":          payment.Amount,
		}
		eventData, err := shared.MarshalToJSON(event)
		if err != nil {
			return err
		}

		_, err = ps.js.Publish("payment.created", eventData)
		return err
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"payment": createPayment,
	})
}
