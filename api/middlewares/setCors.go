package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetCors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		header := w.Header()

		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		next(w, r, p)
	}
}
