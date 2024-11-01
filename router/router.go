package router

import (
	"todoist/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter sets up the router and routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Landing).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	return r
}
