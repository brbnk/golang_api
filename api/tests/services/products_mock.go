package services

import (
	"github.com/brbnk/core/api/models"
)

type MockProductRepository struct{}

func (m *MockProductRepository) GetProducts() ([]*models.Product, error) {
	products := make([]*models.Product, 0)

	return products, nil
}

func (m *MockProductRepository) GetProductById(p *models.Product) (*models.Product, error) {
	return p, nil
}

func (m *MockProductRepository) CreateProduct(p *models.Product) error {
	return nil
}

func (m *MockProductRepository) UpdateProduct(p *models.Product) error {
	return nil
}

func (m *MockProductRepository) DeleteProduct(p *models.Product) error {
	return nil
}
