package logic

import (
	"microservice/entity"
	"microservice/services/auth/model"
	"microservice/shared"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
	js nats.JetStreamContext
}

// NewAuthService initializes the AuthService
func NewAuthService(db *gorm.DB, nc *nats.Conn) (*AuthService, error) {
	js, err := nc.JetStream() // Initialize JetStream context
	if err != nil {
		return nil, err
	}
	return &AuthService{db: db, js: js}, nil
}

// SignUp godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.SignUp	true	"User information"
//	@Success		201		{string}	string			"User created successfully"
//	@Failure		400		{string}	string			"Bad request"
//
//	@Router			/auth/signup [post]
func (us *AuthService) SignUp(ctx *fiber.Ctx) error {
	user := new(model.SignUp)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Convert string to time
	birthDay, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":    err.Error(),
			"Birthday": user.BirthDate,
		})
	}

	hasher := shared.NewLocalConfig(10)

	// Hash password
	hashPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newUser := entity.User{
		ID:        user.ID,
		Name:      user.Name,
		Address:   user.Address,
		Password:  hashPassword,
		BirthDate: birthDay,
	}

	if err := us.db.Create(&newUser).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Publish user created event
	event := map[string]interface{}{
		"id":        newUser.ID,
		"name":      newUser.Name,
		"address":   newUser.Address,
		"birthDate": newUser.BirthDate,
	}

	// Marshal the event to JSON
	eventData, err := shared.MarshalToJSON(event)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not marshal event to JSON",
		})
	}

	// Publish the user created event to JetStream
	if _, err := us.js.Publish("user.created", eventData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not publish event",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.Login	true	"User information"
//	@Success		200		{string}	string		"Token"
//	@Failure		400		{string}	string		"Bad request"
//	@Failure		401		{string}	string		"Unauthorized"
//
//	@Router			/auth/login [post]
func (us *AuthService) Login(ctx *fiber.Ctx) error {
	var login model.Login

	if err := ctx.BodyParser(&login); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"body parser error: ": err.Error(),
		})
	}

	var user entity.User
	if err := us.db.Where("id = ?", login.ID).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !shared.CheckPasswordHash(login.Password, user.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24), // 1-day expiration
		HTTPOnly: false,                          // Prevents client-side JS from accessing the cookie
		SameSite: "None",                         // Controls when the cookie is sent (Lax or None for cross-origin)
		Secure:   true,                           // Set to true in HTTPS environments
	})

	//set auth token from user service to gateway in header
	ctx.Set("Authorization", tokenString)

	//set local
	ctx.Locals("userId", user.ID)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}
