package main

import (
	"context"
	"log"

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
}
