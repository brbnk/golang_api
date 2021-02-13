package main

import (
	"log"

	router "github.com/brbnk/core/api/routers"
	"github.com/brbnk/core/pkg/server"
)

func main() {
	s := server.Get().
		WithAdrress(":5000").
		WithRouter(router.Get())

	log.Fatal(s.Start())
}
