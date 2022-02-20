package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var totalCompare int = 0

func main() {
	// size := 1000
	arr := sampleData // generateRandomData(size)
	fmt.Println(arr)
	// 1. 501 在不在里边？
	// 2. 888 在不在？
	// 3. 900 ？
	// 4. 3 ?
	// ……

	startTime := time.Now()
	for i := 0; i < 1000000; i++ { // 4.838486846s，次数：31 6500 0000
		search(&arr, 501)
		search(&arr, 888)
		search(&arr, 900)
		search(&arr, 3)
	}
	finishTime := time.Now()

	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

func search(arrP *[]int64, targetNum int64) bool {
	for _, v := range *arrP {
		totalCompare++
		if v == targetNum {
			return true
		}
	}
	return false
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(3000))
		arr = append(arr, i.Int64())
	}
	return arr
}
