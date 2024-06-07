package controllers

// func CreateRole(rc *fiber.Ctx) error {
// 	var roleDTO fiber.Map

// 	if err := rc.BodyParser(&roleDTO); err != nil {
// 		return err
// 	}

// 	list := roleDTO["permissions"].([]interface{})

// 	permissions := make([]models.Permissions, len(list))

// 	for i, permissionsId := range list {
// 		id, _ := strconv.Atoi(permissionsId.(string))

// 		permissions[i] = models.Permissions{
// 			Id: uint(id),
// 		}
// 	}

// 	role := models.Role {
// 		Name: roleDTO["name"].(string),
// 		Permissions: permissions,
// 	}

// 	if err := database.DB.Create(&role).Error; err != nil {
// 		return rc.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "could not create role, email already exist",
// 			"success": false,
// 		})
// 	}

// 	return rc.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "role created successfully",
// 		"role": role,
// 		"success": true,
// 	})
// }