// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: query_folder.proto

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

// QueryFolderServiceClient is the client API for QueryFolderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryFolderServiceClient interface {
	Create(ctx context.Context, in *CreateQueryFolderRequest, opts ...grpc.CallOption) (*QueryFolder, error)
	GetAll(ctx context.Context, in *GetAllQueryFolderRequest, opts ...grpc.CallOption) (*GetAllQueryFolderResponse, error)
	GetById(ctx context.Context, in *QueryFolderId, opts ...grpc.CallOption) (*QueryFolder, error)
	Update(ctx context.Context, in *QueryFolder, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *QueryFolderId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type queryFolderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryFolderServiceClient(cc grpc.ClientConnInterface) QueryFolderServiceClient {
	return &queryFolderServiceClient{cc}
}

func (c *queryFolderServiceClient) Create(ctx context.Context, in *CreateQueryFolderRequest, opts ...grpc.CallOption) (*QueryFolder, error) {
	out := new(QueryFolder)
	err := c.cc.Invoke(ctx, "/object_builder_service.QueryFolderService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryFolderServiceClient) GetAll(ctx context.Context, in *GetAllQueryFolderRequest, opts ...grpc.CallOption) (*GetAllQueryFolderResponse, error) {
	out := new(GetAllQueryFolderResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.QueryFolderService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryFolderServiceClient) GetById(ctx context.Context, in *QueryFolderId, opts ...grpc.CallOption) (*QueryFolder, error) {
	out := new(QueryFolder)
	err := c.cc.Invoke(ctx, "/object_builder_service.QueryFolderService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryFolderServiceClient) Update(ctx context.Context, in *QueryFolder, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.QueryFolderService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryFolderServiceClient) Delete(ctx context.Context, in *QueryFolderId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.QueryFolderService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryFolderServiceServer is the server API for QueryFolderService service.
// All implementations must embed UnimplementedQueryFolderServiceServer
// for forward compatibility
type QueryFolderServiceServer interface {
	Create(context.Context, *CreateQueryFolderRequest) (*QueryFolder, error)
	GetAll(context.Context, *GetAllQueryFolderRequest) (*GetAllQueryFolderResponse, error)
	GetById(context.Context, *QueryFolderId) (*QueryFolder, error)
	Update(context.Context, *QueryFolder) (*emptypb.Empty, error)
	Delete(context.Context, *QueryFolderId) (*emptypb.Empty, error)
	mustEmbedUnimplementedQueryFolderServiceServer()
}

// UnimplementedQueryFolderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueryFolderServiceServer struct {
}

func (UnimplementedQueryFolderServiceServer) Create(context.Context, *CreateQueryFolderRequest) (*QueryFolder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedQueryFolderServiceServer) GetAll(context.Context, *GetAllQueryFolderRequest) (*GetAllQueryFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedQueryFolderServiceServer) GetById(context.Context, *QueryFolderId) (*QueryFolder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedQueryFolderServiceServer) Update(context.Context, *QueryFolder) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedQueryFolderServiceServer) Delete(context.Context, *QueryFolderId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedQueryFolderServiceServer) mustEmbedUnimplementedQueryFolderServiceServer() {}

// UnsafeQueryFolderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryFolderServiceServer will
// result in compilation errors.
type UnsafeQueryFolderServiceServer interface {
	mustEmbedUnimplementedQueryFolderServiceServer()
}

func RegisterQueryFolderServiceServer(s grpc.ServiceRegistrar, srv QueryFolderServiceServer) {
	s.RegisterService(&QueryFolderService_ServiceDesc, srv)
}

func _QueryFolderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueryFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryFolderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.QueryFolderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryFolderServiceServer).Create(ctx, req.(*CreateQueryFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryFolderService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllQueryFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryFolderServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.QueryFolderService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryFolderServiceServer).GetAll(ctx, req.(*GetAllQueryFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryFolderService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFolderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryFolderServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.QueryFolderService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryFolderServiceServer).GetById(ctx, req.(*QueryFolderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryFolderService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFolder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryFolderServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.QueryFolderService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryFolderServiceServer).Update(ctx, req.(*QueryFolder))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryFolderService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFolderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryFolderServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.QueryFolderService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryFolderServiceServer).Delete(ctx, req.(*QueryFolderId))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryFolderService_ServiceDesc is the grpc.ServiceDesc for QueryFolderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryFolderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.QueryFolderService",
	HandlerType: (*QueryFolderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _QueryFolderService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _QueryFolderService_GetAll_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _QueryFolderService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _QueryFolderService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _QueryFolderService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query_folder.proto",
}
