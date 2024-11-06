package main

import (
	"github.com/Fachri27/api-ecommerce/configs"
	"github.com/Fachri27/api-ecommerce/handler"
	"github.com/Fachri27/api-ecommerce/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	//db
	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	// migrations
	migrations.Migrate(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Hello World",
		})
	})

	// route category
	r.GET("/categories", handler.ListCategory(db))
	r.GET("/categories/:id", handler.GetListCategory(db))
	r.DELETE("/categories/:id", handler.DeleteCategory(db))
	r.POST("/create-category", handler.CreateCategory(db))
	r.PUT("/update-category/:id", handler.UpdateCategory(db))

	// Route Product
	r.GET("/products", handler.ListProduct(db))
	r.GET("/products/:id", handler.GetListProduct(db))
	r.POST("/create-product", handler.CreateProduct(db))
	r.DELETE("/products/:id", handler.DeleteProduct(db))
	r.PUT("/update-product/:id", handler.UpdateProduct(db))

	r.Run(":3000")
}
