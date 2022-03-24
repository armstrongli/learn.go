package main

import (
	"context"
	"io"
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
	w, err := c.WatchPersons(context.TODO(), &apis.Null{})
	if err != nil {
		log.Fatal("启动watcher失败：", err)
	}
	for {
		pi, err := w.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("服务器告知说完了")
				break
			}
			log.Fatal("接收异常：", err)
		}
		log.Println("收到新变动：", pi.String())
	}
}
