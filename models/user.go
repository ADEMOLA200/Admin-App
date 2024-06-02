package models

type User struct {
	Id		uint
	FirstName 	string 	`json:"name"`
	LastName	string	
	Email		string	`gorm:"unique" json:"email"`
	Password	string
	
}
