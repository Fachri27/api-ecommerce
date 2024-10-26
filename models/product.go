package models

import (
	"gorm.io/gorm"
)

// create model for product table
type Product struct {
	gorm.Model
	Id           uint    `gorm:"primary_key" json:"id"`
	Product_name string  `json:"product_name"`
	Price        float64 `json:"price"`
	Deskripsi    string  `json:"deskripsi"`
	Stock        int     `json:"stock"`
	Category_id  uint    `json:"category_id"`
}
