package controllers

import (
	"basic-auth/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

		var data map[string]string

		password , _:= bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
		
		user := models.User{
			Name: data["name"],
			Email: data["email"],
			Password: password,
		}


		if err := c.BodyParser(&data); err != nil {
			return err
		}

		return c.JSON(user)
	}
