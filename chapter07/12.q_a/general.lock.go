package main

import (
	"sync"
)

var counter *safeCount = &safeCount{}

type safeCount struct {
	totalNumber      int
	totalLetterCount int
	totalWordCount   int

	// .....
	sync.Mutex
}

// 线程安全的加数据
func (c *safeCount) AddNumber(totalNumber int,
	totalLetterCount int,
	totalWordCount int) {
	c.Lock()
	defer c.Unlock()
	c.totalNumber += totalNumber
	// ....
}

func countNumber() {
	counter.AddNumber(100, 5, 500)
}
