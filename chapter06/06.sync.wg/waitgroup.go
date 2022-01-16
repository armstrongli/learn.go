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

func (r Runner) Run(wg *sync.WaitGroup) {
	defer wg.Done()
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

	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d", i),
		})
	}

	for _, runnerItem := range runners {
		go runnerItem.Run(&wg)
	}

	wg.Wait()
	fmt.Println("赛跑结束")
}
