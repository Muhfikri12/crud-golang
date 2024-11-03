package database

import (
	"database/sql"
	"log"
)

func ConnectDb() (*sql.DB, error)  {
	connStr := "user=postgre dbname=golang_database sslmode=disable password=admin host=localhost"
	db, err := sql.Open("postgre", connStr)
	if err := db.Ping(); err != nil {
			db.Close()
			log.Fatal(err)
	}

	log.Println("Successfuly Connected")

	return db, err
}