package main

import (
	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	
	routes.Setup(app)

	app.Listen(":9000")
}
 