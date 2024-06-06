package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserSetup(ur *fiber.App) {
	ur.Get("/api/users", controllers.GetAllUsers)

	user := ur.Group("/api/user")
	user.Use(middlewares.IsAuthenticated)
	{
		user.Post("/create", controllers.CreateUser)
		user.Get("/get-user-id/:id", controllers.GetUserById)
		user.Put("/update/:id", controllers.UpdateUserById)
		user.Delete("/delete/:id", controllers.DeleteUser)
	}
}
