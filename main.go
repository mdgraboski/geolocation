package main

import (
	"io"
	"os"
	"log"
	"fmt"
	"html"
	"strings"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		path := html.EscapeString(r.URL.Path)
		lastSlash := strings.LastIndex(path, "/")
		if lastSlash == -1 {
			log.Fatalf("ERROR: couldn't parse image URL.")
		}
		filename := path[lastSlash+1:]
		
		if strings.Contains(filename, "europe") {
			serveFile(w, "europe.jpg")
		} else if strings.Contains(filename, "usa") {
			serveFile(w, "usa.jpg")
		} else {
			serveFile(w, "earth.jpg")
		}
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func serveFile(w http.ResponseWriter, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	log.Printf("INFO: Served %s", filename)
}