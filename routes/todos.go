package routes

import (
	// Import dumbmerch/handlers here ...
	"dumbmerch/handlers"

	"github.com/gorilla/mux"
)

func TodoRoutes(r *mux.Router) {

	// Create Routes here ...
	r.HandleFunc("/todos", handlers.FindTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", handlers.GetTodo).Methods("GET")
	r.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", handlers.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todo/{id}", handlers.DeleteTodo).Methods("DELETE")
}

func TodoRoutesV2(r *mux.Router) {

	// Create Routes here ...
	r.HandleFunc("/todos", handlers.FindTodosV2).Methods("GET")
}
