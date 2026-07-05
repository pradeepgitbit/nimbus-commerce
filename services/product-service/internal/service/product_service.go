package service

import (
	"github.com/google/uuid"
	"github.com/pradeepgitbit/nimbus-commerce/services/product-service/internal/models"
	"github.com/pradeepgitbit/nimbus-commerce/services/product-service/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		repo: repository.NewProductRepository(),
	}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetByID(id uuid.UUID) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}