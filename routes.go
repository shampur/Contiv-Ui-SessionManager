package main

import "net/http"

// Route General routing struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes handeld by the web server
type Routes []Route

var routes = Routes{
	Route{
		"options",
		"OPTIONS",
		"/",
		preFlight,
	},
	Route{
		"login",
		"POST",
		"/",
		loginUser,
	},
	Route{
		"validateApp",
		"GET",
		"/",
		validateAppState,
	},
	Route{
		"logout",
		"DELETE",
		"/",
		logOutUser,
	},
}
