// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: web_page.proto

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

// WebPageServiceClient is the client API for WebPageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebPageServiceClient interface {
	CreateWebPage(ctx context.Context, in *CreateWebPageReq, opts ...grpc.CallOption) (*WebPage, error)
	GetListWebPage(ctx context.Context, in *GetListWebPageReq, opts ...grpc.CallOption) (*GetListWebPageRes, error)
	GetSingleWebPage(ctx context.Context, in *GetSingleWebPageReq, opts ...grpc.CallOption) (*WebPage, error)
	UpdateWebPage(ctx context.Context, in *WebPage, opts ...grpc.CallOption) (*WebPage, error)
	DeleteWebPage(ctx context.Context, in *DeleteWebPageReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevertWebPage(ctx context.Context, in *RevertWebPageReq, opts ...grpc.CallOption) (*WebPage, error)
	CreateManyWebPage(ctx context.Context, in *ManyVersions, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetWebPageChanges(ctx context.Context, in *GetWebPageHistoryReq, opts ...grpc.CallOption) (*GetWebPageHistoryRes, error)
}

type webPageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebPageServiceClient(cc grpc.ClientConnInterface) WebPageServiceClient {
	return &webPageServiceClient{cc}
}

func (c *webPageServiceClient) CreateWebPage(ctx context.Context, in *CreateWebPageReq, opts ...grpc.CallOption) (*WebPage, error) {
	out := new(WebPage)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/CreateWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageServiceClient) GetListWebPage(ctx context.Context, in *GetListWebPageReq, opts ...grpc.CallOption) (*GetListWebPageRes, error) {
	out := new(GetListWebPageRes)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/GetListWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageServiceClient) GetSingleWebPage(ctx context.Context, in *GetSingleWebPageReq, opts ...grpc.CallOption) (*WebPage, error) {
	out := new(WebPage)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/GetSingleWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageServiceClient) UpdateWebPage(ctx context.Context, in *WebPage, opts ...grpc.CallOption) (*WebPage, error) {
	out := new(WebPage)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/UpdateWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageServiceClient) DeleteWebPage(ctx context.Context, in *DeleteWebPageReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/DeleteWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageServiceClient) RevertWebPage(ctx context.Context, in *RevertWebPageReq, opts ...grpc.CallOption) (*WebPage, error) {
	out := new(WebPage)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/RevertWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageServiceClient) CreateManyWebPage(ctx context.Context, in *ManyVersions, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/CreateManyWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webPageServiceClient) GetWebPageChanges(ctx context.Context, in *GetWebPageHistoryReq, opts ...grpc.CallOption) (*GetWebPageHistoryRes, error) {
	out := new(GetWebPageHistoryRes)
	err := c.cc.Invoke(ctx, "/web_page_service.WebPageService/GetWebPageChanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebPageServiceServer is the server API for WebPageService service.
// All implementations must embed UnimplementedWebPageServiceServer
// for forward compatibility
type WebPageServiceServer interface {
	CreateWebPage(context.Context, *CreateWebPageReq) (*WebPage, error)
	GetListWebPage(context.Context, *GetListWebPageReq) (*GetListWebPageRes, error)
	GetSingleWebPage(context.Context, *GetSingleWebPageReq) (*WebPage, error)
	UpdateWebPage(context.Context, *WebPage) (*WebPage, error)
	DeleteWebPage(context.Context, *DeleteWebPageReq) (*emptypb.Empty, error)
	RevertWebPage(context.Context, *RevertWebPageReq) (*WebPage, error)
	CreateManyWebPage(context.Context, *ManyVersions) (*emptypb.Empty, error)
	GetWebPageChanges(context.Context, *GetWebPageHistoryReq) (*GetWebPageHistoryRes, error)
	mustEmbedUnimplementedWebPageServiceServer()
}

// UnimplementedWebPageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebPageServiceServer struct {
}

func (UnimplementedWebPageServiceServer) CreateWebPage(context.Context, *CreateWebPageReq) (*WebPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWebPage not implemented")
}
func (UnimplementedWebPageServiceServer) GetListWebPage(context.Context, *GetListWebPageReq) (*GetListWebPageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListWebPage not implemented")
}
func (UnimplementedWebPageServiceServer) GetSingleWebPage(context.Context, *GetSingleWebPageReq) (*WebPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleWebPage not implemented")
}
func (UnimplementedWebPageServiceServer) UpdateWebPage(context.Context, *WebPage) (*WebPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWebPage not implemented")
}
func (UnimplementedWebPageServiceServer) DeleteWebPage(context.Context, *DeleteWebPageReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWebPage not implemented")
}
func (UnimplementedWebPageServiceServer) RevertWebPage(context.Context, *RevertWebPageReq) (*WebPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevertWebPage not implemented")
}
func (UnimplementedWebPageServiceServer) CreateManyWebPage(context.Context, *ManyVersions) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManyWebPage not implemented")
}
func (UnimplementedWebPageServiceServer) GetWebPageChanges(context.Context, *GetWebPageHistoryReq) (*GetWebPageHistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebPageChanges not implemented")
}
func (UnimplementedWebPageServiceServer) mustEmbedUnimplementedWebPageServiceServer() {}

// UnsafeWebPageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebPageServiceServer will
// result in compilation errors.
type UnsafeWebPageServiceServer interface {
	mustEmbedUnimplementedWebPageServiceServer()
}

func RegisterWebPageServiceServer(s grpc.ServiceRegistrar, srv WebPageServiceServer) {
	s.RegisterService(&WebPageService_ServiceDesc, srv)
}

func _WebPageService_CreateWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWebPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).CreateWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/CreateWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).CreateWebPage(ctx, req.(*CreateWebPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPageService_GetListWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListWebPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).GetListWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/GetListWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).GetListWebPage(ctx, req.(*GetListWebPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPageService_GetSingleWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleWebPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).GetSingleWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/GetSingleWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).GetSingleWebPage(ctx, req.(*GetSingleWebPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPageService_UpdateWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).UpdateWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/UpdateWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).UpdateWebPage(ctx, req.(*WebPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPageService_DeleteWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWebPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).DeleteWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/DeleteWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).DeleteWebPage(ctx, req.(*DeleteWebPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPageService_RevertWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevertWebPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).RevertWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/RevertWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).RevertWebPage(ctx, req.(*RevertWebPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPageService_CreateManyWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManyVersions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).CreateManyWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/CreateManyWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).CreateManyWebPage(ctx, req.(*ManyVersions))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebPageService_GetWebPageChanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWebPageHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebPageServiceServer).GetWebPageChanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web_page_service.WebPageService/GetWebPageChanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebPageServiceServer).GetWebPageChanges(ctx, req.(*GetWebPageHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WebPageService_ServiceDesc is the grpc.ServiceDesc for WebPageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebPageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "web_page_service.WebPageService",
	HandlerType: (*WebPageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWebPage",
			Handler:    _WebPageService_CreateWebPage_Handler,
		},
		{
			MethodName: "GetListWebPage",
			Handler:    _WebPageService_GetListWebPage_Handler,
		},
		{
			MethodName: "GetSingleWebPage",
			Handler:    _WebPageService_GetSingleWebPage_Handler,
		},
		{
			MethodName: "UpdateWebPage",
			Handler:    _WebPageService_UpdateWebPage_Handler,
		},
		{
			MethodName: "DeleteWebPage",
			Handler:    _WebPageService_DeleteWebPage_Handler,
		},
		{
			MethodName: "RevertWebPage",
			Handler:    _WebPageService_RevertWebPage_Handler,
		},
		{
			MethodName: "CreateManyWebPage",
			Handler:    _WebPageService_CreateManyWebPage_Handler,
		},
		{
			MethodName: "GetWebPageChanges",
			Handler:    _WebPageService_GetWebPageChanges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web_page.proto",
}