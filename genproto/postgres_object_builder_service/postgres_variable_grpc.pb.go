// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: postgres_variable.proto

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

// VariableServiceClient is the client API for VariableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VariableServiceClient interface {
	Create(ctx context.Context, in *CreateVariableRequest, opts ...grpc.CallOption) (*Variable, error)
	GetList(ctx context.Context, in *GetAllVariablesRequest, opts ...grpc.CallOption) (*GetAllVariablesResponse, error)
	GetSingle(ctx context.Context, in *VariablePrimaryKey, opts ...grpc.CallOption) (*Variable, error)
	Update(ctx context.Context, in *Variable, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *VariablePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type variableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVariableServiceClient(cc grpc.ClientConnInterface) VariableServiceClient {
	return &variableServiceClient{cc}
}

func (c *variableServiceClient) Create(ctx context.Context, in *CreateVariableRequest, opts ...grpc.CallOption) (*Variable, error) {
	out := new(Variable)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.VariableService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variableServiceClient) GetList(ctx context.Context, in *GetAllVariablesRequest, opts ...grpc.CallOption) (*GetAllVariablesResponse, error) {
	out := new(GetAllVariablesResponse)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.VariableService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variableServiceClient) GetSingle(ctx context.Context, in *VariablePrimaryKey, opts ...grpc.CallOption) (*Variable, error) {
	out := new(Variable)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.VariableService/GetSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variableServiceClient) Update(ctx context.Context, in *Variable, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.VariableService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variableServiceClient) Delete(ctx context.Context, in *VariablePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.VariableService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VariableServiceServer is the server API for VariableService service.
// All implementations must embed UnimplementedVariableServiceServer
// for forward compatibility
type VariableServiceServer interface {
	Create(context.Context, *CreateVariableRequest) (*Variable, error)
	GetList(context.Context, *GetAllVariablesRequest) (*GetAllVariablesResponse, error)
	GetSingle(context.Context, *VariablePrimaryKey) (*Variable, error)
	Update(context.Context, *Variable) (*emptypb.Empty, error)
	Delete(context.Context, *VariablePrimaryKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedVariableServiceServer()
}

// UnimplementedVariableServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVariableServiceServer struct {
}

func (UnimplementedVariableServiceServer) Create(context.Context, *CreateVariableRequest) (*Variable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVariableServiceServer) GetList(context.Context, *GetAllVariablesRequest) (*GetAllVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedVariableServiceServer) GetSingle(context.Context, *VariablePrimaryKey) (*Variable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingle not implemented")
}
func (UnimplementedVariableServiceServer) Update(context.Context, *Variable) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVariableServiceServer) Delete(context.Context, *VariablePrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedVariableServiceServer) mustEmbedUnimplementedVariableServiceServer() {}

// UnsafeVariableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VariableServiceServer will
// result in compilation errors.
type UnsafeVariableServiceServer interface {
	mustEmbedUnimplementedVariableServiceServer()
}

func RegisterVariableServiceServer(s grpc.ServiceRegistrar, srv VariableServiceServer) {
	s.RegisterService(&VariableService_ServiceDesc, srv)
}

func _VariableService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.VariableService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).Create(ctx, req.(*CreateVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VariableService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.VariableService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).GetList(ctx, req.(*GetAllVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VariableService_GetSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VariablePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).GetSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.VariableService/GetSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).GetSingle(ctx, req.(*VariablePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _VariableService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Variable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.VariableService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).Update(ctx, req.(*Variable))
	}
	return interceptor(ctx, in, info, handler)
}

func _VariableService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VariablePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.VariableService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).Delete(ctx, req.(*VariablePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// VariableService_ServiceDesc is the grpc.ServiceDesc for VariableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VariableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postgres_object_builder_service.VariableService",
	HandlerType: (*VariableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VariableService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _VariableService_GetList_Handler,
		},
		{
			MethodName: "GetSingle",
			Handler:    _VariableService_GetSingle_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VariableService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VariableService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postgres_variable.proto",
}
