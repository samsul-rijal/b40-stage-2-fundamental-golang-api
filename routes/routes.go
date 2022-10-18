package routes

import "github.com/gorilla/mux"

// Import gorilla/mux package here ...

// Create RouteInit function and Call TodoRoutes function here ...
func RouteInit(r *mux.Router) {
	TodoRoutes(r)
	UserRoutes(r)
}

func RouteInitV2(r *mux.Router) {
	TodoRoutesV2(r)
}
