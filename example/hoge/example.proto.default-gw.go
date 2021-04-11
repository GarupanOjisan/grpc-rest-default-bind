// auto generated file
package hoge

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"

	"github.com/garupanojisan/protoc-gen-restize/runtime"
)

// HelloGateway is RESTFul gateway
type ExampleGateway struct{}

func (g *ExampleGateway) Methods() []runtime.Method {
	return []runtime.Method{ 
		&MethodExampleGet{}, 
		&MethodExamplePost{},
	}
}
// MethodExampleGet is a service of Example
type MethodExampleGet struct {
	conn *grpc.ClientConn
}

func (m *MethodExampleGet) SetConn(conn *grpc.ClientConn) {
	m.conn = conn
}

func (m *MethodExampleGet) Path() string {
	return "/garupanojissan.grpc.example.example/Get"
}

func (m *MethodExampleGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	req := &Get_Request{}
	if err := jsonpb.Unmarshal(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client := NewExampleClient(m.conn)
	resp, err := client.Get(ctx, req)
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
// MethodExamplePost is a service of Example
type MethodExamplePost struct {
	conn *grpc.ClientConn
}

func (m *MethodExamplePost) SetConn(conn *grpc.ClientConn) {
	m.conn = conn
}

func (m *MethodExamplePost) Path() string {
	return "/garupanojissan.grpc.example.example/Post"
}

func (m *MethodExamplePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	req := &Post_Request{}
	if err := jsonpb.Unmarshal(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client := NewExampleClient(m.conn)
	resp, err := client.Post(ctx, req)
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
