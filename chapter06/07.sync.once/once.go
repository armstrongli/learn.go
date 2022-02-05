package main

import (
	"fmt"
	"sync"
)

type rank struct {
	standard []string
}

var globalRank = &rank{}
var once sync.Once = sync.Once{}

func initGlobalRankStandard(standard []string) {
	once.Do(func() {
		globalRank.standard = standard
	})
}

var facStore = &dbFactoryStore{}

type dbFactoryStore struct {
	store map[string]DBFactory
}

type Conn struct{}

type DBFactory interface {
	GetConnection() *Conn
}

func initMySqlFac(connStr string) DBFactory {
	return &MySqlDBFactory{}
}

type MySqlDBFactory struct {
	once sync.Once
}

func (MySqlDBFactory) GetConnection() *Conn {
	once.Do(func() {
		initMySqlFac("")
	})
	// todo
	return nil
}

var counter int = 0
var counterOnce sync.Once

type School struct {
	classroomLocation map[string]string
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("第x次：", i)
		counterOnce.Do(func() {
			fmt.Println("初始化")
			counter++
		})
	}
	fmt.Println("最终结果：", counter)
	// standard := []string{"asia"}
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		initGlobalRankStandard(standard)
	// 	}()
	// }
	// connStr := "xxxxxxxx"
}
