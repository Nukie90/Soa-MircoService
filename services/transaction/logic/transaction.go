package logic

import (
	"microservice/entity"
	"microservice/services/transaction/model"

	"github.com/gofiber/fiber/v2"
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
// @Description	Create transaction
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

	transaction := entity.Transaction{
		SourceAccountID:      createTransaction.SourceAccountID,
		DestinationAccountID: createTransaction.DestinationAccountID,
		Amount:               createTransaction.Amount,
	}

	result := ts.db.Create(&transaction)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"transaction": transaction,
	})
}
