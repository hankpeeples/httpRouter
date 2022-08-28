package main

// Things to do for every incoming request:
// 	1. Extract HTTP method and URL path from request
// 	2. Check if any routes exist that match the method and path
// 	3. Invoke the route if there is a match
// 	4. Return a 404 if no match is found

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
)

// Router holds all routes
type Router struct {
	routes []RouteEntry
}

// RouteEntry stores relevant information for each route
type RouteEntry struct {
	Path    *regexp.Regexp
	Method  string
	Handler http.HandlerFunc
}

// Route is a helper function that creates a new RouteEntry and add it to the list of routes
func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")

	e := RouteEntry{
		Method:  method,
		Path:    exactPath,
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
func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		// no match
		return nil
	}

	// Create map to store URL parameters
	params := make(map[string]string)
	groupNames := ent.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}

	return params
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Loop over all routes and check to see if one of them matches
	for _, e := range rtr.routes {
		params := e.Match(r)
		if params == nil {
			continue // no match found
		}

		// Create new request with params stored in context
		ctx := context.WithValue(r.Context(), "params", params)
		e.Handler.ServeHTTP(w, r.WithContext(ctx))
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
