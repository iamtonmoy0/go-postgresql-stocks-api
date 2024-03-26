package route

import (
	controller "go-postgresql-stocks-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// initialize the router
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", controller.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", controller.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", controller.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", controller.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", controller.DeleteStock).Methods("DELETE", "OPTIONS")
}
