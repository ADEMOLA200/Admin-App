package models

type User struct {
	Id		uint
	FirstName 	string 	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Email		string	`gorm:"unique" json:"email"`
	Password	[]byte	`json:"-"`
}
