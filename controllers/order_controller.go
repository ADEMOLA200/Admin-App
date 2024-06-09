package controllers

import (
	"math"
	"strconv"

	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrders(oc *fiber.Ctx) error {
	page, _ := strconv.Atoi(oc.Query("page", "1")) 

	limit := 5

	offset := (page - 1) * limit

	var total int64
	var orders []models.Order

	if err := database.DB.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return oc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get orders",
			"success": false,
		})
	}

	if err := database.DB.Model(&models.Order{}).Count(&total).Error; err != nil {
		return oc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get orders",
			"success": false,
		})
	}

	lastPage := math.Ceil(float64(float64(total) / float64(limit)))

	return oc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all orders successfully",
		"orders": orders,
		"meta": fiber.Map{
			"total": total,
			"page": page,
			"last_page": lastPage,
		},
		"success": true,
	})
}
