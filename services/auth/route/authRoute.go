package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	logic2 "microservice/services/auth/logic"
)

type AuthRoute struct {
	AuthLogic *logic2.AuthService
}

func NewAuthRoute(AuthLogic *logic2.AuthService) *AuthRoute {
	return &AuthRoute{AuthLogic: AuthLogic}
}

func (ur *AuthRoute) SetupAuthRoute(app *fiber.App) {
	app.Get("/monitor", monitor.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.Post("/login", ur.AuthLogic.Login)
		auth.Post("/signup", ur.AuthLogic.SignUp)
	}
}
