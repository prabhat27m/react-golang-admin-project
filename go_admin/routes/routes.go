package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/controllers"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Listen(":3000")
}
