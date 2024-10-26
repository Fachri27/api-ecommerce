package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Id         uint `gorm:"primary_key" json:"id"`
	User_id    uint `json:"user_id"`
	Product_id uint `json:"product_id"`
	Quantity   uint `json:"quantity"`
}
