package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_DB"))
	fmt.Print(db)
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to db")
	return db
}
