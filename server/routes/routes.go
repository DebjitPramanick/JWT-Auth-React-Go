package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	
	app.Get("/healthz", func (c *fiber.Ctx) error {
		return c.SendString("APP IS RUNNING.")
	})
}