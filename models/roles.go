package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"unique;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permission struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`
}

func MigrateRolesAndPermissions(db *gorm.DB) {
	db.AutoMigrate(&Role{}, &Permission{})
}
