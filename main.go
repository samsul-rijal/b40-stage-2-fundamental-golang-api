package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Declare Todos Struct here ...
type Todos struct {
	Id     string
	Title  string
	IsDone bool
}

// Declare Todos Global Variable ...
var todos = []Todos{
	{
		Id:     "1",
		Title:  "Cuci Tangan",
		IsDone: true,
	},
	{
		Id:     "2",
		Title:  "Jaga Jarak",
		IsDone: false,
	},
}

func main() {
	r := mux.NewRouter()

	// Create routes here ...
	r.HandleFunc("/todos", FindTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", GetTodo).Methods("GET")

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

// Create FindTodos Function here ...
func FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// Create GetTodo Function here ...
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var todoDataByID Todos
	var isGetTodo = false

	for _, todo := range todos {
		// jika id dari params sama dengan id dari todos sama, maka
		if id == todo.Id {
			isGetTodo = true
			todoDataByID = todo
		}
		// fmt.Println(todo)
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoDataByID)
}

// Create CreateTodo Function here ...

// Create UpdateTodo Function here ...

// Create DeleteTodo Function here ...
