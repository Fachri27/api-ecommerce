package migrations

import (
	"github.com/Fachri27/api-ecommerce/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Categories{},
		&models.Product{},
		&models.User{},
		&models.Order{},
		&models.Cart{},
		&models.OrderItem{},
	) 
}
