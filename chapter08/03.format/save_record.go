package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/proto"
	"learn.go/pkg/apis"
)

func NewRecord(filePath string) *record {
	return &record{
		filePath:         filePath,
		yamlFilePath:     filePath + ".yaml",
		protobufFilePath: filePath + ".proto.base64",
	}
}

type record struct {
	filePath         string
	yamlFilePath     string
	protobufFilePath string
}

func (r *record) savePersonalInformation(pi *apis.PersonalInformation) error {
	{
		data, err := json.Marshal(pi)
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.writeFileWithAppendJson(data); err != nil {
			log.Println("写入JSON时出错：", err)
			return err
		}
	}
	{
		data, err := yaml.Marshal(pi)
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.writeFileWithAppendYaml(data); err != nil {
			log.Println("写入YAML时出错：", err)
			return err
		}
	}
	{
		data, err := proto.Marshal(pi)
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.writeFileWithAppendProtobuf(data); err != nil {
			log.Println("写入PROTOBUF时出错：", err)
			return err
		}
	}
	return nil
}

func (r *record) writeFileWithAppendJson(data []byte) error {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func (r *record) writeFileWithAppendYaml(data []byte) error {
	file, err := os.OpenFile(r.yamlFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	newData := append([]byte("---\n"), data...)
	_, err = file.Write(newData)
	if err != nil {
		return err
	}
	return nil
}

func (r *record) writeFileWithAppendProtobuf(data []byte) error {
	file, err := os.OpenFile(r.protobufFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write([]byte(base64.StdEncoding.EncodeToString(data)))
	if err != nil {
		return err
	}
	return nil
}
