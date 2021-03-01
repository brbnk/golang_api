package application

import (
	"github.com/brbnk/core/api/models"
	"github.com/brbnk/core/cfg/db"
	"github.com/brbnk/core/cfg/environment"
)

type DbContext struct {
	Product models.ProductModel
}

type Application struct {
	Cfg *environment.Configurations
	Ctx *DbContext
}

func Get() (*Application, error) {
	cfg := environment.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	ctx := &DbContext{
		Product: models.ProductModel{DB: db.Client},
	}

	if err != nil {
		return nil, err
	}

	return &Application{
		Cfg: cfg,
		Ctx: ctx,
	}, nil
}
