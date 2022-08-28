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

// Router holds all routes
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
}

// PrintRoutes simply prints the routes that are being listened for
func (rtr *Router) PrintRoutes() {
	fmt.Println("Listening for:")
	for i, route := range rtr.routes {
		fmt.Printf("[%d] %s request on path '%s'\n", i, route.Method, route.Path)
	}
}

// Match returns whether a match was found or not.
func (re *RouteEntry) Match(r *http.Request) bool {
	if r.Method != re.Method {
		return false // Method mismatch
	}

	if r.URL.Path != re.Path {
		return false // Path mismatch
	}

	return true
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Loop over all routes and check to see if one of them matches
	for _, e := range rtr.routes {
		match := e.Match(r)
		if !match {
			continue
		}

		// Match found, call handler and return
		e.Handler.ServeHTTP(w, r)
		return
	}

	// No matches found, 404
	http.NotFound(w, r)
}

// URLParam extracts a parameter from the URL by name
func URLParam(r *http.Request, name string) string {
	ctx := r.Context()

	// Cast `interface{}` from ctx.Value() to map used to store params
	params := ctx.Value("params").(map[string]string)
	return params[name]
}
