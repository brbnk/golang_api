package services

import (
	"time"

	"github.com/brbnk/core/api/models"
)

type ProductService struct {
	repository models.ProductRepositoryInterface
}

type ProductServiceInterface interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductById(p *models.Product) (*models.Product, error)
	InsertProducts(p *models.Product) error
	UpdateProduct(p *models.Product) error
	DeleteProduct(p *models.Product) error
}

func NewService(ctx models.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		repository: ctx,
	}
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	products, err := s.repository.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) GetProductById(p *models.Product) (*models.Product, error) {
	product, err := s.repository.GetProductById(p)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) InsertProducts(p *models.Product) error {
	p.CreateDate = time.Now()
	p.LastUpdate = time.Now()

	if err := s.repository.CreateProduct(p); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) UpdateProduct(p *models.Product) error {
	p.LastUpdate = time.Now()

	if err := s.repository.UpdateProduct(p); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(p *models.Product) error {
	p.LastUpdate = time.Now()

	if err := s.repository.DeleteProduct(p); err != nil {
		return err
	}

	return nil
}
