package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {

	handler := handlers.NewInventoryHandler()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "inventory-service",
			"status":  "UP",
		})
	})

	router.POST("/inventory", handler.CreateInventory)
	router.GET("/inventory", handler.GetInventory)
	router.GET("/inventory/:id", handler.GetInventoryByID)
	router.PUT("/inventory/:id", handler.UpdateInventory)
	router.DELETE("/inventory/:id", handler.DeleteInventory)
}