package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	port = flag.Int("port", 8080, "Which port to listen")
	path = flag.String("path", "./", "The root path where want to serve.")
)

const VERSION = "0.1"

func main() {
	flag.Parse()
	rootPath, err := filepath.Abs(*path)
	if err != nil {
		log.Fatal("Path Error: %v", err)
	}
	log.Printf("HTTP Static File Server start at http://%s:%d", "localhost", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(rootPath))))
}
