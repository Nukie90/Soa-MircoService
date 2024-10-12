package logic

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
)

// ForwardGetAllAccount godoc
//
// @Summary		Forward get all account request to account service
// @Description	Forward get all account request to account service
// @Tags			account
// @Accept			json
// @Produce		json
//
// @Router			/account/ [get]
func ForwardGetAllAccount(ctx *fiber.Ctx) error {
	fmt.Println("calling account service")
	resp, err := http.Get("http://account-service:3200/api/v1/account/")
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

func ForwardGetAccountByID(ctx *fiber.Ctx) error {
	return nil
}

// ForwardCreateAccount godoc
//
// @Summary		Forward create account request to account service
// @Description	Forward create account request to account service
// @Tags			account
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			account	body		model.CreateAccount	true	"Account information"
// @Success		201		{string}	string			"Account created successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/account/ [post]
func ForwardCreateAccount(ctx *fiber.Ctx) error {
	fmt.Println("calling account service")
	tokenString := ctx.Cookies("Authorization")

	account := &http.Client{}
	req, err := http.NewRequest("POST", "http://account-service:3200/api/v1/account/", bytes.NewReader(ctx.Body()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req.Header)

	resp, err := account.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Forward the response with the same status code
	return ctx.Status(resp.StatusCode).Send(body)
}

// ForwardTopUp godoc
//
// @Summary		Forward top up request to account service
// @Description	Forward top up request to account service
// @Tags			account
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			account	body		model.TopUp	true	"Top up information"
// @Success		200		{string}	string			"Top up successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/account/topup [put]
func ForwardTopUp(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("Authorization")

	account := &http.Client{}
	req, err := http.NewRequest("PUT", "http://account-service:3200/api/v1/account/topup", bytes.NewReader(ctx.Body()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")

	resp, err := account.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ctx.Status(resp.StatusCode).Send(body)
}
