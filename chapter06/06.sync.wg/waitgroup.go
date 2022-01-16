package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	Name string
}

func (r Runner) Run(startPointWg, wg *sync.WaitGroup) {
	defer wg.Done()
	startPointWg.Wait()

	start := time.Now()
	fmt.Println(r.Name, "开始跑@", start)
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	finish := time.Now()
	fmt.Println(r.Name, "跑到终点，用时：", finish.Sub(start))
}

func main() {
	runnerCount := 10
	runners := []Runner{}

	wg := sync.WaitGroup{}
	wg.Add(runnerCount)

	startPointWg := sync.WaitGroup{}
	startPointWg.Add(1)

	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d", i),
		})
	}

	for _, runnerItem := range runners {
		go runnerItem.Run(&startPointWg, &wg)
	}

	fmt.Println("各就位")
	time.Sleep(1 * time.Second)
	fmt.Println("预备：跑")

	startPointWg.Done()

	wg.Wait()
	fmt.Println("赛跑结束")
}
