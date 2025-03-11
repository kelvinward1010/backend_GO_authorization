package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
