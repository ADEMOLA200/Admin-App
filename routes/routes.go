package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(r *fiber.App) {
	r.Post("/api/register", controllers.Register)
	r.Post("/api/login", controllers.Login)
}