package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

// hi :)

type createPaymentIntentReq struct {
	PaymentType string `json:"paymentType"`
	Currency string `json:"currency"`
	Amount int64 `json:"amount"`

}



func CreateCheckoutSession(c *fiber.Ctx) error {
	req := createPaymentIntentReq{}

	// unmarshalling the request into the req 
	err := c.BodyParser(&req)

	if err != nil {
		return c.SendStatus(500)
	}

	fmt.Println(req)
	if req.Amount <= 1998 {
		return c.SendString("amount cannot be less than zero you punk bitch")
	}

	// This is your test secret API key.
	stripe.Key = os.Getenv("STRIPE_API_KEY")

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(req.Amount),
		Currency: stripe.String(req.Currency),
		PaymentMethodTypes: stripe.StringSlice([]string{
			req.PaymentType,
		}),
	}
	pi, err := paymentintent.New(params)

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(pi.ClientSecret)
}