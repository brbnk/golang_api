package application

import (
	"github.com/brbnk/core/api/models/products"
	"github.com/brbnk/core/api/models/skus"
	"github.com/brbnk/core/cfg/db"
	"github.com/brbnk/core/cfg/environment"
)

type DbContext struct {
	Product products.ProductModel
	Sku     skus.SkuModel
}

type Application struct {
	Cfg *environment.Configurations
	Ctx *DbContext
}

func Get() (*Application, error) {
	cfg := environment.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	ctx := &DbContext{
		Product: products.ProductModel{DB: db.Client},
		Sku:     skus.SkuModel{DB: db.Client},
	}

	if err != nil {
		return nil, err
	}

	return &Application{
		Cfg: cfg,
		Ctx: ctx,
	}, nil
}
