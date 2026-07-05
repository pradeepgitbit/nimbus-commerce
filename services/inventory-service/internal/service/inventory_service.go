package service

import (
	"github.com/google/uuid"

	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/models"
	"github.com/pradeepgitbit/nimbus-commerce/services/inventory-service/internal/repository"
)

type InventoryService struct {
	repo *repository.InventoryRepository
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		repo: repository.NewInventoryRepository(),
	}
}

func (s *InventoryService) Create(inventory *models.Inventory) error {
	return s.repo.Create(inventory)
}

func (s *InventoryService) GetAll() ([]models.Inventory, error) {
	return s.repo.GetAll()
}

func (s *InventoryService) GetByID(id uuid.UUID) (*models.Inventory, error) {
	return s.repo.GetByID(id)
}

func (s *InventoryService) Update(inventory *models.Inventory) error {
	return s.repo.Update(inventory)
}

func (s *InventoryService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}