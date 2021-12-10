package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println(finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	arr2 := testPanic()

	arr3 := make([]string, 3, 4)
	copy(arr3, arr2)

	i := 3333
	j := &i
	fmt.Println(reflect.TypeOf(j))
}

func testPanic() []string {
	arr2 := make([]string, 0, 4)
	fmt.Println("len: ", len(arr2), "cap: ", cap(arr2))
	fmt.Println("default:", arr2[0])
	fmt.Println("default:", arr2[1])
	fmt.Println("default:", arr2[2])
	return arr2
}
