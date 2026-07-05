package repository

import (
	"github.com/google/uuid"

	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/database"
	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/models"
)

type InventoryRepository struct{}

func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{}
}

func (r *InventoryRepository) Create(inventory *models.Inventory) error {
	return database.DB.Create(inventory).Error
}

func (r *InventoryRepository) GetAll() ([]models.Inventory, error) {
	var inventory []models.Inventory
	err := database.DB.Find(&inventory).Error
	return inventory, err
}

func (r *InventoryRepository) GetByID(id uuid.UUID) (*models.Inventory, error) {
	var inventory models.Inventory
	err := database.DB.First(&inventory, "id = ?", id).Error
	return &inventory, err
}

func (r *InventoryRepository) Update(inventory *models.Inventory) error {
	return database.DB.Save(inventory).Error
}

func (r *InventoryRepository) Delete(id uuid.UUID) error {
	return database.DB.Delete(&models.Inventory{}, "id = ?", id).Error
}