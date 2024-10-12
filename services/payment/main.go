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
	"log"
	"microservice/services/payment/logic"
	"microservice/services/payment/route"
	"microservice/shared"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	fiberSwagger "github.com/swaggo/fiber-swagger"

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
	dbName := os.Getenv("PAYMENT_NAME")

	fmt.Println("Starting payment service")

	db := shared.NewDatabase(computeID, password, dbName)

	newDB, err := db.PostgresConnection()
	if err != nil {
		return err
	}

	nc, err := shared.ConnectNATS()
	if err != nil {
		return err
	}
	defer nc.Close()

	// Initialize JetStream
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Error initializing JetStream: %v", err)
	}

	// Create a stream for storing messages
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "payment_stream",
		Subjects: []string{"payment.created"},
	})
	if err != nil {
		log.Printf("Stream may already exist: %v", err)
	}

	// Initialize Payment Service with JetStream
	paymentService, err := logic.NewPaymentService(newDB, nc)
	if err != nil {
		log.Fatalf("Error creating Account Service: %v", err)
	}

	// Subscribe to user.created events using JetStream
	if err := paymentService.SubscribeToUserCreated(); err != nil {
		log.Fatalf("Error subscribing to user.created events: %v", err)
	}

	// Subscribe to account.created events using JetStream
	if err := paymentService.SubscribeToAccountCreated(); err != nil {
		log.Fatalf("Error subscribing to account.created events: %v", err)
	}

	if err := paymentService.SubscribeToAccountTopUp(); err != nil {
		log.Fatalf("Error subscribing to account.topup events: %v", err)
	}

	// Subscribe to account.deleted events using JetStream
	if err := paymentService.SubscribeToAccountDeleted(); err != nil {
		log.Fatalf("Error subscribing to account.deleted events: %v", err)
	}

	// Subscribe to account.changedPin events using JetStream
	if err := paymentService.SubscribeToAccountChangedPin(); err != nil {
		log.Fatalf("Error subscribing to account.changedPin events: %v", err)
	}

	paymentLogic, err := logic.NewPaymentService(newDB, nc)

	a.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:3000, http://localhost:3000",
	}))

	a.Use(logger.New())
	a.Get("swagger/*", fiberSwagger.WrapHandler)

	route.NewPaymentRoute(paymentLogic).SetupPaymentRoute(a.App)

	if err := a.Listen(":3400"); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	app := newApp()
	if err := app.startApp(); err != nil {
		fmt.Println(err)
	}
}
