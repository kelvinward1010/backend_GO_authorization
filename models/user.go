package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" binding:"required"`
	Email    string `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password string `json:"password,omitempty" gorm:"not null" binding:"required,min=6"`
	Role     string `json:"role" gorm:"default:user"`
}
