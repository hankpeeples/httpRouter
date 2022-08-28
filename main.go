package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &Router{}

	fmt.Println("Server running on :8000")

	r.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		routeOutput("Router is working!", w, r)
	})

	r.Route("GET", `/hello/(?P<Message>\w+)`, func(w http.ResponseWriter, r *http.Request) {
		message := URLParam(r, "Message")
		routeOutput(fmt.Sprintf("Hello, %s", message), w, r)
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
