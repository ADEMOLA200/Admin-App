package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 	string `json:"name"`
	Email	string	`gorm:"unique" json:"email"`
	Address
}

type Address struct {
	userId		uint
	streetName	string
}