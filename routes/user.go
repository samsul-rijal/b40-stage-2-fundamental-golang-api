package routes

import (
	"dumbmerch/handlers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {

	r.HandleFunc("/users", handlers.GetUser).Methods("GET")

}
