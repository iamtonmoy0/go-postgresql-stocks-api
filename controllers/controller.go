package controller

import (
	"encoding/json"
	"fmt"
	"go-postgresql-stocks-api/models"
	service "go-postgresql-stocks-api/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	ID      int32  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create request
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode incoming request.")
	}

	insertId := service.CreateStockService(stock)
	res := Response{
		ID: insertId, Message: "Stock created!",
	}

	json.NewEncoder(w).Encode(res)
}
func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	stock, err := service.GetStockService(int32(id))
	if err != nil {
		log.Fatal("failed to get stock")
	}
	json.NewEncoder(w).Encode(stock)
}
func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := service.GetAllStocksService()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(stocks)
}
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode incoming request.%v", err)
	}
	updateRows := service.UpdateStockService(int32(id, stock))
	if err != nil {
		log.Fatal("failed to update %v", updateRows)
	}
	msg := fmt.Sprintf("stock updated successfully %v", updateRows)
	res := Response{ID: int32(id), Message: msg}
	json.NewEncoder(w).Encode(res)

}
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	deletedRows := service.DeleteStockService(int32(id))
	msg := fmt.Sprintf("stocks deleted %v", deletedRows)
	res := (Response{ID: int32(id), Message: msg})
	json.NewEncoder(w).Encode(res)
}
