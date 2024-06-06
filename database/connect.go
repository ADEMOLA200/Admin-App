package database

import (
    "fmt"
    "os"

    "github.com/ADEMOLA200/Admin-App.git/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")

    if dbUser == "" || dbPassword == "" {
        panic("DB_USER or DB_PASSWORD environment variables are not set")
    }

    dsn := fmt.Sprintf("%s:%s@/admin_app", dbUser, dbPassword)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("could not connect to the database: " + err.Error())
    }

    DB = db
    DB.AutoMigrate(&models.User{})
}
