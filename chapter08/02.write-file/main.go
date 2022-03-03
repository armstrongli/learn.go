package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/Users/jianqli/Desktop/小强.new"
	// writeFile(filePath)
	writeFileWithAppend(filePath)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	// b := make([]byte, 10)
	// var n int
	_, err = file.Write([]byte("this is first line---"))
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	// handle error
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	// handle error
	fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	// file.Sync() // 告诉操作系统直接把内容都写到磁盘上
	// handle error
	fmt.Println(err)
}

func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	// b := make([]byte, 10)
	// var n int
	_, err = file.Write([]byte("this is first line---"))
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	// handle error
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	// handle error
	fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	// file.Sync() // 告诉操作系统直接把内容都写到磁盘上
	// handle error
	fmt.Println(err)
}
