package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RolesSetup(ur *fiber.App) {
	ur.Get("/api/roles", controllers.GetAllRoles)

	user := ur.Group("/api/role")
	user.Use(middlewares.IsAuthenticated)
	{
		user.Post("/create", controllers.CreateRole)
		user.Get("/get-role/:id", controllers.GetRoleById)
		user.Put("/update/:id", controllers.UpdateRoleById)
		user.Delete("/delete/:id", controllers.DeleteRole)
	}
}