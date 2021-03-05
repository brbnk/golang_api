package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Client *sqlx.DB
}

func Get(connStr string) (*DB, error) {
	db, err := connect(connStr)

	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

func connect(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", connStr)

	if err != nil {
		log.Fatal(err)
	}

	// Check if database is available and accessible
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
