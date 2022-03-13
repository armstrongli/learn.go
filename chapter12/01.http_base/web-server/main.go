package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Body == nil {
			writer.Write([]byte("no body"))
			return
		}
		// data, _ := ioutil.ReadAll(request.Body)
		// defer request.Body.Close()
		// encoded := base64.StdEncoding.EncodeToString(data)
		writer.Write([]byte("append(data, []byte(encoded)...)"))
	}))
}
