package service

import (
	database "go-postgresql-stocks-api/config"
	"go-postgresql-stocks-api/models"
)

func CreateStockService(stock models.Stock) int32 {
	db := database.CreateConnection()
	defer db.Close()
	sqlQuery := `INSERT INTO stocks(name,price,company) VALUES ($1,$2,$3)`
	db.QueryRow(sqlQuery,stock.Name,stock.Price,stock.Company)

}
func GetStockService(id int32) (models.Stock, error)        {}
func GetAllStockService() ([]models.Stock, error)           {}
func UpdateStockService(id int32, stock models.Stock) int32 {}
func DeleteStockService(id int32) int32                     {}
