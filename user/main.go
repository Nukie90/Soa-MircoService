// @title User Service
// @version 1.0
// @description This is the user service API documentation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email nukie.nxk@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3100
// @BasePath /api/v1
package main

import (
	"github.com/gofiber/fiber/v2"
	_ "microservice/user/docs"
)

type App struct {
	*fiber.App
}
