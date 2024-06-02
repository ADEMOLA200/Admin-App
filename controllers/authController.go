package controllers

import (
	_"github.com/ADEMOLA200/Admin-App.git/models"
	"github.com/gofiber/fiber/v2"
)

func Register(ac *fiber.Ctx) error {
	var data map[string]string

	if err := ac.BodyParser(&data); err != nil {
		return err
	}
}
