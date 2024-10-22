package logic

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
)

// ForwardGetAllPayment godoc
//
// @Summary		Forward get all payment request to payment service
// @Description	Forward get all payment request to payment service
// @Tags			payment
// @Accept			json
// @Produce		json
//
// @Router			/payment/ [get]
func ForwardGetAllPayment(ctx *fiber.Ctx) error {
	fmt.Println("calling payment service")
	resp, err := http.Get("http://payment-service:3400/api/v1/payment/")
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

// ForwardCreatePayment godoc
//
// @Summary		Forward create payment request to payment service
// @Description	Forward create payment request to payment service
// @Tags			payment
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			payment	body		model.CreatePayment	true	"Payment information"
// @Success		201		{string}	string			"Payment created successfully"
// @Failure		400		{string}	string			"Bad request"
//
// @Router			/payment/ [post]
func ForwardCreatePayment(ctx *fiber.Ctx) error {
	fmt.Println("calling payment service")
	tokenString := ctx.Cookies("Authorization")

	payment := &http.Client{}
	req, err := http.NewRequest("POST", "http://payment-service:3400/api/v1/payment/", bytes.NewReader(ctx.Body()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req.Header)

	resp, err := payment.Do(req)
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

// ForwardGetPaymentByID godoc
//
// @Summary		Forward get payment by ID request to payment service
// @Description	Forward get payment by ID request to payment service
// @Tags			payment
// @Accept			json
// @Produce		json
// @security		Bearer
// @Param			id	path	string	true	"Payment ID"
// @Success		200		{string}	string	"Payment information"
// @Failure		400		{string}	string	"Bad request"
//
// @Router			/payment/{id} [get]
func ForwardGetPaymentByID(ctx *fiber.Ctx) error {
	fmt.Println("calling payment service")
	tokenString := ctx.Cookies("Authorization")

	payment := &http.Client{}
	req, err := http.NewRequest("GET", "http://payment-service:3400/api/v1/payment/"+ctx.Params("id"), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req.Header)

	resp, err := payment.Do(req)
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
