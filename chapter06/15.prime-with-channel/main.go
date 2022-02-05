package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	maxNum := 200000
	result := make(chan int, maxNum/4)
	wg := sync.WaitGroup{}

	workerNumber := 16
	wg.Add(workerNumber)
	baseNumCh := make(chan int, 10)
	for i := 0; i < workerNumber; i++ {
		go func() {
			defer wg.Done()
			for oNum := range baseNumCh {
				if isPrime(oNum) {
					result <- oNum
				}
			}
		}()
	}

	for num := 2; num <= maxNum; num++ {
		baseNumCh <- num
	}
	close(baseNumCh)
	wg.Wait()

	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
}
