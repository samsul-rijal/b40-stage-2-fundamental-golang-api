package main

import (
	// Import dumbmerch/database here ...
	// Import dumbmerch/pkg/mysql here ...
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB here ...
	// koneksi ke database
	mysql.DatabaseInit()

	// Run migration here ...
	// migrate table
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
