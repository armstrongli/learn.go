package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var p string
	flag.StringVar(&p, "path", ".", "the path to expose as http")
	var port int
	flag.IntVar(&port, "port", 8080, "the port to expose")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(p)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
