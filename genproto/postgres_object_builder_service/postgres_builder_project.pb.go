// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: postgres_builder_project.proto

package postgres_object_builder_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K8SNamespace string                              `protobuf:"bytes,1,opt,name=k8s_namespace,json=k8sNamespace,proto3" json:"k8s_namespace,omitempty"`
	ProjectId    string                              `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	SecretPath   string                              `protobuf:"bytes,3,opt,name=secret_path,json=secretPath,proto3" json:"secret_path,omitempty"`
	Credentials  *RegisterProjectRequest_Credentials `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	UserId       string                              `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceId   string                              `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *RegisterProjectRequest) Reset() {
	*x = RegisterProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterProjectRequest) ProtoMessage() {}

func (x *RegisterProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterProjectRequest.ProtoReflect.Descriptor instead.
func (*RegisterProjectRequest) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterProjectRequest) GetK8SNamespace() string {
	if x != nil {
		return x.K8SNamespace
	}
	return ""
}

func (x *RegisterProjectRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *RegisterProjectRequest) GetSecretPath() string {
	if x != nil {
		return x.SecretPath
	}
	return ""
}

func (x *RegisterProjectRequest) GetCredentials() *RegisterProjectRequest_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *RegisterProjectRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterProjectRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type DeregisterProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K8SNamespace string `protobuf:"bytes,1,opt,name=k8s_namespace,json=k8sNamespace,proto3" json:"k8s_namespace,omitempty"`
	ProjectId    string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DeregisterProjectRequest) Reset() {
	*x = DeregisterProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterProjectRequest) ProtoMessage() {}

func (x *DeregisterProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterProjectRequest.ProtoReflect.Descriptor instead.
func (*DeregisterProjectRequest) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{1}
}

func (x *DeregisterProjectRequest) GetK8SNamespace() string {
	if x != nil {
		return x.K8SNamespace
	}
	return ""
}

func (x *DeregisterProjectRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type RegisterDeregisterProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K8SNamespace string `protobuf:"bytes,1,opt,name=k8s_namespace,json=k8sNamespace,proto3" json:"k8s_namespace,omitempty"`
	HasError     bool   `protobuf:"varint,2,opt,name=has_error,json=hasError,proto3" json:"has_error,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *RegisterDeregisterProjectResponse) Reset() {
	*x = RegisterDeregisterProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDeregisterProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeregisterProjectResponse) ProtoMessage() {}

func (x *RegisterDeregisterProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeregisterProjectResponse.ProtoReflect.Descriptor instead.
func (*RegisterDeregisterProjectResponse) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterDeregisterProjectResponse) GetK8SNamespace() string {
	if x != nil {
		return x.K8SNamespace
	}
	return ""
}

func (x *RegisterDeregisterProjectResponse) GetHasError() bool {
	if x != nil {
		return x.HasError
	}
	return false
}

func (x *RegisterDeregisterProjectResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type RegisterManyProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*RegisterProjectRequest `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *RegisterManyProjectsRequest) Reset() {
	*x = RegisterManyProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterManyProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterManyProjectsRequest) ProtoMessage() {}

func (x *RegisterManyProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterManyProjectsRequest.ProtoReflect.Descriptor instead.
func (*RegisterManyProjectsRequest) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterManyProjectsRequest) GetProjects() []*RegisterProjectRequest {
	if x != nil {
		return x.Projects
	}
	return nil
}

type RegisterManyProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects map[string]*RegisterDeregisterProjectResponse `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegisterManyProjectsResponse) Reset() {
	*x = RegisterManyProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterManyProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterManyProjectsResponse) ProtoMessage() {}

func (x *RegisterManyProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterManyProjectsResponse.ProtoReflect.Descriptor instead.
func (*RegisterManyProjectsResponse) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterManyProjectsResponse) GetProjects() map[string]*RegisterDeregisterProjectResponse {
	if x != nil {
		return x.Projects
	}
	return nil
}

type DeregisterManyProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*DeregisterProjectRequest `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *DeregisterManyProjectsRequest) Reset() {
	*x = DeregisterManyProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterManyProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterManyProjectsRequest) ProtoMessage() {}

func (x *DeregisterManyProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterManyProjectsRequest.ProtoReflect.Descriptor instead.
func (*DeregisterManyProjectsRequest) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{5}
}

func (x *DeregisterManyProjectsRequest) GetProjects() []*DeregisterProjectRequest {
	if x != nil {
		return x.Projects
	}
	return nil
}

type DeregisterManyProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects map[string]*RegisterDeregisterProjectResponse `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeregisterManyProjectsResponse) Reset() {
	*x = DeregisterManyProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterManyProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterManyProjectsResponse) ProtoMessage() {}

func (x *DeregisterManyProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterManyProjectsResponse.ProtoReflect.Descriptor instead.
func (*DeregisterManyProjectsResponse) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{6}
}

func (x *DeregisterManyProjectsResponse) GetProjects() map[string]*RegisterDeregisterProjectResponse {
	if x != nil {
		return x.Projects
	}
	return nil
}

type RegisterProjectRequest_Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Database string `protobuf:"bytes,5,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *RegisterProjectRequest_Credentials) Reset() {
	*x = RegisterProjectRequest_Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_builder_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterProjectRequest_Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterProjectRequest_Credentials) ProtoMessage() {}

func (x *RegisterProjectRequest_Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_builder_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterProjectRequest_Credentials.ProtoReflect.Descriptor instead.
func (*RegisterProjectRequest_Credentials) Descriptor() ([]byte, []int) {
	return file_postgres_builder_project_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RegisterProjectRequest_Credentials) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RegisterProjectRequest_Credentials) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *RegisterProjectRequest_Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterProjectRequest_Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterProjectRequest_Credentials) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

var File_postgres_builder_project_proto protoreflect.FileDescriptor

var file_postgres_builder_project_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa,
	0x03, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x38, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6b, 0x38, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x65,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x1a,
	0x89, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x5e, 0x0a, 0x18, 0x44,
	0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x38, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6b, 0x38, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x21,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x38, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x38, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x72, 0x0a, 0x1b, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x88, 0x02, 0x0a,
	0x1c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x4b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a, 0x7f, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x58, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x76, 0x0a, 0x1d, 0x44, 0x65, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x8c, 0x02, 0x0a, 0x1e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x69, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a, 0x7f, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x58, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xc6,
	0x05, 0x0a, 0x15, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x37, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x37, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x61,
	0x0a, 0x0a, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x39, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x37,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x79, 0x12, 0x3c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e,
	0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x93, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4d, 0x61, 0x6e, 0x79, 0x12, 0x3e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postgres_builder_project_proto_rawDescOnce sync.Once
	file_postgres_builder_project_proto_rawDescData = file_postgres_builder_project_proto_rawDesc
)

func file_postgres_builder_project_proto_rawDescGZIP() []byte {
	file_postgres_builder_project_proto_rawDescOnce.Do(func() {
		file_postgres_builder_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_postgres_builder_project_proto_rawDescData)
	})
	return file_postgres_builder_project_proto_rawDescData
}

var file_postgres_builder_project_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_postgres_builder_project_proto_goTypes = []interface{}{
	(*RegisterProjectRequest)(nil),             // 0: postgres_object_builder_service.RegisterProjectRequest
	(*DeregisterProjectRequest)(nil),           // 1: postgres_object_builder_service.DeregisterProjectRequest
	(*RegisterDeregisterProjectResponse)(nil),  // 2: postgres_object_builder_service.RegisterDeregisterProjectResponse
	(*RegisterManyProjectsRequest)(nil),        // 3: postgres_object_builder_service.RegisterManyProjectsRequest
	(*RegisterManyProjectsResponse)(nil),       // 4: postgres_object_builder_service.RegisterManyProjectsResponse
	(*DeregisterManyProjectsRequest)(nil),      // 5: postgres_object_builder_service.DeregisterManyProjectsRequest
	(*DeregisterManyProjectsResponse)(nil),     // 6: postgres_object_builder_service.DeregisterManyProjectsResponse
	(*RegisterProjectRequest_Credentials)(nil), // 7: postgres_object_builder_service.RegisterProjectRequest.Credentials
	nil,                   // 8: postgres_object_builder_service.RegisterManyProjectsResponse.ProjectsEntry
	nil,                   // 9: postgres_object_builder_service.DeregisterManyProjectsResponse.ProjectsEntry
	(*emptypb.Empty)(nil), // 10: google.protobuf.Empty
}
var file_postgres_builder_project_proto_depIdxs = []int32{
	7,  // 0: postgres_object_builder_service.RegisterProjectRequest.credentials:type_name -> postgres_object_builder_service.RegisterProjectRequest.Credentials
	0,  // 1: postgres_object_builder_service.RegisterManyProjectsRequest.projects:type_name -> postgres_object_builder_service.RegisterProjectRequest
	8,  // 2: postgres_object_builder_service.RegisterManyProjectsResponse.projects:type_name -> postgres_object_builder_service.RegisterManyProjectsResponse.ProjectsEntry
	1,  // 3: postgres_object_builder_service.DeregisterManyProjectsRequest.projects:type_name -> postgres_object_builder_service.DeregisterProjectRequest
	9,  // 4: postgres_object_builder_service.DeregisterManyProjectsResponse.projects:type_name -> postgres_object_builder_service.DeregisterManyProjectsResponse.ProjectsEntry
	2,  // 5: postgres_object_builder_service.RegisterManyProjectsResponse.ProjectsEntry.value:type_name -> postgres_object_builder_service.RegisterDeregisterProjectResponse
	2,  // 6: postgres_object_builder_service.DeregisterManyProjectsResponse.ProjectsEntry.value:type_name -> postgres_object_builder_service.RegisterDeregisterProjectResponse
	0,  // 7: postgres_object_builder_service.BuilderProjectService.Register:input_type -> postgres_object_builder_service.RegisterProjectRequest
	0,  // 8: postgres_object_builder_service.BuilderProjectService.RegisterProjects:input_type -> postgres_object_builder_service.RegisterProjectRequest
	1,  // 9: postgres_object_builder_service.BuilderProjectService.Deregister:input_type -> postgres_object_builder_service.DeregisterProjectRequest
	0,  // 10: postgres_object_builder_service.BuilderProjectService.Reconnect:input_type -> postgres_object_builder_service.RegisterProjectRequest
	3,  // 11: postgres_object_builder_service.BuilderProjectService.RegisterMany:input_type -> postgres_object_builder_service.RegisterManyProjectsRequest
	5,  // 12: postgres_object_builder_service.BuilderProjectService.DeregisterMany:input_type -> postgres_object_builder_service.DeregisterManyProjectsRequest
	10, // 13: postgres_object_builder_service.BuilderProjectService.Register:output_type -> google.protobuf.Empty
	10, // 14: postgres_object_builder_service.BuilderProjectService.RegisterProjects:output_type -> google.protobuf.Empty
	10, // 15: postgres_object_builder_service.BuilderProjectService.Deregister:output_type -> google.protobuf.Empty
	10, // 16: postgres_object_builder_service.BuilderProjectService.Reconnect:output_type -> google.protobuf.Empty
	4,  // 17: postgres_object_builder_service.BuilderProjectService.RegisterMany:output_type -> postgres_object_builder_service.RegisterManyProjectsResponse
	6,  // 18: postgres_object_builder_service.BuilderProjectService.DeregisterMany:output_type -> postgres_object_builder_service.DeregisterManyProjectsResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_postgres_builder_project_proto_init() }
func file_postgres_builder_project_proto_init() {
	if File_postgres_builder_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_postgres_builder_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProjectRequest); i {
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
		file_postgres_builder_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterProjectRequest); i {
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
		file_postgres_builder_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDeregisterProjectResponse); i {
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
		file_postgres_builder_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterManyProjectsRequest); i {
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
		file_postgres_builder_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterManyProjectsResponse); i {
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
		file_postgres_builder_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterManyProjectsRequest); i {
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
		file_postgres_builder_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterManyProjectsResponse); i {
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
		file_postgres_builder_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProjectRequest_Credentials); i {
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
			RawDescriptor: file_postgres_builder_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postgres_builder_project_proto_goTypes,
		DependencyIndexes: file_postgres_builder_project_proto_depIdxs,
		MessageInfos:      file_postgres_builder_project_proto_msgTypes,
	}.Build()
	File_postgres_builder_project_proto = out.File
	file_postgres_builder_project_proto_rawDesc = nil
	file_postgres_builder_project_proto_goTypes = nil
	file_postgres_builder_project_proto_depIdxs = nil
}
