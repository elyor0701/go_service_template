// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: cascading.proto

package object_builder_service

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

// CascadingServiceClient is the client API for CascadingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CascadingServiceClient interface {
	GetCascadings(ctx context.Context, in *GetCascadingRequest, opts ...grpc.CallOption) (*CommonMessage, error)
}

type cascadingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCascadingServiceClient(cc grpc.ClientConnInterface) CascadingServiceClient {
	return &cascadingServiceClient{cc}
}

func (c *cascadingServiceClient) GetCascadings(ctx context.Context, in *GetCascadingRequest, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.CascadingService/GetCascadings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CascadingServiceServer is the server API for CascadingService service.
// All implementations must embed UnimplementedCascadingServiceServer
// for forward compatibility
type CascadingServiceServer interface {
	GetCascadings(context.Context, *GetCascadingRequest) (*CommonMessage, error)
	mustEmbedUnimplementedCascadingServiceServer()
}

// UnimplementedCascadingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCascadingServiceServer struct {
}

func (UnimplementedCascadingServiceServer) GetCascadings(context.Context, *GetCascadingRequest) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCascadings not implemented")
}
func (UnimplementedCascadingServiceServer) mustEmbedUnimplementedCascadingServiceServer() {}

// UnsafeCascadingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CascadingServiceServer will
// result in compilation errors.
type UnsafeCascadingServiceServer interface {
	mustEmbedUnimplementedCascadingServiceServer()
}

func RegisterCascadingServiceServer(s grpc.ServiceRegistrar, srv CascadingServiceServer) {
	s.RegisterService(&CascadingService_ServiceDesc, srv)
}

func _CascadingService_GetCascadings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCascadingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CascadingServiceServer).GetCascadings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.CascadingService/GetCascadings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CascadingServiceServer).GetCascadings(ctx, req.(*GetCascadingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CascadingService_ServiceDesc is the grpc.ServiceDesc for CascadingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CascadingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.CascadingService",
	HandlerType: (*CascadingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCascadings",
			Handler:    _CascadingService_GetCascadings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cascading.proto",
}