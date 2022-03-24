package main

import (
	"io"
	"log"
	"sync"

	context2 "golang.org/x/net/context"
	"learn.go/pkg/apis"
)

var _ apis.RankServiceServer = &rankServer{}

type rankServer struct {
	sync.Mutex
	persons map[string]*apis.PersonalInformation
}

func (r *rankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for {
		pi, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Printf("WARNING: 获取人员注册时失败：%v\n", err)
			return err
		}
		pis.Items = append(pis.Items, pi)
		r.Lock()
		r.persons[pi.Name] = pi
		r.Unlock()
	}
	log.Println("连续得到注册清单：", pis.String())
	return server.SendAndClose(pis)
}

func (r *rankServer) Register(context context2.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.Lock()
	defer r.Unlock()
	r.persons[information.Name] = information
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}
