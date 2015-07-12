package main

import (
	"github.com/cryptojuice/cocwarz/handlers"
	"github.com/cryptojuice/cocwarz/handlers/clans"
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
		"GET",
		"/wars/:warId",
		warHandler.Show,
	},
	Route{
		"POST",
		"/wars",
		warHandler.Create,
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
}

// GET    "/wars"        Get Wars
// POST   "/wars"        Create War
// PUT    "/wars/:warId" Update War
// DELETE "/wars/:warId" Delete War

// GET  "/wars/:warId/targets"           Get War Targets
// POST "/wars/:warId/targets"           Create Target
// PUT  "/wars/:warId/targets/:targetId" Update Target

// GET    "/wars/:warId/targets/:targetId/attackers"
// POST   "/wars/:warId/targets/:targetId/attackers"
// PUT    "/wars/:warId/targets/:targetId/attackers/:attackerId"
// DELETE "/wars/:warId/targets/:targetId/attackers/:attackerId"
