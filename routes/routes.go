package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(r *fiber.App) {
	auth := r.Group("/auth")
	{
		auth.Post("/register", controllers.Register)
		auth.Post("/login", controllers.Login)
	}

	user := r.Group("/api")
	user.Use(middlewares.IsAuthenticated)
	{
		user.Get("/user", controllers.GetUser)
		user.Post("/logout", controllers.Logout)
	}
}
