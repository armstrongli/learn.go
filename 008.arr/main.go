package main

import (
	"fmt"
)

func main() {
	fmt.Println("1st version")
	// 难以长期维护
	xqInfo := [3]string{"小强", "男", "在职"}
	xlInfo := [3]string{"小李", "男", "在职"}
	xsInfo := [3]string{"小苏", "女", "在职"}

	fmt.Println(xqInfo)
	fmt.Println(xlInfo)
	fmt.Println(xsInfo)
	// ....

	fmt.Println("2nd version")
	// 难点：数组长度管理
	newPersonInfos := [3][3]string{
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小苏", "女", "在职"},
	}
	for _, val := range newPersonInfos {
		fmt.Println(val)
	}

	fmt.Println("3rd version")
	// ... 支持动态添加
	newPersonInfos2 := [][3]string{ // [...][3]string{
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小苏", "女", "在职"},
		[3]string{"小苏2", "女", "在职"},
		[3]string{"小苏3", "女", "在职"},
	}
	for _, val := range newPersonInfos2 {
		fmt.Println(val)
	}

	newPersonInfos2 = append(newPersonInfos2, [3]string{"大牛", "不明", "未知"}) // 在数组中，无法append。换成slice后可以更改

	fmt.Println("用降维方式输出：")
	for d1, d1val := range newPersonInfos2 {
		for d2, d2val := range d1val {
			fmt.Println(d1, d1val, d2, "val:", d2val)
		}
	}

	// d1, d2, d3, d4,.....dn v1
	// d1, d2, d3, d4,.....dn v2
	// ....
}
