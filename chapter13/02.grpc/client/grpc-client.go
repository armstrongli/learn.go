package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"learn.go/pkg/apis"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{Name: "tom"})
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	log.Println("注册成功", ret)

	log.Println("开始批量注册")
	regCli, err := c.RegisterPersons(context.TODO())
	if err != nil {
		log.Fatal("获取批量注册客户端失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	resp, err := regCli.CloseAndRecv()
	if err != nil {
		log.Fatal("无法接收结果：", err)
	}
	log.Println("批量注册结果：", resp.String())
}
