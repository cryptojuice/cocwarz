package main

import (
	"github.com/cryptojuice/cocwarz/handlers"
	"github.com/cryptojuice/cocwarz/handlers/clans"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

type router struct {
	*httprouter.Router
}

type Route struct {
	method  string
	path    string
	Handler http.HandlerFunc
}

type Routes []Route

func NewRouter() *httprouter.Router {
	commonHandlers := alice.New(handlers.LoggingHandler)
	router := httprouter.New()

	for _, route := range routes {
		router.Handle(route.method, route.path, handlers.HandlerWrapper(commonHandlers.ThenFunc(route.Handler)))
	}
	return router
}

var clanHandler = clans.NewClanHandler(getSession())

var routes = Routes{
	Route{
		"GET",
		"/",
		handlers.IndexHandler,
	},
	Route{
		"POST",
		"/clans",
		clanHandler.Create,
	},
}
