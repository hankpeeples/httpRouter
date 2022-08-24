package main

import "net/http"

type Router struct{}

func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// return 404 for every request, for now
	http.NotFound(w, r)
}
