package routers

import (
	"github.com/brbnk/core/api/controllers/products"
	"github.com/brbnk/core/api/controllers/skus"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	products.InitController(app, mux)
	skus.InitController(app, mux)

	return mux
}
