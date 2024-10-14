package logic

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

// ForwardGetAccountByID godoc
//
// @Summary		Forward get account by ID request to account service
// @Description	Forward get account by ID request to account service
// @Tags			account
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"User ID"
// @Success		200		{string}	string	"Account"
// @Failure		400		{string}	string	"Bad request"
//
// @Router			/account/{id} [get]
func ForwardGetAccountByID(ctx *fiber.Ctx) error {
	fmt.Println("calling account service")
	resp, err := http.Get("http://account-service:3200/api/v1/account/" + ctx.Params("id"))
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

// ForwardChangePin godoc
//
// @Summary		Forward change pin request to account service
// @Description	Forward change pin request to account service
// @Tags			account
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			account	body		model.ChangePin	true	"Change pin information"
// @Success		200		{string}	string			"Pin changed successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/account/ [put]
func ForwardChangePin(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("Authorization")

	account := &http.Client{}
	req, err := http.NewRequest("PUT", "http://account-service:3200/api/v1/account/", bytes.NewReader(ctx.Body()))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ctx.Status(resp.StatusCode).Send(body)
}

// ForwardDeleteAccount godoc
//
// @Summary		Forward delete account request to account service
// @Description	Forward delete account request to account service
// @Tags			account
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			account	body		model.DeleteAccount	true	"Delete account information"
// @Success		200		{string}	string			"Account deleted successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/account/ [delete]
func ForwardDeleteAccount(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("Authorization")

	account := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://account-service:3200/api/v1/account/", bytes.NewReader(ctx.Body()))
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

// ForwardVerifyAccount godoc
//
// @Summary		Forward verify account request to account service
// @Description	Forward verify account request to account service
// @Tags			account
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			account	body		model.AccountVerify	true	"Account verify information"
// @Success		200		{string}	string			"Account verified successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/account/verify [post]
func ForwardVerifyAccount(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("Authorization")

	account := &http.Client{}
	req, err := http.NewRequest("POST", "http://account-service:3200/api/v1/account/verify", bytes.NewReader(ctx.Body()))
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

// ForwardGetAccountsByUserID godoc
//
// @Summary      Forward get accounts by user ID request to account service
// @Description  Forward get accounts by user ID request to account service
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200      {string}  string  "List of accounts"
// @Failure      400      {string}  string  "Bad request"
// @Failure      404      {string}  string  "Accounts not found"
// @Router       /account/getAccountsByUserID [get]
func ForwardGetAccountsByUserID(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("Authorization")

	account := &http.Client{}
	req, err := http.NewRequest("GET", "http://account-service:3200/api/v1/account/getAccountsByUserID", bytes.NewReader(ctx.Body()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")
	defer req.Body.Close()

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
