package router

import (
	"github.com/brbnk/core/api/handlers/helloworld"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/", helloworld.Do(app))

	return mux
}
