package products

import (
	"time"

	b "github.com/brbnk/core/api/models/base"
	"github.com/brbnk/core/api/models/skus"
	"github.com/brbnk/core/pkg/log"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	b.Base
	Code string `json:"code" db:"Code"`
	Name string `json:"name" db:"Name"`
}

type ProductSkuViewModel struct {
	Product
	Skus []*skus.Sku `json:"skus"`
}

type ProductModel struct {
	DB *sqlx.DB
}

type ProductRepositoryInterface interface {
	GetProducts() ([]*Product, error)
	GetProductById(uint) (*Product, error)
	CreateProduct(*Product) error
	UpdateProduct(*Product) error
	DeleteProduct(uint, time.Time) error
	GetSkusByProductId(uint) (*ProductSkuViewModel, error)
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

func (m ProductModel) GetProductById(id uint) (*Product, error) {
	stmt := GET

	p := &Product{}

	err := m.DB.
		QueryRowx(stmt, id).
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

	_, err = stmt.Exec(p.Code, p.Name, p.Base.IsActive, p.Base.IsDeleted, p.Base.CreateDate, p.Base.LastUpdate)
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

	_, err = stmt.Exec(p.Code, p.Name, p.Base.IsActive, p.Base.IsDeleted, p.Base.LastUpdate, p.Base.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m ProductModel) DeleteProduct(id uint, t time.Time) error {
	stmt, err := m.DB.Prepare(DELETE)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(t, id)
	if err != nil {
		return err
	}

	return nil
}

func (m ProductModel) GetSkusByProductId(productid uint) (*ProductSkuViewModel, error) {
	vm := &ProductSkuViewModel{
		Skus: make([]*skus.Sku, 0),
	}

	err := m.DB.Get(vm, GET, productid)
	if err != nil {
		return nil, log.LogMethodError("GetSkusByProductId (Product GET)", err)
	}

	err = m.DB.Select(&vm.Skus, GET_SKUS_BY_PRODUCTID, productid)
	if err != nil {
		return nil, log.LogMethodError("GetSkusByProductId (Skus SELECT)", err)
	}

	return vm, nil
}
