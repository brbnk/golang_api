package routers

import (
	"github.com/brbnk/core/api/controllers/products"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func InitProductRoutes(app *application.Application, h *httprouter.Router) {
	h.GET("/products", products.GetAll(app))
	h.POST("/products", products.Create(app))
	h.GET("/products/:id", products.Get(app))
	h.PUT("/products/:id", products.Update(app))
	h.DELETE("/products/:id", products.Delete(app))
}
