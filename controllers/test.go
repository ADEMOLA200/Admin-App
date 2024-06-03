package controllers


// func User (ac *fiber.Ctx) error {
// 	cookie := ac.Cookies("jwt")

// 	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(t *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return ac.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "unauthorized user",
// 			"success": false,
// 		})
// 	}
	
// 	claims := token.Claims.(*Claims)

// 	var user models.User

// 	database.DB.Where("id = ?", claims.Issuer).First(&user)

// 	return ac.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "authenticated user",
// 		"user": user,
// 		"success": true,
// 	})
// }


// func Login(ac *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := ac.BodyParser(&data); err != nil {
// 		return err
// 	}

// 	var user models.User

// 	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {
// 		return ac.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "user not found",
// 			"success": false,
// 		})
// 	}

// 	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
// 		return ac.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "incorrect password",
// 			"success": false,
// 		})
// 	}

// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		Issuer: strconv.Itoa(int(user.Id)),
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 	})

// 	token, err := claims.SignedString([]byte("secret"))
// 	if err != nil {
// 		return ac.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	cookie := fiber.Cookie {
// 		Name: "jwt",
// 		Value: token,
// 		Expires: time.Now().Add(time.Hour * 24),
// 		HTTPOnly: true,
// 	}

// 	ac.Cookie(&cookie)

// 	return ac.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "user successfully logged in",
// 		"user": user,
// 		"success": true,
// 	})
// }