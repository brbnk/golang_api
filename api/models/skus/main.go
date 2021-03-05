package skus

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Sku struct {
	Id         uint      `json:"id" db:"Id"`
	Code       string    `json:"code" db:"Code"`
	Name       string    `json:"name" db:"Name"`
	ProductId  uint      `json:"productid" db:"ProductId"`
	IsActive   bool      `json:"isactive" db:"IsActive"`
	IsDeleted  bool      `json:"isdeleted" db:"IsDeleted"`
	CreateDate time.Time `json:"createdate" db:"CreateDate"`
	LastUpdate time.Time `json:"lastupdate" db:"LastUpdate"`
}

type SkuModel struct {
	DB *sqlx.DB
}

type SkuRepositoryInterface interface {
	GetSkus() ([]*Sku, error)
	GetSkuById(p *Sku) (*Sku, error)
	CreateSku(p *Sku) error
	UpdateSku(p *Sku) error
	DeleteSku(p *Sku) error
}

func NewSkuRepository(db *sqlx.DB) SkuRepositoryInterface {
	return &SkuModel{DB: db}
}

func (m SkuModel) GetSkus() ([]*Sku, error) {
	skus := make([]*Sku, 0)

	stmt := GETALL

	rows, err := m.DB.Queryx(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		sku := new(Sku)

		if err := rows.StructScan(&sku); err != nil {
			return nil, err
		}

		skus = append(skus, sku)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return skus, nil
}

func (m SkuModel) GetSkuById(s *Sku) (*Sku, error) {
	stmt := GET

	err := m.DB.
		QueryRowx(stmt, s.Id).
		StructScan(s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (m SkuModel) CreateSku(s *Sku) error {
	stmt, err := m.DB.Prepare(CREATE)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(s.Code, s.Name, s.ProductId, s.IsActive, s.IsDeleted, s.CreateDate, s.LastUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (m SkuModel) UpdateSku(s *Sku) error {
	stmt, err := m.DB.Prepare(UPDATE)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(s.Code, s.Name, s.IsActive, s.IsDeleted, s.LastUpdate, s.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m SkuModel) DeleteSku(s *Sku) error {
	stmt, err := m.DB.Prepare(DELETE)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(s.LastUpdate, s.Id)
	if err != nil {
		return err
	}

	return nil
}
