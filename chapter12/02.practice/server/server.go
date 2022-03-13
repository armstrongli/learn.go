package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"learn.go/chapter12/02.practice/frinterface"
	"learn.go/pkg/apis"
)

func main() {
	var rankServer frinterface.ServeInterface = NewFatRateRank()

	m := http.NewServeMux()
	m.Handle("/register", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据：%s", err)))
			return
		}
		var pi *apis.PersonalInformation
		if err := json.Unmarshal(payload, &pi); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据：%s", err)))
			return
		}
		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("注册失败：%s", err)))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`success`))
	}))
	m.Handle("/personalinfo", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据：%s", err)))
			return
		}
		var pi *apis.PersonalInformation
		if err := json.Unmarshal(payload, &pi); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据：%s", err)))
			return
		}
		if fr, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("更新失败：%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(fr)
			writer.Write(data)
		}
	}))
	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		name := request.URL.Query().Get("name")
		if name == "" {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("name参数未设置"))
			return
		}
		if fr, err := rankServer.GetFatRate(name); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("获取排行数据失败：%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(fr)
			writer.Write(data)
		}
	}))
	m.Handle("/ranktop", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if frTop, err := rankServer.GetTop(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("获取排行数据失败：%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(frTop)
			writer.Write(data)
		}
	}))
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatal(err)
	}
}
