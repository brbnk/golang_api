package helloworld

import (
	"encoding/json"
	"net/http"

	middleware "github.com/brbnk/core/api/middlewares"
	"github.com/julienschmidt/httprouter"
)

type HttpResponse struct {
	HttpCode int
	Message  string
}

func sayHelloWorld() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		resp := HttpResponse{
			HttpCode: 200,
			Message:  "Hello, World!",
		}

		response, _ := json.Marshal(resp)
		w.Write(response)
	}
}

func Do() httprouter.Handle {
	mw := []middleware.Middleware{
		middleware.SetHeader,
	}
	return middleware.Chain(sayHelloWorld(), mw...)
}
