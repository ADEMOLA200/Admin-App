package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(ar *fiber.App) {
	auth := ar.Group("/auth")
	{
		auth.Post("/register", controllers.Register)
		auth.Post("/login", controllers.Login)
	}

	user := ar.Group("/api")
	user.Use(middlewares.IsAuthenticated)
	{
		user.Put("/user/profile", controllers.UpdateProfile)
		user.Put("/user/change/password", controllers.ChangePassword)

		user.Get("/user", controllers.GetUser)
		user.Post("/user/logout", controllers.Logout)
	}
}
