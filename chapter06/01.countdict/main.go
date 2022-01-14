package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	for p := 0; p <= 5000; p++ {
		fmt.Print("正在统计第", p, "页，")
		time.Sleep(1 * time.Second)
		r, _ := rand.Int(rand.Reader, big.NewInt(800))
		fmt.Println("有", r.Int64(), "字")
		totalCount += r.Int64()
	}
	fmt.Println("总共有：", totalCount, "字")
}
