package repository

import (
	"github.com/google/uuid"
	"github.com/pradeepgitbit/nimbus-commerce/services/product-service/internal/database"
	"github.com/pradeepgitbit/nimbus-commerce/services/product-service/internal/models"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return database.DB.Create(product).Error
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := database.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return database.DB.Save(product).Error
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
	return database.DB.Delete(&models.Product{}, "id = ?", id).Error
}