package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	const (
		host    = "roach"
		dbUser  = "root"
		dbPort  = "26257"
		dbTable = "postgres"
	)
	dbOpts := "?sslmode=disable"
	dbcreds := fmt.Sprintf("postgresql://%s@%s:%s/%s%s",
		dbUser, host, dbPort, dbTable, dbOpts)
	db, dberr := sqlx.Connect("pgx", dbcreds)
	if dberr != nil {
		log.Fatal(dberr)
	}
	tables := [5]string{"Users", "Networks", "Hosts", "Groups", "Domains"}
	count := 1
	checkInit, err := db.Prepare("select count(*) as count from information_schema.tables where table_name = $1;")
	if err != nil {
		log.Fatal(err)
	}
	initialized := true
	for _, table := range tables {
		err := checkInit.QueryRow(table).Scan(&count)
		initialized = initialized && (count == 1)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = db.Exec("select 0")
	if err != nil {
		log.Fatal(err)
	}
	/*if !initialized {
		fmt.Print("Initializing database")
		file, ferr := ioutil.ReadFile("init.sql")
		if ferr != nil {
			log.Fatal(ferr)
		}
		requests := strings.Split(string(file), ";\n")
		for _, request := range requests {
			_, err := db.Exec(request)
			if err != nil {
				log.Fatal(err)
			}
		}
	}*/
}
