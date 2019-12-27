package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

func initDB() {
	dbcreds := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	DB_USER, DB_PASSWORD, DB_NAME)
	db, dberr := sql.Open("roach", dbcreds)
	if dberr != nil {
		log.Fatal(dberr)
	}
	defer db.Close()
}