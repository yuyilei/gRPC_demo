package main

import (
	"golang.org/x/net/context"
	hello "github.com/gRPC_demo/greeter/helloworld"
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
)
const (
	port = "30001"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloworldServer(s,&server{})
	s.Serve(lis)
}
