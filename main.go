package main

// import fmt, net/http, github.com/gorilla/mux here ...
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Create main function to show "hello world" here ...
func main() {

	fmt.Println("Hallo World")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", r)
}
