package logic

import (
	"fmt"
	"microservice/entity"
	"microservice/services/user/model"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
	js nats.JetStreamContext
}

func NewUserService(db *gorm.DB, nc *nats.Conn) (*UserService, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return &UserService{db: db, js: js}, nil
}

// GetAllUser godoc
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/users/all [get]
func (us *UserService) GetAllUser(ctx *fiber.Ctx) error {
	var users []entity.User
	if err := us.db.Find(&users).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []model.UserModel
	for _, user := range users {
		response = append(response, model.UserModel{
			Name:      user.Name,
			Address:   user.Address,
			BirthDate: user.BirthDate.Format("2006-01-02"),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": response,
	})
}

// GetUserByID godoc
//
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Router			/users/all/{id} [get]
func (us *UserService) GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user entity.User
	if err := us.db.Where("id = ?", id).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": model.UserModel{
			Name:      user.Name,
			Address:   user.Address,
			BirthDate: user.BirthDate.Format("2006-01-02"),
		},
	})
}

// GetMe godoc
//
//	@Summary		Get user by token
//	@Description	Get user by token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/users/me [get]
func (us *UserService) GetMe(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token from gateway",
		})
	}

	fmt.Println("Token: ", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var user entity.User
	if err := us.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error1": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": model.UserModel{
			Name:      user.Name,
			Address:   user.Address,
			BirthDate: user.BirthDate.Format("2006-01-02"),
		},
	})
}
