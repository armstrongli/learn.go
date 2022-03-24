package main

import (
	"context"
	"encoding/json"
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
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{
		Name:   "tom",
		Sex:    "男",
		Tall:   1.7,
		Weight: 70,
		Age:    30,
	})
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	log.Println("注册成功", ret)

	top, err := c.GetTop(context.TODO(), &apis.Null{})
	if err != nil {
		log.Fatal("获取榜单失败：", err)
	}
	log.Println("榜单：", top.String())

	log.Println("开始批量注册")
	regCli, err := c.RegisterPersons(context.TODO())
	if err := regCli.Send(&apis.PersonalInformation{
		Name:   "tom-3",
		Sex:    "男",
		Tall:   1.7,
		Weight: 70,
		Age:    30,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{
		Name:   "tom-1",
		Sex:    "男",
		Tall:   1.7,
		Weight: 70,
		Age:    30,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{
		Name:   "tom-2",
		Sex:    "男",
		Tall:   1.7,
		Weight: 70,
		Age:    30,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	regCli.CloseAndRecv()

	{
		top, err := c.GetTop(context.TODO(), &apis.Null{})
		if err != nil {
			log.Fatal("获取榜单失败：", err)
		}
		data, _ := json.Marshal(top)
		log.Println("榜单：", string(data))
	}

}
