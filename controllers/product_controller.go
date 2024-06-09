package controllers

import (
	"math"
	"strconv"

	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(uc *fiber.Ctx) error {
	page, _ := strconv.Atoi(uc.Query("page", "1")) 

	limit := 5

	offset := (page - 1) * limit

	var total int64
	var products []models.Product

	if err := database.DB.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get products",
			"success": false,
		})
	}

	if err := database.DB.Model(&models.Product{}).Count(&total).Error; err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get products",
			"success": false,
		})
	}

	lastPage := math.Ceil(float64(float64(total) / float64(limit)))

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all products successfully",
		"products": products,
		"meta": fiber.Map{
			"total": total,
			"page": page,
			"last_page": lastPage,
		},
		"success": true,
	})
}

func CreateProduct(uc *fiber.Ctx) error {
	var product models.Product

	if err := uc.BodyParser(&product); err != nil {
		return err
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not create Product, email already exist",
			"success": false,
		})
	}

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product created successfully",
		"Product": product,
		"success": true,
	})
}

func GetProductById(uc *fiber.Ctx) error {
    id, err := strconv.Atoi(uc.Params("id"))
    if err != nil {
        return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "invalid id, id must be an integer value",
            "success": false,
        })
    }

    var product models.Product
    result := database.DB.Find(&product, uint(id))
    if result.Error != nil {
        return uc.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Product not found with the id of " + strconv.Itoa(id),
            "success": false,
        })
    }

    return uc.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Product successfully found",
        "Product": product,
        "success": true,
    })
}

func UpdateProductById(uc *fiber.Ctx) error {
	id, err := strconv.Atoi(uc.Params("id"))
    if err!= nil {
        return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "invalid id, id must be an integer value",
            "success": false,
        })
    }

	var product models.Product
	result := database.DB.First(&product, uint(id))
    if result.Error!= nil {
        return uc.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Product not found with the id of " + strconv.Itoa(id),
            "success": false,
        })
    }

	if err := uc.BodyParser(&product); err!= nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "updated Product successfully",
        "Product": product,
        "success": true,
    })
}

func DeleteProduct(uc *fiber.Ctx) error {
	id, err := strconv.Atoi(uc.Params("id"))
	if err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id, id must be an integer value",
			"success": false,
		})
	}

	var product models.Product
	result := database.DB.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return uc.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error occurred while deleting Product",
			"success": false,
		})
	}

	if result.RowsAffected == 0 {
		return uc.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "could not delete Product with the id of " + strconv.Itoa(id),
			"success": false,
		})
	}

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted Product with id " + strconv.Itoa(id) + " successfully",
		"success": true,
	})
}
