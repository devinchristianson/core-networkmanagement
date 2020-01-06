package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host 	   = "roach"
	dbUser     = "postgres"
	dbPassword = "postgres"
)

func initDB() {
	dbcreds := fmt.Sprintf("host=%s user=%s "+
    "password=%s sslmode=disable",
	host, dbUser, dbPassword)
	db, dberr := sql.Open("postgres", dbcreds)
	fmt.Println("DB status:", db.Ping())
	if dberr != nil {
		log.Fatal(dberr)
	}
}