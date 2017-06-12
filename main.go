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
  originsOk := handlers.AllowedOrigins([]string{"*"})
  headersOk := handlers.AllowedHeaders([]string{"Origin, X-Requested-With, Content-Type, Accept"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PATCH", "OPTIONS", "DELETE"})
  ignoreOptions := handlers.IgnoreOptions()

  // Start server
  log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk, ignoreOptions)(r)))
}
