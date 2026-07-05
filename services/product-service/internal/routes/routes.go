package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pradeepgitbit/nimbus-commerce/services/product-service/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {

	handler := handlers.NewProductHandler()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "product-service",
			"status": "UP",
		})
	})

	router.POST("/products", handler.CreateProduct)
	router.GET("/products", handler.GetProducts)
	router.GET("/products/:id", handler.GetProduct)
	router.PUT("/products/:id", handler.UpdateProduct)
	router.DELETE("/products/:id", handler.DeleteProduct)
}