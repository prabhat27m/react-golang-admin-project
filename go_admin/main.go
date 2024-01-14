package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/database"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.SetUp(app)
	
}
