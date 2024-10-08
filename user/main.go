// @title			User Service
// @version		1.0
// @description	This is the user service API documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	nukie.nxk@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3100
// @BasePath		/api/v1
package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"microservice/shared"
	_ "microservice/user/docs"
	"os"
)

type App struct {
	*fiber.App
}

func NewApp() *App {
	return &App{fiber.New()}
}

func (a *App) StartApp() error {
	err := godotenv.Load("../env/.env")
	if err != nil {
		return err
	}

	computeID := os.Getenv("DB_COMPUTE_ID")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("Starting user service")

	db := shared.NewDatabase(computeID, password, dbName)

	_, err = db.PostgresConnection()
	if err != nil {
		return err
	}

	a.Get("swagger/*", fiberSwagger.WrapHandler)

	err = a.Listen(":3100")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := NewApp()
	err := app.StartApp()
	if err != nil {
		panic(err)
	}
}
