package main

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/jackc/pgx/stdlib"
)

func addHost() {
	fmt.Printf("adding host\n")
	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	hostid := -1
	err = tx.QueryRow("INSERT INTO \"Hosts\" (currenttimestamp) VALUES (current_timestamp) RETURNING (hostid)").Scan(&hostid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", hostid)
	mac := "ffffffffffff"
	int, err := strconv.ParseUint(mac, 16, 48)
	bits := fmt.Sprintf("%48b", int)
	_, err = tx.Exec("INSERT INTO \"HostDetails\" (timestamp, hostid, mac, description) VALUES (current_timestamp, $1, $2, 'test host')", hostid, bits)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
func modifyHost() {
	//modify old record adding end time, return all values
	//make changes
	//insert new record
}
