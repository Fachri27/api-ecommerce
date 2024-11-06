package handler

import (
	"net/http"

	"github.com/Fachri27/api-ecommerce/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product

		db.Find(&products)
		c.JSON(http.StatusOK, gin.H{
			"data": products,
		})
	}
}

func GetListProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product

		if err := db.First(&product, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "id not found",
			})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product

		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"messege": "invalid input",
			})
			return
		}

		if input.Product_name == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"messege": "required",
			})
			return
		}

		if input.Image_id == nil {
			db.Create(&input)
			c.JSON(http.StatusOK, gin.H{
				"Messege": "Upload without image success",
			})
			return
		}

		db.Create(&input)
		c.JSON(http.StatusOK, gin.H{
			"Messege": "Input success",
		})

	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product

		if err := db.First(&product, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "id not found",
			})
			return
		}

		var input models.Product
		if err := c.ShouldBindBodyWithJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "update failed",
			})
			return
		}

		db.Model(&product).Updates(input)
		c.JSON(200, gin.H{
			"message": "update success",
		})
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		id := c.Param("id")
		var product models.Product

		if err := db.First(&product, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "id not found",
			})
			return
		}

		db.Delete(&product)
		c.JSON(200, gin.H{
			"message" : "Delete success",
		})
	}
}
