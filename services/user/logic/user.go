package logic

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"microservice/entity"
	"microservice/services/user/model"
	"microservice/shared"
	"os"
	"time"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
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
func (us *UserService) SignUp(ctx *fiber.Ctx) error {
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

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
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
func (us *UserService) Login(ctx *fiber.Ctx) error {
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
		Name:    "jwt",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
	})

	//set auth token from user service to gateway in header
	ctx.Set("Authorization", tokenString)

	//set local
	ctx.Locals("userId", user.ID)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
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
