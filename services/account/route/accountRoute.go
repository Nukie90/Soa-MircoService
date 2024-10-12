package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"microservice/services/account/logic"
)

type AccountRoute struct {
	accountLogic *logic.AccountService
}

func NewAccountRoute(accountLogic *logic.AccountService) *AccountRoute {
	return &AccountRoute{accountLogic: accountLogic}
}

func (ar *AccountRoute) SetupAccountRoute(app *fiber.App) {
	app.Get("/monitor", monitor.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	account := v1.Group("/account")
	{
		account.Get("/", ar.accountLogic.GetAllAccount)
		account.Get("/:id", ar.accountLogic.GetAccountByID)
		account.Post("/", ar.accountLogic.CreateAccount)
		account.Post("/topup", ar.accountLogic.TopUp)
		//account.Put("/:id", ar.accountLogic.ChangePin)
		account.Delete("/", ar.accountLogic.DeleteAccount)
		//account.Get("/clientAcc", ar.accountLogic.GetClientAccount)
	}

}
