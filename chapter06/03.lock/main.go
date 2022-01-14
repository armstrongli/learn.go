package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// countDict()
		// countDictGoroutineSafe()
		// countDictForgetUnlock()
		countDictLockPrice()
	}
}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			// fmt.Print("正在统计第", p, "页，")
			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictGoroutineSafe() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			// fmt.Print("正在统计第", p, "页，")
			defer totalCountLock.Unlock()
			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
			// totalCountLock.Unlock() // 也可以用defer来解锁
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictLockPrice() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func(pageNum int) {
			defer wg.Done()
			// fmt.Print("正在统计第", p, "页，")
			totalCountLock.Lock()
			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
			if pageNum == 3 {
				time.Sleep(3 * time.Second)
			}
			totalCountLock.Unlock()

			// //////
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictForgetUnlock() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func() {
			defer wg.Done()
			// fmt.Print("正在统计第", p, "页，")
			totalCountLock.Lock()
			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
			// totalCountLock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}
