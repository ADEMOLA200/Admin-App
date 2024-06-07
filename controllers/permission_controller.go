package controllers

import (
	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetPermissions(pc *fiber.Ctx) error {
	var permissions []models.Permissions

	if err := database.DB.Find(&permissions).Error; err != nil {
		return pc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get users",
			"success": false,
		})
	}

	return pc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "permissions fetched successfully",
		"users": permissions,
		"success": true,
	})
}
 