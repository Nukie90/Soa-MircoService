package logic

import (
	"bytes"
	"fmt"
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
	resp, err := http.Post("http://auth-service:3500/api/v1/auth/signup", "application/json", body)
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
	fmt.Println("call login")
	body := bytes.NewReader(ctx.Body())
	resp, err := http.Post("http://auth-service:3500/api/v1/auth/login", "application/json", body)
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
// @Router			/users/all [get]
func ForwardGetAllUser(ctx *fiber.Ctx) error {
	fmt.Println("call all user")
	resp, err := http.Get("http://user-service:3100/api/v1/users")
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

// ForwardGetUserByID godoc
//
// @Summary		Forward get user by ID request to user service
// @Description	Forward get user by ID request to user service
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"User ID"
// @Success		200		{string} model.UserModel	"User"
// @Failure		400		{string}	string		"Bad request"
//
// @Router			/users/all/{id} [get]
func ForwardGetUserByID(ctx *fiber.Ctx) error {
	resp, err := http.Get("http://user-service:3100/api/v1/users/" + ctx.Params("id"))
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

// ForwardGetMe godoc
//
// @Summary		Forward get me request to user service
// @Description	Forward get me request to user service
// @Tags			user
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200		{string} model.UserModel	"User"
// @Failure		400		{string}	string		"Bad request"
//
// @Router			/users/me [get]
func ForwardGetMe(ctx *fiber.Ctx) error {
	fmt.Println("call me")
	cookie := ctx.Cookies("Authorization")

	userService := &http.Client{}
	req, err := http.NewRequest("GET", "http://user-service:3100/api/v1/users/me", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", cookie)

	resp, err := userService.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return ctx.Status(resp.StatusCode).Send(body)
	}

	return ctx.Status(resp.StatusCode).Send(body)
}
