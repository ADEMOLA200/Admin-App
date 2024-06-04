package database

import (
	"github.com/ADEMOLA200/Admin-App.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/admin_app"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = db

	DB.AutoMigrate(&models.User{})
}
