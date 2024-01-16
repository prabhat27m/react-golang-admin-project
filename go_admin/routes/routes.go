package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/controllers"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/middleware"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Use(middleware.IsAuthenticated)
	app.Get("/api/authenticate", controllers.User)
	app.Get("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.GetAllUsers)
	app.Post("/api/create/user", controllers.CreateUser)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Listen(":3000")
}
