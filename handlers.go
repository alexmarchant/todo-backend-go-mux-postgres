package main

import (
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
  "database/sql"
  "encoding/json"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  addDefaultHeaders(w)
  tasks, err := getTasks()
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  responseTasks := []responseTask{}
  for _, t := range tasks {
    rt := makeResponseTask(&t, r)
    responseTasks = append(responseTasks, *rt)
  }
  respondWithJSON(w, http.StatusOK, responseTasks)
}

func DeleteAllHandler(w http.ResponseWriter, r *http.Request) {
  addDefaultHeaders(w)
  if err := deleteTasks(); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  addDefaultHeaders(w)
  var t task
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&t); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  defer r.Body.Close()

  if err := t.createTask(); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, makeResponseTask(&t, r))
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
  addDefaultHeaders(w)
  id, err := getID(w, r)
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid task ID")
    return
  }
  
  t := &task{ID: id}
  if err := t.getTask(); err != nil {
    switch err {
    case sql.ErrNoRows:
      respondWithError(w, http.StatusNotFound, "Task not found")
    default:
      respondWithError(w, http.StatusInternalServerError, err.Error())
    }
    return
  }

  respondWithJSON(w, http.StatusOK, makeResponseTask(t, r))
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
  addDefaultHeaders(w)
  id, err := getID(w, r)
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid task ID")
    return
  }

  var t task
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&t); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  defer r.Body.Close()
  t.ID = id

  if err := t.updateTask(); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, makeResponseTask(&t, r))
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
  addDefaultHeaders(w)
  id, err := getID(w, r)
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid task ID")
    return
  }

  t:= &task{ID: id}
  if err := t.deleteTask(); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
  addDefaultHeaders(w)
  w.WriteHeader(http.StatusOK)
  return
}

// Helpers

func getID(w http.ResponseWriter, r *http.Request) (id int, err error) {
  vars := mux.Vars(r)
  id, err = strconv.Atoi(vars["id"])
  return
}

func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

func addDefaultHeaders(w http.ResponseWriter) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
  w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
}

type responseTask struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
  Order int `json:"order"`
  URL string `json:"url"`
}

func makeResponseTask(t *task, r *http.Request) *responseTask {
  return &responseTask{
    ID: t.ID,
    Title: t.Title,
    Completed: t.Completed,
    Order: t.Order,
    URL: t.url(r)}
}

