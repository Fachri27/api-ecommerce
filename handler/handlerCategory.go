package handler

import (
	"net/http"

	"github.com/Fachri27/api-ecommerce/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		var category models.Categories

		if err := db.First(&category, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": "This id not found",
			})
			return
		}

		c.JSON(http.StatusOK, category)

	}
}

func ListCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var categories []models.Categories

		db.Find(&categories)
		c.JSON(200, categories)
	}
}

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var input models.Categories

		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": "Invalid input",
			})
			return
		}

		if input.Category_name == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": "required",
			})
			return
		}

		db.Create(&input)
		c.JSON(200, input)

	}

}

func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var category models.Categories

		if err := db.First(&category, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message" : "id not found",
			})
			return
		}

		var input models.Categories
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message" : "Update filed",
			})
			return
		}

		db.Model(&category).Updates(input)
		c.JSON(http.StatusOK, gin.H{
			"message" : "Update success",
		})
	}
}

func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		var category models.Categories

		if err := db.First(&category, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Category not found",
			})
			return
		}

		db.Delete(&category)
		c.JSON(http.StatusOK, gin.H{
			"message" : "Category has been successfully delete",
		})
	}
}
