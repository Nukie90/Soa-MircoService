package logic

import (
	"microservice/entity"
	"microservice/services/payment/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaymentService struct {
	db *gorm.DB
}

func NewPaymentService(db *gorm.DB) *PaymentService {
	return &PaymentService{
		db: db,
	}
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
// @Description	Create payment
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

	payment := entity.Payment{
		SourceAccountID: createPayment.SourceAccountID,
		ReferenceCode:   createPayment.ReferenceCode,
		Amount:          createPayment.Amount,
	}

	result := ps.db.Create(&payment)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"payment": payment,
	})
}
