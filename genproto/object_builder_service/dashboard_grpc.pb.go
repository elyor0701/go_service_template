// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: dashboard.proto

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

// DashboardServiceClient is the client API for DashboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardServiceClient interface {
	Create(ctx context.Context, in *CreateDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error)
	GetList(ctx context.Context, in *GetAllDashboardsRequest, opts ...grpc.CallOption) (*GetAllDashboadsResponse, error)
	GetSingle(ctx context.Context, in *DashboardPrimaryKey, opts ...grpc.CallOption) (*Dashboard, error)
	Update(ctx context.Context, in *Dashboard, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DashboardPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dashboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardServiceClient(cc grpc.ClientConnInterface) DashboardServiceClient {
	return &dashboardServiceClient{cc}
}

func (c *dashboardServiceClient) Create(ctx context.Context, in *CreateDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error) {
	out := new(Dashboard)
	err := c.cc.Invoke(ctx, "/object_builder_service.DashboardService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) GetList(ctx context.Context, in *GetAllDashboardsRequest, opts ...grpc.CallOption) (*GetAllDashboadsResponse, error) {
	out := new(GetAllDashboadsResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.DashboardService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) GetSingle(ctx context.Context, in *DashboardPrimaryKey, opts ...grpc.CallOption) (*Dashboard, error) {
	out := new(Dashboard)
	err := c.cc.Invoke(ctx, "/object_builder_service.DashboardService/GetSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) Update(ctx context.Context, in *Dashboard, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.DashboardService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) Delete(ctx context.Context, in *DashboardPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.DashboardService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServiceServer is the server API for DashboardService service.
// All implementations must embed UnimplementedDashboardServiceServer
// for forward compatibility
type DashboardServiceServer interface {
	Create(context.Context, *CreateDashboardRequest) (*Dashboard, error)
	GetList(context.Context, *GetAllDashboardsRequest) (*GetAllDashboadsResponse, error)
	GetSingle(context.Context, *DashboardPrimaryKey) (*Dashboard, error)
	Update(context.Context, *Dashboard) (*emptypb.Empty, error)
	Delete(context.Context, *DashboardPrimaryKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedDashboardServiceServer()
}

// UnimplementedDashboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardServiceServer struct {
}

func (UnimplementedDashboardServiceServer) Create(context.Context, *CreateDashboardRequest) (*Dashboard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDashboardServiceServer) GetList(context.Context, *GetAllDashboardsRequest) (*GetAllDashboadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedDashboardServiceServer) GetSingle(context.Context, *DashboardPrimaryKey) (*Dashboard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingle not implemented")
}
func (UnimplementedDashboardServiceServer) Update(context.Context, *Dashboard) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDashboardServiceServer) Delete(context.Context, *DashboardPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDashboardServiceServer) mustEmbedUnimplementedDashboardServiceServer() {}

// UnsafeDashboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardServiceServer will
// result in compilation errors.
type UnsafeDashboardServiceServer interface {
	mustEmbedUnimplementedDashboardServiceServer()
}

func RegisterDashboardServiceServer(s grpc.ServiceRegistrar, srv DashboardServiceServer) {
	s.RegisterService(&DashboardService_ServiceDesc, srv)
}

func _DashboardService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.DashboardService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).Create(ctx, req.(*CreateDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDashboardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.DashboardService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).GetList(ctx, req.(*GetAllDashboardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_GetSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DashboardPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).GetSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.DashboardService/GetSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).GetSingle(ctx, req.(*DashboardPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dashboard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.DashboardService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).Update(ctx, req.(*Dashboard))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DashboardPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.DashboardService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).Delete(ctx, req.(*DashboardPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// DashboardService_ServiceDesc is the grpc.ServiceDesc for DashboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DashboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.DashboardService",
	HandlerType: (*DashboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DashboardService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _DashboardService_GetList_Handler,
		},
		{
			MethodName: "GetSingle",
			Handler:    _DashboardService_GetSingle_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DashboardService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DashboardService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboard.proto",
}
