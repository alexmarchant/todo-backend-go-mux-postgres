package main

import (
  "database/sql"
  "log"
  "os"
  _ "github.com/lib/pq"
)

func database() *sql.DB {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    log.Fatal(err)
  }

  ensureTables(db)

  return db
}

func ensureTables(db *sql.DB) {
  query := `CREATE TABLE IF NOT EXISTS tasks
(
id SERIAL,
title TEXT NOT NULL,
completed BOOL NOT NULL DEFAULT FALSE,
"order" INT NOT NULL DEFAULT 0,
CONSTRAINT tasks_pkey PRIMARY KEY (id)
)`
 if _, err := db.Exec(query); err != nil {
    log.Fatal(err)
  }
}
