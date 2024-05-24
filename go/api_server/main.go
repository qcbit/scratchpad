package main

import(
  "fmt"
  "net/http"
  "log"
  "os"
)

func main(){
  http.HandleFunc("/", helloWorld)
  fmt.Sprintf("Server started and listening on %s:9003", os.Getenv("CODESANDBOX_HOST"))
  log.Fatal(http.ListenAndServe(":9003", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, fmt.Sprintf("Server started and listening on %s:9003", os.Getenv("CODESANDBOX_HOST"))) 
}