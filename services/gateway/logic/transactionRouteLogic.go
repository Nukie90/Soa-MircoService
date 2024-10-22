package logic

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
)

// ForwardGetAllTransaction godoc
//
// @Summary		Forward get all transaction request to transaction service
// @Description	Forward get all transaction request to transaction service
// @Tags			transaction
// @Accept			json
// @Produce		json
//
// @Router			/transaction/ [get]
func ForwardGetAllTransaction(ctx *fiber.Ctx) error {
	fmt.Println("calling transaction service")
	resp, err := http.Get("http://transaction-service:3300/api/v1/transaction/")
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

// ForwardCreateTransaction godoc
//
// @Summary		Forward create transaction request to transaction service
// @Description	Forward create transaction request to transaction service
// @Tags			transaction
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			transaction	body		model.CreateTransaction	true	"Transaction information"
// @Success		201		{string}	string			"Transaction created successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/transaction/ [post]
func ForwardCreateTransaction(ctx *fiber.Ctx) error {
	fmt.Println("calling transaction service")
	tokenString := ctx.Cookies("Authorization")

	transaction := &http.Client{}
	req, err := http.NewRequest("POST", "http://transaction-service:3300/api/v1/transaction/", bytes.NewReader(ctx.Body()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req.Header)

	resp, err := transaction.Do(req)
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

// ForwardGetTransactionByID godoc
//
// @Summary		Forward get transaction by ID request to transaction service
// @Description	Forward get transaction by ID request to transaction service
// @Tags			transaction
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			id	path	string	true	"Transaction ID"
// @Success		200		{string}	string	"Transaction information"
// @Failure		400		{string}	string	"Bad request"
//
// @Router			/transaction/{id} [get]
func ForwardGetTransactionByID(ctx *fiber.Ctx) error {
	fmt.Println("calling transaction service")
	tokenString := ctx.Cookies("Authorization")

	transaction := &http.Client{}
	req, err := http.NewRequest("GET", "http://transaction-service:3300/api/v1/transaction/"+ctx.Params("id"), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req.Header)

	resp, err := transaction.Do(req)
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
