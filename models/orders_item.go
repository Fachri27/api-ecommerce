package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Id         uint    `gorm:"primary_key" json:"id"`
	OrderId    uint    `json:"order_id"`
	Product_id uint    `json:"product_id"`
	Quantity   uint    `json:"quantity"`
	Price      float64 `json:"price"`
}