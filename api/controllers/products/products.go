package products

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	middleware "github.com/brbnk/core/api/middlewares"
	"github.com/brbnk/core/api/services/products"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func GetAll(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		products, err := products.GetAllProducts(app)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "Empty List")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		w.Write(products)
	})
}
