// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: builder_project.proto

package object_builder_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BuilderProjectServiceClient is the client API for BuilderProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuilderProjectServiceClient interface {
	Register(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RegisterProjects(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Deregister(ctx context.Context, in *DeregisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Reconnect(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RegisterMany(ctx context.Context, in *RegisterManyProjectsRequest, opts ...grpc.CallOption) (*RegisterManyProjectsResponse, error)
	DeregisterMany(ctx context.Context, in *DeregisterManyProjectsRequest, opts ...grpc.CallOption) (*DeregisterManyProjectsResponse, error)
}

type builderProjectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuilderProjectServiceClient(cc grpc.ClientConnInterface) BuilderProjectServiceClient {
	return &builderProjectServiceClient{cc}
}

func (c *builderProjectServiceClient) Register(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.BuilderProjectService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderProjectServiceClient) RegisterProjects(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.BuilderProjectService/RegisterProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderProjectServiceClient) Deregister(ctx context.Context, in *DeregisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.BuilderProjectService/Deregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderProjectServiceClient) Reconnect(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.BuilderProjectService/Reconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderProjectServiceClient) RegisterMany(ctx context.Context, in *RegisterManyProjectsRequest, opts ...grpc.CallOption) (*RegisterManyProjectsResponse, error) {
	out := new(RegisterManyProjectsResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.BuilderProjectService/RegisterMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderProjectServiceClient) DeregisterMany(ctx context.Context, in *DeregisterManyProjectsRequest, opts ...grpc.CallOption) (*DeregisterManyProjectsResponse, error) {
	out := new(DeregisterManyProjectsResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.BuilderProjectService/DeregisterMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuilderProjectServiceServer is the server API for BuilderProjectService service.
// All implementations must embed UnimplementedBuilderProjectServiceServer
// for forward compatibility
type BuilderProjectServiceServer interface {
	Register(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error)
	RegisterProjects(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error)
	Deregister(context.Context, *DeregisterProjectRequest) (*emptypb.Empty, error)
	Reconnect(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error)
	RegisterMany(context.Context, *RegisterManyProjectsRequest) (*RegisterManyProjectsResponse, error)
	DeregisterMany(context.Context, *DeregisterManyProjectsRequest) (*DeregisterManyProjectsResponse, error)
	mustEmbedUnimplementedBuilderProjectServiceServer()
}

// UnimplementedBuilderProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBuilderProjectServiceServer struct {
}

func (UnimplementedBuilderProjectServiceServer) Register(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedBuilderProjectServiceServer) RegisterProjects(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProjects not implemented")
}
func (UnimplementedBuilderProjectServiceServer) Deregister(context.Context, *DeregisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedBuilderProjectServiceServer) Reconnect(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reconnect not implemented")
}
func (UnimplementedBuilderProjectServiceServer) RegisterMany(context.Context, *RegisterManyProjectsRequest) (*RegisterManyProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMany not implemented")
}
func (UnimplementedBuilderProjectServiceServer) DeregisterMany(context.Context, *DeregisterManyProjectsRequest) (*DeregisterManyProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterMany not implemented")
}
func (UnimplementedBuilderProjectServiceServer) mustEmbedUnimplementedBuilderProjectServiceServer() {}

// UnsafeBuilderProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuilderProjectServiceServer will
// result in compilation errors.
type UnsafeBuilderProjectServiceServer interface {
	mustEmbedUnimplementedBuilderProjectServiceServer()
}

func RegisterBuilderProjectServiceServer(s grpc.ServiceRegistrar, srv BuilderProjectServiceServer) {
	s.RegisterService(&BuilderProjectService_ServiceDesc, srv)
}

func _BuilderProjectService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderProjectServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.BuilderProjectService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderProjectServiceServer).Register(ctx, req.(*RegisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderProjectService_RegisterProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderProjectServiceServer).RegisterProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.BuilderProjectService/RegisterProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderProjectServiceServer).RegisterProjects(ctx, req.(*RegisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderProjectService_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderProjectServiceServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.BuilderProjectService/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderProjectServiceServer).Deregister(ctx, req.(*DeregisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderProjectService_Reconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderProjectServiceServer).Reconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.BuilderProjectService/Reconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderProjectServiceServer).Reconnect(ctx, req.(*RegisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderProjectService_RegisterMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterManyProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderProjectServiceServer).RegisterMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.BuilderProjectService/RegisterMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderProjectServiceServer).RegisterMany(ctx, req.(*RegisterManyProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderProjectService_DeregisterMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterManyProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderProjectServiceServer).DeregisterMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.BuilderProjectService/DeregisterMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderProjectServiceServer).DeregisterMany(ctx, req.(*DeregisterManyProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BuilderProjectService_ServiceDesc is the grpc.ServiceDesc for BuilderProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuilderProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.BuilderProjectService",
	HandlerType: (*BuilderProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _BuilderProjectService_Register_Handler,
		},
		{
			MethodName: "RegisterProjects",
			Handler:    _BuilderProjectService_RegisterProjects_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _BuilderProjectService_Deregister_Handler,
		},
		{
			MethodName: "Reconnect",
			Handler:    _BuilderProjectService_Reconnect_Handler,
		},
		{
			MethodName: "RegisterMany",
			Handler:    _BuilderProjectService_RegisterMany_Handler,
		},
		{
			MethodName: "DeregisterMany",
			Handler:    _BuilderProjectService_DeregisterMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "builder_project.proto",
}
