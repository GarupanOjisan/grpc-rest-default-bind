package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/garupanojisan/protoc-gen-restize/example/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	proto.RegisterHelloServer(server, &HelloServer{})
	proto.RegisterExampleServer(server, &ExampleServer{})

	fmt.Printf("start listening on :8080\n")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type HelloServer struct {
}

func (s *HelloServer) SayBye(ctx context.Context, request *proto.SayBye_Request) (*proto.SayBye_Response, error) {
	return &proto.SayBye_Response{}, nil
}

func (s *HelloServer) SayHello(ctx context.Context, req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	return &proto.SayHelloResponse{Message: req.GetMessage()}, nil
}

type ExampleServer struct {
}

func (s *ExampleServer) Get(ctx context.Context, req *proto.Get_Request) (*proto.Get_Response, error) {
	return &proto.Get_Response{
		Data: []byte("test"),
	}, nil
}

func (s *ExampleServer) Post(ctx context.Context, req *proto.Post_Request) (*proto.Post_Response, error) {
	return &proto.Post_Response{
		Ok: true,
	}, nil
}
