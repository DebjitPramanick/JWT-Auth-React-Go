package main

import (
	"server/database"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":4000")
}