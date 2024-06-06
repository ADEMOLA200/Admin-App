package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserSetup(ur *fiber.App) {
	users := ur.Group("/api")
	users.Use(middlewares.IsAuthenticated)
	{
		users.Get("/users", controllers.GetAllUsers)
		users.Post("/create-user", controllers.CreateUser)
		//users.Post("/test-route", controllers.TestRoute)
	}
}
