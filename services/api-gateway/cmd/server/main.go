package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func newReverseProxy(target string) *httputil.ReverseProxy {

	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalDirector := proxy.Director

	proxy.Director = func(req *http.Request) {
		originalDirector(req)
	}

	return proxy
}

func main() {

	// These will later come from Docker Compose / Kubernetes
	productServiceURL := os.Getenv("PRODUCT_SERVICE_URL")
	if productServiceURL == "" {
		productServiceURL = "http://localhost:8081"
	}

	inventoryServiceURL := os.Getenv("INVENTORY_SERVICE_URL")
	if inventoryServiceURL == "" {
		inventoryServiceURL = "http://localhost:8082"
	}

	productProxy := newReverseProxy(productServiceURL)
	inventoryProxy := newReverseProxy(inventoryServiceURL)

	router := gin.Default()

	// Health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "api-gateway",
			"status":  "UP",
		})
	})

	// Product Service
	router.Any("/api/products", func(c *gin.Context) {
		c.Request.URL.Path = "/products"
		productProxy.ServeHTTP(c.Writer, c.Request)
	})

	router.Any("/api/products/*path", func(c *gin.Context) {
		c.Request.URL.Path = "/products" + c.Param("path")
		productProxy.ServeHTTP(c.Writer, c.Request)
	})

	// Inventory Service
	router.Any("/api/inventory", func(c *gin.Context) {
		c.Request.URL.Path = "/inventory"
		inventoryProxy.ServeHTTP(c.Writer, c.Request)
	})

	router.Any("/api/inventory/*path", func(c *gin.Context) {
		c.Request.URL.Path = "/inventory" + c.Param("path")
		inventoryProxy.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("🚀 API Gateway running on :8083")

	if err := router.Run(":8083"); err != nil {
		log.Fatal(err)
	}
}