package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:learngo@tcp(127.0.0.1:3306)/learngo"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func createNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Tall:   1.70,
		Name:   "小强20220306",
		Sex:    "男",
		Weight: 71,
		Age:    35,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建人***时失败：", err)
		return err
	}
	fmt.Println("创建人***成功")
	return nil
}

func ormSelect(conn *gorm.DB) {
	output := make([]*PersonalInformation, 0)
	resp := conn.Where(&PersonalInformation{Name: "小强"}).Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}

	data, _ := json.Marshal(output)
	fmt.Printf("查询结果：%+v\n", string(data))
}

func ormSelectWithInaccurateQuery(conn *gorm.DB) {
	output := make([]*PersonalInformation, 0)
	resp := conn.Where("tall > ?", 1.7).Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}

	data, _ := json.Marshal(output)
	fmt.Printf("查询结果：%+v\n", string(data))
}

func ormSelectWithInaccurateQueryHack(conn *gorm.DB) {
	output := make([]*PersonalInformation, 0)
	resp := conn.Where("name = ? and sex = ?", "小贝' -- ", "男").Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}

	data, _ := json.Marshal(output)
	fmt.Printf("查询结果：%+v\n", string(data))
}

func updateExistingPerson(conn *gorm.DB) error {
	resp := conn.Updates(&PersonalInformation{
		Id:     23,
		Tall:   1.70,
		Name:   "小强20220306",
		Sex:    "男",
		Weight: 71,
		Age:    1000,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新人***时失败：", err)
		return err
	}
	fmt.Println("更新人***成功")
	return nil
}

func updateExistingPersonSelectedFields(conn *gorm.DB) error {
	p := &PersonalInformation{
		Id:     23,
		Tall:   1.80,
		Name:   "小强",
		Sex:    "男",
		Weight: 91,
		Age:    18,
	}
	resp := conn.Model(p).Select("name", "tall").Updates(p)
	if err := resp.Error; err != nil {
		fmt.Println("更新人***时失败：", err)
		return err
	}
	fmt.Println("更新人***成功")
	return nil
}

func deletePerson(conn *gorm.DB) {
	p := &PersonalInformation{
		Id: 23,
	}
	resp := conn.Delete(p)
	if err := resp.Error; err != nil {
		fmt.Println("删除***时失败：", err)
		return
	}
	fmt.Println("删除***成功")
}

func main() {
	conn := connectDb()
	fmt.Println(conn)
	// if err := createNewPerson(conn); err != nil {
	// 	// todo handle error
	// }
	// ormSelect(conn)

	// ormSelectWithInaccurateQuery(conn)
	// ormSelectWithInaccurateQueryHack(conn)
	// updateExistingPerson(conn)
	// updateExistingPersonSelectedFields(conn)
	deletePerson(conn)
}
