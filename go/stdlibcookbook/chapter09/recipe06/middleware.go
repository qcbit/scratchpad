// Creating HTTP middleware layer
package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", Secure(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `[{"id":"1","login":"ffghi"},{"id":"2","login":"ffghj"}]`)
	}))

	http.ListenAndServe(":8080", mux)
}

func Secure(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sec := r.Header.Get("X-Auth")
		if sec != "authenticated" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(w, r)
	}
}
