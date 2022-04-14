package main

import (
	"fmt"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", greet)

	http.ListenAndServe(port, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprintf(w, msg)
}
