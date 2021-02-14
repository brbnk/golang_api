package models

import (
	"log"

	"github.com/brbnk/core/cfg/application"
)

type Products struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (p *Products) GetProducts(app *application.Application) ([]*Products, error) {
	products := make([]*Products, 0)

	stmt := `
		SELECT Id, Code, Name
		FROM Products
	`

	rows, err := app.DB.Client.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		product := new(Products)

		if err := rows.Scan(&product.Id, &product.Code, &product.Name); err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return products, nil
}
