package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB) //dari repo
	// masuk ke handler
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
}
