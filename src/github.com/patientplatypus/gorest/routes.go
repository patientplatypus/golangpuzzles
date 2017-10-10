package main

import (
	"net/http"

	"github.com/patientplatypus/gorest/users"

	"github.com/patientplatypus/gorest/character"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"ClassType",
		"POST",
		"/character/class",
		character.ClassType,
	},
	Route{
		"RaceType",
		"POST",
		"/character/race",
		character.RaceType,
	},
	Route{
		"BackgroundType",
		"POST",
		"/character/background",
		character.BackgroundType,
	},
	Route{
		"RegisterUser",
		"POST",
		"/users/registeruser",
		users.RegisterUser,
	},
	Route{
		"LoginUser",
		"GET",
		"/users/loginuser/{username}/{password}",
		users.LoginUser,
	},
}
