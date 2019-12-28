package main

import (
	"log"
	"net/http"
	"strconv"
	_ "github.com/lib/pq"
	"core-networkmanager/frontend/plugins"
	_ "core-networkmanager/frontend/plugins/root"
	_ "core-networkmanager/frontend/plugins/assets"
)

func main() {
	initDB()
	var port = 8080
	mux := http.NewServeMux()
	plugins.Setup(mux)
	plugins.Plugins["root"].Setup(mux)
	plugins.Plugins["assets"].Setup(mux)
	err := http.ListenAndServe(":" + strconv.Itoa(port), mux)
	if(err != nil) {
		log.Fatal("ListenAndServe failed with error: ", err)
	}
}
