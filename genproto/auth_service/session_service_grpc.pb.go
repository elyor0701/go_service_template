// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: session_service.proto

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

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	HasAccess(ctx context.Context, in *HasAccessRequest, opts ...grpc.CallOption) (*HasAccessResponse, error)
	HasAccessSuperAdmin(ctx context.Context, in *HasAccessSuperAdminReq, opts ...grpc.CallOption) (*HasAccessSuperAdminRes, error)
	V2Login(ctx context.Context, in *V2LoginRequest, opts ...grpc.CallOption) (*V2LoginResponse, error)
	V2LoginSuperAdmin(ctx context.Context, in *V2LoginSuperAdminReq, opts ...grpc.CallOption) (*V2LoginSuperAdminRes, error)
	V2HasAccess(ctx context.Context, in *HasAccessRequest, opts ...grpc.CallOption) (*HasAccessResponse, error)
	V2HasAccessUser(ctx context.Context, in *V2HasAccessUserReq, opts ...grpc.CallOption) (*V2HasAccessUserRes, error)
	V2RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*V2RefreshTokenResponse, error)
	V2RefreshTokenSuperAdmin(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*V2RefreshTokenSuperAdminResponse, error)
	SessionAndTokenGenerator(ctx context.Context, in *SessionAndTokenRequest, opts ...grpc.CallOption) (*V2LoginResponse, error)
	UpdateSessionsByRoleId(ctx context.Context, in *UpdateSessionByRoleIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MultiCompanyLogin(ctx context.Context, in *MultiCompanyLoginRequest, opts ...grpc.CallOption) (*MultiCompanyLoginResponse, error)
	V2MultiCompanyLogin(ctx context.Context, in *V2MultiCompanyLoginReq, opts ...grpc.CallOption) (*V2MultiCompanyLoginRes, error)
	V2MultiCompanyOneLogin(ctx context.Context, in *V2MultiCompanyLoginReq, opts ...grpc.CallOption) (*V2MultiCompanyOneLoginRes, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) HasAccess(ctx context.Context, in *HasAccessRequest, opts ...grpc.CallOption) (*HasAccessResponse, error) {
	out := new(HasAccessResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/HasAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) HasAccessSuperAdmin(ctx context.Context, in *HasAccessSuperAdminReq, opts ...grpc.CallOption) (*HasAccessSuperAdminRes, error) {
	out := new(HasAccessSuperAdminRes)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/HasAccessSuperAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2Login(ctx context.Context, in *V2LoginRequest, opts ...grpc.CallOption) (*V2LoginResponse, error) {
	out := new(V2LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2LoginSuperAdmin(ctx context.Context, in *V2LoginSuperAdminReq, opts ...grpc.CallOption) (*V2LoginSuperAdminRes, error) {
	out := new(V2LoginSuperAdminRes)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2LoginSuperAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2HasAccess(ctx context.Context, in *HasAccessRequest, opts ...grpc.CallOption) (*HasAccessResponse, error) {
	out := new(HasAccessResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2HasAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2HasAccessUser(ctx context.Context, in *V2HasAccessUserReq, opts ...grpc.CallOption) (*V2HasAccessUserRes, error) {
	out := new(V2HasAccessUserRes)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2HasAccessUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*V2RefreshTokenResponse, error) {
	out := new(V2RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2RefreshTokenSuperAdmin(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*V2RefreshTokenSuperAdminResponse, error) {
	out := new(V2RefreshTokenSuperAdminResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2RefreshTokenSuperAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) SessionAndTokenGenerator(ctx context.Context, in *SessionAndTokenRequest, opts ...grpc.CallOption) (*V2LoginResponse, error) {
	out := new(V2LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/SessionAndTokenGenerator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) UpdateSessionsByRoleId(ctx context.Context, in *UpdateSessionByRoleIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/UpdateSessionsByRoleId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) MultiCompanyLogin(ctx context.Context, in *MultiCompanyLoginRequest, opts ...grpc.CallOption) (*MultiCompanyLoginResponse, error) {
	out := new(MultiCompanyLoginResponse)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/MultiCompanyLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2MultiCompanyLogin(ctx context.Context, in *V2MultiCompanyLoginReq, opts ...grpc.CallOption) (*V2MultiCompanyLoginRes, error) {
	out := new(V2MultiCompanyLoginRes)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2MultiCompanyLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) V2MultiCompanyOneLogin(ctx context.Context, in *V2MultiCompanyLoginReq, opts ...grpc.CallOption) (*V2MultiCompanyOneLoginRes, error) {
	out := new(V2MultiCompanyOneLoginRes)
	err := c.cc.Invoke(ctx, "/auth_service.SessionService/V2MultiCompanyOneLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations must embed UnimplementedSessionServiceServer
// for forward compatibility
type SessionServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	HasAccess(context.Context, *HasAccessRequest) (*HasAccessResponse, error)
	HasAccessSuperAdmin(context.Context, *HasAccessSuperAdminReq) (*HasAccessSuperAdminRes, error)
	V2Login(context.Context, *V2LoginRequest) (*V2LoginResponse, error)
	V2LoginSuperAdmin(context.Context, *V2LoginSuperAdminReq) (*V2LoginSuperAdminRes, error)
	V2HasAccess(context.Context, *HasAccessRequest) (*HasAccessResponse, error)
	V2HasAccessUser(context.Context, *V2HasAccessUserReq) (*V2HasAccessUserRes, error)
	V2RefreshToken(context.Context, *RefreshTokenRequest) (*V2RefreshTokenResponse, error)
	V2RefreshTokenSuperAdmin(context.Context, *RefreshTokenRequest) (*V2RefreshTokenSuperAdminResponse, error)
	SessionAndTokenGenerator(context.Context, *SessionAndTokenRequest) (*V2LoginResponse, error)
	UpdateSessionsByRoleId(context.Context, *UpdateSessionByRoleIdRequest) (*emptypb.Empty, error)
	MultiCompanyLogin(context.Context, *MultiCompanyLoginRequest) (*MultiCompanyLoginResponse, error)
	V2MultiCompanyLogin(context.Context, *V2MultiCompanyLoginReq) (*V2MultiCompanyLoginRes, error)
	V2MultiCompanyOneLogin(context.Context, *V2MultiCompanyLoginReq) (*V2MultiCompanyOneLoginRes, error)
	mustEmbedUnimplementedSessionServiceServer()
}

// UnimplementedSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (UnimplementedSessionServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSessionServiceServer) Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSessionServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedSessionServiceServer) HasAccess(context.Context, *HasAccessRequest) (*HasAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasAccess not implemented")
}
func (UnimplementedSessionServiceServer) HasAccessSuperAdmin(context.Context, *HasAccessSuperAdminReq) (*HasAccessSuperAdminRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasAccessSuperAdmin not implemented")
}
func (UnimplementedSessionServiceServer) V2Login(context.Context, *V2LoginRequest) (*V2LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2Login not implemented")
}
func (UnimplementedSessionServiceServer) V2LoginSuperAdmin(context.Context, *V2LoginSuperAdminReq) (*V2LoginSuperAdminRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2LoginSuperAdmin not implemented")
}
func (UnimplementedSessionServiceServer) V2HasAccess(context.Context, *HasAccessRequest) (*HasAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2HasAccess not implemented")
}
func (UnimplementedSessionServiceServer) V2HasAccessUser(context.Context, *V2HasAccessUserReq) (*V2HasAccessUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2HasAccessUser not implemented")
}
func (UnimplementedSessionServiceServer) V2RefreshToken(context.Context, *RefreshTokenRequest) (*V2RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2RefreshToken not implemented")
}
func (UnimplementedSessionServiceServer) V2RefreshTokenSuperAdmin(context.Context, *RefreshTokenRequest) (*V2RefreshTokenSuperAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2RefreshTokenSuperAdmin not implemented")
}
func (UnimplementedSessionServiceServer) SessionAndTokenGenerator(context.Context, *SessionAndTokenRequest) (*V2LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionAndTokenGenerator not implemented")
}
func (UnimplementedSessionServiceServer) UpdateSessionsByRoleId(context.Context, *UpdateSessionByRoleIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSessionsByRoleId not implemented")
}
func (UnimplementedSessionServiceServer) MultiCompanyLogin(context.Context, *MultiCompanyLoginRequest) (*MultiCompanyLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCompanyLogin not implemented")
}
func (UnimplementedSessionServiceServer) V2MultiCompanyLogin(context.Context, *V2MultiCompanyLoginReq) (*V2MultiCompanyLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2MultiCompanyLogin not implemented")
}
func (UnimplementedSessionServiceServer) V2MultiCompanyOneLogin(context.Context, *V2MultiCompanyLoginReq) (*V2MultiCompanyOneLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V2MultiCompanyOneLogin not implemented")
}
func (UnimplementedSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_HasAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).HasAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/HasAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).HasAccess(ctx, req.(*HasAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_HasAccessSuperAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasAccessSuperAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).HasAccessSuperAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/HasAccessSuperAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).HasAccessSuperAdmin(ctx, req.(*HasAccessSuperAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V2LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2Login(ctx, req.(*V2LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2LoginSuperAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V2LoginSuperAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2LoginSuperAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2LoginSuperAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2LoginSuperAdmin(ctx, req.(*V2LoginSuperAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2HasAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2HasAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2HasAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2HasAccess(ctx, req.(*HasAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2HasAccessUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V2HasAccessUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2HasAccessUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2HasAccessUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2HasAccessUser(ctx, req.(*V2HasAccessUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2RefreshTokenSuperAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2RefreshTokenSuperAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2RefreshTokenSuperAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2RefreshTokenSuperAdmin(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_SessionAndTokenGenerator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionAndTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).SessionAndTokenGenerator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/SessionAndTokenGenerator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).SessionAndTokenGenerator(ctx, req.(*SessionAndTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_UpdateSessionsByRoleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionByRoleIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).UpdateSessionsByRoleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/UpdateSessionsByRoleId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).UpdateSessionsByRoleId(ctx, req.(*UpdateSessionByRoleIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_MultiCompanyLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCompanyLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).MultiCompanyLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/MultiCompanyLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).MultiCompanyLogin(ctx, req.(*MultiCompanyLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2MultiCompanyLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V2MultiCompanyLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2MultiCompanyLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2MultiCompanyLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2MultiCompanyLogin(ctx, req.(*V2MultiCompanyLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_V2MultiCompanyOneLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V2MultiCompanyLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).V2MultiCompanyOneLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.SessionService/V2MultiCompanyOneLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).V2MultiCompanyOneLogin(ctx, req.(*V2MultiCompanyLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SessionService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SessionService_Logout_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _SessionService_RefreshToken_Handler,
		},
		{
			MethodName: "HasAccess",
			Handler:    _SessionService_HasAccess_Handler,
		},
		{
			MethodName: "HasAccessSuperAdmin",
			Handler:    _SessionService_HasAccessSuperAdmin_Handler,
		},
		{
			MethodName: "V2Login",
			Handler:    _SessionService_V2Login_Handler,
		},
		{
			MethodName: "V2LoginSuperAdmin",
			Handler:    _SessionService_V2LoginSuperAdmin_Handler,
		},
		{
			MethodName: "V2HasAccess",
			Handler:    _SessionService_V2HasAccess_Handler,
		},
		{
			MethodName: "V2HasAccessUser",
			Handler:    _SessionService_V2HasAccessUser_Handler,
		},
		{
			MethodName: "V2RefreshToken",
			Handler:    _SessionService_V2RefreshToken_Handler,
		},
		{
			MethodName: "V2RefreshTokenSuperAdmin",
			Handler:    _SessionService_V2RefreshTokenSuperAdmin_Handler,
		},
		{
			MethodName: "SessionAndTokenGenerator",
			Handler:    _SessionService_SessionAndTokenGenerator_Handler,
		},
		{
			MethodName: "UpdateSessionsByRoleId",
			Handler:    _SessionService_UpdateSessionsByRoleId_Handler,
		},
		{
			MethodName: "MultiCompanyLogin",
			Handler:    _SessionService_MultiCompanyLogin_Handler,
		},
		{
			MethodName: "V2MultiCompanyLogin",
			Handler:    _SessionService_V2MultiCompanyLogin_Handler,
		},
		{
			MethodName: "V2MultiCompanyOneLogin",
			Handler:    _SessionService_V2MultiCompanyOneLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session_service.proto",
}