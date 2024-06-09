package models

import (
	"time"
)

type Product struct {
    Id          uint       `json:"id"`
    UUID        string     `gorm:"null"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Image       string     `json:"image"`
    Price       float64    `json:"price"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
