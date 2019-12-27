package main 

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

//Page struct to hold per-page data
type Page struct {
	Location string
	Name string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("index.gohtml")
		if err != nil {
			log.Fatal("Parse failed: ", err)
		}
		home := Page{Location: "/", Name: "Home"}
		search := Page{Location: "/search", Name: "Search"}
		data := struct { Pages []*Page }{ []*Page{ &home, &search } }
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