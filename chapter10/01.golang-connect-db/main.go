package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"learn.go/pkg/apis"
)

func main() {
	learnDB, err := getDbConnection()
	defer learnDB.Close()
	err = testDbConnection(err, learnDB)

	queryAllDataWithHack(err, learnDB)

	// queryAllData(err, learnDB)

	// insertData(learnDB) // todo handle error

	// queryAllData(err, learnDB)
}

func insertData(learnDB *sql.DB) error {
	_, err := learnDB.Exec(fmt.Sprintf("insert into personal_info(name,sex,tall,weight,age) values('%s','%s',%f,%f,%d)",
		"小贝",
		"女",
		1.75,
		58.0,
		18,
	))
	if err != nil {
		fmt.Printf("新增数据失败：%v\n", err)
		return err
	}
	return nil
}

func queryAllDataWithHack(err error, learnDB *sql.DB) {
	_sql := fmt.Sprintf(`select name,age from personal_info where name='%s' and sex = '%s'`, "小贝' -- ", "男")

	rows, err := learnDB.Query(_sql)
	// rows, err := learnDB.Query("select age,name from personal_info ") // 更换列的顺序，因为数据类型不匹配导致失败。
	// 警告：如果不巧，数据类型兼容，会引入更大的灾难。
	if err != nil {
		log.Fatal(err)
	}
	list := &apis.PersonalInformationList{}
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			// todo handle error
			log.Printf("扫描数据失败，跳过该行: %v", err)
			continue
		}
		fmt.Printf("name: %s\tage: %d\n", name, age)
		list.Items = append(list.Items, &apis.PersonalInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func queryAllData(err error, learnDB *sql.DB) {
	rows, err := learnDB.Query("select name,age from personal_info ")
	// rows, err := learnDB.Query("select age,name from personal_info ") // 更换列的顺序，因为数据类型不匹配导致失败。
	// 警告：如果不巧，数据类型兼容，会引入更大的灾难。
	if err != nil {
		log.Fatal(err)
	}
	list := &apis.PersonalInformationList{}
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			// todo handle error
			log.Printf("扫描数据失败，跳过该行: %v", err)
			continue
		}
		fmt.Printf("name: %s\tage: %d\n", name, age)
		list.Items = append(list.Items, &apis.PersonalInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func testDbConnection(err error, learnDB *sql.DB) error {
	if err = learnDB.Ping(); nil != err {
		log.Fatalf("DB 测试失败: %v", err)
	}
	fmt.Println("数据库连接成功，可以执行命令")
	return err
}

func getDbConnection() (*sql.DB, error) {
	learnDB, err := sql.Open("mysql", "root:learngo@tcp(127.0.0.1:3306)/learngo")
	if err != nil {
		log.Fatalf("链接数据库失败: %v", err)
	}
	return learnDB, err
}
