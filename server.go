package mosquito


import (
	"net/http"
)

var server	*Server

type Method		string

type Server struct {
	middlewares			[]Middleware
	routes				[]*Route
}

func (server *Server) Use(middleware Middleware) *Server {
	server.middlewares = append(server.middlewares, middleware)
	return server
}

func (server *Server) Get(path string, handler Handler) *Server {
	route := NewRoute(path, Method("GET"), handler)

	server.routes = append(server.routes, route)

	return server
}

func (server *Server) Post(path string, handler Handler) *Server {
	route := NewRoute(path, Method("POST"), handler)

	server.routes = append(server.routes, route)

	return server
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var next func()

	req := &Request{Request:r}
	res := Response{ResponseWriter:w}

	index := 0
	middlewares := server.middlewares

    next = func() {
	    if len(middlewares) > index {
		    nextMiddleware := middlewares[index]
		    index++

			nextMiddleware.ServeHTTP(res, req, next)
		}
    }

	next();
}

