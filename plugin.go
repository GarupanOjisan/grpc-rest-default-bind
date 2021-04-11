package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	pluginpb "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type DefaultRestfulBind struct {
}

func NewDefaultRestfulBind() *DefaultRestfulBind {
	return &DefaultRestfulBind{}
}

func (d *DefaultRestfulBind) Run(req *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	files := make(map[string]*descriptor.FileDescriptorProto)
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}

	var resp pluginpb.CodeGeneratorResponse
	for _, fname := range req.FileToGenerate {
		f := files[fname]
		out := fname + ".dump"
		resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(out),
			Content: proto.String(proto.MarshalTextString(f)),
		})
	}
	return &resp
}
