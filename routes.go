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
	commonHandlers := alice.New(handlers.LoggingHandler)
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
	Route{
		"GET",
		"/",
		handlers.IndexHandler,
	},
	Route{
		"GET",
		"/wars",
		warHandler.Index,
	},
	Route{
		"POST",
		"/wars",
		warHandler.Create,
	},
	Route{
		"GET",
		"/wars/:warId",
		warHandler.Show,
	},
	Route{
		"PUT",
		"/wars/:warId",
		warHandler.Update,
	},
	Route{
		"DELETE",
		"/wars/:warId",
		warHandler.Destroy,
	},
	Route{
		"GET",
		"/wars/:warId/targets",
		targetHandler.Index,
	},
	Route{
		"POST",
		"/wars/:warId/targets",
		targetHandler.Create,
	},
	Route{
		"PUT",
		"/wars/:warId/targets/:targetId",
		targetHandler.Update,
	},
	Route{
		"GET",
		"/wars/:warId/targets/:targetId/attackers",
		attackerHandler.Index,
	},
	Route{
		"POST",
		"/wars/:warId/targets/:targetId/attackers",
		attackerHandler.Create,
	},
	Route{
		"PUT",
		"/wars/:warId/targets/:targetId/attackers/:attackerId",
		attackerHandler.Update,
	},
	Route{
		"DELETE",
		"/wars/:warId/targets/:targetId/attackers/:attackerId",
		attackerHandler.Destroy,
	},
}
