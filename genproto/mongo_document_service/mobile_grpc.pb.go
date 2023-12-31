// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: mobile.proto

package mongo_document_service

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

// MobileServiceClient is the client API for MobileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobileServiceClient interface {
	Create(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Mobile, error)
	GetAll(ctx context.Context, in *ListMobileReq, opts ...grpc.CallOption) (*ListMobileResp, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mobileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileServiceClient(cc grpc.ClientConnInterface) MobileServiceClient {
	return &mobileServiceClient{cc}
}

func (c *mobileServiceClient) Create(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.MobileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Mobile, error) {
	out := new(Mobile)
	err := c.cc.Invoke(ctx, "/mongo_document_service.MobileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) GetAll(ctx context.Context, in *ListMobileReq, opts ...grpc.CallOption) (*ListMobileResp, error) {
	out := new(ListMobileResp)
	err := c.cc.Invoke(ctx, "/mongo_document_service.MobileService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.MobileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileServiceClient) Update(ctx context.Context, in *Mobile, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.MobileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileServiceServer is the server API for MobileService service.
// All implementations must embed UnimplementedMobileServiceServer
// for forward compatibility
type MobileServiceServer interface {
	Create(context.Context, *Mobile) (*emptypb.Empty, error)
	Get(context.Context, *ById) (*Mobile, error)
	GetAll(context.Context, *ListMobileReq) (*ListMobileResp, error)
	Delete(context.Context, *ById) (*emptypb.Empty, error)
	Update(context.Context, *Mobile) (*emptypb.Empty, error)
	mustEmbedUnimplementedMobileServiceServer()
}

// UnimplementedMobileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMobileServiceServer struct {
}

func (UnimplementedMobileServiceServer) Create(context.Context, *Mobile) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMobileServiceServer) Get(context.Context, *ById) (*Mobile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMobileServiceServer) GetAll(context.Context, *ListMobileReq) (*ListMobileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMobileServiceServer) Delete(context.Context, *ById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMobileServiceServer) Update(context.Context, *Mobile) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMobileServiceServer) mustEmbedUnimplementedMobileServiceServer() {}

// UnsafeMobileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobileServiceServer will
// result in compilation errors.
type UnsafeMobileServiceServer interface {
	mustEmbedUnimplementedMobileServiceServer()
}

func RegisterMobileServiceServer(s grpc.ServiceRegistrar, srv MobileServiceServer) {
	s.RegisterService(&MobileService_ServiceDesc, srv)
}

func _MobileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mobile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.MobileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).Create(ctx, req.(*Mobile))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.MobileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).Get(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMobileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.MobileService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).GetAll(ctx, req.(*ListMobileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.MobileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mobile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.MobileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServiceServer).Update(ctx, req.(*Mobile))
	}
	return interceptor(ctx, in, info, handler)
}

// MobileService_ServiceDesc is the grpc.ServiceDesc for MobileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MobileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongo_document_service.MobileService",
	HandlerType: (*MobileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MobileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MobileService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MobileService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MobileService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MobileService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mobile.proto",
}
