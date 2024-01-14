package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/models"
)

func Register(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data);  err != nil{
		return err
	}
	if data["password"]!= data["password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"password do not match",
		})
	}
	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
	}
	return c.JSON(user)
}
