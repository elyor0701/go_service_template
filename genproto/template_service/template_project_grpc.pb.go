// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: template_project.proto

package template_service

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

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Register(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RegisterProjects(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Deregister(ctx context.Context, in *DeregisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Reconnect(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RegisterMany(ctx context.Context, in *RegisterManyProjectsRequest, opts ...grpc.CallOption) (*RegisterManyProjectsResponse, error)
	DeregisterMany(ctx context.Context, in *DeregisterManyProjectsRequest, opts ...grpc.CallOption) (*DeregisterManyProjectsResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Register(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/template_service.ProjectService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) RegisterProjects(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/template_service.ProjectService/RegisterProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Deregister(ctx context.Context, in *DeregisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/template_service.ProjectService/Deregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Reconnect(ctx context.Context, in *RegisterProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/template_service.ProjectService/Reconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) RegisterMany(ctx context.Context, in *RegisterManyProjectsRequest, opts ...grpc.CallOption) (*RegisterManyProjectsResponse, error) {
	out := new(RegisterManyProjectsResponse)
	err := c.cc.Invoke(ctx, "/template_service.ProjectService/RegisterMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeregisterMany(ctx context.Context, in *DeregisterManyProjectsRequest, opts ...grpc.CallOption) (*DeregisterManyProjectsResponse, error) {
	out := new(DeregisterManyProjectsResponse)
	err := c.cc.Invoke(ctx, "/template_service.ProjectService/DeregisterMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	Register(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error)
	RegisterProjects(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error)
	Deregister(context.Context, *DeregisterProjectRequest) (*emptypb.Empty, error)
	Reconnect(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error)
	RegisterMany(context.Context, *RegisterManyProjectsRequest) (*RegisterManyProjectsResponse, error)
	DeregisterMany(context.Context, *DeregisterManyProjectsRequest) (*DeregisterManyProjectsResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) Register(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedProjectServiceServer) RegisterProjects(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProjects not implemented")
}
func (UnimplementedProjectServiceServer) Deregister(context.Context, *DeregisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedProjectServiceServer) Reconnect(context.Context, *RegisterProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reconnect not implemented")
}
func (UnimplementedProjectServiceServer) RegisterMany(context.Context, *RegisterManyProjectsRequest) (*RegisterManyProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMany not implemented")
}
func (UnimplementedProjectServiceServer) DeregisterMany(context.Context, *DeregisterManyProjectsRequest) (*DeregisterManyProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterMany not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_service.ProjectService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Register(ctx, req.(*RegisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_RegisterProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).RegisterProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_service.ProjectService/RegisterProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).RegisterProjects(ctx, req.(*RegisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_service.ProjectService/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Deregister(ctx, req.(*DeregisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Reconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Reconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_service.ProjectService/Reconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Reconnect(ctx, req.(*RegisterProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_RegisterMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterManyProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).RegisterMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_service.ProjectService/RegisterMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).RegisterMany(ctx, req.(*RegisterManyProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeregisterMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterManyProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeregisterMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_service.ProjectService/DeregisterMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeregisterMany(ctx, req.(*DeregisterManyProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template_service.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ProjectService_Register_Handler,
		},
		{
			MethodName: "RegisterProjects",
			Handler:    _ProjectService_RegisterProjects_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _ProjectService_Deregister_Handler,
		},
		{
			MethodName: "Reconnect",
			Handler:    _ProjectService_Reconnect_Handler,
		},
		{
			MethodName: "RegisterMany",
			Handler:    _ProjectService_RegisterMany_Handler,
		},
		{
			MethodName: "DeregisterMany",
			Handler:    _ProjectService_DeregisterMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template_project.proto",
}
