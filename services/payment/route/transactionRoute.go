package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"microservice/services/payment/logic"
)

type PaymentRoute struct {
	paymentLogic *logic.PaymentService
}

func NewPaymentRoute(paymentLogic *logic.PaymentService) *PaymentRoute {
	return &PaymentRoute{paymentLogic: paymentLogic}
}

func (pr *PaymentRoute) SetupPaymentRoute(app *fiber.App) {
	app.Get("/monitor", monitor.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	account := v1.Group("/payment")
	{
		account.Get("/", pr.paymentLogic.GetAllPayment)
		account.Get("/:id", pr.paymentLogic.GetPaymentByID)
		account.Post("/", pr.paymentLogic.CreatePayment)
		//account.Put("/:id", ar.accountLogic.ChangePin)
		//account.Delete("/:id", ar.accountLogic.DeleteAccount)
		//account.Get("/clientAcc", ar.accountLogic.GetClientAccount)
	}

}
