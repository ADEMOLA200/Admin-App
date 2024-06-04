package controllers

import (
	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(uc *fiber.Ctx) error {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get users",
			"success": false,
		})
	}

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all users successfully",
		"users": users,
		"success": true,
	})
}