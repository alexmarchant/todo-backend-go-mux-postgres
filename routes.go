package main

import (
  "github.com/gorilla/mux"
)

func router() *mux.Router {
  r := mux.NewRouter()
  r.Methods("OPTIONS").HandlerFunc(OptionsHandler)
  r.Path("/").Methods("GET").HandlerFunc(IndexHandler)
  r.Path("/").Methods("POST").HandlerFunc(CreateHandler)
  r.Path("/").Methods("DELETE").HandlerFunc(DeleteAllHandler)
	r.Path("/tasks/{id}").Methods("GET").HandlerFunc(ReadHandler)
	r.Path("/tasks/{id}").Methods("PATCH").HandlerFunc(UpdateHandler)
	r.Path("/tasks/{id}").Methods("DELETE").HandlerFunc(DeleteHandler)
  return r
}

