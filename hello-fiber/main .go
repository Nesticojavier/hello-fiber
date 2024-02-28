package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	Firstname string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "John",
		Lastname:  "Doe",
	}
	return c.Status(200).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	return c.Status(200).JSON(user)
}

func main() {
	app := fiber.New()

	// middlewares
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	userGroup := app.Group("/user")
	userGroup.Get("", handleUser)
	userGroup.Post("", handleCreateUser)

	app.Listen(":3000")
}
