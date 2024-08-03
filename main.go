package main

import (
	"basic-auth/database"
	"basic-auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connection()

	app := fiber.New()

	routes.SetUp(app)

	app.Listen(":3000")
}
