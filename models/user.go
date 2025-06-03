package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string       `json:"username" gorm:"unique;not null"`
	Email       string       `json:"email" gorm:"unique;not null"`
	Password    string       `json:"password,omitempty" gorm:"not null"`
	Roles       []Role       `gorm:"many2many:user_roles;" json:"roles"`
	Permissions []Permission `gorm:"many2many:user_permissions;" json:"permissions"`
}
