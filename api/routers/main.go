package routers

import (
	"net/http"

	"github.com/brbnk/core/api/controllers/products"
	"github.com/brbnk/core/api/controllers/skus"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()

		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		header.Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	products.InitController(app, mux)
	skus.InitController(app, mux)

	return mux
}
