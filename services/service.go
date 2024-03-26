package service

import (
	"database/sql"
	"fmt"
	database "go-postgresql-stocks-api/config"
	"go-postgresql-stocks-api/models"
	"log"
)

// create new stock
func CreateStockService(stock models.Stock) int64 {
	db := database.CreateConnection()
	defer db.Close()
	var id int64
	sqlQuery := `INSERT INTO stocks(name,price,company) VALUES ($1,$2,$3)`
	err := db.QueryRow(sqlQuery, stock.Name, stock.Price, stock.Company).Scan(id)
	if err != nil {
		log.Fatalf("failed to create stocks %v", err)
	}
	fmt.Printf("new stock created")
	return id

}

// get stock by id
func GetStockService(id int32) (models.Stock, error) {
	db := database.CreateConnection()
	defer db.Close()
	var stock models.Stock
	sqlStatement := `SELECT * FROM stocks WHERE stockid = $1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("no rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		fmt.Println("Unable to scan row %v", err)
	}
	return stock, err

}

// get all stocks
func GetAllStockService() ([]models.Stock, error) {
	db := database.CreateConnection()
	defer db.Close()
	var stocks []models.Stock
	sqlStatement := `SELECT * FROM stocks`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf(err)
	}

	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatal(err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, err
}

// update stocks
func UpdateStockService(id int32, stock models.Stock) int64 {
	db := database.CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE stocks SET name=$1,price=$2,company=$3 WHERE  stockid=$4;`
	res, err := db.Exec(sqlStatement, stock.Name, stock.Price, stock.Company, id)
	if err != nil {
		log.Fatal(err)
	}

	data, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return data

}
func DeleteStockService(id int32) int32 {
	db := database.CreateConnection()
	defer db.Close()
}
