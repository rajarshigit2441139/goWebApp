package router

import (
	"github.com/rajarshigit2441139/goWebApp/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	return r
}
