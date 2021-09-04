// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_template_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OcpTemplateApiServiceClient is the client API for OcpTemplateApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpTemplateApiServiceClient interface {
	// CreateTemplateV1 - Create an template
	CreateTemplateV1(ctx context.Context, in *CreateTemplateV1Request, opts ...grpc.CallOption) (*CreateTemplateV1Response, error)
}

type ocpTemplateApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpTemplateApiServiceClient(cc grpc.ClientConnInterface) OcpTemplateApiServiceClient {
	return &ocpTemplateApiServiceClient{cc}
}

func (c *ocpTemplateApiServiceClient) CreateTemplateV1(ctx context.Context, in *CreateTemplateV1Request, opts ...grpc.CallOption) (*CreateTemplateV1Response, error) {
	out := new(CreateTemplateV1Response)
	err := c.cc.Invoke(ctx, "/ozoncp.ocp_template_api.v1.OcpTemplateApiService/CreateTemplateV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpTemplateApiServiceServer is the server API for OcpTemplateApiService service.
// All implementations must embed UnimplementedOcpTemplateApiServiceServer
// for forward compatibility
type OcpTemplateApiServiceServer interface {
	// CreateTemplateV1 - Create an template
	CreateTemplateV1(context.Context, *CreateTemplateV1Request) (*CreateTemplateV1Response, error)
	mustEmbedUnimplementedOcpTemplateApiServiceServer()
}

// UnimplementedOcpTemplateApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOcpTemplateApiServiceServer struct {
}

func (UnimplementedOcpTemplateApiServiceServer) CreateTemplateV1(context.Context, *CreateTemplateV1Request) (*CreateTemplateV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplateV1 not implemented")
}
func (UnimplementedOcpTemplateApiServiceServer) mustEmbedUnimplementedOcpTemplateApiServiceServer() {}

// UnsafeOcpTemplateApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpTemplateApiServiceServer will
// result in compilation errors.
type UnsafeOcpTemplateApiServiceServer interface {
	mustEmbedUnimplementedOcpTemplateApiServiceServer()
}

func RegisterOcpTemplateApiServiceServer(s grpc.ServiceRegistrar, srv OcpTemplateApiServiceServer) {
	s.RegisterService(&OcpTemplateApiService_ServiceDesc, srv)
}

func _OcpTemplateApiService_CreateTemplateV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpTemplateApiServiceServer).CreateTemplateV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozoncp.ocp_template_api.v1.OcpTemplateApiService/CreateTemplateV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpTemplateApiServiceServer).CreateTemplateV1(ctx, req.(*CreateTemplateV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpTemplateApiService_ServiceDesc is the grpc.ServiceDesc for OcpTemplateApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpTemplateApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozoncp.ocp_template_api.v1.OcpTemplateApiService",
	HandlerType: (*OcpTemplateApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplateV1",
			Handler:    _OcpTemplateApiService_CreateTemplateV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozoncp/ocp_template_api/v1/ocp_template_api.proto",
}
