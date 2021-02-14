package main

import (
	"log"

	router "github.com/brbnk/core/api/routers"
	"github.com/brbnk/core/cfg/config"
	"github.com/brbnk/core/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	// make sure env variables are loaded
	if err := godotenv.Load(); err != nil {
		log.Fatalln("It was not possible to load .env file!")
	}

	env := config.Get()

	srv := server.Get().
		WithAdrress(":" + env.Srv.APIPort).
		WithRouter(router.Get())

	log.Fatal(srv.Start())
}
