package models

import (
	"gorm.io/gorm"
)

// create model for product table
type Product struct {
	gorm.Model
	Id           uint    `gorm:"primary_key" json:"id"`
	Product_name string  `json:"product_name" binding:"required"`
	Price        float64 `json:"price"`
	Deskripsi    string  `json:"deskripsi" binding:"required"`
	Stock        int     `json:"stock"`
	Category_id  uint    `json:"category_id"`
	Image_id     *uint   `json:"image_id"`
}
