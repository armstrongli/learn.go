package main

import (
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(2 * time.Second)
		writer.Write([]byte("hello,您好"))
	}))
}
