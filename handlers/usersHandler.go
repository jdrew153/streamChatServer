package handlers

import (
	"github.com/gofiber/fiber/v2"
	"streamChatServer/services"
	"streamChatServer/types"
)

func ListAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()

	if err != nil {
		return err
	}

	return c.JSON(users)
}

func CreateNewUser(c *fiber.Ctx) error {
	var user types.User
	err := c.BodyParser(&user)

	if err != nil {
		return err
	}
	result, err := services.CreateNewUser(user)

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func LogInUser(c *fiber.Ctx) error {

	var user types.User

	err := c.BodyParser(&user)

	if err != nil {
		return err
	}

	result, err := services.LoginUser(user)

	if err != nil {
		return err
	}

	return c.JSON(result)
}
