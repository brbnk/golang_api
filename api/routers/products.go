package routers

import (
	"github.com/brbnk/core/api/controllers/products"
	"github.com/brbnk/core/cfg/application"
	"github.com/julienschmidt/httprouter"
)

func InitProductRoutes(app *application.Application, h *httprouter.Router) {
	controller := products.NewController(app.Ctx)

	h.GET("/products", controller.GetAll())
	h.POST("/products", controller.Create())
	h.GET("/products/:id", controller.Get())
	h.PUT("/products/:id", controller.Update())
	h.DELETE("/products/:id", controller.Delete())
}
