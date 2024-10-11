package logic

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"microservice/entity"
	"microservice/services/user/model"
	"os"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
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
