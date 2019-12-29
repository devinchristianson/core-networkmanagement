package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "test"
)

func initDB() {
	dbcreds := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	dbUser, dbPassword, dbName)
	db, dberr := sql.Open("postgres", dbcreds)
	fmt.Println("DB status:", db.Ping())
	if dberr != nil {
		log.Fatal(dberr)
	}
}