package main

import (
	"context"
	"fmt"

	"protobuf_grpc_study/proto/pb/proto_demo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const Address = "127.0.0.1:1234"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	c := proto_demo.NewHelloClient(conn)
	req := &proto_demo.HelloRequest{Name: "grpc"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(res.Message)

	req2 := &proto_demo.HiRequest{
		Name:   "grpc",
		Grade:  3,
		Age:    18,
		Status: 1,
		School: "dongyou",
	}
	res2, err := c.SayHi(context.Background(), req2)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(res2.Message)
}
