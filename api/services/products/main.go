package products

import (
	"time"

	"github.com/brbnk/core/api/models/products"
)

type ProductService struct {
	repository products.ProductRepositoryInterface
}

type ProductServiceInterface interface {
	GetAllProducts() ([]*products.Product, error)
	GetProductById(uint) (*products.Product, error)
	InsertProducts(p *products.Product) error
	UpdateProduct(p *products.Product) error
	DeleteProduct(uint) error
}

func NewService(ctx products.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		repository: ctx,
	}
}

func (s *ProductService) GetAllProducts() ([]*products.Product, error) {
	products, err := s.repository.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) GetProductById(id uint) (*products.Product, error) {
	product, err := s.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) InsertProducts(p *products.Product) error {
	p.CreateDate = time.Now()
	p.LastUpdate = time.Now()

	if err := s.repository.CreateProduct(p); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) UpdateProduct(p *products.Product) error {
	p.LastUpdate = time.Now()

	if err := s.repository.UpdateProduct(p); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	time := time.Now()

	if err := s.repository.DeleteProduct(id, time); err != nil {
		return err
	}

	return nil
}
