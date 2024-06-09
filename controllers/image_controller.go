package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func EnsureUploadsDir() error {
	uploadsDir := "./uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		return os.Mkdir(uploadsDir, os.ModePerm)
	}
	return nil
}

func UploadImage(c *fiber.Ctx) error {
	if err := EnsureUploadsDir(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create uploads directory",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form",
		})
	}

	files := form.File["image"]
	if files == nil || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No files uploaded",
		})
	}

	var filename string

	for _, file := range files {
		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Unsupported file type",
			})
		}

		filename = file.Filename
		savePath := filepath.Join("./uploads", filename)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Could not save file: %v", err),
			})
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:9000/api/upload/image/" + filename,
	})
}
