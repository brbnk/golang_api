package products

import (
	"database/sql"
	"errors"
	"net/http"

	middleware "github.com/brbnk/core/api/middlewares"
	"github.com/brbnk/core/api/models/base"
	"github.com/brbnk/core/api/models/products"
	s "github.com/brbnk/core/api/services/products"
	"github.com/brbnk/core/cfg/application"
	"github.com/brbnk/core/pkg/http/parser"
	httpresponse "github.com/brbnk/core/pkg/http/response"
	"github.com/julienschmidt/httprouter"
)

type ProductController struct {
	service s.ProductServiceInterface
}

func newController(ctx *application.DbContext) *ProductController {
	return &ProductController{
		service: s.NewService(ctx.Product),
	}
}

func InitController(app *application.Application, h *httprouter.Router) {
	controller := newController(app.Ctx)

	h.GET("/products", controller.GetAll())
	h.POST("/products", controller.Create())
	h.GET("/products/:id", controller.Get())
	h.PUT("/products/:id", controller.Update())
	h.DELETE("/products/:id", controller.Delete())
}

func (c *ProductController) GetAll() httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()
		products, err := c.service.GetAllProducts()

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.SetMessage("Empty List").Ok(w, r)
				return
			}

			response.SetMessage(err.Error()).InternalServerError(w, r)
			return
		}

		response.SetResult(products).Ok(w, r)
	})
}

func (c *ProductController) Get() httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid Paramater 'id'").BadRequest(w, r)
			return
		}

		product, err := c.service.GetProductById(uint(id))
		if err != nil {
			response.SetMessage(err.Error()).NotFound(w, r)
			return
		}

		response.SetResult(product).Ok(w, r)
	})
}

func (c *ProductController) Create() httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		product := products.Product{}
		response := httpresponse.New()

		if err := parser.ParseBody(r.Body, &product); err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		if err := c.service.InsertProducts(&product); err != nil {
			response.SetMessage(err.Error()).BadRequest(w, r)
			return
		}

		response.SetResult(product).SetMessage("Product created with success").Ok(w, r)
	})
}

func (c *ProductController) Update() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid paramater 'id'").BadRequest(w, r)
			return
		}

		product := products.Product{Base: base.Base{Id: uint(id)}}

		if err := parser.ParseBody(r.Body, &product); err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		if err := c.service.UpdateProduct(&product); err != nil {
			response.SetMessage(err.Error()).BadRequest(w, r)
			return
		}

		response.SetResult(product).SetMessage("Product updated with success").Ok(w, r)
	}
}

func (c *ProductController) Delete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		if err := c.service.DeleteProduct(uint(id)); err != nil {
			response.SetMessage(err.Error()).BadRequest(w, r)
			return
		}

		response.SetMessage("Product deleted with success").Ok(w, r)
	}
}
