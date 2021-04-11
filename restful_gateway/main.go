package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/garupanojisan/protoc-gen-restize/example/proto"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal(err)
	}

	gs := []proto.Gateway{
		&proto.HelloGateway{},
	}
	for _, g := range gs {
		for _, m := range g.Methods() {
			m.SetConn(conn)
			http.Handle(m.Path(), m)
		}
	}

	fmt.Println("start gateway on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
