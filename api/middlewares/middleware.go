package middleware

import "github.com/julienschmidt/httprouter"

// httprouter.Hanle is a type --> type Handle func(http.ResponseWriter, *http.Request, Params)
type Middleware func(httprouter.Handle) httprouter.Handle

func Chain(f httprouter.Handle, m ...Middleware) httprouter.Handle {
	if len(m) == 0 {
		return f
	}

	return m[0](Chain(f, m[1:]...))
}
