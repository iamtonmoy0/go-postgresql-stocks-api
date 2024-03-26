package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

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
