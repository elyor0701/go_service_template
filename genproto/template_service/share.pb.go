// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: share.proto

package template_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId   string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ProjectId  string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Role       string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	IsPrivate  bool   `protobuf:"varint,5,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	ResourceId string `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *CreateTokenReq) Reset() {
	*x = CreateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenReq) ProtoMessage() {}

func (x *CreateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenReq.ProtoReflect.Descriptor instead.
func (*CreateTokenReq) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTokenReq) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *CreateTokenReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateTokenReq) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateTokenReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CreateTokenReq) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *CreateTokenReq) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type CreateTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ObjectId  string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Role      string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	IsPrivate bool   `protobuf:"varint,6,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
}

func (x *CreateTokenRes) Reset() {
	*x = CreateTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRes) ProtoMessage() {}

func (x *CreateTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRes.ProtoReflect.Descriptor instead.
func (*CreateTokenRes) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokenRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateTokenRes) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *CreateTokenRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateTokenRes) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateTokenRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CreateTokenRes) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type UpdateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId   string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ProjectId  string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Role       string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	IsPrivate  bool   `protobuf:"varint,5,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	ResourceId string `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *UpdateTokenReq) Reset() {
	*x = UpdateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenReq) ProtoMessage() {}

func (x *UpdateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenReq.ProtoReflect.Descriptor instead.
func (*UpdateTokenReq) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTokenReq) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *UpdateTokenReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateTokenReq) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *UpdateTokenReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UpdateTokenReq) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *UpdateTokenReq) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ObjectId  string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Role      string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	IsPrivate bool   `protobuf:"varint,6,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
}

func (x *UpdateTokenRes) Reset() {
	*x = UpdateTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenRes) ProtoMessage() {}

func (x *UpdateTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenRes.ProtoReflect.Descriptor instead.
func (*UpdateTokenRes) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTokenRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateTokenRes) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *UpdateTokenRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateTokenRes) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *UpdateTokenRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UpdateTokenRes) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type GetObjectTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VersionId  string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	ResourceId string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *GetObjectTokenReq) Reset() {
	*x = GetObjectTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectTokenReq) ProtoMessage() {}

func (x *GetObjectTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectTokenReq.ProtoReflect.Descriptor instead.
func (*GetObjectTokenReq) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{4}
}

func (x *GetObjectTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetObjectTokenReq) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *GetObjectTokenReq) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type GetObjectTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Role string           `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *GetObjectTokenRes) Reset() {
	*x = GetObjectTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectTokenRes) ProtoMessage() {}

func (x *GetObjectTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectTokenRes.ProtoReflect.Descriptor instead.
func (*GetObjectTokenRes) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{5}
}

func (x *GetObjectTokenRes) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetObjectTokenRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_share_proto protoreflect.FileDescriptor

var file_share_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0xb4, 0x01,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x22, 0x69, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x32, 0x84, 0x02, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x1d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x20, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_share_proto_rawDescOnce sync.Once
	file_share_proto_rawDescData = file_share_proto_rawDesc
)

func file_share_proto_rawDescGZIP() []byte {
	file_share_proto_rawDescOnce.Do(func() {
		file_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_share_proto_rawDescData)
	})
	return file_share_proto_rawDescData
}

var file_share_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_share_proto_goTypes = []interface{}{
	(*CreateTokenReq)(nil),    // 0: share_service.CreateTokenReq
	(*CreateTokenRes)(nil),    // 1: share_service.CreateTokenRes
	(*UpdateTokenReq)(nil),    // 2: share_service.UpdateTokenReq
	(*UpdateTokenRes)(nil),    // 3: share_service.UpdateTokenRes
	(*GetObjectTokenReq)(nil), // 4: share_service.GetObjectTokenReq
	(*GetObjectTokenRes)(nil), // 5: share_service.GetObjectTokenRes
	(*structpb.Struct)(nil),   // 6: google.protobuf.Struct
}
var file_share_proto_depIdxs = []int32{
	6, // 0: share_service.GetObjectTokenRes.data:type_name -> google.protobuf.Struct
	0, // 1: share_service.ShareService.CreateToken:input_type -> share_service.CreateTokenReq
	2, // 2: share_service.ShareService.UpdateToken:input_type -> share_service.UpdateTokenReq
	4, // 3: share_service.ShareService.GetObjectToken:input_type -> share_service.GetObjectTokenReq
	1, // 4: share_service.ShareService.CreateToken:output_type -> share_service.CreateTokenRes
	3, // 5: share_service.ShareService.UpdateToken:output_type -> share_service.UpdateTokenRes
	5, // 6: share_service.ShareService.GetObjectToken:output_type -> share_service.GetObjectTokenRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_share_proto_init() }
func file_share_proto_init() {
	if File_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenReq); i {
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
		file_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRes); i {
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
		file_share_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTokenReq); i {
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
		file_share_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTokenRes); i {
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
		file_share_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectTokenReq); i {
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
		file_share_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectTokenRes); i {
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
			RawDescriptor: file_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_share_proto_goTypes,
		DependencyIndexes: file_share_proto_depIdxs,
		MessageInfos:      file_share_proto_msgTypes,
	}.Build()
	File_share_proto = out.File
	file_share_proto_rawDesc = nil
	file_share_proto_goTypes = nil
	file_share_proto_depIdxs = nil
}