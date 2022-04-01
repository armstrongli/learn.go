package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/docker/docker/client"
	clientv3 "go.etcd.io/etcd/client/v3"
	"learn.go/chapter15/01.model-driven/models"
	"learn.go/pkg/dockertool"
)

func main() {
	backendName := "/rankservice/backend"
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		watcher := cli.Watch(ctx, backendName)
		for respData := range watcher {
			evs := respData.Events
			for _, ev := range evs {
				rawBackend := ev.Kv.Value

				backend := &models.RankServiceBackend{}
				json.Unmarshal(rawBackend, backend) // todo handle error

				ids, _ := dockertool.List(ctx, dockerCli, map[string]string{"backend": backendName})
				backend.Status.RunningCount = len(ids)

				if backend.Expected.Count != backend.Status.RunningCount {
					fmt.Println("开始创建实例")
					// todo 调用docker创建新的容器实例 或 删除一些，并更新status
					ip, err := dockertool.Run(ctx, dockerCli, map[string]string{"backend": backendName}, "nginx:stable-alpine", nil)
					if err != nil {
						// todo handle error
					} else {
						backend.Status.RunningCount++
						backend.Status.InstanceIPs = append(backend.Status.InstanceIPs, ip)
						rawData, _ := json.Marshal(backend)
						cli.Put(ctx, backendName, string(rawData)) // todo handle error
					}
				} else {
					fmt.Println("已经满足预期：", backend.Expected.Count)
				}
			}
		}
	}()
	<-ctx.Done()
}
