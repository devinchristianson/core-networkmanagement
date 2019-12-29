package root 

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"core-networkmanager/frontend/plugins"
)

func setup() bool {
	plugins.RegisterEndpoint("/", homePage)
	return true
}

//Page struct to hold per-page data
type page struct {
	Location string
	Name string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("index.gohtml")
		if err != nil {
			log.Fatal("Parse failed: ", err)
		}
		home := page{Location: "/", Name: "Home"}
		search := page{Location: "/search", Name: "Search"}
		data := struct { Pages []*page }{ []*page{ &home, &search } }
		if err := t.ExecuteTemplate(w, "index", data); err != nil {
			log.Fatal("ExecuteTemplate failed:", err)
		}
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}