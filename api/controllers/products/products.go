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
				response.SetMessage("Empty List").Ok(w, r)
				return
			}

			response.InternalServerError(w, r)
			return
		}

		response.SetResult(products).Ok(w, r)
	})
}

func Get(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid Paramater 'id'").BadRequest(w, r)
			return
		}

		model := &models.Products{Id: uint(id)}

		product, err := products.GetProductById(app, model)
		if err != nil {
			response.SetMessage("This product doesn't exist!").NotFound(w, r)
			return
		}

		response.SetResult(product).Ok(w, r)
	})
}

func Create(app *application.Application) httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		product := models.Products{}
		response := httpresponse.New()

		if err := parser.ParseBody(r.Body, &product); err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		if err := products.InsertProducts(app, &product); err != nil {
			response.SetMessage("The given product already exists!").BadRequest(w, r)
			return
		}

		response.SetResult(product).SetMessage("Product created with success").Ok(w, r)
	})
}

func Update(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid paramater 'id'").BadRequest(w, r)
			return
		}

		product := models.Products{Id: uint(id)}

		if err := parser.ParseBody(r.Body, &product); err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		if err := products.UpdateProduct(app, &product); err != nil {
			response.SetMessage("It was not possible to update product").BadRequest(w, r)
			return
		}

		response.SetResult(product).SetMessage("Product updated with success").Ok(w, r)
	}
}

func Delete(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		product := &models.Products{Id: uint(id)}
		if err := products.DeleteProduct(app, product); err != nil {
			response.SetMessage("It was not possible to delete product").BadRequest(w, r)
			return
		}

		response.SetResult(product).SetMessage("Product deleted with success").Ok(w, r)
	}
}
