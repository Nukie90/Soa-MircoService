package logic

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
)

// ForwardSignup godoc
//
// @Summary		Forward signup request to user service
// @Description	Forward signup request to user service
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			user	body		model.SignUp	true	"User information"
// @Success		201		{string}	string			"User created successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/auth/signup [post]
func ForwardSignup(ctx *fiber.Ctx) error {
	body := bytes.NewReader(ctx.Body())
	resp, err := http.Post("http://localhost:3100/api/v1/auth/signup", "application/json", body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	// Return response from user service
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Status(resp.StatusCode)
	ctx.Set("Content-Type", resp.Header.Get("Content-Type"))
	return ctx.Send(respBody)
}

// ForwardLogin godoc
//
// @Summary		Forward login request to user service
// @Description	Forward login request to user service
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			user	body		model.Login	true	"User information"
// @Success		200		{string}	string			"User logged in successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/auth/login [post]
func ForwardLogin(ctx *fiber.Ctx) error {
	body := bytes.NewReader(ctx.Body())
	resp, err := http.Post("http://localhost:3100/api/v1/auth/login", "application/json", body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	// Return response from user service
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Status(resp.StatusCode)

	token := resp.Header.Get("Authorization")
	ctx.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    token,
		HTTPOnly: true,
	})

	return ctx.Send(respBody)
}

// GetToken godoc
//
// @Summary		Get token
// @Description	Get token from header
// @Tags			auth
// @Accept			json
// @Produce		json
// @Success		200		{string}	string		"Token"
// @Failure		400		{string}	string		"Bad request"
//
// @Router			/auth/token [get]
func GetToken(ctx *fiber.Ctx) error {
	token := ctx.Cookies("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// ForwardGetAllUser godoc
//
// @Summary		Forward get all users request to user service
// @Description	Forward get all users request to user service
// @Tags			user
// @Accept			json
// @Produce		json
// @Success		200		{string} model.UserModel	"Users"
// @Failure		400		{string}	string		"Bad request"
//
// @Router			/users [get]
func ForwardGetAllUser(ctx *fiber.Ctx) error {
	resp, err := http.Get("http://localhost:3100/api/v1/users")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	// Return response from user service
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Status(resp.StatusCode)
	ctx.Set("Content-Type", resp.Header.Get("Content-Type"))
	return ctx.Send(respBody)
}
