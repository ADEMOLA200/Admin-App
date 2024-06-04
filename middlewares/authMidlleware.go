package middlewares

import (
	"github.com/ADEMOLA200/Admin-App.git/utils"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(am *fiber.Ctx) error {
	cookie := am.Cookies("jwt")

	if _, err := utils.ParseJwt(cookie); err != nil  {
		return am.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized user",
			"success": false,
		})
	}

	return am.Next()
}
