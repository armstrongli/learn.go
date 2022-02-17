package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(arr []int) []int {
	left, right := (arr)[:len(arr)/2], (arr)[len(arr)/2:]
	if len(arr) <= 2 {
		return mergeArr(left, right)
	} else {
		return mergeArr(mergeSort(left), mergeSort(right))
	}
}

func mergeArr(left, right []int) []int {
	out := []int{}
	leftI, rightI := 0, 0
	for {
		if leftI == len(left) || rightI == len(right) {
			break
		}

		if left[leftI] < right[rightI] {
			out = append(out, left[leftI])
			leftI++
			continue
		} else {
			out = append(out, right[rightI])
			rightI++
			continue
		}
	}
	for ; leftI < len(left); leftI++ {
		out = append(out, left[leftI])
	}
	for ; rightI < len(right); rightI++ {
		out = append(out, right[rightI])
	}
	return out
}

func main() {
	arrSize := 10000
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	// fmt.Println(arr)

	start := time.Now()
	mergeSort(arr)
	finish := time.Now()
	fmt.Println(finish.Sub(start))
	// fmt.Println(sorted)
}
