package logic

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"microservice/entity"
	"microservice/services/account/model"
	"microservice/shared"
	"os"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

// GetAllAccount godoc
//
// @Summary		Get all account
// @Description	Get all account
// @Tags			account
// @Accept			json
// @Produce		json
//
// @Router			/account/ [get]
func (as *AccountService) GetAllAccount(ctx *fiber.Ctx) error {
	var accounts []entity.Account
	result := as.db.Find(&accounts)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	var accountModel []model.AccountInfo
	for _, account := range accounts {
		accountModel = append(accountModel, model.AccountInfo{
			ID:      account.ID,
			UserID:  account.UserID,
			Type:    account.Type,
			Balance: account.Balance,
			Pin:     account.Pin,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"accounts": accountModel,
	})
}

// GetAccountByID godoc
//
// @Summary		Get account by ID
// @Description	Get account by ID
// @Tags			account
// @Accept			json
// @Produce		json
//
// @Param			id path string true "Account ID"
// @Router			/account/{id} [get]
func (as *AccountService) GetAccountByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var account entity.Account
	result := as.db.Where("id = ?", id).First(&account)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})
}

// CreateAccount godoc
//
// @Summary		Create account
// @Description	Create account
// @Tags			account
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			account body model.CreateAccount true "Account information"
// @Router			/account/ [post]
func (as *AccountService) CreateAccount(ctx *fiber.Ctx) error {
	var createAccount model.CreateAccount
	if err := ctx.BodyParser(&createAccount); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//get from header
	tokenHeader := ctx.Get("Authorization")
	if tokenHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	fmt.Println("Token: ", tokenHeader)

	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	userID := claims["id"].(string)

	hasher := shared.NewLocalConfig(10)
	hashedPin, err := hasher.HashPassword(createAccount.Pin)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	account := entity.Account{
		UserID:  userID,
		Type:    createAccount.Type,
		Balance: 0.00,
		Pin:     hashedPin,
	}

	result := as.db.Create(&account)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})
}
