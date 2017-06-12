package main

import (
  "github.com/gorilla/mux"
)

func router() *mux.Router {
  r := mux.NewRouter()
  r.Methods("OPTIONS").HandlerFunc(OptionsHandler)
  r.HandleFunc("/", IndexHandler).Methods("GET")
  r.HandleFunc("/", CreateHandler).Methods("POST")
  r.HandleFunc("/", DeleteAllHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", ReadHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", UpdateHandler).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", DeleteHandler).Methods("DELETE")
  return r
}

