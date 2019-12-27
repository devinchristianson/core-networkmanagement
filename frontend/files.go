package main

import (
	"fmt"
	"net/http"
	"strings"
)

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