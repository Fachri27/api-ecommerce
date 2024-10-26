package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id           uint        `gorm:"primary_key" json:"id"`
	User_id      uint        `gorm:"foreign_key" json:"user_id"`
	Order_date   time.Time   `json:"order_date"`
	Order_status string      `json:"order_status"`
	Amount       float64     `json:"amount"`
	Price        float64     `json:"price"`
	Items        []OrderItem `gorm:"foreignkey:OrderId" json:"items"`
}
