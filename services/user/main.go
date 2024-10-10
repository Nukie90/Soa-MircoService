// @title			User Service
// @version		1.0
// @description	This is the user service API documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	nukie.nxk@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			127.0.0.1:3100
// @BasePath		/api/v1
package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "microservice/services/user/docs"
	"microservice/services/user/logic"
	"microservice/services/user/route"
	"microservice/shared"
	"os"
)

type App struct {
	*fiber.App
}

func NewApp() *App {
	return &App{fiber.New()}
}

func (a *App) StartApp() error {

	computeID := os.Getenv("DB_COMPUTE_ID")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("Starting user service")

	db := shared.NewDatabase(computeID, password, dbName)

	newDB, err := db.PostgresConnection()
	if err != nil {
		return err
	}
	userlogic := logic.NewUserService(newDB)

	a.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:3000, http://localhost:3000",
	}))

	a.Use(logger.New())
	a.Get("swagger/*", fiberSwagger.WrapHandler)

	route.NewUserRoute(userlogic).SetupUserRoute(a.App)

	err = a.Listen(":3100")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load("../../env/.env")
	if err != nil {
		panic(err)
	}

	app := NewApp()
	err = app.StartApp()
	if err != nil {
		panic(err)
	}
}
