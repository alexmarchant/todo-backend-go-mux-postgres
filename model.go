package main

import (
  "strconv"
  "net/http"
)

type task struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
  Order int `json:"order"`
}

func (t *task) createTask() error {
  err := db.QueryRow(
    "INSERT INTO tasks(title, completed, \"order\") VALUES($1, $2, $3) RETURNING id",
    t.Title, t.Completed, t.Order).Scan(&t.ID)

  if err != nil {
    return err
  }

  return nil
}

func (t *task) getTask() error {
  return db.
    QueryRow("SELECT title, completed, \"order\" FROM tasks WHERE id=$1", t.ID).
    Scan(&t.Title, &t.Completed, &t.Order)
}

func (t *task) updateTask() error {
  _, err := db.Exec("UPDATE tasks SET title=$1, completed=$2, \"order\"=$3 WHERE id=$4", t.Title, t.Completed, t.Order, t.ID)
  return err
}

func (t *task) deleteTask() error {
  _, err := db.Exec("DELETE FROM tasks WHERE id=$1", t.ID)
  return err
}

func (t *task) url(r *http.Request) string {
  var protocol string

  if r.Host == "localhost:3000" {
    protocol = "http"
  } else {
    protocol = "https"
  }
    
  return protocol + "://" + r.Host + "/tasks/" + strconv.Itoa(t.ID)
}

func getTasks() ([]task, error) {
  rows, err := db.Query("SELECT id, title, completed, \"order\" FROM tasks")

  if err != nil {
      return nil, err
  }

  defer rows.Close()

  tasks := []task{}

  for rows.Next() {
    var t task
    if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.Order); err != nil {
      return nil, err
    }
    tasks = append(tasks, t)
  }

  return tasks, nil
}

func deleteTasks() error {
  _, err := db.Exec("DELETE FROM tasks")
  return err
}

