// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: user_token.proto

package notification_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token      string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	PlatformId string `protobuf:"bytes,5,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
}

func (x *UserToken) Reset() {
	*x = UserToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserToken) ProtoMessage() {}

func (x *UserToken) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserToken.ProtoReflect.Descriptor instead.
func (*UserToken) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{0}
}

func (x *UserToken) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserToken) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserToken) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserToken) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

type CreateUserTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	PlatformId string `protobuf:"bytes,3,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
}

func (x *CreateUserTokenRequest) Reset() {
	*x = CreateUserTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTokenRequest) ProtoMessage() {}

func (x *CreateUserTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateUserTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateUserTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateUserTokenRequest) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

type CreateUserTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTokenId string `protobuf:"bytes,1,opt,name=user_token_id,json=userTokenId,proto3" json:"user_token_id,omitempty"`
}

func (x *CreateUserTokenResponse) Reset() {
	*x = CreateUserTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTokenResponse) ProtoMessage() {}

func (x *CreateUserTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateUserTokenResponse) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserTokenResponse) GetUserTokenId() string {
	if x != nil {
		return x.UserTokenId
	}
	return ""
}

type GetListUserTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetListUserTokenRequest) Reset() {
	*x = GetListUserTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListUserTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListUserTokenRequest) ProtoMessage() {}

func (x *GetListUserTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListUserTokenRequest.ProtoReflect.Descriptor instead.
func (*GetListUserTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{3}
}

func (x *GetListUserTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListUserToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTokens []*UserToken `protobuf:"bytes,1,rep,name=user_tokens,json=userTokens,proto3" json:"user_tokens,omitempty"`
}

func (x *ListUserToken) Reset() {
	*x = ListUserToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserToken) ProtoMessage() {}

func (x *ListUserToken) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserToken.ProtoReflect.Descriptor instead.
func (*ListUserToken) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserToken) GetUserTokens() []*UserToken {
	if x != nil {
		return x.UserTokens
	}
	return nil
}

type DeleteUserTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTokenId string `protobuf:"bytes,1,opt,name=user_token_id,json=userTokenId,proto3" json:"user_token_id,omitempty"`
}

func (x *DeleteUserTokenRequest) Reset() {
	*x = DeleteUserTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTokenRequest) ProtoMessage() {}

func (x *DeleteUserTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserTokenRequest) GetUserTokenId() string {
	if x != nil {
		return x.UserTokenId
	}
	return ""
}

type GetUsersTokenListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *GetUsersTokenListRequest) Reset() {
	*x = GetUsersTokenListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersTokenListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersTokenListRequest) ProtoMessage() {}

func (x *GetUsersTokenListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersTokenListRequest.ProtoReflect.Descriptor instead.
func (*GetUsersTokenListRequest) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersTokenListRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetUsersTokenListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTokens []*UserToken `protobuf:"bytes,1,rep,name=user_tokens,json=userTokens,proto3" json:"user_tokens,omitempty"`
}

func (x *GetUsersTokenListResponse) Reset() {
	*x = GetUsersTokenListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_token_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersTokenListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersTokenListResponse) ProtoMessage() {}

func (x *GetUsersTokenListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_token_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersTokenListResponse.ProtoReflect.Descriptor instead.
func (*GetUsersTokenListResponse) Descriptor() ([]byte, []int) {
	return file_user_token_proto_rawDescGZIP(), []int{7}
}

func (x *GetUsersTokenListResponse) GetUserTokens() []*UserToken {
	if x != nil {
		return x.UserTokens
	}
	return nil
}

var File_user_token_proto protoreflect.FileDescriptor

var file_user_token_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x49, 0x64, 0x22, 0x68, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x51, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x40, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64,
	0x22, 0x35, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x5d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_token_proto_rawDescOnce sync.Once
	file_user_token_proto_rawDescData = file_user_token_proto_rawDesc
)

func file_user_token_proto_rawDescGZIP() []byte {
	file_user_token_proto_rawDescOnce.Do(func() {
		file_user_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_token_proto_rawDescData)
	})
	return file_user_token_proto_rawDescData
}

var file_user_token_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_token_proto_goTypes = []interface{}{
	(*UserToken)(nil),                 // 0: notification_service.UserToken
	(*CreateUserTokenRequest)(nil),    // 1: notification_service.CreateUserTokenRequest
	(*CreateUserTokenResponse)(nil),   // 2: notification_service.CreateUserTokenResponse
	(*GetListUserTokenRequest)(nil),   // 3: notification_service.GetListUserTokenRequest
	(*ListUserToken)(nil),             // 4: notification_service.ListUserToken
	(*DeleteUserTokenRequest)(nil),    // 5: notification_service.DeleteUserTokenRequest
	(*GetUsersTokenListRequest)(nil),  // 6: notification_service.GetUsersTokenListRequest
	(*GetUsersTokenListResponse)(nil), // 7: notification_service.GetUsersTokenListResponse
}
var file_user_token_proto_depIdxs = []int32{
	0, // 0: notification_service.ListUserToken.user_tokens:type_name -> notification_service.UserToken
	0, // 1: notification_service.GetUsersTokenListResponse.user_tokens:type_name -> notification_service.UserToken
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_token_proto_init() }
func file_user_token_proto_init() {
	if File_user_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListUserTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersTokenListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_token_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersTokenListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_token_proto_goTypes,
		DependencyIndexes: file_user_token_proto_depIdxs,
		MessageInfos:      file_user_token_proto_msgTypes,
	}.Build()
	File_user_token_proto = out.File
	file_user_token_proto_rawDesc = nil
	file_user_token_proto_goTypes = nil
	file_user_token_proto_depIdxs = nil
}
