package models

import (
	"time"

	"github.com/brbnk/core/cfg/application"
)

type Products struct {
	Id         uint      `json:"id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	IsActive   bool      `json:"isactive"`
	IsDeleted  bool      `json:"isdeleted"`
	CreateDate time.Time `json:"createdate"`
	LastUpdate time.Time `json:"lastupdate"`
}

func (p *Products) GetProducts(app *application.Application) ([]*Products, error) {
	products := make([]*Products, 0)

	stmt := `
		SELECT
			p.Id, p.Code, p.Name, p.IsActive,
			p.IsDeleted, p.CreateDate, p.LastUpdate
		FROM Products p
		ORDER BY p.Code;
	`

	rows, err := app.DB.Client.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		product := new(Products)

		if err := rows.Scan(
			&product.Id,
			&product.Code,
			&product.Name,
			&product.IsActive,
			&product.IsDeleted,
			&product.CreateDate,
			&product.LastUpdate,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *Products) GetProductById(app *application.Application) (*Products, error) {
	stmt := `
		SELECT
			p.Id, p.Code, p.IsActive, p.IsDeleted, p.Name,
			p.CreateDate, p.LastUpdate
		FROM Products p
		WHERE Id = ?
	`

	err := app.DB.Client.
		QueryRow(stmt, p.Id).
		Scan(&p.Id, &p.Code, &p.IsActive, &p.IsDeleted, &p.Name, &p.CreateDate, &p.LastUpdate)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Products) CreateProduct(app *application.Application) error {
	stmt, err := app.DB.Client.Prepare(`
		INSERT INTO Products (code, name, isactive, isdeleted, createdate, lastupdate)
		VALUES (?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.Code, p.Name, p.IsActive, p.IsDeleted, p.CreateDate, p.LastUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (p *Products) UpdateProduct(app *application.Application) error {
	stmt, err := app.DB.Client.Prepare(`
    UPDATE Products
		SET
		  Code = ?,
			Name = ?,
			IsActive = ?,
			IsDeleted = ?,
			LastUpdate = ?
		WHERE Id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.Code, p.Name, p.IsActive, p.IsDeleted, p.LastUpdate, p.Id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Products) DeleteProduct(app *application.Application) error {
	stmt, err := app.DB.Client.Prepare(`
		UPDATE Products
		SET
			IsDeleted = 1,
			IsActive = 0,
			LastUpdate = ?
		WHERE Id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.LastUpdate, p.Id)
	if err != nil {
		return err
	}

	return nil
}
