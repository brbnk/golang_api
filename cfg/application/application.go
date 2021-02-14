package application

import (
	"github.com/brbnk/core/cfg/db"
	"github.com/brbnk/core/cfg/environment"
)

type Application struct {
	Cfg *environment.Configurations
	DB  *db.DB
}

func Get() (*Application, error) {
	cfg := environment.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application{
		Cfg: cfg,
		DB:  db,
	}, nil
}
