// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: release_service.proto

package versioning_service

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

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	Create(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*ReleaseWithCommit, error)
	GetByID(ctx context.Context, in *ReleasePrimaryKey, opts ...grpc.CallOption) (*ReleaseWithCommit, error)
	GetList(ctx context.Context, in *GetReleaseListRequest, opts ...grpc.CallOption) (*GetReleaseListResponse, error)
	Update(ctx context.Context, in *UpdateReleaseRequest, opts ...grpc.CallOption) (*ReleaseWithCommit, error)
	Delete(ctx context.Context, in *ReleasePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCurrentActive(ctx context.Context, in *GetCurrentReleaseRequest, opts ...grpc.CallOption) (*GetCurrentReleaseResponse, error)
	SetCurrentActive(ctx context.Context, in *SetCurrentReleaseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMultipleVersionInfo(ctx context.Context, in *GetMultipleVersionInfoRequest, opts ...grpc.CallOption) (*GetMultipleVersionInfoResponse, error)
}

type releaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReleaseServiceClient(cc grpc.ClientConnInterface) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) Create(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*ReleaseWithCommit, error) {
	out := new(ReleaseWithCommit)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) GetByID(ctx context.Context, in *ReleasePrimaryKey, opts ...grpc.CallOption) (*ReleaseWithCommit, error) {
	out := new(ReleaseWithCommit)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) GetList(ctx context.Context, in *GetReleaseListRequest, opts ...grpc.CallOption) (*GetReleaseListResponse, error) {
	out := new(GetReleaseListResponse)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Update(ctx context.Context, in *UpdateReleaseRequest, opts ...grpc.CallOption) (*ReleaseWithCommit, error) {
	out := new(ReleaseWithCommit)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Delete(ctx context.Context, in *ReleasePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) GetCurrentActive(ctx context.Context, in *GetCurrentReleaseRequest, opts ...grpc.CallOption) (*GetCurrentReleaseResponse, error) {
	out := new(GetCurrentReleaseResponse)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/GetCurrentActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) SetCurrentActive(ctx context.Context, in *SetCurrentReleaseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/SetCurrentActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) GetMultipleVersionInfo(ctx context.Context, in *GetMultipleVersionInfoRequest, opts ...grpc.CallOption) (*GetMultipleVersionInfoResponse, error) {
	out := new(GetMultipleVersionInfoResponse)
	err := c.cc.Invoke(ctx, "/versioning_service.ReleaseService/GetMultipleVersionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
// All implementations must embed UnimplementedReleaseServiceServer
// for forward compatibility
type ReleaseServiceServer interface {
	Create(context.Context, *CreateReleaseRequest) (*ReleaseWithCommit, error)
	GetByID(context.Context, *ReleasePrimaryKey) (*ReleaseWithCommit, error)
	GetList(context.Context, *GetReleaseListRequest) (*GetReleaseListResponse, error)
	Update(context.Context, *UpdateReleaseRequest) (*ReleaseWithCommit, error)
	Delete(context.Context, *ReleasePrimaryKey) (*emptypb.Empty, error)
	GetCurrentActive(context.Context, *GetCurrentReleaseRequest) (*GetCurrentReleaseResponse, error)
	SetCurrentActive(context.Context, *SetCurrentReleaseRequest) (*emptypb.Empty, error)
	GetMultipleVersionInfo(context.Context, *GetMultipleVersionInfoRequest) (*GetMultipleVersionInfoResponse, error)
	mustEmbedUnimplementedReleaseServiceServer()
}

// UnimplementedReleaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReleaseServiceServer struct {
}

func (UnimplementedReleaseServiceServer) Create(context.Context, *CreateReleaseRequest) (*ReleaseWithCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReleaseServiceServer) GetByID(context.Context, *ReleasePrimaryKey) (*ReleaseWithCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedReleaseServiceServer) GetList(context.Context, *GetReleaseListRequest) (*GetReleaseListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedReleaseServiceServer) Update(context.Context, *UpdateReleaseRequest) (*ReleaseWithCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReleaseServiceServer) Delete(context.Context, *ReleasePrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedReleaseServiceServer) GetCurrentActive(context.Context, *GetCurrentReleaseRequest) (*GetCurrentReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentActive not implemented")
}
func (UnimplementedReleaseServiceServer) SetCurrentActive(context.Context, *SetCurrentReleaseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCurrentActive not implemented")
}
func (UnimplementedReleaseServiceServer) GetMultipleVersionInfo(context.Context, *GetMultipleVersionInfoRequest) (*GetMultipleVersionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipleVersionInfo not implemented")
}
func (UnimplementedReleaseServiceServer) mustEmbedUnimplementedReleaseServiceServer() {}

// UnsafeReleaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReleaseServiceServer will
// result in compilation errors.
type UnsafeReleaseServiceServer interface {
	mustEmbedUnimplementedReleaseServiceServer()
}

func RegisterReleaseServiceServer(s grpc.ServiceRegistrar, srv ReleaseServiceServer) {
	s.RegisterService(&ReleaseService_ServiceDesc, srv)
}

func _ReleaseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Create(ctx, req.(*CreateReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).GetByID(ctx, req.(*ReleasePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleaseListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).GetList(ctx, req.(*GetReleaseListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Update(ctx, req.(*UpdateReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Delete(ctx, req.(*ReleasePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_GetCurrentActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).GetCurrentActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/GetCurrentActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).GetCurrentActive(ctx, req.(*GetCurrentReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_SetCurrentActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCurrentReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).SetCurrentActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/SetCurrentActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).SetCurrentActive(ctx, req.(*SetCurrentReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_GetMultipleVersionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipleVersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).GetMultipleVersionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versioning_service.ReleaseService/GetMultipleVersionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).GetMultipleVersionInfo(ctx, req.(*GetMultipleVersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReleaseService_ServiceDesc is the grpc.ServiceDesc for ReleaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReleaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "versioning_service.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReleaseService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _ReleaseService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ReleaseService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReleaseService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReleaseService_Delete_Handler,
		},
		{
			MethodName: "GetCurrentActive",
			Handler:    _ReleaseService_GetCurrentActive_Handler,
		},
		{
			MethodName: "SetCurrentActive",
			Handler:    _ReleaseService_SetCurrentActive_Handler,
		},
		{
			MethodName: "GetMultipleVersionInfo",
			Handler:    _ReleaseService_GetMultipleVersionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "release_service.proto",
}
