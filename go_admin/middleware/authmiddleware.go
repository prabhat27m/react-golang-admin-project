package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")
	
	if _, err := utils.ParseJwtToken(cookie);err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated!!",
		})
	}
	return c.Next()
}