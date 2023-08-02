// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: postgres_permission.proto

package postgres_object_builder_service

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

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	UpsertPermissionsByAppId(ctx context.Context, in *UpsertPermissionsByAppIdRequest, opts ...grpc.CallOption) (*UpsertPermissionsByAppIdResponse, error)
	GetAllPermissionsByRoleId(ctx context.Context, in *GetAllPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error)
	GetFieldPermissions(ctx context.Context, in *GetFieldPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error)
	GetActionPermissions(ctx context.Context, in *GetActionPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error)
	GetViewRelationPermissions(ctx context.Context, in *GetActionPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error)
	GetListWithRoleAppTablePermissions(ctx context.Context, in *GetListWithRoleAppTablePermissionsRequest, opts ...grpc.CallOption) (*GetListWithRoleAppTablePermissionsResponse, error)
	UpdateRoleAppTablePermissions(ctx context.Context, in *UpdateRoleAppTablePermissionsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) UpsertPermissionsByAppId(ctx context.Context, in *UpsertPermissionsByAppIdRequest, opts ...grpc.CallOption) (*UpsertPermissionsByAppIdResponse, error) {
	out := new(UpsertPermissionsByAppIdResponse)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.PermissionService/UpsertPermissionsByAppId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetAllPermissionsByRoleId(ctx context.Context, in *GetAllPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.PermissionService/GetAllPermissionsByRoleId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetFieldPermissions(ctx context.Context, in *GetFieldPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.PermissionService/GetFieldPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetActionPermissions(ctx context.Context, in *GetActionPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.PermissionService/GetActionPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetViewRelationPermissions(ctx context.Context, in *GetActionPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.PermissionService/GetViewRelationPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetListWithRoleAppTablePermissions(ctx context.Context, in *GetListWithRoleAppTablePermissionsRequest, opts ...grpc.CallOption) (*GetListWithRoleAppTablePermissionsResponse, error) {
	out := new(GetListWithRoleAppTablePermissionsResponse)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.PermissionService/GetListWithRoleAppTablePermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdateRoleAppTablePermissions(ctx context.Context, in *UpdateRoleAppTablePermissionsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.PermissionService/UpdateRoleAppTablePermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations must embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	UpsertPermissionsByAppId(context.Context, *UpsertPermissionsByAppIdRequest) (*UpsertPermissionsByAppIdResponse, error)
	GetAllPermissionsByRoleId(context.Context, *GetAllPermissionRequest) (*CommonMessage, error)
	GetFieldPermissions(context.Context, *GetFieldPermissionRequest) (*CommonMessage, error)
	GetActionPermissions(context.Context, *GetActionPermissionRequest) (*CommonMessage, error)
	GetViewRelationPermissions(context.Context, *GetActionPermissionRequest) (*CommonMessage, error)
	GetListWithRoleAppTablePermissions(context.Context, *GetListWithRoleAppTablePermissionsRequest) (*GetListWithRoleAppTablePermissionsResponse, error)
	UpdateRoleAppTablePermissions(context.Context, *UpdateRoleAppTablePermissionsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPermissionServiceServer()
}

// UnimplementedPermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) UpsertPermissionsByAppId(context.Context, *UpsertPermissionsByAppIdRequest) (*UpsertPermissionsByAppIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertPermissionsByAppId not implemented")
}
func (UnimplementedPermissionServiceServer) GetAllPermissionsByRoleId(context.Context, *GetAllPermissionRequest) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPermissionsByRoleId not implemented")
}
func (UnimplementedPermissionServiceServer) GetFieldPermissions(context.Context, *GetFieldPermissionRequest) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFieldPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) GetActionPermissions(context.Context, *GetActionPermissionRequest) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActionPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) GetViewRelationPermissions(context.Context, *GetActionPermissionRequest) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetViewRelationPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) GetListWithRoleAppTablePermissions(context.Context, *GetListWithRoleAppTablePermissionsRequest) (*GetListWithRoleAppTablePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListWithRoleAppTablePermissions not implemented")
}
func (UnimplementedPermissionServiceServer) UpdateRoleAppTablePermissions(context.Context, *UpdateRoleAppTablePermissionsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoleAppTablePermissions not implemented")
}
func (UnimplementedPermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_UpsertPermissionsByAppId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPermissionsByAppIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpsertPermissionsByAppId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.PermissionService/UpsertPermissionsByAppId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpsertPermissionsByAppId(ctx, req.(*UpsertPermissionsByAppIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetAllPermissionsByRoleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetAllPermissionsByRoleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.PermissionService/GetAllPermissionsByRoleId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetAllPermissionsByRoleId(ctx, req.(*GetAllPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetFieldPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFieldPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetFieldPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.PermissionService/GetFieldPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetFieldPermissions(ctx, req.(*GetFieldPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetActionPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetActionPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.PermissionService/GetActionPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetActionPermissions(ctx, req.(*GetActionPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetViewRelationPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetViewRelationPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.PermissionService/GetViewRelationPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetViewRelationPermissions(ctx, req.(*GetActionPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetListWithRoleAppTablePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListWithRoleAppTablePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetListWithRoleAppTablePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.PermissionService/GetListWithRoleAppTablePermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetListWithRoleAppTablePermissions(ctx, req.(*GetListWithRoleAppTablePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdateRoleAppTablePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleAppTablePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdateRoleAppTablePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.PermissionService/UpdateRoleAppTablePermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdateRoleAppTablePermissions(ctx, req.(*UpdateRoleAppTablePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postgres_object_builder_service.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertPermissionsByAppId",
			Handler:    _PermissionService_UpsertPermissionsByAppId_Handler,
		},
		{
			MethodName: "GetAllPermissionsByRoleId",
			Handler:    _PermissionService_GetAllPermissionsByRoleId_Handler,
		},
		{
			MethodName: "GetFieldPermissions",
			Handler:    _PermissionService_GetFieldPermissions_Handler,
		},
		{
			MethodName: "GetActionPermissions",
			Handler:    _PermissionService_GetActionPermissions_Handler,
		},
		{
			MethodName: "GetViewRelationPermissions",
			Handler:    _PermissionService_GetViewRelationPermissions_Handler,
		},
		{
			MethodName: "GetListWithRoleAppTablePermissions",
			Handler:    _PermissionService_GetListWithRoleAppTablePermissions_Handler,
		},
		{
			MethodName: "UpdateRoleAppTablePermissions",
			Handler:    _PermissionService_UpdateRoleAppTablePermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postgres_permission.proto",
}
