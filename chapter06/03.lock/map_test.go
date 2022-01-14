package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 100; i++ {
		go func() {
			for {
				v := m[i]
				m[i] = v + 1
				fmt.Println("i=", m[i])
			}
		}()
	}
	time.Sleep(10 * time.Second)
}
