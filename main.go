package main

import (
	"basic-auth/database"
	"basic-auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connection()

	app := fiber.New()

	app.Use(cors.New())

	routes.SetUp(app)

	app.Listen(":3000")
}
