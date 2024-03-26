package service

import "go-postgresql-stocks-api/models"

func CreateStockService(stock models.Stock) int32           {}
func GetStockService(id int32) (models.Stock, error)        {}
func GetAllStockService() ([]models.Stock, error)           {}
func UpdateStockService(id int32, stock models.Stock) int32 {}
func DeleteStockService(id int32) int32                     {}
