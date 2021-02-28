package products

import (
	"time"

	"github.com/brbnk/core/api/models"
	"github.com/brbnk/core/cfg/application"
)

func GetAllProducts(app *application.Application) ([]*models.Products, error) {
	model := &models.Products{}

	products, err := model.GetProducts(app)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(app *application.Application, p *models.Products) (*models.Products, error) {
	product, err := p.GetProductById(app)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func InsertProducts(app *application.Application, p *models.Products) error {
	p.CreateDate = time.Now()
	p.LastUpdate = time.Now()

	if err := p.CreateProduct(app); err != nil {
		return err
	}

	return nil
}

func UpdateProduct(app *application.Application, p *models.Products) error {
	p.LastUpdate = time.Now()

	if err := p.UpdateProduct(app); err != nil {
		return err
	}

	return nil
}

func DeleteProduct(app *application.Application, p *models.Products) error {
	p.LastUpdate = time.Now()

	if err := p.DeleteProduct(app); err != nil {
		return err
	}

	return nil
}
