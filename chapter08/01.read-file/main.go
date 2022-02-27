package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "/Users/jianqli/Desktop/小强"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	b := make([]byte, 10)
	var n int
	for i := 0; i < 2; i++ {
		n, err = file.Read(b)
		if err != nil {
			fmt.Println("无法读取文件：", err)
			os.Exit(2)
		}
		// fmt.Println("读出的内容：", b) // 如果不转换string类型，输出的是一长串数字
		// fmt.Println("读出的内容：", string(b))
		fmt.Println("n的大小：", n)
		fmt.Println("读出的内容：", string(b[:n])) // 一定要记得给后续程序使用时，截取到实际读取到的数据，而不是全部。否则后续程序会把无效读取也作为正常数据处理
		// file.Seek(0, io.SeekStart) // seekstart 表示从文件头开始数数，转移游标
		file.Seek(3, io.SeekCurrent) // SeekCurrent 表示从目前游标的位置开始，转移相对位置
		// file.Seek(0, io.SeekCurrent) // 改行的返回值的ret表示当前游标的位置
		// todo handle error
	}
}
