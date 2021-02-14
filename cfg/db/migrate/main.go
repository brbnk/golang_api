package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/brbnk/core/cfg/environment"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := environment.Get()

	direction := cfg.GetMigration()

	db, err := sql.Open("mysql", cfg.GetDBConnStr())
	if err != nil {
		log.Fatalf("could not connect to the MySQL database... %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if direction != "down" && direction != "up" {
		log.Println("-migrate accepts [up, down] values only")
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://cfg/db/migrations", "mysql", driver)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	if direction == "up" {
		log.Printf("Migrating...")
		if err := m.Up(); err != nil {
			log.Printf("failed migrate up: %s", err)
			return
		}
	}

	if direction == "down" {
		if err := m.Down(); err != nil {
			log.Printf("failed migrate down: %s", err)
			return
		}
	}

	log.Printf("Finished!")

	os.Exit(0)
}
