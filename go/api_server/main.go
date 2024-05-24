package main

import(
  "fmt"
  "net/http"
  "log"
  "os"
  
  "github.com/gorilla/mux"
)

func main(){
  r := mux.NewRouter()
  http.HandleFunc("/", helloWorld)
  log.Fatal(http.ListenAndServe(":9003", nil))
  http.Handle("/", r)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, fmt.Sprintf("Server started and listening on %s:9003", os.Getenv("CODESANDBOX_HOST"))) 
}