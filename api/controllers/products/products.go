package products

import (
	"database/sql"
	"errors"
	"net/http"

	middleware "github.com/brbnk/core/api/middlewares"
	"github.com/brbnk/core/api/models"
	"github.com/brbnk/core/api/services/products"
	"github.com/brbnk/core/cfg/application"
	"github.com/brbnk/core/pkg/http/parser"
	httpresponse "github.com/brbnk/core/pkg/http/response"
	"github.com/julienschmidt/httprouter"
)

func GetAll(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()
		products, err := products.GetAllProducts(app)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.SetStatus(http.StatusOK).SetMessage("Empty List")
				response.Write(w, r)
				return
			}

			response.SetStatus(http.StatusPreconditionFailed).SetMessage("Internal Server Error").SetSuccess(false)
			response.Write(w, r)
			return
		}

		response.SetStatus(http.StatusOK).SetResult(products)
		response.Write(w, r)
	})
}

func Get(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("Invalid Paramater 'id'").SetSuccess(false)
			response.Write(w, r)
			return
		}

		model := &models.Products{Id: uint(id)}

		product, err := products.GetProductById(app, model)
		if err != nil {
			response.SetStatus(http.StatusNotFound).SetMessage("This product doesn't exist!").SetSuccess(false)
			response.Write(w, r)
			return
		}

		response.SetStatus(http.StatusOK).SetResult(product)
		response.Write(w, r)
	})
}

func Create(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		product := models.Products{}
		response := httpresponse.New()

		if err := parser.ParseBody(r.Body, &product); err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("Invalid payload!").SetSuccess(false)
			response.Write(w, r)
			return
		}

		if err := products.InsertProducts(app, &product); err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("The given product already exists!").SetSuccess(false)
			response.Write(w, r)
			return
		}

		response.
			SetStatus(http.StatusOK).
			SetMessage("Product created with success!").
			SetResult(product).
			Write(w, r)
	})
}

func Update(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("Invalid Paramater 'id'").SetSuccess(false)
			response.Write(w, r)
			return
		}

		product := models.Products{Id: uint(id)}

		if err := parser.ParseBody(r.Body, &product); err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("Invalid payload!").SetSuccess(false)
			response.Write(w, r)
			return
		}

		if err := products.UpdateProduct(app, &product); err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("It was not possible to update product").SetSuccess(false)
			response.Write(w, r)
			return
		}

		response.
			SetStatus(http.StatusOK).
			SetMessage("Product updated with success").
			SetResult(product).
			Write(w, r)
	}
}

func Delete(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("Invalid Paramater 'id'").SetSuccess(false)
			response.Write(w, r)
			return
		}

		product := &models.Products{Id: uint(id)}
		if err := products.DeleteProduct(app, product); err != nil {
			response.SetStatus(http.StatusBadRequest).SetMessage("It was not possible to delete product").SetSuccess(false)
			response.Write(w, r)
			return
		}

		response.
			SetStatus(http.StatusOK).
			SetMessage("Product deleted with success").
			SetResult(product).
			Write(w, r)
	}
}
