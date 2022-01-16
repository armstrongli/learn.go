package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWebServerV2(t *testing.T) {
	v2 := &WebServerV2{
		config: &Config{Content: "a"},
	}
	go v2.ReloadWorker()

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				v2.Visit()
			}
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("总时间：", finish.Sub(start))
}
