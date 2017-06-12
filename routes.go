package main

import (
  "github.com/gorilla/mux"
  "net/http"
)

func router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/", options).Methods("OPTIONS")
  r.HandleFunc("/", IndexHandler).Methods("GET")
  r.HandleFunc("/", CreateHandler).Methods("POST")
  r.HandleFunc("/", DeleteAllHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", ReadHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", UpdateHandler).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", DeleteHandler).Methods("DELETE")
  return r
}

func options(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  return
}
