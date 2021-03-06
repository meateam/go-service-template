// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package template

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/template.template/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
// All implementations must embed UnimplementedTemplateServer
// for forward compatibility
type TemplateServer interface {
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	mustEmbedUnimplementedTemplateServer()
}

// UnimplementedTemplateServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (UnimplementedTemplateServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedTemplateServer) mustEmbedUnimplementedTemplateServer() {}

// UnsafeTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServer will
// result in compilation errors.
type UnsafeTemplateServer interface {
	mustEmbedUnimplementedTemplateServer()
}

func RegisterTemplateServer(s grpc.ServiceRegistrar, srv TemplateServer) {
	s.RegisterService(&_Template_serviceDesc, srv)
}

func _Template_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.template/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Template_serviceDesc = grpc.ServiceDesc{
	ServiceName: "template.template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _Template_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/template.proto",
}
