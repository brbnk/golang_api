package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetHeader(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r, p)
	}
}
