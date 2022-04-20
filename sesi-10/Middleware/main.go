package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware1(endpoint))

	fmt.Println("Listening on port 9000")

	err := http.ListenAndServe(":9000", mux)

	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware pertama")
		role := r.Header.Get("role")
		if role != "admin" {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Forbidden",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func middleware2(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		role := r.Header.Get("role")
// 		if r.Method != "GET" {
// 			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 			return
// 		}

// 		fmt.Println("middleware kedua, role", role)

// 		if role != "admin" {
// 			http.Error(w, "role is undefined", http.StatusForbidden)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
