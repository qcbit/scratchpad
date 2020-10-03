// Beaking HTTP handlers into groups

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")
	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", restModule()))
	mainMux.Handle("/ui/", http.StripPrefix("/ui", uiModule()))

	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		panic(err)
	}
}

func restModule() http.Handler {
	restApi := http.NewServeMux()
	restApi.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1, "name":"john"}]`)
	})
	return restApi
}

func uiModule() http.Handler {
	ui := http.NewServeMux()
	ui.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<html><body>Hello from UI!</body></html>`)
	})
	return ui
}
