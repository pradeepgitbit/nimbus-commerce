package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/models"
	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/service"
)

type InventoryHandler struct {
	service *service.InventoryService
}

func NewInventoryHandler() *InventoryHandler {
	return &InventoryHandler{
		service: service.NewInventoryService(),
	}
}

func (h *InventoryHandler) CreateInventory(c *gin.Context) {

	var inventory models.Inventory

	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&inventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, inventory)
}

func (h *InventoryHandler) GetInventory(c *gin.Context) {

	inventory, err := h.service.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, inventory)
}

func (h *InventoryHandler) GetInventoryByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	inventory, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Inventory not found"})
		return
	}

	c.JSON(200, inventory)
}

func (h *InventoryHandler) UpdateInventory(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	inventory, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Inventory not found"})
		return
	}

	if err := c.ShouldBindJSON(inventory); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(inventory); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, inventory)
}

func (h *InventoryHandler) DeleteInventory(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}