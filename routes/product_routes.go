package routes

import (
	"github.com/ADEMOLA200/Admin-App.git/controllers"
	"github.com/ADEMOLA200/Admin-App.git/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProductSetup(ur *fiber.App) {
	ur.Get("/api/products", controllers.GetAllProducts, middlewares.IsAuthenticated)

	user := ur.Group("/api/product")
	user.Use(middlewares.IsAuthenticated)
	{
		user.Post("/create", controllers.CreateProduct)
		user.Get("/get-product/:id", controllers.GetProductById)
		user.Put("/update/:id", controllers.UpdateProductById)
		user.Delete("/delete/:id", controllers.DeleteProduct)
	}
}
