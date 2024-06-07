package controllers

import (
	"strconv"
	"time"

	"github.com/ADEMOLA200/Admin-App.git/database"
	"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/ADEMOLA200/Admin-App.git/utils"
	"github.com/gofiber/fiber/v2"
)

func Register(ac *fiber.Ctx) error {
	var data map[string]string

	if err := ac.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return ac.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "passwords do not match",
			"success": false,
		})
	}

	user := models.User {
		FirstName: 	data["first_name"],
		LastName: 	data["last_name"],
		Email: 		data["email"],
		RoleId: 	1,
	}

	user.SetPassword(data["password"])

	if err := database.DB.Create(&user).Error; err != nil {
		return ac.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create user",
			"success": false,
		})
	}

	return ac.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user registered successfully",
		"user": user,
		"success": true,
	})
}

func Login(ac *fiber.Ctx) error {
	var data map[string]string

	if err := ac.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {
		return ac.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user not found",
			"success": false,
		})
	}

	if err := user.ComparePasswords(data["password"]); err != nil {
		return ac.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect password",
			"success": false,
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return ac.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie {
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ac.Cookie(&cookie)

	return ac.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged in successfully",
		"user": user,
		"success": true,
	})
}

func GetUser (ac *fiber.Ctx) error {
	cookie := ac.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return ac.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get user",
			"success": false,
		})
	}

	return ac.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully get user ",
		"user": user,
		"success": true,
	})
}

func Logout (ac *fiber.Ctx) error {
	cookie := fiber.Cookie {
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ac.Cookie(&cookie)

	return ac.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user logged out successfully",
	})
}
