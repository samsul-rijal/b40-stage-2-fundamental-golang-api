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
	r.HandleFunc("/todos", CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

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
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var data Todos

	json.NewDecoder(r.Body).Decode(&data)

	todos = append(todos, data)
	// todos.push(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// Create UpdateTodo Function here ...
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var data Todos
	var isGetTodo = false

	json.NewDecoder(r.Body).Decode(&data)

	for idx, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			todos[idx] = data
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// Create DeleteTodo Function here ...
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var isGetTodo = false
	var index = 0

	for idx, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			index = idx
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found")
		return
	}

	todos = append(todos[:index], todos[index+1:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ID: " + id + " delete success")
}
