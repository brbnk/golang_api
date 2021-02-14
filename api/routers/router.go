package router

import (
	"github.com/brbnk/core/api/controllers/products"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/products", products.GetAll(app))

	return mux
}
