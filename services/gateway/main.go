// @title			API Gateway Service
// @version		1.0
// @description	This is the API Gateway service API documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	nukie.nxk@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			127.0.0.1:3000
// @BasePath		/api/v1
package main

import (
	"fmt"
	_ "microservice/services/gateway/docs"
	"microservice/services/gateway/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupGateway() error {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:3100, http://localhost:3100, http://127.0.0.1:5000, http://localhost:5000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	// Setup routes
	route.SetupRoute(app)

	app.Get("/monitor", monitor.New())
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	err := app.Listen(":3000")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading ..env file")
	}

	err = SetupGateway()
	if err != nil {
		fmt.Println("Error starting gateway")
	}
}
