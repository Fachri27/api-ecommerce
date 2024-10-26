package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id           uint   `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Alamat       string `json:"alamat"`
	Number_phone uint   `json:"phone"`
}
