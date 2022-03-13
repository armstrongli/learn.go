package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		qp := request.URL.Query()
		data, _ := json.Marshal(qp)
		writer.Write([]byte("hello,您好:" + string(data)))
	}))
}
