package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rajarshigit2441139/goWebApp/models"
)

var todos []models.Todo

func init() {
	todos = append(todos, models.Todo{ID: "1", Title: "Learn Go", Status: "pending"})
	todos = append(todos, models.Todo{ID: "2", Title: "Read a book", Status: "completed"})
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(1000000))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			var todo models.Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
