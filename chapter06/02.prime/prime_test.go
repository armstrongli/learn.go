package _2_prime

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRunPrime(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	for num := 2; num <= 200000; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
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

func TestRunPrime2(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	go func() {
		fmt.Println("第一个worker开始计算", time.Now())
		result = append(result, collectPrime(2, 100000)...)
		fmt.Println("第一个worker完成计算", time.Now())
	}()
	go func() {
		fmt.Println("第二个worker开始计算", time.Now())
		result = append(result, collectPrime(100001, 200000)...)
		fmt.Println("第二个worker完成计算", time.Now())
	}()
	time.Sleep(15 * time.Second)
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func TestRunPrime3(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("第一个worker开始计算", time.Now())
		result = append(result, collectPrime(2, 100000)...)
		fmt.Println("第一个worker完成计算", time.Now())
	}()
	go func() {
		defer wg.Done()
		fmt.Println("第二个worker开始计算", time.Now())
		result = append(result, collectPrime(100001, 200000)...)
		fmt.Println("第二个worker完成计算", time.Now())
	}()
	wg.Wait()
	finishTime := time.Now()
	fmt.Println("finishTime: ", finishTime)
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func collectPrime(start int, end int) (result []int) {
	for num := start; num <= end; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	return
}

func TestHelloGoroutine(t *testing.T) {
	go fmt.Println("hello, goroutine")
}

func TestHelloGoroutine2(t *testing.T) {
	go fmt.Println("hello, goroutine")
	time.Sleep(1 * time.Second)
}

func TestForLoop(t *testing.T) {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 100; i < 120; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func TestForeverGoroutine(t *testing.T) {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			go func() {
				fmt.Println("启动新的goroutine@", time.Now())
				time.Sleep(1 * time.Hour)
			}()
		}
	}()
	for {
		fmt.Println(runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
