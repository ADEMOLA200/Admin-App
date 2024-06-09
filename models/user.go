package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id			uint	`json:"id"`
	UUID        string  `gorm:"null"`
	FirstName 	string 	`json:"first_name" binding:"required"`
	LastName	string	`json:"last_name" binding:"required"`
	Email		string	`gorm:"unique" json:"email" binding:"required"`
	Password	[]byte	`json:"-"`
	RoleId		uint	`json:"role_id"`
	Role 		Role	`json:"role" gorm:"foreignKey:RoleId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user *User) BeforeCreate(u *gorm.DB) (err error) {
	user.UUID = uuid.NewString()
	return
}

func (user *User) SetPassword(password string) {
	hashedPassword , _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
