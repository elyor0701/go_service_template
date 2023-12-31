// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: web_page_app.proto

package web_page_service

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

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	CreateApp(ctx context.Context, in *CreateAppReq, opts ...grpc.CallOption) (*App, error)
	GetListApp(ctx context.Context, in *GetListAppReq, opts ...grpc.CallOption) (*GetListAppRes, error)
	GetSingleApp(ctx context.Context, in *GetSingleAppReq, opts ...grpc.CallOption) (*App, error)
	UpdateApp(ctx context.Context, in *UpdateAppReq, opts ...grpc.CallOption) (*App, error)
	DeleteApp(ctx context.Context, in *DeleteAppReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) CreateApp(ctx context.Context, in *CreateAppReq, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, "/web_page_service.AppService/CreateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetListApp(ctx context.Context, in *GetListAppReq, opts ...grpc.CallOption) (*GetListAppRes, error) {
	out := new(GetListAppRes)
	err := c.cc.Invoke(ctx, "/web_page_service.AppService/GetListApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetSingleApp(ctx context.Context, in *GetSingleAppReq, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, "/web_page_service.AppService/GetSingleApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) UpdateApp(ctx context.Context, in *UpdateAppReq, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, "/web_page_service.AppService/UpdateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteApp(ctx context.Context, in *DeleteAppReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/web_page_service.AppService/DeleteApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	CreateApp(context.Context, *CreateAppReq) (*App, error)
	GetListApp(context.Context, *GetListAppReq) (*GetListAppRes, error)
	GetSingleApp(context.Context, *GetSingleAppReq) (*App, error)
	UpdateApp(context.Context, *UpdateAppReq) (*App, error)
	DeleteApp(context.Context, *DeleteAppReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) CreateApp(context.Context, *CreateAppReq) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (UnimplementedAppServiceServer) GetListApp(context.Context, *GetListAppReq) (*GetListAppRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListApp not implemented")
}
func (UnimplementedAppServiceServer) GetSingleApp(context.Context, *GetSingleAppReq) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleApp not implemented")
}
func (UnimplementedAppServiceServer) UpdateApp(context.Context, *UpdateAppReq) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApp not implemented")
}
func (UnimplementedAppServiceServer) DeleteApp(context.Context, *DeleteAppReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApp not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.AppService/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CreateApp(ctx, req.(*CreateAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetListApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetListApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.AppService/GetListApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetListApp(ctx, req.(*GetListAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetSingleApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetSingleApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.AppService/GetSingleApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetSingleApp(ctx, req.(*GetSingleAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_UpdateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).UpdateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.AppService/UpdateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).UpdateApp(ctx, req.(*UpdateAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.AppService/DeleteApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteApp(ctx, req.(*DeleteAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "web_page_service.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _AppService_CreateApp_Handler,
		},
		{
			MethodName: "GetListApp",
			Handler:    _AppService_GetListApp_Handler,
		},
		{
			MethodName: "GetSingleApp",
			Handler:    _AppService_GetSingleApp_Handler,
		},
		{
			MethodName: "UpdateApp",
			Handler:    _AppService_UpdateApp_Handler,
		},
		{
			MethodName: "DeleteApp",
			Handler:    _AppService_DeleteApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web_page_app.proto",
}
