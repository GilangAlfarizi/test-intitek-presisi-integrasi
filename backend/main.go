package main

import (
	"fullstack-test/handler"
	"fullstack-test/product"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/backend-intitek?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&product.Product{})

	router := gin.Default()

	// Configuring CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Replace with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	productRepository := product.NewRepositoryProduct(db)
	productService := product.NewServiceProduct(productRepository)
	productHandler := handler.NewProductHandler(productService)

	product := router.Group("/api/product")
	product.GET("/", productHandler.ReadProducts)
	product.POST("/", productHandler.CreateProduct)
	product.PUT("/:id", productHandler.UpdateProductById)
	product.DELETE("/:id", productHandler.DeleteProductById)
	product.GET("/:id", productHandler.ReadOneProduct)

	router.Run(":8080")
}
