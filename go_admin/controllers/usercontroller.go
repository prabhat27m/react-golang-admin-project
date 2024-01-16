package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/database"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/models"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	user.Password = password
	database.DB.Create(&user)
	return c.JSON(user)

}
func GetUser(c *fiber.Ctx) error{
	id,_ := strconv.Atoi(c.Params("id"))
	user:= models.User{
		Id: uint(id),
	}
	database.DB.Find(&user)
	return c.JSON(user)
}