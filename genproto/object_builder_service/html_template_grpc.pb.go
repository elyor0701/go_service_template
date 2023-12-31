// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: html_template.proto

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

// HtmlTemplateServiceClient is the client API for HtmlTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HtmlTemplateServiceClient interface {
	Create(ctx context.Context, in *CreateHtmlTemplateRequest, opts ...grpc.CallOption) (*HtmlTemplate, error)
	GetList(ctx context.Context, in *GetAllHtmlTemplateRequest, opts ...grpc.CallOption) (*GetAllHtmlTemplateResponse, error)
	GetSingle(ctx context.Context, in *HtmlTemplatePrimaryKey, opts ...grpc.CallOption) (*HtmlTemplate, error)
	Update(ctx context.Context, in *HtmlTemplate, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *HtmlTemplatePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type htmlTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHtmlTemplateServiceClient(cc grpc.ClientConnInterface) HtmlTemplateServiceClient {
	return &htmlTemplateServiceClient{cc}
}

func (c *htmlTemplateServiceClient) Create(ctx context.Context, in *CreateHtmlTemplateRequest, opts ...grpc.CallOption) (*HtmlTemplate, error) {
	out := new(HtmlTemplate)
	err := c.cc.Invoke(ctx, "/object_builder_service.HtmlTemplateService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *htmlTemplateServiceClient) GetList(ctx context.Context, in *GetAllHtmlTemplateRequest, opts ...grpc.CallOption) (*GetAllHtmlTemplateResponse, error) {
	out := new(GetAllHtmlTemplateResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.HtmlTemplateService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *htmlTemplateServiceClient) GetSingle(ctx context.Context, in *HtmlTemplatePrimaryKey, opts ...grpc.CallOption) (*HtmlTemplate, error) {
	out := new(HtmlTemplate)
	err := c.cc.Invoke(ctx, "/object_builder_service.HtmlTemplateService/GetSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *htmlTemplateServiceClient) Update(ctx context.Context, in *HtmlTemplate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.HtmlTemplateService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *htmlTemplateServiceClient) Delete(ctx context.Context, in *HtmlTemplatePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.HtmlTemplateService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HtmlTemplateServiceServer is the server API for HtmlTemplateService service.
// All implementations must embed UnimplementedHtmlTemplateServiceServer
// for forward compatibility
type HtmlTemplateServiceServer interface {
	Create(context.Context, *CreateHtmlTemplateRequest) (*HtmlTemplate, error)
	GetList(context.Context, *GetAllHtmlTemplateRequest) (*GetAllHtmlTemplateResponse, error)
	GetSingle(context.Context, *HtmlTemplatePrimaryKey) (*HtmlTemplate, error)
	Update(context.Context, *HtmlTemplate) (*emptypb.Empty, error)
	Delete(context.Context, *HtmlTemplatePrimaryKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedHtmlTemplateServiceServer()
}

// UnimplementedHtmlTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHtmlTemplateServiceServer struct {
}

func (UnimplementedHtmlTemplateServiceServer) Create(context.Context, *CreateHtmlTemplateRequest) (*HtmlTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedHtmlTemplateServiceServer) GetList(context.Context, *GetAllHtmlTemplateRequest) (*GetAllHtmlTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedHtmlTemplateServiceServer) GetSingle(context.Context, *HtmlTemplatePrimaryKey) (*HtmlTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingle not implemented")
}
func (UnimplementedHtmlTemplateServiceServer) Update(context.Context, *HtmlTemplate) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedHtmlTemplateServiceServer) Delete(context.Context, *HtmlTemplatePrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedHtmlTemplateServiceServer) mustEmbedUnimplementedHtmlTemplateServiceServer() {}

// UnsafeHtmlTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HtmlTemplateServiceServer will
// result in compilation errors.
type UnsafeHtmlTemplateServiceServer interface {
	mustEmbedUnimplementedHtmlTemplateServiceServer()
}

func RegisterHtmlTemplateServiceServer(s grpc.ServiceRegistrar, srv HtmlTemplateServiceServer) {
	s.RegisterService(&HtmlTemplateService_ServiceDesc, srv)
}

func _HtmlTemplateService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHtmlTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HtmlTemplateServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.HtmlTemplateService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HtmlTemplateServiceServer).Create(ctx, req.(*CreateHtmlTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HtmlTemplateService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllHtmlTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HtmlTemplateServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.HtmlTemplateService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HtmlTemplateServiceServer).GetList(ctx, req.(*GetAllHtmlTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HtmlTemplateService_GetSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HtmlTemplatePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HtmlTemplateServiceServer).GetSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.HtmlTemplateService/GetSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HtmlTemplateServiceServer).GetSingle(ctx, req.(*HtmlTemplatePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _HtmlTemplateService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HtmlTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HtmlTemplateServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.HtmlTemplateService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HtmlTemplateServiceServer).Update(ctx, req.(*HtmlTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _HtmlTemplateService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HtmlTemplatePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HtmlTemplateServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.HtmlTemplateService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HtmlTemplateServiceServer).Delete(ctx, req.(*HtmlTemplatePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// HtmlTemplateService_ServiceDesc is the grpc.ServiceDesc for HtmlTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HtmlTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.HtmlTemplateService",
	HandlerType: (*HtmlTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _HtmlTemplateService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _HtmlTemplateService_GetList_Handler,
		},
		{
			MethodName: "GetSingle",
			Handler:    _HtmlTemplateService_GetSingle_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _HtmlTemplateService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _HtmlTemplateService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "html_template.proto",
}
