package main

import (
	"core-networkmanager/frontend/plugins"
	_ "core-networkmanager/frontend/plugins/root"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	var port = 8080
	mux := echo.New()
	activePlugins := []string{"root"}
	plugins.SetupPlugins(mux, nil, activePlugins)
	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	if err != nil {
		log.Fatal("ListenAndServe failed with error: ", err)
	}

}
