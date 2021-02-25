package products

import (
	"encoding/json"
	"log"

	"github.com/brbnk/core/api/models"
	"github.com/brbnk/core/cfg/application"
)

func GetAllProducts(app *application.Application) ([]byte, error) {
	model := &models.Products{}

	products, err := model.GetProducts(app)
	if err != nil {
		log.Fatal(err)
	}

	response, _ := json.Marshal(products)
	return response, nil
}

func GetProductById(app *application.Application, p *models.Products) ([]byte, error) {
	product, err := p.GetProductById(app)
	if err != nil {
		log.Fatal(err)
	}
	response, _ := json.Marshal(product)
	return response, nil
}

func InsertProducts(app *application.Application, p *models.Products) error {
	if err := p.CreateProduct(app); err != nil {
		log.Fatal(err)
	}
	return nil
}

func UpdateProduct(app *application.Application, p *models.Products) error {
	if err := p.UpdateProduct(app); err != nil {
		log.Fatal(err)
	}
	return nil
}

func DeleteProduct(app *application.Application, p *models.Products) error {
	if err := p.DeleteProduct(app); err != nil {
		log.Fatal(err)
	}
	return nil
}
