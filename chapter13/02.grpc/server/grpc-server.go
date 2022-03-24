package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"learn.go/pkg/apis"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	startGRPCServer(ctx)
}

func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	apis.RegisterRankServiceServer(s, &rankServer{
		persons:  map[string]*apis.PersonalInformation{},
		personCh: make(chan *apis.PersonalInformation),
	})
	go func() {
		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
