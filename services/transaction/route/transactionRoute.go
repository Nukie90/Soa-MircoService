package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"microservice/services/transaction/logic"
)

type TransactionRoute struct {
	transactionLogic *logic.TransactionService
}

func NewTransactionRoute(transactionLogic *logic.TransactionService) *TransactionRoute {
	return &TransactionRoute{transactionLogic: transactionLogic}
}

func (tr *TransactionRoute) SetupTransactionRoute(app *fiber.App) {
	app.Get("/monitor", monitor.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	account := v1.Group("/transaction")
	{
		account.Get("/", tr.transactionLogic.GetAllTransaction)
		account.Get("/:id", tr.transactionLogic.GetTransactionByID)
		account.Post("/", tr.transactionLogic.CreateTransaction)
	}

}
