package main

import (
	"fmt"
	"sync"
)

func main() {
	iarr := []int{1, 2, 3, 4, 5, 6, 7}
	wg := sync.WaitGroup{}
	wg.Add(len(iarr))
	for _, item := range iarr {
		go func(newItem int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(newItem)
			}
		}(item)
	}
	wg.Wait()
}
