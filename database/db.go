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
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
        panic("One or more environment variables (DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT) are not set")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", dbUser, dbPassword, dbHost, dbPort, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("could not connect to the database: " + err.Error())
    }

    DB = db
    DB.AutoMigrate(
        &models.User{},
        &models.Role{},
        &models.Permissions{},
        &models.Product{},
        &models.Order{},
        &models.OrderItem{},
    )
}
