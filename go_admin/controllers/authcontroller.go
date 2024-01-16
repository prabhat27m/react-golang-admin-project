package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/database"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/models"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}
	

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	user.GeneratePassword([]byte(data["password"]))
	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User Not Found!",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password!",
		})
	}

	token, _ := utils.GenerateJwtToken(fmt.Sprint(user.Id))

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success!!",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	Issuer, _ := utils.ParseJwtToken(cookie)

	var user models.User
	database.DB.Where("id =?", Issuer).First(&user)

	return c.JSON(fiber.Map{
		"message": "Authenticated!!",
		"user_id": Issuer,
		"user":    user,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "token",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Successfully Logged Out!!",
	})
}
