// auto generated file
package proto

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
)

type Gateway interface {
	Methods() []Method
}

type Method interface {
	SetConn(conn *grpc.ClientConn)
	Path() string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// HelloGateway is RESTFul gateway
type HelloGateway struct{}

func (g *HelloGateway) Methods() []Method {
	return []Method{ 
		&MethodHelloSayHello{}, 
		&MethodHelloSayBye{},
	}
}
// MethodHelloSayHello is a service of Hello
type MethodHelloSayHello struct {
	conn *grpc.ClientConn
}

func (m *MethodHelloSayHello) SetConn(conn *grpc.ClientConn) {
	m.conn = conn
}

func (m *MethodHelloSayHello) Path() string {
	return "/garupanojissan.grpc.example.hello/SayHello"
}

func (m *MethodHelloSayHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	req := &SayHelloRequest{}
	if err := jsonpb.Unmarshal(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client := NewHelloClient(m.conn)
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jm := jsonpb.Marshaler{}
	if err := jm.Marshal(w, resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
// MethodHelloSayBye is a service of Hello
type MethodHelloSayBye struct {
	conn *grpc.ClientConn
}

func (m *MethodHelloSayBye) SetConn(conn *grpc.ClientConn) {
	m.conn = conn
}

func (m *MethodHelloSayBye) Path() string {
	return "/garupanojissan.grpc.example.hello/SayBye"
}

func (m *MethodHelloSayBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	req := &SayBye_Request{}
	if err := jsonpb.Unmarshal(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client := NewHelloClient(m.conn)
	resp, err := client.SayBye(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jm := jsonpb.Marshaler{}
	if err := jm.Marshal(w, resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
