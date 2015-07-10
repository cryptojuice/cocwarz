package main

import "github.com/julienschmidt/httprouter"
import "github.com/justinas/alice"
import "github.com/cryptojuice/cocwarz/handlers"
import "net/http"

type router struct {
	*httprouter.Router
}

type Route struct {
	method  string
	path    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		handlers.IndexHandler,
	},
}

func NewRouter() *httprouter.Router {
	commonHandlers := alice.New(handlers.LoggingHandler)
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, handlers.HandlerWrapper(commonHandlers.ThenFunc(route.Handler)))
	}
	return router
}
