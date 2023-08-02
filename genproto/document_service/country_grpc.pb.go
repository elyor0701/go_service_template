// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: country.proto

package document_service

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

// CountryServiceClient is the client API for CountryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountryServiceClient interface {
	Create(ctx context.Context, in *CreateCountry, opts ...grpc.CallOption) (*Country, error)
	GetList(ctx context.Context, in *GetListCountriesRequest, opts ...grpc.CallOption) (*GetListCountriesResponse, error)
	Get(ctx context.Context, in *CountryId, opts ...grpc.CallOption) (*Country, error)
	Reload(ctx context.Context, in *ReloadCountriesListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountry, error)
}

type countryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountryServiceClient(cc grpc.ClientConnInterface) CountryServiceClient {
	return &countryServiceClient{cc}
}

func (c *countryServiceClient) Create(ctx context.Context, in *CreateCountry, opts ...grpc.CallOption) (*Country, error) {
	out := new(Country)
	err := c.cc.Invoke(ctx, "/document_service.CountryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) GetList(ctx context.Context, in *GetListCountriesRequest, opts ...grpc.CallOption) (*GetListCountriesResponse, error) {
	out := new(GetListCountriesResponse)
	err := c.cc.Invoke(ctx, "/document_service.CountryService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) Get(ctx context.Context, in *CountryId, opts ...grpc.CallOption) (*Country, error) {
	out := new(Country)
	err := c.cc.Invoke(ctx, "/document_service.CountryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) Reload(ctx context.Context, in *ReloadCountriesListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/document_service.CountryService/Reload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) Delete(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountry, error) {
	out := new(DeleteCountry)
	err := c.cc.Invoke(ctx, "/document_service.CountryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryServiceServer is the server API for CountryService service.
// All implementations must embed UnimplementedCountryServiceServer
// for forward compatibility
type CountryServiceServer interface {
	Create(context.Context, *CreateCountry) (*Country, error)
	GetList(context.Context, *GetListCountriesRequest) (*GetListCountriesResponse, error)
	Get(context.Context, *CountryId) (*Country, error)
	Reload(context.Context, *ReloadCountriesListRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteCountryRequest) (*DeleteCountry, error)
	mustEmbedUnimplementedCountryServiceServer()
}

// UnimplementedCountryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCountryServiceServer struct {
}

func (UnimplementedCountryServiceServer) Create(context.Context, *CreateCountry) (*Country, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCountryServiceServer) GetList(context.Context, *GetListCountriesRequest) (*GetListCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedCountryServiceServer) Get(context.Context, *CountryId) (*Country, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCountryServiceServer) Reload(context.Context, *ReloadCountriesListRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reload not implemented")
}
func (UnimplementedCountryServiceServer) Delete(context.Context, *DeleteCountryRequest) (*DeleteCountry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCountryServiceServer) mustEmbedUnimplementedCountryServiceServer() {}

// UnsafeCountryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountryServiceServer will
// result in compilation errors.
type UnsafeCountryServiceServer interface {
	mustEmbedUnimplementedCountryServiceServer()
}

func RegisterCountryServiceServer(s grpc.ServiceRegistrar, srv CountryServiceServer) {
	s.RegisterService(&CountryService_ServiceDesc, srv)
}

func _CountryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCountry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.CountryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Create(ctx, req.(*CreateCountry))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.CountryService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).GetList(ctx, req.(*GetListCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.CountryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Get(ctx, req.(*CountryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadCountriesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.CountryService/Reload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Reload(ctx, req.(*ReloadCountriesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.CountryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Delete(ctx, req.(*DeleteCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CountryService_ServiceDesc is the grpc.ServiceDesc for CountryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "document_service.CountryService",
	HandlerType: (*CountryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CountryService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _CountryService_GetList_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CountryService_Get_Handler,
		},
		{
			MethodName: "Reload",
			Handler:    _CountryService_Reload_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CountryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "country.proto",
}
