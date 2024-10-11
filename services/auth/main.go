// @title			Auth Service
// @version		1.0
// @description	This is the user service API documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	nukie.nxk@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			127.0.0.1:3500
// @BasePath		/api/v1
package main

import (
	"fmt"
	"microservice/services/auth/logic"
	"microservice/services/auth/route"
	"microservice/shared"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "microservice/services/auth/docs"
)

type App struct {
	*fiber.App
}

func NewApp() *App {
	return &App{fiber.New()}
}

func (a *App) StartApp() error {

	computeID := os.Getenv("AUTH_COMPUTE_ID")
	password := os.Getenv("AUTH_PASSWORD")
	dbName := os.Getenv("AUTH_NAME")

	fmt.Println("Starting Auth service")

	// Establish connection to NATS
	nc, err := shared.ConnectNATS()
	if err != nil {
		return err
	}
	defer nc.Close()

	db := shared.NewDatabase(computeID, password, dbName)

	newDB, err := db.PostgresConnection()
	if err != nil {
		return err
	}
	authLogic := logic.NewAuthService(newDB, nc)

	a.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:3000, http://localhost:3000",
	}))

	a.Use(logger.New())
	a.Get("swagger/*", fiberSwagger.WrapHandler)

	route.NewAuthRoute(authLogic).SetupAuthRoute(a.App)

	err = a.Listen(":3500")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := NewApp()
	err = app.StartApp()
	if err != nil {
		fmt.Println(err)
	}
}
