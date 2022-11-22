package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

// hi :)

type createPaymentIntentReq struct {
	PaymentType string `json:"paymentType"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
}

const (
	unDollares    string = "unDollares"
	cincoDollares string = "cincoDollares"
	diazDollares  string = "diazDollares"
	manyDollares  string = "manyDollares"
)

func CreateCheckoutSession(c *fiber.Ctx) error {
	req := createPaymentIntentReq{}

	// unmarshalling the request into the req
	err := c.BodyParser(&req)

	if err != nil {
		return c.JSON(err)
	}

	amount := 0

	fmt.Println(req)

	switch req.Amount {
	case unDollares:
		amount = 1000
		break
	case cincoDollares:
		amount = 5000
		break
	case diazDollares:
		amount = 10000
		break
	case manyDollares:
		amount = 20000
		break
	}

	// This is your test secret API key.
	stripe.Key = "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"

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
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
