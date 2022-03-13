package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`hello`))
	}))
	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`rank`))
	}))
	m.Handle("/history/xiaoqiang", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`xiaoqiang的黑历史`))
	}))
	m.Handle("/history/jesse", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`jesse的健身馆`))
	}))
	// .....
	m.Handle("/history", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		qp := request.URL.Query()
		name := qp.Get("name")
		// if/else
		// switch/case
		// switch request.Method{
		// case "get"://....
		// case "post"://.....
		// }
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf(`%s: %s的健身馆`, request.Method, name)))
	}))
	http.ListenAndServe(":8080", m)
}
