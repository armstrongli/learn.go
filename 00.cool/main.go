package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/jianqli/tmp")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
