package controllers

import (
	"strconv"

	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllRoles(rc *fiber.Ctx) error {
	var roles []models.Role

	if err := database.DB.Find(&roles).Error; err != nil {
		return rc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get roles",
			"success": false,
		})
	}

	return rc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all roles successfully",
		"roles": roles,
		"success": true,
	})
}

func CreateRole(rc *fiber.Ctx) error {
	var role models.Role

	if err := rc.BodyParser(&role); err != nil {
		return err
	}

	if err := database.DB.Create(&role).Error; err != nil {
		return rc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not create role, email already exist",
			"success": false,
		})
	}

	return rc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "role created successfully",
		"role": role,
		"success": true,
	})
}

func GetRoleById(rc *fiber.Ctx) error {
    id, err := strconv.Atoi(rc.Params("id"))
    if err != nil {
        return rc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "invalid id, id must be an integer value",
            "success": false,
        })
    }

    var role models.Role
    result := database.DB.Find(&role, uint(id))
    if result.Error != nil {
        return rc.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "role not found with the id of " + strconv.Itoa(id),
            "success": false,
        })
    }

    return rc.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "role successfully found",
        "role": role,
        "success": true,
    })
}

func UpdateRoleById(rc *fiber.Ctx) error {
	id, err := strconv.Atoi(rc.Params("id"))
    if err!= nil {
        return rc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "invalid id, id must be an integer value",
            "success": false,
        })
    }

	var role models.Role
	result := database.DB.First(&role, uint(id))
    if result.Error!= nil {
        return rc.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "role not found with the id of " + strconv.Itoa(id),
            "success": false,
        })
    }

	if err := rc.BodyParser(&role); err!= nil {
		return err
	}

	database.DB.Model(&role).Updates(role)

	return rc.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "updated role successfully",
        "role": role,
        "success": true,
    })
}

func DeleteRole(rc *fiber.Ctx) error {
	id, err := strconv.Atoi(rc.Params("id"))
	if err != nil {
		return rc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id, id must be an integer value",
			"success": false,
		})
	}

	var role models.Role
	result := database.DB.Where("id = ?", id).Delete(&role)
	if result.Error != nil {
		return rc.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error occurred while deleting role",
			"success": false,
		})
	}

	if result.RowsAffected == 0 {
		return rc.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "could not delete role with the id of " + strconv.Itoa(id),
			"success": false,
		})
	}

	return rc.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted role with id " + strconv.Itoa(id) + " successfully",
		"success": true,
	})
}
