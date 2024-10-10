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
			users.Get("/", logic.ForwardGetAllUser)
			users.Get("/:id", logic.ForwardGetUserByID)
		}
	}

}
