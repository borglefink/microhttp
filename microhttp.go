// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file. package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"config"
)

var fullhostname = "localhost:80"

// ------------------------------------------
// init
// ------------------------------------------
func init() {
	var c = config.Load()
	fullhostname = fmt.Sprintf("%s:%d", c.Hostname, c.Port)
}

// ------------------------------------------
// exists
// ------------------------------------------
func exists(path string) (os.FileInfo, bool) {
	fileinfo, err := os.Stat(path)
	if err == nil {
		return fileinfo, true
	}
	if os.IsNotExist(err) {
		return fileinfo, false
	}
	return fileinfo, false
}

// ------------------------------------------
// mainHandler
// ------------------------------------------
func mainHandler(w http.ResponseWriter, r *http.Request) {
	var url = r.URL.Path[1:]

	var fileinfo, exists = exists(url)

	if exists && fileinfo.IsDir() {
		url += "/index.html"
	}

	http.ServeFile(w, r, url)
}

// ------------------------------------------
// main
// ------------------------------------------
func main() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(fullhostname, nil))
}
