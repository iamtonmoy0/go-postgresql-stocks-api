package main

import (
	"fmt"
	route "go-postgresql-stocks-api/routes"
	"log"
	"net/http"
)

func main() {
	r := route.Router()
	fmt.Println("Server is starting on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
