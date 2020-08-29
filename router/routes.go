package router

import (
	"net/http"

	handler "LoginServer/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"register",
		"POST",
		"/register",
		handler.Register,
	},
	Route{
		"login",
		"POST",
		"/login",
		handler.Login,
	},
	Route{
		"profile",
		"GET",
		"/profile",
		handler.Profile,
	},
	Route{
		"gethost",
		"POST",
		"/gethost",
		handler.GetHostID,
	},
}
