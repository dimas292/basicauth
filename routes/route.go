package routes

import (
	"basic-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App){
	app.Get("/", controllers.HandlelerIndex)
}