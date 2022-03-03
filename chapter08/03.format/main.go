package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/pkg/apis"
)

func main() {
	input := &inputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	records := NewRecord("/Users/jianqli/Desktop/records.txt")

	for {
		pi := input.GetInput()
		if err := records.savePersonalInformation(pi); err != nil {
			log.Fatal("保存失败：", err)
		}
		fr, err := calc.FatRate(pi)
		if err != nil {
			log.Fatal("计算体脂率失败：", err)
		}
		rk.inputRecord(pi.Name, fr)

		rankResult, _ := rk.getRank(pi.Name)
		fmt.Println("排名结果：", rankResult)
	}
}

func caseStudy() {
	filePath := "/Users/jianqli/Desktop/小强.self.information.json"

	personalInformation := apis.PersonalInformation{
		// Name:   `"小"强""`,
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%+v\n", personalInformation)

	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal 的结果是(原生)：", data)
	fmt.Println("marshal 的结果是（string）：", string(data))

	writeFile(filePath, data)

	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取出来的内容是：", string(data))

	personalInformation := apis.PersonalInformation{}
	json.Unmarshal(data, &personalInformation) // todo handle error

	fmt.Println("开始计算体脂信息：", personalInformation)
	bmi, _ := gobmi.BMI(float64(personalInformation.Weight), float64(personalInformation.Tall)) // todo handle error
	fmt.Printf("%s 的 BMI是：%v\n", personalInformation.Name, bmi)
	fatRate := gobmi.CalcFatRate(bmi, int(personalInformation.Age), personalInformation.Sex)
	fmt.Printf("%s 的 体脂率是：%v\n", personalInformation.Name, fatRate)
}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	// b := make([]byte, 10)
	// var n int
	_, err = file.Write(data)
	fmt.Println(err)
}
