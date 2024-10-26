package models

import "gorm.io/gorm"

// create model for categories table
type Categories struct {
	gorm.Model
	Id            uint   `gorm:"primary_key" json:"id"`
	Category_name string `json:"category_name" binding:"required"`
}
