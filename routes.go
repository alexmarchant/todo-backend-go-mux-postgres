package main

import (
  "github.com/gorilla/mux"
)

func makeRoutes() *mux.Router {
  r := mux.NewRouter()
  r.Methods("OPTIONS").Handler(commonHandlersHeadersOnly())
  r.Path("/").Methods("GET").Handler(
    commonHandlers(indexHandler))
  r.Path("/").Methods("POST").Handler(
    commonHandlers(createHandler))
  r.Path("/").Methods("DELETE").Handler(
    commonHandlers(deleteAllHandler))
	r.Path("/tasks/{id}").Methods("GET").Handler(
    commonHandlers(readHandler))
	r.Path("/tasks/{id}").Methods("PATCH").Handler(
    commonHandlers(updateHandler))
	r.Path("/tasks/{id}").Methods("DELETE").Handler(
    commonHandlers(deleteHandler))
  return r
}

