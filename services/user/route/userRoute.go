package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"microservice/services/user/logic"
)

type UserRoute struct {
	userLogic *logic.UserService
}

func NewUserRoute(userLogic *logic.UserService) *UserRoute {
	return &UserRoute{userLogic: userLogic}
}

func (ur *UserRoute) SetupUserRoute(app *fiber.App) {
	app.Get("/monitor", monitor.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.Post("/login", ur.userLogic.Login)
		auth.Post("/signup", ur.userLogic.SignUp)
	}
	users := v1.Group("/users")
	{
		users.Get("/all", ur.userLogic.GetAllUser)
		users.Get("/me", ur.userLogic.GetMe)
		users.Get("/all/:id", ur.userLogic.GetUserByID)
	}
}
