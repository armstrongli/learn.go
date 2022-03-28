package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	pi := PersonalInformation{
		Id:     0,
		Name:   "小强",
		Sex:    "男",
		Tall:   1.7,
		Weight: 66,
		Age:    35,
	}
	jsonData, _ := json.Marshal(pi)
	fmt.Println("=============")
	fmt.Println(string(jsonData))
	fmt.Println("=============")
	protoData, _ := proto.Marshal(&pi)
	fmt.Println(string(protoData))
	fmt.Println("=============编码后：")
	fmt.Println(base64.StdEncoding.EncodeToString(protoData))
	fmt.Println("=============")
}
