package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/database"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	
	routes.SetUp(app)
	
}
