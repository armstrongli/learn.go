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

	m := sync.Map{}

	start := time.Now()
	concurrency := 100
	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			changed := 0
			pre := v2.config.Content
			for j := 0; j < 1000000000; j++ {
				if v2.Visit() != pre {
					changed++
					pre = v2.config.Content
				}
			}
			m.Store(time.Now().UnixNano(), changed)
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("总时间：", finish.Sub(start))
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
