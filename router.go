package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns a global mux router
func NewRouter() *mux.Router {
	fmt.Println("Creating mux Router")
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		fmt.Println("Tieing handler = ", route.HandlerFunc, "for path = ", route.Pattern, " to mux router")
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
