package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(r *fiber.App) {
	auth := r.Group("/auth")
	{
		auth.Post("/register", controllers.Register)
		auth.Post("/login", controllers.Login)
	}

	r.Get("/api/user", controllers.User)
}