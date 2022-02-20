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
	// arr := sampleData // generateRandomData(size)
	// fmt.Println(arr)
	// 1. 501 在不在里边？
	// 2. 888 在不在？
	// 3. 900 ？
	// 4. 3 ?
	// ……

	// 复制一份新的array，避免碰到原始数组
	newArr := append([]int64{}, sampleData...)

	startTime := time.Now()
	// todo 排序
	quickSort(&newArr, 0, len(newArr)-1)

	// fmt.Println(search(&newArr, 501)) // 证明结果无误用
	// fmt.Println(search(&newArr, 888)) // 证明结果无误用
	// fmt.Println(search(&newArr, 900)) // 证明结果无误用
	// fmt.Println(search(&newArr, 3)) // 证明结果无误用

	for i := 0; i < 1000000; i++ { // 时间：82.424268ms，次数：3700 0000
		search(&newArr, 501)
		search(&newArr, 888)
		search(&newArr, 900)
		search(&newArr, 3)
	}
	finishTime := time.Now()

	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

func search(arrP *[]int64, targetNum int64) bool {
	return searchHarf(arrP, 0, len(*arrP)-1, targetNum)
}

func searchHarf(arrP *[]int64, left, right int, targetNum int64) bool {
	middleIndex := (left + right) / 2
	data := (*arrP)[middleIndex]

	totalCompare++ // 增加计数器数值

	if data < targetNum {
		if left == right {
			return false
		}
		return searchHarf(arrP, middleIndex+1, right, targetNum)
	} else if data > targetNum {
		if left == right {
			return false
		}
		return searchHarf(arrP, left, middleIndex-1, targetNum)
	} else {
		return true
	}
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(3000))
		arr = append(arr, i.Int64())
	}
	return arr
}
