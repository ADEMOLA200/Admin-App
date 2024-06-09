package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id			uint		`json:"id"`
	UUID        string  	`gorm:"null"`
	FirstName 	string 		`json:"first_name" binding:"required"`
	LastName	string		`json:"last_name" binding:"required"`
	Email		string		`gorm:"unique" json:"email" binding:"required"`
	OrderItems	[]OrderItem	`json:"order_items" gorm:"foreignKey:OrderId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderItem struct {
	Id					uint	`json:"id"`
	OrderId				uint	`json:"order_id"`
	ProductTitle		string	`json:"product_title"`
	Price				float32	`json:"price"`
	Quantity			uint	`json:"quantity"`
}

func (order *Order) BeforeCreate(o *gorm.DB) (err error) {
	order.UUID = uuid.NewString()
	return
}
