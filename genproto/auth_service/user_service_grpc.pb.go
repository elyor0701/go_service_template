// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: user_service.proto

package auth_service

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*User, error)
	GetUserListByIDs(ctx context.Context, in *UserPrimaryKeyList, opts ...grpc.CallOption) (*GetUserListResponse, error)
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*User, error)
	SendMessageToEmail(ctx context.Context, in *SendMessageToEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddUserRelation(ctx context.Context, in *AddUserRelationRequest, opts ...grpc.CallOption) (*UserRelation, error)
	RemoveUserRelation(ctx context.Context, in *UserRelationPrimaryKey, opts ...grpc.CallOption) (*UserRelation, error)
	UpsertUserInfo(ctx context.Context, in *UpsertUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error)
	V2CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	V2GetUserByID(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*User, error)
	V2GetUserListByIDs(ctx context.Context, in *UserPrimaryKeyList, opts ...grpc.CallOption) (*GetUserListResponse, error)
	V2GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
	V2UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	V2DeleteUser(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddUserToProject(ctx context.Context, in *AddUserToProjectReq, opts ...grpc.CallOption) (*AddUserToProjectRes, error)
	GetProjectsByUserId(ctx context.Context, in *GetProjectsByUserIdReq, opts ...grpc.CallOption) (*GetProjectsByUserIdRes, error)
	V2GetUserByLoginTypes(ctx context.Context, in *GetUserByLoginTypesRequest, opts ...grpc.CallOption) (*GetUserByLoginTypesResponse, error)
	RegisterUserViaEmail(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	RegisterWithGoogle(ctx context.Context, in *RegisterWithGoogleRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByID(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserListByIDs(ctx context.Context, in *UserPrimaryKeyList, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/GetUserListByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendMessageToEmail(ctx context.Context, in *SendMessageToEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/SendMessageToEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUserRelation(ctx context.Context, in *AddUserRelationRequest, opts ...grpc.CallOption) (*UserRelation, error) {
	out := new(UserRelation)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/AddUserRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveUserRelation(ctx context.Context, in *UserRelationPrimaryKey, opts ...grpc.CallOption) (*UserRelation, error) {
	out := new(UserRelation)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/RemoveUserRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpsertUserInfo(ctx context.Context, in *UpsertUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/UpsertUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) V2CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/V2CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) V2GetUserByID(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/V2GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) V2GetUserListByIDs(ctx context.Context, in *UserPrimaryKeyList, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/V2GetUserListByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) V2GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/V2GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) V2UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/V2UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) V2DeleteUser(ctx context.Context, in *UserPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/V2DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUserToProject(ctx context.Context, in *AddUserToProjectReq, opts ...grpc.CallOption) (*AddUserToProjectRes, error) {
	out := new(AddUserToProjectRes)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/AddUserToProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProjectsByUserId(ctx context.Context, in *GetProjectsByUserIdReq, opts ...grpc.CallOption) (*GetProjectsByUserIdRes, error) {
	out := new(GetProjectsByUserIdRes)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/GetProjectsByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) V2GetUserByLoginTypes(ctx context.Context, in *GetUserByLoginTypesRequest, opts ...grpc.CallOption) (*GetUserByLoginTypesResponse, error) {
	out := new(GetUserByLoginTypesResponse)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/V2GetUserByLoginTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterUserViaEmail(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/RegisterUserViaEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterWithGoogle(ctx context.Context, in *RegisterWithGoogleRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth_service.UserService/RegisterWithGoogle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	GetUserByID(context.Context, *UserPrimaryKey) (*User, error)
	GetUserListByIDs(context.Context, *UserPrimaryKeyList) (*GetUserListResponse, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	DeleteUser(context.Context, *UserPrimaryKey) (*emptypb.Empty, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*User, error)
	SendMessageToEmail(context.Context, *SendMessageToEmailRequest) (*emptypb.Empty, error)
	AddUserRelation(context.Context, *AddUserRelationRequest) (*UserRelation, error)
	RemoveUserRelation(context.Context, *UserRelationPrimaryKey) (*UserRelation, error)
	UpsertUserInfo(context.Context, *UpsertUserInfoRequest) (*UserInfo, error)
	V2CreateUser(context.Context, *CreateUserRequest) (*User, error)
	V2GetUserByID(context.Context, *UserPrimaryKey) (*User, error)
	V2GetUserListByIDs(context.Context, *UserPrimaryKeyList) (*GetUserListResponse, error)
	V2GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	V2UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	V2DeleteUser(context.Context, *UserPrimaryKey) (*emptypb.Empty, error)
	AddUserToProject(context.Context, *AddUserToProjectReq) (*AddUserToProjectRes, error)
	GetProjectsByUserId(context.Context, *GetProjectsByUserIdReq) (*GetProjectsByUserIdRes, error)
	V2GetUserByLoginTypes(context.Context, *GetUserByLoginTypesRequest) (*GetUserByLoginTypesResponse, error)
	RegisterUserViaEmail(context.Context, *CreateUserRequest) (*User, error)
	RegisterWithGoogle(context.Context, *RegisterWithGoogleRequest) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByID(context.Context, *UserPrimaryKey) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedUserServiceServer) GetUserListByIDs(context.Context, *UserPrimaryKeyList) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserListByIDs not implemented")
}
func (UnimplementedUserServiceServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *UserPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) SendMessageToEmail(context.Context, *SendMessageToEmailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToEmail not implemented")
}
func (UnimplementedUserServiceServer) AddUserRelation(context.Context, *AddUserRelationRequest) (*UserRelation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserRelation not implemented")
}
func (UnimplementedUserServiceServer) RemoveUserRelation(context.Context, *UserRelationPrimaryKey) (*UserRelation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserRelation not implemented")
}
func (UnimplementedUserServiceServer) UpsertUserInfo(context.Context, *UpsertUserInfoRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUserInfo not implemented")
}
func (UnimplementedUserServiceServer) V2CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2CreateUser not implemented")
}
func (UnimplementedUserServiceServer) V2GetUserByID(context.Context, *UserPrimaryKey) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2GetUserByID not implemented")
}
func (UnimplementedUserServiceServer) V2GetUserListByIDs(context.Context, *UserPrimaryKeyList) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2GetUserListByIDs not implemented")
}
func (UnimplementedUserServiceServer) V2GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2GetUserList not implemented")
}
func (UnimplementedUserServiceServer) V2UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) V2DeleteUser(context.Context, *UserPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) AddUserToProject(context.Context, *AddUserToProjectReq) (*AddUserToProjectRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToProject not implemented")
}
func (UnimplementedUserServiceServer) GetProjectsByUserId(context.Context, *GetProjectsByUserIdReq) (*GetProjectsByUserIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectsByUserId not implemented")
}
func (UnimplementedUserServiceServer) V2GetUserByLoginTypes(context.Context, *GetUserByLoginTypesRequest) (*GetUserByLoginTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2GetUserByLoginTypes not implemented")
}
func (UnimplementedUserServiceServer) RegisterUserViaEmail(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserViaEmail not implemented")
}
func (UnimplementedUserServiceServer) RegisterWithGoogle(context.Context, *RegisterWithGoogleRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWithGoogle not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByID(ctx, req.(*UserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserListByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKeyList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserListByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/GetUserListByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserListByIDs(ctx, req.(*UserPrimaryKeyList))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*UserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendMessageToEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendMessageToEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/SendMessageToEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendMessageToEmail(ctx, req.(*SendMessageToEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddUserRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUserRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/AddUserRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUserRelation(ctx, req.(*AddUserRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveUserRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRelationPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveUserRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/RemoveUserRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveUserRelation(ctx, req.(*UserRelationPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpsertUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpsertUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/UpsertUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpsertUserInfo(ctx, req.(*UpsertUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_V2CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).V2CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/V2CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).V2CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_V2GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).V2GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/V2GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).V2GetUserByID(ctx, req.(*UserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_V2GetUserListByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKeyList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).V2GetUserListByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/V2GetUserListByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).V2GetUserListByIDs(ctx, req.(*UserPrimaryKeyList))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_V2GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).V2GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/V2GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).V2GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_V2UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).V2UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/V2UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).V2UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_V2DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).V2DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/V2DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).V2DeleteUser(ctx, req.(*UserPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddUserToProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUserToProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/AddUserToProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUserToProject(ctx, req.(*AddUserToProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProjectsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectsByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProjectsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/GetProjectsByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProjectsByUserId(ctx, req.(*GetProjectsByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_V2GetUserByLoginTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByLoginTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).V2GetUserByLoginTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/V2GetUserByLoginTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).V2GetUserByLoginTypes(ctx, req.(*GetUserByLoginTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterUserViaEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUserViaEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/RegisterUserViaEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUserViaEmail(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterWithGoogle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWithGoogleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterWithGoogle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.UserService/RegisterWithGoogle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterWithGoogle(ctx, req.(*RegisterWithGoogleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserService_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserListByIDs",
			Handler:    _UserService_GetUserListByIDs_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _UserService_GetUserList_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "SendMessageToEmail",
			Handler:    _UserService_SendMessageToEmail_Handler,
		},
		{
			MethodName: "AddUserRelation",
			Handler:    _UserService_AddUserRelation_Handler,
		},
		{
			MethodName: "RemoveUserRelation",
			Handler:    _UserService_RemoveUserRelation_Handler,
		},
		{
			MethodName: "UpsertUserInfo",
			Handler:    _UserService_UpsertUserInfo_Handler,
		},
		{
			MethodName: "V2CreateUser",
			Handler:    _UserService_V2CreateUser_Handler,
		},
		{
			MethodName: "V2GetUserByID",
			Handler:    _UserService_V2GetUserByID_Handler,
		},
		{
			MethodName: "V2GetUserListByIDs",
			Handler:    _UserService_V2GetUserListByIDs_Handler,
		},
		{
			MethodName: "V2GetUserList",
			Handler:    _UserService_V2GetUserList_Handler,
		},
		{
			MethodName: "V2UpdateUser",
			Handler:    _UserService_V2UpdateUser_Handler,
		},
		{
			MethodName: "V2DeleteUser",
			Handler:    _UserService_V2DeleteUser_Handler,
		},
		{
			MethodName: "AddUserToProject",
			Handler:    _UserService_AddUserToProject_Handler,
		},
		{
			MethodName: "GetProjectsByUserId",
			Handler:    _UserService_GetProjectsByUserId_Handler,
		},
		{
			MethodName: "V2GetUserByLoginTypes",
			Handler:    _UserService_V2GetUserByLoginTypes_Handler,
		},
		{
			MethodName: "RegisterUserViaEmail",
			Handler:    _UserService_RegisterUserViaEmail_Handler,
		},
		{
			MethodName: "RegisterWithGoogle",
			Handler:    _UserService_RegisterWithGoogle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}