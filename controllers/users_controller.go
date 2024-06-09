package controllers

import (
	"math"
	"strconv"

	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(uc *fiber.Ctx) error {
	page, _ := strconv.Atoi(uc.Query("page", "1")) 

	limit := 5

	offset := (page - 1) * limit

	var total int64
	var users []models.User

	if err := database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get users",
			"success": false,
		})
	}

	if err := database.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get users",
			"success": false,
		})
	}

	lastPage := math.Ceil(float64(float64(total) / float64(limit)))

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all users successfully",
		"users": users,
		"meta": fiber.Map{
			"total": total,
			"page": page,
			"last_page": lastPage,
		},
		"success": true,
	})
}

func CreateUser(uc *fiber.Ctx) error {
	var user models.User

	if err := uc.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	if err := database.DB.Preload("Role").Create(&user).Error; err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not create user, email already exist",
			"success": false,
		})
	}

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user created successfully",
		"user": user,
		"success": true,
	})
}

func GetUserById(uc *fiber.Ctx) error {
    id, err := strconv.Atoi(uc.Params("id"))
    if err != nil {
        return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "invalid id, id must be an integer value",
            "success": false,
        })
    }

    var user models.User
    result := database.DB.Preload("Role").Find(&user, uint(id))
    if result.Error != nil {
        return uc.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "user not found with the id of " + strconv.Itoa(id),
            "success": false,
        })
    }

    return uc.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "user successfully found",
        "user": user,
        "success": true,
    })
}

func UpdateUserById(uc *fiber.Ctx) error {
	id, err := strconv.Atoi(uc.Params("id"))
    if err!= nil {
        return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "invalid id, id must be an integer value",
            "success": false,
        })
    }

	var user models.User
	result := database.DB.Preload("Role").First(&user, uint(id))
    if result.Error!= nil {
        return uc.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "user not found with the id of " + strconv.Itoa(id),
            "success": false,
        })
    }

	if err := uc.BodyParser(&user); err!= nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "updated user successfully",
        "user": user,
        "success": true,
    })
}

func DeleteUser(uc *fiber.Ctx) error {
	id, err := strconv.Atoi(uc.Params("id"))
	if err != nil {
		return uc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id, id must be an integer value",
			"success": false,
		})
	}

	var user models.User
	result := database.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return uc.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error occurred while deleting user",
			"success": false,
		})
	}

	if result.RowsAffected == 0 {
		return uc.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "could not delete user with the id of " + strconv.Itoa(id),
			"success": false,
		})
	}

	return uc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted user with id " + strconv.Itoa(id) + " successfully",
		"success": true,
	})
}
