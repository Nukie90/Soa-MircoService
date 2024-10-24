package route

import (
	"github.com/gofiber/fiber/v2"
	"microservice/services/gateway/logic"
	"microservice/shared"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.Post("/signup", logic.ForwardSignup)
			auth.Post("/login", logic.ForwardLogin)
			auth.Get("/token", logic.GetToken)
		}
		users := v1.Group("/users", shared.AuthenticateUser)
		{
			users.Get("/all", logic.ForwardGetAllUser)
			users.Get("/me", logic.ForwardGetMe)
			users.Get("/all/:id", logic.ForwardGetUserByID)
		}
		account := v1.Group("/account", shared.AuthenticateUser)
		{
			account.Get("/", logic.ForwardGetAllAccount)
			account.Post("/", logic.ForwardCreateAccount)
			account.Get("/id/:id", logic.ForwardGetAccountByID)
			account.Put("/topup", logic.ForwardTopUp)
			account.Put("/", logic.ForwardChangePin)
			account.Delete("/", logic.ForwardDeleteAccount)
			account.Post("/verify", logic.ForwardVerifyAccount)
			account.Get("/getAccountsByUserID", logic.ForwardGetAccountsByUserID)
		}
		transaction := v1.Group("/transaction", shared.AuthenticateUser)
		{
			transaction.Post("/", logic.ForwardCreateTransaction)
			transaction.Get("/", logic.ForwardGetAllTransaction)
			transaction.Get("/:id", logic.ForwardGetTransactionByID)
		}
		payment := v1.Group("/payment", shared.AuthenticateUser)
		{
			payment.Get("/", logic.ForwardGetAllPayment)
			payment.Post("/", logic.ForwardCreatePayment)
			payment.Get("/:id", logic.ForwardGetPaymentByID)
		}
	}

}
