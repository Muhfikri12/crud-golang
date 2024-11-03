package main

import (
	"log"

	"github.com/Muhfikri12/crud-golang/database"
	_ "github.com/lib/pq"
)
	
func main() {
	
	db, err := database.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}