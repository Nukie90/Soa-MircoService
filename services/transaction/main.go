// @title			Transaction Service
// @version		1.0
// @description	This is the Transaction service API documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	nukie.nxk@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			127.0.0.1:3300
// @BasePath		/api/v1
package main

import (
	"fmt"
	"microservice/services/transaction/logic"
	"microservice/services/transaction/route"
	"microservice/shared"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "microservice/services/transaction/docs"
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
	dbName := os.Getenv("TRANSACTION_NAME")

	fmt.Println("Starting transaction service")

	db := shared.NewDatabase(computeID, password, dbName)

	newDB, err := db.PostgresConnection()
	if err != nil {
		return err
	}
	accountLogic := logic.NewTransactionService(newDB)

	a.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:3000, http://localhost:3000",
	}))

	a.Use(logger.New())
	a.Get("swagger/*", fiberSwagger.WrapHandler)

	route.NewTransactionRoute(accountLogic).SetupTransactionRoute(a.App)

	if err := a.Listen(":3300"); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No ..env file found")
	}

	nc, err := shared.ConnectNATS()
	if err != nil {
		return
	}
	defer nc.Close()

	// Initialize UserService with the DB connection
	db := shared.NewDatabase(os.Getenv("DB_COMPUTE_ID"), os.Getenv("DB_PASSWORD"), os.Getenv("TRANSACTION_NAME"))
	newDB, err := db.PostgresConnection()
	if err != nil {
		return
	}
	userService := logic.NewTransactionService(newDB)

	// Subscribe to user.created events
	userService.SubscribeToUserCreated(nc)

	app := newApp()
	if err := app.startApp(); err != nil {
		fmt.Println(err)
	}
}
