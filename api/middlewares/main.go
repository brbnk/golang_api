package middleware

import "github.com/julienschmidt/httprouter"

// httprouter.Hanle is a type --> type Handle func(http.ResponseWriter, *http.Request, Params)
type Middleware func(httprouter.Handle) httprouter.Handle

func chain(f httprouter.Handle, m ...Middleware) httprouter.Handle {
	if len(m) == 0 {
		return f
	}

	return m[0](chain(f, m[1:]...))
}

func Apply(f httprouter.Handle) httprouter.Handle {
	mw := []Middleware{
		SetHeader,
		SetCors,
	}
	return chain(f, mw...)
}
