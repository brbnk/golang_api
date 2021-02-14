package main

import (
	"log"

	router "github.com/brbnk/core/api/routers"
	"github.com/brbnk/core/cfg/application"
	"github.com/brbnk/core/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("It was not possible to load .env file!")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err)
	}

	srv := server.Get().
		WithAdrress(":" + app.Cfg.APIPort).
		WithRouter(router.Get(app))

	log.Fatal(srv.Start())
}
