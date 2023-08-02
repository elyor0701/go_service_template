// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: postgres_excel.proto

package postgres_object_builder_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExcelServiceClient is the client API for ExcelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExcelServiceClient interface {
	ExcelRead(ctx context.Context, in *ExcelReadRequest, opts ...grpc.CallOption) (*ExcelReadResponse, error)
	ExcelToDb(ctx context.Context, in *ExcelToDbRequest, opts ...grpc.CallOption) (*ExcelToDbResponse, error)
}

type excelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExcelServiceClient(cc grpc.ClientConnInterface) ExcelServiceClient {
	return &excelServiceClient{cc}
}

func (c *excelServiceClient) ExcelRead(ctx context.Context, in *ExcelReadRequest, opts ...grpc.CallOption) (*ExcelReadResponse, error) {
	out := new(ExcelReadResponse)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ExcelService/ExcelRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *excelServiceClient) ExcelToDb(ctx context.Context, in *ExcelToDbRequest, opts ...grpc.CallOption) (*ExcelToDbResponse, error) {
	out := new(ExcelToDbResponse)
	err := c.cc.Invoke(ctx, "/postgres_object_builder_service.ExcelService/ExcelToDb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExcelServiceServer is the server API for ExcelService service.
// All implementations must embed UnimplementedExcelServiceServer
// for forward compatibility
type ExcelServiceServer interface {
	ExcelRead(context.Context, *ExcelReadRequest) (*ExcelReadResponse, error)
	ExcelToDb(context.Context, *ExcelToDbRequest) (*ExcelToDbResponse, error)
	mustEmbedUnimplementedExcelServiceServer()
}

// UnimplementedExcelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExcelServiceServer struct {
}

func (UnimplementedExcelServiceServer) ExcelRead(context.Context, *ExcelReadRequest) (*ExcelReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExcelRead not implemented")
}
func (UnimplementedExcelServiceServer) ExcelToDb(context.Context, *ExcelToDbRequest) (*ExcelToDbResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExcelToDb not implemented")
}
func (UnimplementedExcelServiceServer) mustEmbedUnimplementedExcelServiceServer() {}

// UnsafeExcelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExcelServiceServer will
// result in compilation errors.
type UnsafeExcelServiceServer interface {
	mustEmbedUnimplementedExcelServiceServer()
}

func RegisterExcelServiceServer(s grpc.ServiceRegistrar, srv ExcelServiceServer) {
	s.RegisterService(&ExcelService_ServiceDesc, srv)
}

func _ExcelService_ExcelRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExcelReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExcelServiceServer).ExcelRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ExcelService/ExcelRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExcelServiceServer).ExcelRead(ctx, req.(*ExcelReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExcelService_ExcelToDb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExcelToDbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExcelServiceServer).ExcelToDb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgres_object_builder_service.ExcelService/ExcelToDb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExcelServiceServer).ExcelToDb(ctx, req.(*ExcelToDbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExcelService_ServiceDesc is the grpc.ServiceDesc for ExcelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExcelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postgres_object_builder_service.ExcelService",
	HandlerType: (*ExcelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExcelRead",
			Handler:    _ExcelService_ExcelRead_Handler,
		},
		{
			MethodName: "ExcelToDb",
			Handler:    _ExcelService_ExcelToDb_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postgres_excel.proto",
}
