// Handling HTTP requests
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "User GET")
		}
		if r.Method == http.MethodPost {
			fmt.Fprintln(w, "User POST")
		}
	})

	itemMux := http.NewServeMux()
	itemMux.HandleFunc("/items/clothes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Clothes")
	})
	mux.Handle("/items/", itemMux)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/ports", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ports")
	})

	mux.Handle("/admin/", http.StripPrefix("/admin", adminMux))

	http.ListenAndServe(":8080", mux)
}
