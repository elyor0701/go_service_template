// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: actMongo.proto

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

// ActMongoServiceClient is the client API for ActMongoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActMongoServiceClient interface {
	Upsert(ctx context.Context, in *MongoActResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*MongoActResponse, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type actMongoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActMongoServiceClient(cc grpc.ClientConnInterface) ActMongoServiceClient {
	return &actMongoServiceClient{cc}
}

func (c *actMongoServiceClient) Upsert(ctx context.Context, in *MongoActResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.ActMongoService/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actMongoServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*MongoActResponse, error) {
	out := new(MongoActResponse)
	err := c.cc.Invoke(ctx, "/mongo_document_service.ActMongoService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actMongoServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.ActMongoService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actMongoServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.ActMongoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActMongoServiceServer is the server API for ActMongoService service.
// All implementations must embed UnimplementedActMongoServiceServer
// for forward compatibility
type ActMongoServiceServer interface {
	Upsert(context.Context, *MongoActResponse) (*emptypb.Empty, error)
	GetById(context.Context, *ById) (*MongoActResponse, error)
	UpdateStatus(context.Context, *UpdateStatusReq) (*emptypb.Empty, error)
	Delete(context.Context, *ById) (*emptypb.Empty, error)
	mustEmbedUnimplementedActMongoServiceServer()
}

// UnimplementedActMongoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActMongoServiceServer struct {
}

func (UnimplementedActMongoServiceServer) Upsert(context.Context, *MongoActResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedActMongoServiceServer) GetById(context.Context, *ById) (*MongoActResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedActMongoServiceServer) UpdateStatus(context.Context, *UpdateStatusReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedActMongoServiceServer) Delete(context.Context, *ById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedActMongoServiceServer) mustEmbedUnimplementedActMongoServiceServer() {}

// UnsafeActMongoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActMongoServiceServer will
// result in compilation errors.
type UnsafeActMongoServiceServer interface {
	mustEmbedUnimplementedActMongoServiceServer()
}

func RegisterActMongoServiceServer(s grpc.ServiceRegistrar, srv ActMongoServiceServer) {
	s.RegisterService(&ActMongoService_ServiceDesc, srv)
}

func _ActMongoService_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MongoActResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActMongoServiceServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.ActMongoService/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActMongoServiceServer).Upsert(ctx, req.(*MongoActResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActMongoService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActMongoServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.ActMongoService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActMongoServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActMongoService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActMongoServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.ActMongoService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActMongoServiceServer).UpdateStatus(ctx, req.(*UpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActMongoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActMongoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.ActMongoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActMongoServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// ActMongoService_ServiceDesc is the grpc.ServiceDesc for ActMongoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActMongoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongo_document_service.ActMongoService",
	HandlerType: (*ActMongoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _ActMongoService_Upsert_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ActMongoService_GetById_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _ActMongoService_UpdateStatus_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ActMongoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "actMongo.proto",
}
