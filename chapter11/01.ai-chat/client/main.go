package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("拨号失败：", err)
	}
	defer conn.Close()
	fmt.Println("连接成功，开始聊天吧：")
	for {
		r := bufio.NewReader(os.Stdin)
		input, _, _ := r.ReadLine() // todo handle error
		if len(input) != 0 {
			talk(conn, string(input))
		}
	}
	// talk(conn, "你好")
	// talk(conn, "你好")
	// talk(conn, "你好")
	// talk(conn, "你是谁？")
	// talk(conn, "再见")
}

func talk(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("发送消息失败:", err)
	} else {
		data := make([]byte, 1024)
		validLen, err := conn.Read(data)
		if err != nil {
			log.Println("WARNING:读取服务器返回数据时出错：", err)
		} else {
			validData := data[:validLen]
			log.Println("去：", message, "回：", string(validData))
		}
	}
}
