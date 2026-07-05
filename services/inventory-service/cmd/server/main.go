package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/config"
	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/database"
	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/routes"
)

func main() {

	config.Load()

	database.Connect()

	router := gin.Default()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	routes.RegisterRoutes(router)

	log.Println("🚀 Inventory Service running on :8082")

	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}