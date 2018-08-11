package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  r := customRouter()
  http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world")
}

func customRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/hello", handler).Methods("GET")
  return r
}
