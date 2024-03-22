package main

import (
	"context"
	"fmt"
	"net"

	"protobuf_grpc_study/proto/pb/proto_demo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const Address = "127.0.0.1:1234"

type HelloService struct{}

func (h *HelloService) SayHello(ctx context.Context, in *proto_demo.HelloRequest) (*proto_demo.HelloResponse, error) {
	resp := new(proto_demo.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func (h *HelloService) SayHi(ctx context.Context, in *proto_demo.HiRequest) (*proto_demo.HiResponse, error) {
	resp := new(proto_demo.HiResponse)
	resp.Message = fmt.Sprintf("Hi %s grade:%d school:%s age:%d status:%d", in.Name, in.Grade, in.School, in.Age, in.Status)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatal(err)
	}
	s := grpc.NewServer()
	proto_demo.RegisterHelloServer(s, &HelloService{})
	fmt.Println("Listen on ", Address)
	grpclog.Infof("Listen on " + Address)
	s.Serve(listen)
}
