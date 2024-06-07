package models

type Role struct {
	Id 			uint			`json:"id"`	
	Name		string			`json:"name"`
	Permissions	[]Permissions	`json:"permissions" gorm:"many2many:role_permissions"`
}