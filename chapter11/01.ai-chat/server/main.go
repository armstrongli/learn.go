package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var qa = map[string]string{
	"你好":      "你好",
	"你是谁？":    "我是小S",
	"你是男是女？":  "你猜猜看",
	"今天天气怎样？": "今天天气不错哦！",
	"再见":      "再见",
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("wanring: 建立连接失败：", err)
			continue
		}
		fmt.Println(conn)

		// talk(conn)
		go talk(conn)
	}
}

func talk(conn net.Conn) {
	defer fmt.Println("结束链接：", conn)
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("WARNING：读数据失败：", err)
			continue
		}
		content := buf[:valid]
		resp, ok := qa[string(content)]
		if !ok {
			log.Println("没有找到回答，问他说了什么")
			conn.Write([]byte(`我听不懂你在讲什么`)) // handle error
			continue
		}
		conn.Write([]byte(resp)) // handle error
		if string(content) == "再见" {
			break
		}
	}
}
