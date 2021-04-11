package main

import (
	"bytes"
	"strings"
	"text/template"

	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/golang/protobuf/proto"

	pluginpb "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type DefaultRestfulBind struct {
}

func NewDefaultRestfulBind() *DefaultRestfulBind {
	return &DefaultRestfulBind{}
}

func (d *DefaultRestfulBind) Run(req *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	gwTmpl := template.Must(template.New("gateway").Parse(gatewayTemplate))

	fileToServices := make(map[string][]*descriptorpb.ServiceDescriptorProto)
	fileToPackage := make(map[string]string)
	fileToGoPkg := make(map[string]string)
	for _, file := range req.ProtoFile {
		fileToGoPkg[file.GetName()] = file.GetOptions().GetGoPackage()
		fileToServices[file.GetName()] = file.GetService()
		fileToPackage[file.GetName()] = file.GetPackage()
	}

	var resp pluginpb.CodeGeneratorResponse
	for _, fname := range req.FileToGenerate {
		pkgName := fileToPackage[fname]
		goPkg := fileToGoPkg[fname]
		goPkgName := strings.Split(goPkg, "/")

		for _, svc := range fileToServices[fname] {
			ms := make([]map[string]string, len(svc.GetMethod()))
			for i, m := range svc.GetMethod() {
				in := m.GetInputType()
				inMsg := strings.TrimPrefix(in, "."+pkgName+".")
				out := m.GetOutputType()
				outMsg := strings.TrimPrefix(out, "."+pkgName+".")

				inMsg = strings.Replace(inMsg, ".", "_", -1)
				outMsg = strings.Replace(outMsg, ".", "_", -1)

				ms[i] = map[string]string{
					"Name":   m.GetName(),
					"Input":  inMsg,
					"Output": outMsg,
				}
			}

			// gateway
			gwOut := fname + ".default-gw.go"
			out := bytes.Buffer{}
			err := gwTmpl.Execute(&out, map[string]interface{}{
				"Service":   svc.GetName(),
				"Methods":   ms,
				"Package":   pkgName,
				"GoPackage": goPkgName[len(goPkgName)-1],
			})
			if err != nil {
				msg := err.Error()
				return &pluginpb.CodeGeneratorResponse{Error: &msg}
			}
			content := out.String()
			resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
				Name:    proto.String(gwOut),
				Content: &content,
			})
		}
	}

	return &resp
}

var gatewayTemplate = `// auto generated file
package {{ .GoPackage }}

{{- $service := .Service }}
{{- $package := .Package }}

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"

	"github.com/garupanojisan/protoc-gen-restize/runtime"
)

// HelloGateway is RESTFul gateway
type HelloGateway struct{}

func (g *{{ $service }}Gateway) Methods() []Method {
	return []Method{
	{{- range $i, $method := .Methods }} 
		&Method{{ $service }}{{ $method.Name }}{},
	{{- end }}
	}
}

{{- range $i, $method := .Methods }}
// Method{{ $service }}{{ $method.Name }} is a service of {{ $service }}
type Method{{ $service }}{{ $method.Name }} struct {
	conn *grpc.ClientConn
}

func (m *Method{{ $service }}{{ $method.Name }}) SetConn(conn *grpc.ClientConn) {
	m.conn = conn
}

func (m *Method{{ $service }}{{ $method.Name }}) Path() string {
	return "/{{ $package }}/{{ $method.Name }}"
}

func (m *Method{{ $service }}{{ $method.Name }}) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	req := &{{ $method.Input }}{}
	if err := jsonpb.Unmarshal(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client := New{{ $service }}Client(m.conn)
	resp, err := client.{{ $method.Name }}(ctx, req)
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
{{- end }}
`
