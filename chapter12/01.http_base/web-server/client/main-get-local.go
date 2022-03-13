package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	go httpDirectGet()
	go httpGetWithContext(ctx)
	<-ctx.Done()
}

func httpDirectGet() {
	resp, err := http.Get("http://localhost:8088")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func httpGetWithContext(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	req, err := http.NewRequest("get", "http://localhost:8088", nil)
	if err != nil {
		log.Fatal("无法生成请求：", err)
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("无法发送请求：", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("无法读取返回内容：", err)
	}
	fmt.Println(string(data))
}
