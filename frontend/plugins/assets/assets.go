package assets

import (
	"fmt"
	"net/http"
	"strings"
	"core-networkmanager/frontend/plugins"
)
//Page struct to hold per-page data
type page struct {
	Location string
	Name string
}
//Plugin implements plugin interface
type Plugin struct {}

func init() {
	plugins.RegisterPlugin("assets", &Plugin{})
}

//Setup sets up endpoints and such
func (p Plugin ) Setup (mux *http.ServeMux) {
	fileserver := http.FileServer(FileSystem{http.Dir("./assets")})
	mux.Handle("/assets/", http.StripPrefix("/assets", fileserver))

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