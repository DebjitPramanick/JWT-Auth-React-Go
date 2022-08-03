package main

import "github.com/gofiber/fiber/v2"

const CONNECTION_URL = "mongodb+srv://debjit:Gyb6XLHxpxvHsRgn@cluster0.c4dqrmw.mongodb.net/?retryWrites=true&w=majority"
const DB_NAME = "jwtAuth"
const COLLECTION_NAME = "users"

func main() {
	app := fiber.New()
	app.Get("/healthz", func (c *fiber.Ctx) error {
		return c.SendString("APP IS RUNNING.")
	})

	app.Listen(":4000")
}