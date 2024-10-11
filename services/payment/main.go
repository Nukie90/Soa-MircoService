// @title			Payment Service
// @version		1.0
// @description	This is the Payment service API documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	nukie.nxk@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			127.0.0.1:3400
// @BasePath		/api/v1
package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"microservice/services/payment/logic"
	"microservice/services/payment/route"
	"microservice/shared"
	"os"

	_ "microservice/services/payment/docs"
)

type app struct {
	*fiber.App
}

func newApp() *app {
	return &app{fiber.New()}
}

func (a *app) startApp() error {
	computeID := os.Getenv("DB_COMPUTE_ID")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("Starting payment service")

	db := shared.NewDatabase(computeID, password, dbName)

	newDB, err := db.PostgresConnection()
	if err != nil {
		return err
	}
	accountLogic := logic.NewPaymentService(newDB)

	a.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:3000, http://localhost:3000",
	}))

	a.Use(logger.New())
	a.Get("swagger/*", fiberSwagger.WrapHandler)

	route.NewPaymentRoute(accountLogic).SetupPaymentRoute(a.App)

	if err := a.Listen(":3400"); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := godotenv.Load("../../env/.env"); err != nil {
		fmt.Println("No .env file found")
	}

	app := newApp()
	if err := app.startApp(); err != nil {
		fmt.Println(err)
	}
}
