package main

import (
  "net/http"
  "log"
  "os"
  "database/sql"
  "github.com/gorilla/handlers"
)

var db *sql.DB

func main() {
  db = database()
  r := router()
  port := ":" + os.Getenv("PORT")
  log.Printf("Starting server on port %s", port)

  // CORS
  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PATCH", "OPTIONS"})

  // Start server
  log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
