package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgresql-stocks-api/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Response struct {
	ID      int32  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		log.Fatal("failed to connect to db")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to db")
	return db
}

// create request
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode incoming request.")
	}

	insertId := insertStock(stock)
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
	stock, err := getStock(int32(id))
	if err != nil {
		log.Fatal("failed to get stock")
	}
	json.NewEncoder(w).Encode(stock)
}
func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
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
	updateRows := updateStock(int32(id, stock))
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
	deletedRows := deleteStock(int32(id))
	msg := fmt.Sprintf("stocks deleted %v", deletedRows)
	res := (Response{ID: int32(id), Message: msg})
	json.NewEncoder(w).Encode(res)
}
