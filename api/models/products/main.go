package products

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id         uint      `json:"id" db:"Id"`
	Code       string    `json:"code" db:"Code"`
	Name       string    `json:"name" db:"Name"`
	IsActive   bool      `json:"isactive" db:"IsActive"`
	IsDeleted  bool      `json:"isdeleted" db:"IsDeleted"`
	CreateDate time.Time `json:"createdate" db:"CreateDate"`
	LastUpdate time.Time `json:"lastupdate" db:"LastUpdate"`
}

type ProductModel struct {
	DB *sqlx.DB
}

type ProductRepositoryInterface interface {
	GetProducts() ([]*Product, error)
	GetProductById(p *Product) (*Product, error)
	CreateProduct(p *Product) error
	UpdateProduct(p *Product) error
	DeleteProduct(p *Product) error
}

func NewProductRepository(db *sqlx.DB) ProductRepositoryInterface {
	return &ProductModel{DB: db}
}

func (m ProductModel) GetProducts() ([]*Product, error) {
	products := make([]*Product, 0)

	stmt := GETALL

	rows, err := m.DB.Queryx(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		product := new(Product)

		if err := rows.StructScan(&product); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (m ProductModel) GetProductById(p *Product) (*Product, error) {
	stmt := GET

	err := m.DB.
		QueryRowx(stmt, p.Id).
		StructScan(p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (m ProductModel) CreateProduct(p *Product) error {
	stmt, err := m.DB.Prepare(CREATE)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.Code, p.Name, p.IsActive, p.IsDeleted, p.CreateDate, p.LastUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (m ProductModel) UpdateProduct(p *Product) error {
	stmt, err := m.DB.Prepare(UPDATE)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.Code, p.Name, p.IsActive, p.IsDeleted, p.LastUpdate, p.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m ProductModel) DeleteProduct(p *Product) error {
	stmt, err := m.DB.Prepare(DELETE)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.LastUpdate, p.Id)
	if err != nil {
		return err
	}

	return nil
}
