package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// etcdGetDemo()
	etcdWatchAndReactDemo()
}

func etcdGetDemo() {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Get(context.TODO(), "a")
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range resp.Kvs {
		log.Printf("key: %s, value: %s", kv.Key, kv.Value)
	}
}

func etcdWatchAndReactDemo() {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}

	dataCh := make(chan int)
	go func() {
		watcher := cli.Watch(context.TODO(), "a")
		for respData := range watcher {
			evs := respData.Events
			whetherBreak := false
			for _, ev := range evs {
				i, err := strconv.Atoi(string(ev.Kv.Value))
				if err != nil {
					fmt.Println("不是数字，结束")
					whetherBreak = true
					break
				}
				dataCh <- i
			}
			if whetherBreak {
				break
			}
		}
	}()

	go func() {
		for i := range dataCh {
			_, err := cli.Put(context.TODO(), "a", fmt.Sprintf("%d", i+1))
			if err != nil {
				fmt.Println("WARNING: 更新失败")
			}
		}
	}()

	time.Sleep(2 * time.Second)
}
