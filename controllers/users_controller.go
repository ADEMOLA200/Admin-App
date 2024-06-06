package controllers

import (
	"strconv"

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

func CreateUser(uc *fiber.Ctx) error {
	var user models.User

	if err := uc.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

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
    result := database.DB.First(&user, uint(id))
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
	result := database.DB.First(&user, uint(id))
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
