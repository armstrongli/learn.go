package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestDecode(t *testing.T) {
	jsonData := `{"name":"小强","sex":"男","tall":1.7,"weight":66,"age":35}`
	protoDataBase64 := `EgblsI/lvLoaA+eUtyWamdk/LQAAhEIwIw==`

	pi1 := &PersonalInformation{}
	json.Unmarshal([]byte(jsonData), pi1)
	fmt.Println("解析JSON：")
	fmt.Printf("%+v\n", *pi1)

	pi2 := &PersonalInformation{}
	protData, _ := base64.StdEncoding.DecodeString(protoDataBase64)
	proto.Unmarshal(protData, pi2)
	fmt.Println("解析protobuf：")
	fmt.Printf("%+v\n", *pi2)
}
