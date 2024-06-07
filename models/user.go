package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id			uint	`json:"id"`
	FirstName 	string 	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Email		string	`gorm:"unique" json:"email"`
	Password	[]byte	`json:"-"`
	RoleId		uint	`json:"role_id"`
	Role 		Role	`json:"role" gorm:"foreignKey:RoleId"`
}

func (user *User) SetPassword(password string) {
	hashedPassword , _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
