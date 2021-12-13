package main

import (
	"fmt"
	"time"
)

var tall float64
var weight float64

func main() {
	panicAndRecover()

	deferGuess()
	fmt.Println("sleep somewhile")
	time.Sleep(10 * time.Second)

	openFile()
	fmt.Println("sleep somewhile")
	time.Sleep(10 * time.Second)

	closureMain()
	fmt.Println("sleep somewhile")
	time.Sleep(10 * time.Second)

	// close call
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))

	showUsedTimes()

	guess(1, 100)
	fmt.Println("done guess, sleep somewhile")
	time.Sleep(10 * time.Second)

	fmt.Println(fib(1))

	fmt.Println("done calc, sleep somewhile")
	time.Sleep(10 * time.Second)

	// sampleSubdomain2()
	fmt.Print("全局变量赋值前：")
	calcAdd() // 0
	tall, weight = 1.70, 70.0
	fmt.Print("全局变量赋值后：")
	calcAdd() // 71.7

	// 重新定义重名的局部变量
	tall, weight := 100.00, 70.0
	calcAdd() // ?? >> 71.7
	tall, weight = 200, 70.0
	calcAdd() // ?? >> 71.7

	calculatorAdd := func(a, b float64) float64 {
		return a + b
	}
	result := calculatorAdd(1, 3)
	fmt.Println(result)

	{
		// fmt.Scanln....
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
		// suggestions
	}

	{
		// fmt.Scanln....
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
		// suggestions
	}
	// fmt.Println(personTall) // personTall 的有效范围为 {} 内部。外部无效

	fmt.Println(tall, weight)
}

func funcDef(nums ...int) (addResult int) {
	for _, item := range nums {
		addResult += item
	}
	return
}

func funcDef1(nums ...int) int {
	sum := 0
	for _, item := range nums {
		sum += item
	}
	return sum
}

func calcAdd() float64 {
	fmt.Println(tall + weight)
	return tall + weight
}

func sampleSubdomain() {
	name := "小强"              // 声明变量 name，值是"小强"
	fmt.Println("名字是：", name) // 小强
	{
		fmt.Println("名字是：", name) // 小强
		name = "Kr"               // 重新赋值，给了 name。name的值是Kr
		fmt.Println("名字是：", name)
	}
	fmt.Println(">>>>名字是：", name) // Kr ?? 小强?? ==》Kr
}

func sampleSubdomain2() {
	name := "小强"              // 声明变量 name，值是"小强"
	fmt.Println("名字是：", name) // 小强
	{
		name = "小强-Update"
		fmt.Println("名字是：", name) // 小强
		name := "Kr"
		fmt.Println("名字是：", name) // Kr
	}
	fmt.Println(">>>>名字是：", name) // Kr ?? 小强?? ==》

	if name == "小强" {
		a := 3
		fmt.Println(a)
	} else {
		a := 4
		fmt.Println(a)
	}
}

func sampleSubdomainIf() {
	if v := calcAdd(); v == 0 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
	// fmt.Println(v) // 无效 。v的有效范围为 if block
}

func sampleSubdomainFor() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang, ", i)
	}
	// fmt.Println(i) // i 的有效范围为for block
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
