package main

// Things to do for every incoming request:
// 	1. Extract HTTP method and URL path from request
// 	2. Check if any routes exist that match the method and path
// 	3. Invoke the route if there is a match
// 	4. Return a 404 if no match is found

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes []RouteEntry
}

// RouteEntry stores relevant information for each route
type RouteEntry struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Route is a helper function that creates a new RouteEntry and add it to the list of routes
func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	e := RouteEntry{
		Method:  method,
		Path:    path,
		Handler: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
	fmt.Println("routes: ", rtr.routes)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// return 404 for every request, for now
	http.NotFound(w, r)
}
