package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type frClient struct {
	handRing Interface
}

func (f frClient) register() {
	pi, _ := f.handRing.ReadPersonalInformation()
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/register", "application/json", r)
	if err != nil {
		log.Println("WARNING: register fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回：", string(data))
	}
}
func (f frClient) update() {
	pi, _ := f.handRing.ReadPersonalInformation()
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/personalinfo", "application/json", r)
	if err != nil {
		log.Println("WARNING: register fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回：", string(data))
	}
}

func main() {
	frCli := &frClient{handRing: &fakeInterface{
		name:       fmt.Sprintf("小强%d", time.Now().UnixNano()),
		sex:        "男",
		baseWeight: 71.0,
		baseTall:   1.70,
		baseAge:    35,
	}}

	frCli.register()
	for {
		frCli.update()
		time.Sleep(1 * time.Second)
	}
}
