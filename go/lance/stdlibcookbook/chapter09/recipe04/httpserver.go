// Creating the HTTP server
package main

import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct{}

func (s SimpleHTTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello world")
}

func main() {
	fmt.Println("Starting HTTP server on port 8080")
	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}
	s.ListenAndServe()
}
