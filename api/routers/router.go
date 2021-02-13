package router

import (
	"github.com/brbnk/core/api/handlers/helloworld"
	"github.com/julienschmidt/httprouter"
)

func Get() *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/", helloworld.Do())

	return mux
}
