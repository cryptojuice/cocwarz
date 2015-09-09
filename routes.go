package main

import (
	"github.com/cryptojuice/cocwarz/handlers"
	"github.com/cryptojuice/cocwarz/handlers/attackers"
	"github.com/cryptojuice/cocwarz/handlers/clans"
	"github.com/cryptojuice/cocwarz/handlers/targets"
	"github.com/cryptojuice/cocwarz/handlers/wars"
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
	commonHandlers := alice.New(handlers.LoggingMiddleware, handlers.CORSMiddleware)
	router := httprouter.New()

	for _, route := range routes {
		router.Handle(route.method, route.path, handlers.HandlerWrapper(commonHandlers.ThenFunc(route.Handler)))
	}
	return router
}

var clanHandler = clans.NewClanHandler(getSession())
var warHandler = wars.NewWarHandler(getSession())
var targetHandler = targets.NewTargetHandler(getSession())
var attackerHandler = attackers.NewAttackerHandler(getSession())

var routes = Routes{
	Route{"GET", "/", handlers.IndexHandler},

	Route{"GET", "/wars", warHandler.Index},
	Route{"POST", "/wars", warHandler.Create},
	Route{"GET", "/wars/:id", warHandler.Show},
	Route{"PUT", "/wars/:id", warHandler.Update},
	Route{"DELETE", "/wars/:id", warHandler.Destroy},

	Route{"GET", "/wars/:id/targets", targetHandler.Index},
	Route{"POST", "/wars/:id/targets", targetHandler.Create},
	Route{"PUT", "/wars/:id/targets/:targetId", targetHandler.Update},

	Route{"GET", "/wars/:id/targets/:targetId/attackers", attackerHandler.Index},
	Route{"POST", "/wars/:id/targets/:targetId/attackers", attackerHandler.Create},
	Route{"PUT", "/wars/:id/targets/:targetId/attackers/:attackerId", attackerHandler.Update},
	Route{"DELETE", "/wars/:id/targets/:targetId/attackers/:attackerId", attackerHandler.Destroy},
}
