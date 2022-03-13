package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/pkg/apis"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	rank := NewFatRateRank()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("wanring: 建立连接失败：", err)
			continue
		}
		fmt.Println(conn)

		// talk(conn)
		go talk(conn, rank)
	}
}

func talk(conn net.Conn, rank *FatRateRank) {
	defer fmt.Println("结束链接：", conn)
	defer conn.Close()
	for {
		finalReceivedMessage := make([]byte, 0, 1024)

		encounterError := false
		for {
			buf := make([]byte, 1024)
			valid, err := conn.Read(buf)
			if err != nil {
				log.Println("WARNING: 读取数据时出问题：", err)
				log.Println("重新读取", err)
				encounterError = true
				time.Sleep(1 * time.Second)
				break
			}
			if valid == 0 {
				break
			}
			validContent := buf[:valid]
			finalReceivedMessage = append(finalReceivedMessage, validContent...)
			if valid < len(buf) {
				break
			}
		}
		if encounterError {
			conn.Write([]byte(`服务器获取数据失败，请重新输入`)) // handle error
			continue
		}

		pi := &apis.PersonalInformation{}
		if err := json.Unmarshal(finalReceivedMessage, pi); err != nil {
			conn.Write([]byte(`输入数据无法解析，请重新输入`)) // handle error
			continue
		}

		bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
		if err != nil {
			conn.Write([]byte(`无法计算您的BMI，请重新输入`)) // handle error
			continue
		}
		fr := gobmi.CalcFatRate(bmi, int(pi.Age), pi.Sex)

		rank.inputRecord(pi.Name, fr)
		rankId, _ := rank.getRank(pi.Name)

		conn.Write([]byte(fmt.Sprintf("姓名：%s, BMI：%v，体脂率：%v，排名：%d", pi.Name, bmi, fr, rankId))) // handle error
		break
	}
}
