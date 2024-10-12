package logic

import (
	"microservice/entity"
	"microservice/services/transaction/model"
	"microservice/shared"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
	js nats.JetStreamContext
}

func NewTransactionService(db *gorm.DB, nc *nats.Conn) (*TransactionService, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return &TransactionService{db: db, js: js}, nil
}

// GetAllTransaction godoc
//
// @Summary		Get all transaction
// @Description	Get all transaction
// @Tags			transaction
// @Accept			json
// @Produce		json
//
// @Router			/transaction/ [get]
func (ts *TransactionService) GetAllTransaction(ctx *fiber.Ctx) error {
	var transactions []entity.Transaction
	result := ts.db.Find(&transactions)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	var transactionModel []model.TransactionInfo
	for _, transaction := range transactions {
		transactionModel = append(transactionModel, model.TransactionInfo{
			ID:                   transaction.ID,
			SourceAccountID:      transaction.SourceAccountID,
			DestinationAccountID: transaction.DestinationAccountID,
			Amount:               transaction.Amount,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactionModel,
	})
}

// GetTransactionByID godoc
//
// @Summary		Get transaction by ID
// @Description	Get transaction by ID
// @Tags			transaction
// @Accept			json
// @Produce		json
//
// @Param			id path string true "Transaction ID"
// @Router			/transaction/{id} [get]
func (ts *TransactionService) GetTransactionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var transaction entity.Transaction
	result := ts.db.Where("id = ?", id).First(&transaction)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"transaction": transaction,
	})
}

// CreateTransaction godoc
//
// @Summary		Create transaction
// @Description	Create a new transaction between accounts
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			transaction body model.CreateTransaction true "Transaction information"
// @Router			/transaction/ [post]
func (ts *TransactionService) CreateTransaction(ctx *fiber.Ctx) error {
	var createTransaction model.CreateTransaction
	if err := ctx.BodyParser(&createTransaction); err != nil {
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

	// Validate transaction amount
	if createTransaction.Amount <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Transaction amount must be greater than zero.",
		})
	}

	// Fetch source and destination accounts
	var sourceAccount, destinationAccount entity.Account
	if err := ts.db.Where("id = ?", createTransaction.SourceAccountID).First(&sourceAccount).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Source account not found.",
		})
	}

	if err := ts.db.Where("id = ?", createTransaction.DestinationAccountID).First(&destinationAccount).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Destination account not found.",
		})
	}

	// Check if source account has enough balance
	if sourceAccount.Balance < createTransaction.Amount {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Insufficient balance.",
		})
	}

	// Process transaction within a database transaction
	err = ts.db.Transaction(func(tx *gorm.DB) error {
		sourceAccount.Balance -= createTransaction.Amount
		destinationAccount.Balance += createTransaction.Amount

		if err := tx.Save(&sourceAccount).Error; err != nil {
			return err
		}
		if err := tx.Save(&destinationAccount).Error; err != nil {
			return err
		}

		// Create transaction record
		transaction := entity.Transaction{
			SourceAccountID:      createTransaction.SourceAccountID,
			DestinationAccountID: createTransaction.DestinationAccountID,
			Amount:               createTransaction.Amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// Publish transaction event
		event := map[string]interface{}{
			"ID":                   transaction.ID,
			"sourceAccountID":      transaction.SourceAccountID,
			"destinationAccountID": transaction.DestinationAccountID,
			"amount":               transaction.Amount,
		}
		eventData, err := shared.MarshalToJSON(event)
		if err != nil {
			return err
		}

		_, err = ts.js.Publish("transaction.created", eventData)
		return err
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Transaction completed successfully",
	})
}
