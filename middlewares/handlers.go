package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgresql-stocks-api/models"
	"log"
	"net/http"
	"os"

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

	json.NewDecoder(r.Body).Decode(&stock)
}
func GetStock()     {}
func GetAllStocks() {}
func UpdateStock()  {}
func DeleteStock()  {}
