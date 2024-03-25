package route

import "github.com/gorilla/mux"

func Router() *mux.Router {
	// initialize the router
	router := mux.NewRouter()

	router.HandleFunc("/")
}
