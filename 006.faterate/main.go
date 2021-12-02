package main

import (
	"fmt"
)

func main() {
	var weight float64 = 65.0
	var tall float64 = 1.70
	var bmi float64 = weight / (tall * tall)
	var age int = 35
	var sexWeight int
	var sex string = "男"
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是：", fatRate)

	if sex == "男" {
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
	} else {
		// todo 编写女性的体脂率与体脂状态表
	}
}
