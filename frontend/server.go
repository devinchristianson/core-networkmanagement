package main

import (
	"log"
	"net/http"
	"strconv"
	"core-networkmanager/frontend/plugins"
	_ "core-networkmanager/frontend/plugins/root"
	_ "core-networkmanager/frontend/plugins/assets"
)

func main() {
	initDB()
	defer db.Close()
	var port = 8080
	mux := http.NewServeMux()
	activePlugins := []string{"root", "assets"}
	plugins.SetupPlugins(mux, nil, activePlugins)
	err := http.ListenAndServe(":" + strconv.Itoa(port), mux)
	if(err != nil) {
		log.Fatal("ListenAndServe failed with error: ", err)
	}

}
