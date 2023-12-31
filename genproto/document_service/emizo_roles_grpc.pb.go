// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: emizo_roles.proto

package document_service

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

// EimzoRolesClient is the client API for EimzoRoles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EimzoRolesClient interface {
	Create(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *RoleGetReq, opts ...grpc.CallOption) (*RoleGetRes, error)
	Reload(ctx context.Context, in *ReloadReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type eimzoRolesClient struct {
	cc grpc.ClientConnInterface
}

func NewEimzoRolesClient(cc grpc.ClientConnInterface) EimzoRolesClient {
	return &eimzoRolesClient{cc}
}

func (c *eimzoRolesClient) Create(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/document_service.EimzoRoles/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eimzoRolesClient) Get(ctx context.Context, in *RoleGetReq, opts ...grpc.CallOption) (*RoleGetRes, error) {
	out := new(RoleGetRes)
	err := c.cc.Invoke(ctx, "/document_service.EimzoRoles/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eimzoRolesClient) Reload(ctx context.Context, in *ReloadReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/document_service.EimzoRoles/Reload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EimzoRolesServer is the server API for EimzoRoles service.
// All implementations must embed UnimplementedEimzoRolesServer
// for forward compatibility
type EimzoRolesServer interface {
	Create(context.Context, *CreateRoleReq) (*emptypb.Empty, error)
	Get(context.Context, *RoleGetReq) (*RoleGetRes, error)
	Reload(context.Context, *ReloadReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedEimzoRolesServer()
}

// UnimplementedEimzoRolesServer must be embedded to have forward compatible implementations.
type UnimplementedEimzoRolesServer struct {
}

func (UnimplementedEimzoRolesServer) Create(context.Context, *CreateRoleReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEimzoRolesServer) Get(context.Context, *RoleGetReq) (*RoleGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEimzoRolesServer) Reload(context.Context, *ReloadReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reload not implemented")
}
func (UnimplementedEimzoRolesServer) mustEmbedUnimplementedEimzoRolesServer() {}

// UnsafeEimzoRolesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EimzoRolesServer will
// result in compilation errors.
type UnsafeEimzoRolesServer interface {
	mustEmbedUnimplementedEimzoRolesServer()
}

func RegisterEimzoRolesServer(s grpc.ServiceRegistrar, srv EimzoRolesServer) {
	s.RegisterService(&EimzoRoles_ServiceDesc, srv)
}

func _EimzoRoles_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EimzoRolesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.EimzoRoles/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EimzoRolesServer).Create(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EimzoRoles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EimzoRolesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.EimzoRoles/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EimzoRolesServer).Get(ctx, req.(*RoleGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EimzoRoles_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EimzoRolesServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.EimzoRoles/Reload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EimzoRolesServer).Reload(ctx, req.(*ReloadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EimzoRoles_ServiceDesc is the grpc.ServiceDesc for EimzoRoles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EimzoRoles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "document_service.EimzoRoles",
	HandlerType: (*EimzoRolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EimzoRoles_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EimzoRoles_Get_Handler,
		},
		{
			MethodName: "Reload",
			Handler:    _EimzoRoles_Reload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emizo_roles.proto",
}
