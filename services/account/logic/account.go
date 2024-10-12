package logic

import (
	"errors"
	"fmt"
	"microservice/entity"
	"microservice/services/account/model"
	"microservice/shared"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
	js nats.JetStreamContext
}

func NewAccountService(db *gorm.DB, nc *nats.Conn) (*AccountService, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return &AccountService{db: db, js: js}, nil
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

	//get the account by time created
	var newAccount entity.Account
	result = as.db.Where("created_at = ?", account.CreatedAt).First(&newAccount)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	event := map[string]interface{}{
		"ID":      newAccount.ID,
		"userID":  account.UserID,
		"type":    account.Type,
		"balance": account.Balance,
		"pin":     account.Pin,
	}

	// Publish message to NATS
	eventData, err := shared.MarshalToJSON(event)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not marshal event to JSON",
		})
	}

	if _, err := as.js.Publish("account.created", eventData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not publish event",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Account created successfully",
	})
}

// TopUp godoc
//
// @Summary		Top up account
// @Description	Top up account
// @Tags			account
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			topUp body model.TopUp true "Top up information"
// @Router			/account/topup [put]
func (as *AccountService) TopUp(ctx *fiber.Ctx) error {
	var topUp model.TopUp
	if err := ctx.BodyParser(&topUp); err != nil {
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

	var account entity.Account
	result := as.db.Where("id = ? and user_id = ?", topUp.ID, userID).First(&account)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	account.Balance += topUp.Amount

	result = as.db.Save(&account)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	event := map[string]interface{}{
		"ID":      account.ID,
		"userID":  account.UserID,
		"type":    account.Type,
		"balance": account.Balance,
		"pin":     account.Pin,
	}

	// Publish message to NATS
	eventData, err := shared.MarshalToJSON(event)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not marshal event to JSON",
		})
	}

	if _, err := as.js.Publish("account.updated", eventData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not publish event",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Top up success",
	})
}

// DeleteAccount godoc
//
// @Summary      Delete account
// @Description  Delete account by verifying ID and PIN
// @Tags         account
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        deleteAccount body model.DeleteAccount true "Delete account information"
// @Router       /account/delete [delete]
func (as *AccountService) DeleteAccount(ctx *fiber.Ctx) error {
	var deleteRequest model.DeleteAccount

	if err := ctx.BodyParser(&deleteRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get the token from the header
	tokenHeader := ctx.Get("Authorization")
	if tokenHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

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

	// Fetch the account from the database
	var account entity.Account
	result := as.db.Where("id = ? AND user_id = ?", deleteRequest.ID, userID).First(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Account not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Verify the PIN
	if !shared.CheckPasswordHash(deleteRequest.Pin, account.Pin) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid PIN",
		})
	}

	// Delete the account
	if err := as.db.Delete(&account).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Publish event to NATS
	event := map[string]interface{}{
		"ID":     account.ID,
		"userID": account.UserID,
		"type":   account.Type,
	}

	eventData, err := shared.MarshalToJSON(event)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not marshal event to JSON",
		})
	}

	if _, err := as.js.Publish("account.deleted", eventData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not publish event",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Account deleted successfully",
	})
}

// ChangePin godoc
//
// @Summary     Change account PIN
// @Description Change the PIN of the account
// @Tags        account
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       changePin body model.ChangePinRequest true "Change PIN request"
// @Router      /account/change-pin [put]
func (as *AccountService) ChangePin(ctx *fiber.Ctx) error {
	var changePinRequest model.ChangePinRequest

	if err := ctx.BodyParser(&changePinRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get the token from the header
	tokenHeader := ctx.Get("Authorization")
	if tokenHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

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

	// Fetch the account
	var account entity.Account
	result := as.db.Where("id = ? AND user_id = ?", changePinRequest.ID, userID).First(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Account not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Verify the old PIN
	if !shared.CheckPasswordHash(changePinRequest.OldPin, account.Pin) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid old PIN",
		})
	}

	// Hash the new PIN
	hasher := shared.NewLocalConfig(10)
	hashedNewPin, err := hasher.HashPassword(changePinRequest.NewPin)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update the account's PIN
	account.Pin = hashedNewPin
	result = as.db.Save(&account)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Publish NATS event
	event := map[string]interface{}{
		"ID":     account.ID,
		"userID": account.UserID,
		"pin":    account.Pin,
		"type":   account.Type,
	}

	eventData, err := shared.MarshalToJSON(event)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not marshal event to JSON",
		})
	}

	if _, err := as.js.Publish("account.changedPin", eventData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not publish event",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "PIN changed successfully",
	})
}

// VerifyAccount godoc
//
// @Summary     Verify account
// @Description Verify the account by ID and PIN
// @Tags        account
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       verifyAccount body model.AccountVerify true "Verify account information"
// @Router      /account/verify [post]
func (as *AccountService) VerifyAccount(ctx *fiber.Ctx) error {
	var verifyRequest model.AccountVerify

	if err := ctx.BodyParser(&verifyRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get the token from the header
	tokenHeader := ctx.Get("Authorization")
	if tokenHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

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

	// Fetch the account
	var account entity.Account
	result := as.db.Where("id = ? AND user_id = ?", verifyRequest.ID, userID).First(&account)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Account not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Verify the PIN
	if !shared.CheckPasswordHash(verifyRequest.Pin, account.Pin) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid PIN",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Account verified",
	})
}
