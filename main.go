package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &Router{}

	fmt.Println("Server running on :8000")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Error serving server: ", err)
	}
}
