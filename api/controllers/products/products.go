package products

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	middleware "github.com/brbnk/core/api/middlewares"
	"github.com/brbnk/core/api/models"
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

		w.WriteHeader(200)
		w.Write(products)
	})
}

func Get(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id, err := strconv.ParseUint(p.ByName("id"), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		model := &models.Products{Id: uint(id)}

		product, err := products.GetProductById(app, model)

		w.WriteHeader(200)
		w.Write(product)
	})
}

func Create(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		product := models.Products{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&product); err != nil {
			response, _ := json.Marshal("Invalid request payload")
			w.Write(response)
			return
		}

		product.CreateDate = time.Now()
		product.LastUpdate = time.Now()

		if err := products.InsertProducts(app, &product); err != nil {
			response, _ := json.Marshal("The given product already exists!")
			w.Write(response)
			return
		}

		w.WriteHeader(200)
		response, _ := json.Marshal(product)
		w.Write(response)
	})
}

func Update(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id, err := strconv.ParseUint(p.ByName("id"), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		product := models.Products{Id: uint(id)}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&product); err != nil {
			response, _ := json.Marshal("Invalid request payload")
			w.Write(response)
			return
		}

		product.LastUpdate = time.Now()

		if err := products.UpdateProduct(app, &product); err != nil {
			response, _ := json.Marshal("Error to update Product!")
			w.Write(response)
			return
		}

		w.WriteHeader(200)
		response, _ := json.Marshal(product)
		w.Write(response)
	}
}

func Delete(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id, err := strconv.ParseUint(p.ByName("id"), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		model := &models.Products{Id: uint(id)}
		if err := products.DeleteProduct(app, model); err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(200)
	}
}
