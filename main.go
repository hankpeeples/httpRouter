package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &Router{}

	fmt.Println("Server running on :8000")

	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		routeOutput("Router is working!", w, r)
	})

	r.Route(http.MethodGet, `/hello/(?P<Message>\w+)`, func(w http.ResponseWriter, r *http.Request) {
		message := URLParam(r, "Message")
		routeOutput(fmt.Sprintf("Hello, %s", message), w, r)
	})

	r.Route(http.MethodGet, "/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong, panicking...")
	})

	r.PrintRoutes()

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Error serving server: ", err)
	}
}

func routeOutput(output string, w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(output))
	if err != nil {
		fmt.Println("Error writing response: ", err)
	}
}
