package router

import (
	"github.com/brbnk/core/api/controllers/products"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/products", products.GetAll(app))
	mux.POST("/products", products.Create(app))
	mux.GET("/products/:id", products.Get(app))
	mux.PUT("/products/:id", products.Update(app))
	mux.DELETE("/products/:id", products.Delete(app))

	return mux
}
