package helloworld

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	middleware "github.com/brbnk/core/api/middlewares"
	"github.com/brbnk/core/api/models"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func sayHelloWorld(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		model := &models.Products{}

		products, err := model.GetProducts(app)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "products do not exist")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		response, _ := json.Marshal(products)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	mw := []middleware.Middleware{
		middleware.SetHeader,
	}
	return middleware.Chain(sayHelloWorld(app), mw...)
}
