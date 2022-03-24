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
	persons  map[string]*apis.PersonalInformation
	personCh chan *apis.PersonalInformation
}

func (r *rankServer) regPerson(pi *apis.PersonalInformation) {
	r.Lock()
	defer r.Unlock()
	r.persons[pi.Name] = pi
	r.personCh <- pi
}

func (r *rankServer) WatchPersons(null *apis.Null, server apis.RankService_WatchPersonsServer) error {
	for pi := range r.personCh {
		if err := server.Send(pi); err != nil {
			log.Println("发送失败，结束:", err)
			return err
		}
	}
	return nil
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
		r.regPerson(pi)
	}
	log.Println("连续得到注册清单：", pis.String())
	return server.SendAndClose(pis)
}

func (r *rankServer) Register(context context2.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.regPerson(information)
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}
