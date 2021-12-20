package main

import (
	"fmt"
	"runtime/debug"

	calc "learn.go/chapter02/015.fatrate.refactor/calc"
)

func init() {
	fmt.Println("我是init函数--1")
}

func main() {
	for {
		mainFatRateBody()

		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func init() {
	fmt.Println("我是init函数--2")
}

func recoverMainBody() {
	if re := recover(); re != nil {
		fmt.Printf("warning: catch critical error: %v\n", re)
		debug.PrintStack()
	}
}

func mainFatRateBody() {
	defer recoverMainBody() // 成功捕获
	weight, tall, age, _, sex := getMaterialsFromInput()

	fatRate, err := calcFatRate(weight, tall, age, sex)
	if err != nil {
		fmt.Println("warning: 计算 体脂率 出错，不能再继续", err)
		return
	}
	if fatRate <= 0 {
		panic("fat rate is not allowed to be negative")
	}

	var checkHealthinessFunc func(age int, fatRate float64)
	if sex == "男" {
		checkHealthinessFunc = getHealthinessSuggestionsForMale
	} else {
		checkHealthinessFunc = getHealthinessSuggestionsForFemale
	}
	getHealthinessSuggestions(age, fatRate, checkHealthinessFunc)
}

func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate)
}

func generateCheckHealthinessForMale() func(age int, fatRate float64) {
	// 定制功能

	return func(age int, fatRate float64) {

	}
}

func getHealthinessSuggestionsForMale(age int, fatRate float64) {
	//  编写男性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("目前是：肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		// todo
	} else if age >= 60 {
		// todo
	} else {
		fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
	}
}

func getHealthinessSuggestionsForFemale(age int, fatRate float64) {
	//  编写男性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("目前是：肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		// todo
	} else if age >= 60 {
		// todo
	} else {
		fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
	}
}

func calcFatRate(weight float64, tall float64, age int, sex string) (fatRate float64, err error) {
	// 计算体脂率
	bmi, err := calc.CalcBMI(tall, weight)
	if err != nil {
		return 0, err
	}

	fatRate = calc.CalcFatRate(bmi, age, sex)
	fmt.Println("体脂率是：", fatRate)
	return
}

func getMaterialsFromInput() (float64, float64, int, int, string) {
	// 录入各项
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sexWeight int
	sex := "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return weight, tall, age, sexWeight, sex
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n）？")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}
