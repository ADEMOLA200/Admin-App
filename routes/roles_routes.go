package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RolesSetup(rr *fiber.App) {
	rr.Get("/api/roles", controllers.GetAllRoles, middlewares.IsAuthenticated)

	user := rr.Group("/api/role")
	user.Use(middlewares.IsAuthenticated)
	{
		user.Post("/create", controllers.CreateRole)
		user.Get("/get-role/:id", controllers.GetRoleById)
		user.Put("/update/:id", controllers.UpdateRoleById)
		user.Delete("/delete/:id", controllers.DeleteRole)
	}
}