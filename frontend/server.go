package main

import (
	"log"
	"net/http"
	"strconv"
	_ "github.com/lib/pq"
)

func main() {
	initDB()
	var port = 8080
	mux := http.NewServeMux()
	fileserver := http.FileServer(FileSystem{http.Dir("./assets")})
	mux.Handle("/assets/", http.StripPrefix("/assets", fileserver))
	mux.HandleFunc("/", homePage)
	err := http.ListenAndServe(":" + strconv.Itoa(port), mux)
	if(err != nil) {
		log.Fatal("ListenAndServe failed with error: ", err)
	}
}
