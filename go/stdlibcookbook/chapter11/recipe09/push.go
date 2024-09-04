// Utiliting HTTP/2 server push
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if p, ok := w.(http.Pusher); ok {
			if err := p.Push("/app.css", nil); err != nil {
				log.Printf("Push err: %v", err)
			}
		}
		io.WriteString(w,
			`<html>
			<head>
				<link rel="stylesheet" type="text/css" href="app.css">
			</head>
			<body>
				<p>Hello</p>
			</body>
		</html>`)
	})

	http.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `p {
								text-align: center;
								color: red,
							}`)
	})

	if err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil); err != nil {
		panic(err)
	}
}
