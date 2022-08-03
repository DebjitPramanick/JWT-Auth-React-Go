package main

import (
	"server/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	database.Connect()

	app := fiber.New()
	app.Get("/healthz", func (c *fiber.Ctx) error {
		return c.SendString("APP IS RUNNING.")
	})

	app.Listen(":4000")
}