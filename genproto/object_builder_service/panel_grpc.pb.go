// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: panel.proto

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

// PanelServiceClient is the client API for PanelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PanelServiceClient interface {
	Create(ctx context.Context, in *CreatePanelRequest, opts ...grpc.CallOption) (*Panel, error)
	GetList(ctx context.Context, in *GetAllPanelsRequest, opts ...grpc.CallOption) (*GetAllPanelsResponse, error)
	GetSingle(ctx context.Context, in *PanelPrimaryKey, opts ...grpc.CallOption) (*Panel, error)
	Update(ctx context.Context, in *Panel, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCoordinates(ctx context.Context, in *UpdatePanelCoordinatesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *PanelPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type panelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPanelServiceClient(cc grpc.ClientConnInterface) PanelServiceClient {
	return &panelServiceClient{cc}
}

func (c *panelServiceClient) Create(ctx context.Context, in *CreatePanelRequest, opts ...grpc.CallOption) (*Panel, error) {
	out := new(Panel)
	err := c.cc.Invoke(ctx, "/object_builder_service.PanelService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) GetList(ctx context.Context, in *GetAllPanelsRequest, opts ...grpc.CallOption) (*GetAllPanelsResponse, error) {
	out := new(GetAllPanelsResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.PanelService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) GetSingle(ctx context.Context, in *PanelPrimaryKey, opts ...grpc.CallOption) (*Panel, error) {
	out := new(Panel)
	err := c.cc.Invoke(ctx, "/object_builder_service.PanelService/GetSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) Update(ctx context.Context, in *Panel, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.PanelService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) UpdateCoordinates(ctx context.Context, in *UpdatePanelCoordinatesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.PanelService/UpdateCoordinates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) Delete(ctx context.Context, in *PanelPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.PanelService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PanelServiceServer is the server API for PanelService service.
// All implementations must embed UnimplementedPanelServiceServer
// for forward compatibility
type PanelServiceServer interface {
	Create(context.Context, *CreatePanelRequest) (*Panel, error)
	GetList(context.Context, *GetAllPanelsRequest) (*GetAllPanelsResponse, error)
	GetSingle(context.Context, *PanelPrimaryKey) (*Panel, error)
	Update(context.Context, *Panel) (*emptypb.Empty, error)
	UpdateCoordinates(context.Context, *UpdatePanelCoordinatesRequest) (*emptypb.Empty, error)
	Delete(context.Context, *PanelPrimaryKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedPanelServiceServer()
}

// UnimplementedPanelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPanelServiceServer struct {
}

func (UnimplementedPanelServiceServer) Create(context.Context, *CreatePanelRequest) (*Panel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPanelServiceServer) GetList(context.Context, *GetAllPanelsRequest) (*GetAllPanelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedPanelServiceServer) GetSingle(context.Context, *PanelPrimaryKey) (*Panel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingle not implemented")
}
func (UnimplementedPanelServiceServer) Update(context.Context, *Panel) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPanelServiceServer) UpdateCoordinates(context.Context, *UpdatePanelCoordinatesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoordinates not implemented")
}
func (UnimplementedPanelServiceServer) Delete(context.Context, *PanelPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPanelServiceServer) mustEmbedUnimplementedPanelServiceServer() {}

// UnsafePanelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PanelServiceServer will
// result in compilation errors.
type UnsafePanelServiceServer interface {
	mustEmbedUnimplementedPanelServiceServer()
}

func RegisterPanelServiceServer(s grpc.ServiceRegistrar, srv PanelServiceServer) {
	s.RegisterService(&PanelService_ServiceDesc, srv)
}

func _PanelService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePanelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PanelService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).Create(ctx, req.(*CreatePanelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPanelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PanelService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).GetList(ctx, req.(*GetAllPanelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_GetSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PanelPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).GetSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PanelService/GetSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).GetSingle(ctx, req.(*PanelPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Panel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PanelService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).Update(ctx, req.(*Panel))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_UpdateCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePanelCoordinatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).UpdateCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PanelService/UpdateCoordinates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).UpdateCoordinates(ctx, req.(*UpdatePanelCoordinatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PanelPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PanelService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).Delete(ctx, req.(*PanelPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// PanelService_ServiceDesc is the grpc.ServiceDesc for PanelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PanelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.PanelService",
	HandlerType: (*PanelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PanelService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _PanelService_GetList_Handler,
		},
		{
			MethodName: "GetSingle",
			Handler:    _PanelService_GetSingle_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PanelService_Update_Handler,
		},
		{
			MethodName: "UpdateCoordinates",
			Handler:    _PanelService_UpdateCoordinates_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PanelService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "panel.proto",
}
