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
