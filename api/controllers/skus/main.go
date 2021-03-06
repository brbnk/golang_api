package skus

import (
	"database/sql"
	"errors"
	"net/http"

	middleware "github.com/brbnk/core/api/middlewares"
	b "github.com/brbnk/core/api/models/base"
	"github.com/brbnk/core/api/models/skus"
	s "github.com/brbnk/core/api/services/skus"
	"github.com/brbnk/core/cfg/application"
	"github.com/brbnk/core/pkg/http/parser"
	httpresponse "github.com/brbnk/core/pkg/http/response"
	"github.com/julienschmidt/httprouter"
)

type SkuController struct {
	service s.SkuServiceInterface
}

func newController(ctx *application.DbContext) *SkuController {
	return &SkuController{
		service: s.NewService(ctx.Sku),
	}
}

func InitController(app *application.Application, h *httprouter.Router) {
	controller := newController(app.Ctx)

	h.GET("/skus", controller.GetAll())
	h.POST("/skus", controller.Create())
	h.GET("/skus/:id", controller.Get())
	h.PUT("/skus/:id", controller.Update())
	h.DELETE("/skus/:id", controller.Delete())
}

func (c *SkuController) GetAll() httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()
		products, err := c.service.GetAllSkus()

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

func (c *SkuController) Get() httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid Paramater 'id'").BadRequest(w, r)
			return
		}

		sku, err := c.service.GetSkuById(uint(id))
		if err != nil {
			response.SetMessage(err.Error()).NotFound(w, r)
			return
		}

		response.SetResult(sku).Ok(w, r)
	})
}

func (c *SkuController) Create() httprouter.Handle {
	return middleware.Apply(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		sku := skus.Sku{}
		response := httpresponse.New()

		if err := parser.ParseBody(r.Body, &sku); err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		if err := c.service.InsertSku(&sku); err != nil {
			response.SetMessage(err.Error()).BadRequest(w, r)
			return
		}

		response.SetResult(sku).SetMessage("Sku created with success").Ok(w, r)
	})
}

func (c *SkuController) Update() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid paramater 'id'").BadRequest(w, r)
			return
		}

		sku := skus.Sku{Base: b.Base{Id: uint(id)}}

		if err := parser.ParseBody(r.Body, &sku); err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		if err := c.service.UpdateSku(&sku); err != nil {
			response.SetMessage(err.Error()).BadRequest(w, r)
			return
		}

		response.SetResult(sku).SetMessage("Sku updated with success").Ok(w, r)
	}
}

func (c *SkuController) Delete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		response := httpresponse.New()

		id, err := parser.ParseId(p.ByName("id"))
		if err != nil {
			response.SetMessage("Invalid payload!").BadRequest(w, r)
			return
		}

		sku := &skus.Sku{Base: b.Base{Id: uint(id)}}
		if err := c.service.DeleteSku(uint(id)); err != nil {
			response.SetMessage(err.Error()).BadRequest(w, r)
			return
		}

		response.SetResult(sku).SetMessage("Sku deleted with success").Ok(w, r)
	}
}
