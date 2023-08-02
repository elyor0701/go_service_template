// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: postgres_view.proto

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

// ViewServiceClient is the client API for ViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewServiceClient interface {
	Create(ctx context.Context, in *CreateViewRequest, opts ...grpc.CallOption) (*View, error)
	GetList(ctx context.Context, in *GetAllViewsRequest, opts ...grpc.CallOption) (*GetAllViewsResponse, error)
	GetSingle(ctx context.Context, in *ViewPrimaryKey, opts ...grpc.CallOption) (*View, error)
	Update(ctx context.Context, in *View, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ViewPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ConvertHtmlToPdf(ctx context.Context, in *HtmlBody, opts ...grpc.CallOption) (*PdfBody, error)
	ConvertTemplateToHtml(ctx context.Context, in *HtmlBody, opts ...grpc.CallOption) (*HtmlBody, error)
}

type viewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViewServiceClient(cc grpc.ClientConnInterface) ViewServiceClient {
	return &viewServiceClient{cc}
}

func (c *viewServiceClient) Create(ctx context.Context, in *CreateViewRequest, opts ...grpc.CallOption) (*View, error) {
	out := new(View)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ViewService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetList(ctx context.Context, in *GetAllViewsRequest, opts ...grpc.CallOption) (*GetAllViewsResponse, error) {
	out := new(GetAllViewsResponse)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ViewService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetSingle(ctx context.Context, in *ViewPrimaryKey, opts ...grpc.CallOption) (*View, error) {
	out := new(View)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ViewService/GetSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) Update(ctx context.Context, in *View, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ViewService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) Delete(ctx context.Context, in *ViewPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ViewService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) ConvertHtmlToPdf(ctx context.Context, in *HtmlBody, opts ...grpc.CallOption) (*PdfBody, error) {
	out := new(PdfBody)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ViewService/ConvertHtmlToPdf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) ConvertTemplateToHtml(ctx context.Context, in *HtmlBody, opts ...grpc.CallOption) (*HtmlBody, error) {
	out := new(HtmlBody)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ViewService/ConvertTemplateToHtml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewServiceServer is the server API for ViewService service.
// All implementations must embed UnimplementedViewServiceServer
// for forward compatibility
type ViewServiceServer interface {
	Create(context.Context, *CreateViewRequest) (*View, error)
	GetList(context.Context, *GetAllViewsRequest) (*GetAllViewsResponse, error)
	GetSingle(context.Context, *ViewPrimaryKey) (*View, error)
	Update(context.Context, *View) (*emptypb.Empty, error)
	Delete(context.Context, *ViewPrimaryKey) (*emptypb.Empty, error)
	ConvertHtmlToPdf(context.Context, *HtmlBody) (*PdfBody, error)
	ConvertTemplateToHtml(context.Context, *HtmlBody) (*HtmlBody, error)
	mustEmbedUnimplementedViewServiceServer()
}

// UnimplementedViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedViewServiceServer struct {
}

func (UnimplementedViewServiceServer) Create(context.Context, *CreateViewRequest) (*View, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedViewServiceServer) GetList(context.Context, *GetAllViewsRequest) (*GetAllViewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedViewServiceServer) GetSingle(context.Context, *ViewPrimaryKey) (*View, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingle not implemented")
}
func (UnimplementedViewServiceServer) Update(context.Context, *View) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedViewServiceServer) Delete(context.Context, *ViewPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedViewServiceServer) ConvertHtmlToPdf(context.Context, *HtmlBody) (*PdfBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertHtmlToPdf not implemented")
}
func (UnimplementedViewServiceServer) ConvertTemplateToHtml(context.Context, *HtmlBody) (*HtmlBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertTemplateToHtml not implemented")
}
func (UnimplementedViewServiceServer) mustEmbedUnimplementedViewServiceServer() {}

// UnsafeViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewServiceServer will
// result in compilation errors.
type UnsafeViewServiceServer interface {
	mustEmbedUnimplementedViewServiceServer()
}

func RegisterViewServiceServer(s grpc.ServiceRegistrar, srv ViewServiceServer) {
	s.RegisterService(&ViewService_ServiceDesc, srv)
}

func _ViewService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ViewService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).Create(ctx, req.(*CreateViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllViewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ViewService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetList(ctx, req.(*GetAllViewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ViewService/GetSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetSingle(ctx, req.(*ViewPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ViewService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).Update(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ViewService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).Delete(ctx, req.(*ViewPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_ConvertHtmlToPdf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HtmlBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).ConvertHtmlToPdf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ViewService/ConvertHtmlToPdf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).ConvertHtmlToPdf(ctx, req.(*HtmlBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_ConvertTemplateToHtml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HtmlBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).ConvertTemplateToHtml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ViewService/ConvertTemplateToHtml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).ConvertTemplateToHtml(ctx, req.(*HtmlBody))
	}
	return interceptor(ctx, in, info, handler)
}

// ViewService_ServiceDesc is the grpc.ServiceDesc for ViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postgres_object_builder_service.ViewService",
	HandlerType: (*ViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ViewService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ViewService_GetList_Handler,
		},
		{
			MethodName: "GetSingle",
			Handler:    _ViewService_GetSingle_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ViewService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ViewService_Delete_Handler,
		},
		{
			MethodName: "ConvertHtmlToPdf",
			Handler:    _ViewService_ConvertHtmlToPdf_Handler,
		},
		{
			MethodName: "ConvertTemplateToHtml",
			Handler:    _ViewService_ConvertTemplateToHtml_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postgres_view.proto",
}
