package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 	string 	`json:"name"`
	Email	string	`gorm:"unique" json:"email"`
	Address	Address	`gorm:"foreignKey:UserId"`
}

type Address struct {
	UserId		uint
	StreetName	string
}