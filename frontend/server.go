package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"html/template"
	"github.com/lib/pq"
)

func main() {
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

//FileSystem implements http.Filesystem, with Open() that doesnt allow Dirs
type FileSystem struct {
	fs http.FileSystem
}
//Open overrides http.FileSystem.Open in order to prevent Directory access
func (cfs FileSystem) Open(path string) (http.File, error) {
	fmt.Println("filepath",path)
	f, err := cfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	st, err := f.Stat()
	if st.IsDir() {
		fmt.Println("Dir", path, "is not allowed")
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := cfs.fs.Open(index); err != nil {
			return nil, err
		}
	}
	return f, nil
}