// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: chat.proto

package chat_service

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	GetChatByChatId(ctx context.Context, in *GetChatByChatIdRequest, opts ...grpc.CallOption) (*GetChatByChatIdResponse, error)
	GetChatByUserId(ctx context.Context, in *GetChatByUserIdRequest, opts ...grpc.CallOption) (*GetChatByUserIdResponse, error)
	GetChatList(ctx context.Context, in *GetChatListRequest, opts ...grpc.CallOption) (*GetChatListResponse, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateChat(ctx context.Context, in *UpdateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateChatProfielPhotoUrl(ctx context.Context, in *UpdateProfilePhotoUrlRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateChatMessage(ctx context.Context, in *UpdateChatMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateBot(ctx context.Context, in *CreateBotRequest, opts ...grpc.CallOption) (*CreateBotResponse, error)
	UpdateBotStatus(ctx context.Context, in *UpdateBotStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateBotToken(ctx context.Context, in *UpdateBotTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetBotTokenList(ctx context.Context, in *GetBotTokenListRequest, opts ...grpc.CallOption) (*GetBotTokenListResponse, error)
	GetBotTokenByBotId(ctx context.Context, in *GetBotByBotIdRequest, opts ...grpc.CallOption) (*GetBotByBotIdResponse, error)
	GetBotIDByChatId(ctx context.Context, in *GetBotIDByChatIDRequest, opts ...grpc.CallOption) (*GetBotIDByChatIDResponse, error)
	DeleteBotToken(ctx context.Context, in *DeleteBotTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, "/ChatService/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatByChatId(ctx context.Context, in *GetChatByChatIdRequest, opts ...grpc.CallOption) (*GetChatByChatIdResponse, error) {
	out := new(GetChatByChatIdResponse)
	err := c.cc.Invoke(ctx, "/ChatService/GetChatByChatId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatByUserId(ctx context.Context, in *GetChatByUserIdRequest, opts ...grpc.CallOption) (*GetChatByUserIdResponse, error) {
	out := new(GetChatByUserIdResponse)
	err := c.cc.Invoke(ctx, "/ChatService/GetChatByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatList(ctx context.Context, in *GetChatListRequest, opts ...grpc.CallOption) (*GetChatListResponse, error) {
	out := new(GetChatListResponse)
	err := c.cc.Invoke(ctx, "/ChatService/GetChatList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateChat(ctx context.Context, in *UpdateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatService/UpdateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateChatProfielPhotoUrl(ctx context.Context, in *UpdateProfilePhotoUrlRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatService/UpdateChatProfielPhotoUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateChatMessage(ctx context.Context, in *UpdateChatMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatService/UpdateChatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateBot(ctx context.Context, in *CreateBotRequest, opts ...grpc.CallOption) (*CreateBotResponse, error) {
	out := new(CreateBotResponse)
	err := c.cc.Invoke(ctx, "/ChatService/CreateBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateBotStatus(ctx context.Context, in *UpdateBotStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatService/UpdateBotStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateBotToken(ctx context.Context, in *UpdateBotTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatService/UpdateBotToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetBotTokenList(ctx context.Context, in *GetBotTokenListRequest, opts ...grpc.CallOption) (*GetBotTokenListResponse, error) {
	out := new(GetBotTokenListResponse)
	err := c.cc.Invoke(ctx, "/ChatService/GetBotTokenList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetBotTokenByBotId(ctx context.Context, in *GetBotByBotIdRequest, opts ...grpc.CallOption) (*GetBotByBotIdResponse, error) {
	out := new(GetBotByBotIdResponse)
	err := c.cc.Invoke(ctx, "/ChatService/GetBotTokenByBotId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetBotIDByChatId(ctx context.Context, in *GetBotIDByChatIDRequest, opts ...grpc.CallOption) (*GetBotIDByChatIDResponse, error) {
	out := new(GetBotIDByChatIDResponse)
	err := c.cc.Invoke(ctx, "/ChatService/GetBotIDByChatId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteBotToken(ctx context.Context, in *DeleteBotTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatService/DeleteBotToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	GetChatByChatId(context.Context, *GetChatByChatIdRequest) (*GetChatByChatIdResponse, error)
	GetChatByUserId(context.Context, *GetChatByUserIdRequest) (*GetChatByUserIdResponse, error)
	GetChatList(context.Context, *GetChatListRequest) (*GetChatListResponse, error)
	CreateMessage(context.Context, *CreateMessageRequest) (*emptypb.Empty, error)
	UpdateChat(context.Context, *UpdateChatRequest) (*emptypb.Empty, error)
	UpdateChatProfielPhotoUrl(context.Context, *UpdateProfilePhotoUrlRequest) (*emptypb.Empty, error)
	UpdateChatMessage(context.Context, *UpdateChatMessageRequest) (*emptypb.Empty, error)
	CreateBot(context.Context, *CreateBotRequest) (*CreateBotResponse, error)
	UpdateBotStatus(context.Context, *UpdateBotStatusRequest) (*emptypb.Empty, error)
	UpdateBotToken(context.Context, *UpdateBotTokenRequest) (*emptypb.Empty, error)
	GetBotTokenList(context.Context, *GetBotTokenListRequest) (*GetBotTokenListResponse, error)
	GetBotTokenByBotId(context.Context, *GetBotByBotIdRequest) (*GetBotByBotIdResponse, error)
	GetBotIDByChatId(context.Context, *GetBotIDByChatIDRequest) (*GetBotIDByChatIDResponse, error)
	DeleteBotToken(context.Context, *DeleteBotTokenRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) GetChatByChatId(context.Context, *GetChatByChatIdRequest) (*GetChatByChatIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatByChatId not implemented")
}
func (UnimplementedChatServiceServer) GetChatByUserId(context.Context, *GetChatByUserIdRequest) (*GetChatByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatByUserId not implemented")
}
func (UnimplementedChatServiceServer) GetChatList(context.Context, *GetChatListRequest) (*GetChatListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatList not implemented")
}
func (UnimplementedChatServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedChatServiceServer) UpdateChat(context.Context, *UpdateChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChat not implemented")
}
func (UnimplementedChatServiceServer) UpdateChatProfielPhotoUrl(context.Context, *UpdateProfilePhotoUrlRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatProfielPhotoUrl not implemented")
}
func (UnimplementedChatServiceServer) UpdateChatMessage(context.Context, *UpdateChatMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatMessage not implemented")
}
func (UnimplementedChatServiceServer) CreateBot(context.Context, *CreateBotRequest) (*CreateBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBot not implemented")
}
func (UnimplementedChatServiceServer) UpdateBotStatus(context.Context, *UpdateBotStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBotStatus not implemented")
}
func (UnimplementedChatServiceServer) UpdateBotToken(context.Context, *UpdateBotTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBotToken not implemented")
}
func (UnimplementedChatServiceServer) GetBotTokenList(context.Context, *GetBotTokenListRequest) (*GetBotTokenListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBotTokenList not implemented")
}
func (UnimplementedChatServiceServer) GetBotTokenByBotId(context.Context, *GetBotByBotIdRequest) (*GetBotByBotIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBotTokenByBotId not implemented")
}
func (UnimplementedChatServiceServer) GetBotIDByChatId(context.Context, *GetBotIDByChatIDRequest) (*GetBotIDByChatIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBotIDByChatId not implemented")
}
func (UnimplementedChatServiceServer) DeleteBotToken(context.Context, *DeleteBotTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBotToken not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatByChatId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatByChatIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatByChatId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/GetChatByChatId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatByChatId(ctx, req.(*GetChatByChatIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/GetChatByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatByUserId(ctx, req.(*GetChatByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/GetChatList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatList(ctx, req.(*GetChatListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/UpdateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateChat(ctx, req.(*UpdateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateChatProfielPhotoUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfilePhotoUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateChatProfielPhotoUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/UpdateChatProfielPhotoUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateChatProfielPhotoUrl(ctx, req.(*UpdateProfilePhotoUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/UpdateChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateChatMessage(ctx, req.(*UpdateChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/CreateBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateBot(ctx, req.(*CreateBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateBotStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBotStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateBotStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/UpdateBotStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateBotStatus(ctx, req.(*UpdateBotStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateBotToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBotTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateBotToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/UpdateBotToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateBotToken(ctx, req.(*UpdateBotTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetBotTokenList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotTokenListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetBotTokenList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/GetBotTokenList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetBotTokenList(ctx, req.(*GetBotTokenListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetBotTokenByBotId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotByBotIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetBotTokenByBotId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/GetBotTokenByBotId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetBotTokenByBotId(ctx, req.(*GetBotByBotIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetBotIDByChatId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotIDByChatIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetBotIDByChatId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/GetBotIDByChatId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetBotIDByChatId(ctx, req.(*GetBotIDByChatIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteBotToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBotTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteBotToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/DeleteBotToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteBotToken(ctx, req.(*DeleteBotTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "GetChatByChatId",
			Handler:    _ChatService_GetChatByChatId_Handler,
		},
		{
			MethodName: "GetChatByUserId",
			Handler:    _ChatService_GetChatByUserId_Handler,
		},
		{
			MethodName: "GetChatList",
			Handler:    _ChatService_GetChatList_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _ChatService_CreateMessage_Handler,
		},
		{
			MethodName: "UpdateChat",
			Handler:    _ChatService_UpdateChat_Handler,
		},
		{
			MethodName: "UpdateChatProfielPhotoUrl",
			Handler:    _ChatService_UpdateChatProfielPhotoUrl_Handler,
		},
		{
			MethodName: "UpdateChatMessage",
			Handler:    _ChatService_UpdateChatMessage_Handler,
		},
		{
			MethodName: "CreateBot",
			Handler:    _ChatService_CreateBot_Handler,
		},
		{
			MethodName: "UpdateBotStatus",
			Handler:    _ChatService_UpdateBotStatus_Handler,
		},
		{
			MethodName: "UpdateBotToken",
			Handler:    _ChatService_UpdateBotToken_Handler,
		},
		{
			MethodName: "GetBotTokenList",
			Handler:    _ChatService_GetBotTokenList_Handler,
		},
		{
			MethodName: "GetBotTokenByBotId",
			Handler:    _ChatService_GetBotTokenByBotId_Handler,
		},
		{
			MethodName: "GetBotIDByChatId",
			Handler:    _ChatService_GetBotIDByChatId_Handler,
		},
		{
			MethodName: "DeleteBotToken",
			Handler:    _ChatService_DeleteBotToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
