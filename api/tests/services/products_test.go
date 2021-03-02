package services

import (
	"testing"

	"github.com/brbnk/core/api/services"
)

var (
	mock    = &MockProductRepository{}
	service = services.NewProductService(mock)
)

func TestGetAllProducts(t *testing.T) {
	t.Error()
}

func TestGetProductById(t *testing.T) {
	t.Error()
}

func TestCreateProduct(t *testing.T) {
	t.Error()
}

func TestUpdateProduct(t *testing.T) {
	t.Error()
}

func TestDeleteProduct(t *testing.T) {
	t.Error()
}
