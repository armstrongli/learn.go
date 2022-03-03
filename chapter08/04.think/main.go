package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/proto"
	"learn.go/pkg/apis"
)

func main() {
	// 5万条记录
	counter := 50000
	persons := make([]*apis.PersonalInformation, 0, counter)

	pLister := &apis.PersonalInformationList{}
	start := time.Now()
	data, _ := ioutil.ReadFile("/Users/jianqli/Desktop/data.protobuf")
	proto.Unmarshal(data, pLister)
	finish := time.Now()
	fmt.Println("json unmarshal: ", finish.Sub(start))

	for i := 0; i < counter; i++ {
		persons = append(persons, &apis.PersonalInformation{
			Name:   "小强",
			Sex:    "男",
			Tall:   1.7,
			Weight: 71,
			Age:    35,
		})
	}
	// JSON、YAML、Protobuf分别序列化，记录序列化耗时
	// 保存文件，记录耗时
	{
		fmt.Println("序列化JSON")
		startTime := time.Now()
		data, err := json.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("/Users/jianqli/Desktop/data.json", data, 0777) // todo handle error
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishMarshalTime.Sub(startTime))
		fmt.Println("写文件耗时：", finishWriteFileTime.Sub(finishMarshalTime))
	}
	{
		fmt.Println("序列化YAML")
		startTime := time.Now()
		data, err := yaml.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("/Users/jianqli/Desktop/data.yaml", data, 0777) // todo handle error
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishMarshalTime.Sub(startTime))
		fmt.Println("写文件耗时：", finishWriteFileTime.Sub(finishMarshalTime))
	}
	{
		fmt.Println("序列化PROTOBUF")
		pLister := &apis.PersonalInformationList{
			Items: persons,
		}
		startTime := time.Now()
		data, err := proto.Marshal(pLister)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("/Users/jianqli/Desktop/data.protobuf", data, 0777) // todo handle error
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishMarshalTime.Sub(startTime))
		fmt.Println("写文件耗时：", finishWriteFileTime.Sub(finishMarshalTime))
	}
}
