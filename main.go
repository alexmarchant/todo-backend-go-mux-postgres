package main

import (
  "net/http"
  "log"
  "os"
  "database/sql"
)

var db *sql.DB

func main() {
  db = connectToDB()
  r := makeRoutes()
  port := ":" + os.Getenv("PORT")

  log.Printf("Starting server on port %s", port)
  log.Fatal(http.ListenAndServe(port, r))
}
