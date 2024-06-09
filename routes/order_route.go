package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func OrderSetup(ur *fiber.App) {
	ur.Get("/api/orders", controllers.GetAllOrders, middlewares.IsAuthenticated)
}
