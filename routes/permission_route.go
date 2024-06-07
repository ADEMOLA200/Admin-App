package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PermissionsSetup(pr *fiber.App) {
	pr.Get("/api/get/permissions", controllers.GetPermissions, middlewares.IsAuthenticated)
}