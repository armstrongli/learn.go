package main

import (
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

func (r *rankServer) Register(context context2.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.Lock()
	defer r.Unlock()
	r.persons[information.Name] = information
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}
