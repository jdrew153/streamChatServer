package main

import (
	"fmt"
	stream "github.com/GetStream/stream-chat-go/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"streamChatServer/handlers"
	"time"
)

type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server is up and running")
	})
	app.Get("/get-all-users", handlers.ListAllUsers)

	app.Post("/create-new-user", handlers.CreateNewUser)
	app.Post("/login-user", handlers.LogInUser)
	app.Post("/create-payment-intent", handlers.CreateCheckoutSession)

	app.Post("/create-stream-user", func(c *fiber.Ctx) error {
		client, err := stream.NewClient("yt87ffbuwxzy", "wvjvm4rbhkteq4sx4d2rx5cvg82e7zcau5huubwght3m6x5mfvs9us6ku3rxsgf8")
		if err != nil {
			return err
		}
		var user = User{}

		err = c.BodyParser(&user)

		if err != nil {
			return err
		}
		fmt.Println(user)
		token, err := client.CreateToken(user.UserId, time.Time{})

		if err != nil {
			return err
		}
		fmt.Println(token)
		return c.SendString(token)
	})

	log.Fatal(app.Listen(":8000"))
}
