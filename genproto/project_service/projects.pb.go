// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: projects.proto

package project_service

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

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace                string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ObjectBuilderServiceHost string `protobuf:"bytes,3,opt,name=object_builder_service_host,json=objectBuilderServiceHost,proto3" json:"object_builder_service_host,omitempty"`
	ObjectBuilderServicePort string `protobuf:"bytes,4,opt,name=object_builder_service_port,json=objectBuilderServicePort,proto3" json:"object_builder_service_port,omitempty"`
	AuthServiceHost          string `protobuf:"bytes,5,opt,name=auth_service_host,json=authServiceHost,proto3" json:"auth_service_host,omitempty"`
	AuthServicePort          string `protobuf:"bytes,6,opt,name=auth_service_port,json=authServicePort,proto3" json:"auth_service_port,omitempty"`
	AnalyticsServiceHost     string `protobuf:"bytes,7,opt,name=analytics_service_host,json=analyticsServiceHost,proto3" json:"analytics_service_host,omitempty"`
	AnalyticsServicePort     string `protobuf:"bytes,8,opt,name=analytics_service_port,json=analyticsServicePort,proto3" json:"analytics_service_port,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateProjectRequest) GetObjectBuilderServiceHost() string {
	if x != nil {
		return x.ObjectBuilderServiceHost
	}
	return ""
}

func (x *CreateProjectRequest) GetObjectBuilderServicePort() string {
	if x != nil {
		return x.ObjectBuilderServicePort
	}
	return ""
}

func (x *CreateProjectRequest) GetAuthServiceHost() string {
	if x != nil {
		return x.AuthServiceHost
	}
	return ""
}

func (x *CreateProjectRequest) GetAuthServicePort() string {
	if x != nil {
		return x.AuthServicePort
	}
	return ""
}

func (x *CreateProjectRequest) GetAnalyticsServiceHost() string {
	if x != nil {
		return x.AnalyticsServiceHost
	}
	return ""
}

func (x *CreateProjectRequest) GetAnalyticsServicePort() string {
	if x != nil {
		return x.AnalyticsServicePort
	}
	return ""
}

type CreateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace                string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ObjectBuilderServiceHost string `protobuf:"bytes,4,opt,name=object_builder_service_host,json=objectBuilderServiceHost,proto3" json:"object_builder_service_host,omitempty"`
	ObjectBuilderServicePort string `protobuf:"bytes,5,opt,name=object_builder_service_port,json=objectBuilderServicePort,proto3" json:"object_builder_service_port,omitempty"`
	AuthServiceHost          string `protobuf:"bytes,6,opt,name=auth_service_host,json=authServiceHost,proto3" json:"auth_service_host,omitempty"`
	AuthServicePort          string `protobuf:"bytes,7,opt,name=auth_service_port,json=authServicePort,proto3" json:"auth_service_port,omitempty"`
	AnalyticsServiceHost     string `protobuf:"bytes,8,opt,name=analytics_service_host,json=analyticsServiceHost,proto3" json:"analytics_service_host,omitempty"`
	AnalyticsServicePort     string `protobuf:"bytes,9,opt,name=analytics_service_port,json=analyticsServicePort,proto3" json:"analytics_service_port,omitempty"`
}

func (x *CreateProjectResponse) Reset() {
	*x = CreateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectResponse) ProtoMessage() {}

func (x *CreateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectResponse.ProtoReflect.Descriptor instead.
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProjectResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateProjectResponse) GetObjectBuilderServiceHost() string {
	if x != nil {
		return x.ObjectBuilderServiceHost
	}
	return ""
}

func (x *CreateProjectResponse) GetObjectBuilderServicePort() string {
	if x != nil {
		return x.ObjectBuilderServicePort
	}
	return ""
}

func (x *CreateProjectResponse) GetAuthServiceHost() string {
	if x != nil {
		return x.AuthServiceHost
	}
	return ""
}

func (x *CreateProjectResponse) GetAuthServicePort() string {
	if x != nil {
		return x.AuthServicePort
	}
	return ""
}

func (x *CreateProjectResponse) GetAnalyticsServiceHost() string {
	if x != nil {
		return x.AnalyticsServiceHost
	}
	return ""
}

func (x *CreateProjectResponse) GetAnalyticsServicePort() string {
	if x != nil {
		return x.AnalyticsServicePort
	}
	return ""
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace                string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ObjectBuilderServiceHost string `protobuf:"bytes,4,opt,name=object_builder_service_host,json=objectBuilderServiceHost,proto3" json:"object_builder_service_host,omitempty"`
	ObjectBuilderServicePort string `protobuf:"bytes,5,opt,name=object_builder_service_port,json=objectBuilderServicePort,proto3" json:"object_builder_service_port,omitempty"`
	AuthServiceHost          string `protobuf:"bytes,6,opt,name=auth_service_host,json=authServiceHost,proto3" json:"auth_service_host,omitempty"`
	AuthServicePort          string `protobuf:"bytes,7,opt,name=auth_service_port,json=authServicePort,proto3" json:"auth_service_port,omitempty"`
	AnalyticsServiceHost     string `protobuf:"bytes,8,opt,name=analytics_service_host,json=analyticsServiceHost,proto3" json:"analytics_service_host,omitempty"`
	AnalyticsServicePort     string `protobuf:"bytes,9,opt,name=analytics_service_port,json=analyticsServicePort,proto3" json:"analytics_service_port,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{2}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Project) GetObjectBuilderServiceHost() string {
	if x != nil {
		return x.ObjectBuilderServiceHost
	}
	return ""
}

func (x *Project) GetObjectBuilderServicePort() string {
	if x != nil {
		return x.ObjectBuilderServicePort
	}
	return ""
}

func (x *Project) GetAuthServiceHost() string {
	if x != nil {
		return x.AuthServiceHost
	}
	return ""
}

func (x *Project) GetAuthServicePort() string {
	if x != nil {
		return x.AuthServicePort
	}
	return ""
}

func (x *Project) GetAnalyticsServiceHost() string {
	if x != nil {
		return x.AnalyticsServiceHost
	}
	return ""
}

func (x *Project) GetAnalyticsServicePort() string {
	if x != nil {
		return x.AnalyticsServicePort
	}
	return ""
}

type UpdateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace                string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ObjectBuilderServiceHost string `protobuf:"bytes,4,opt,name=object_builder_service_host,json=objectBuilderServiceHost,proto3" json:"object_builder_service_host,omitempty"`
	ObjectBuilderServicePort string `protobuf:"bytes,5,opt,name=object_builder_service_port,json=objectBuilderServicePort,proto3" json:"object_builder_service_port,omitempty"`
	AuthServiceHost          string `protobuf:"bytes,6,opt,name=auth_service_host,json=authServiceHost,proto3" json:"auth_service_host,omitempty"`
	AuthServicePort          string `protobuf:"bytes,7,opt,name=auth_service_port,json=authServicePort,proto3" json:"auth_service_port,omitempty"`
	AnalyticsServiceHost     string `protobuf:"bytes,8,opt,name=analytics_service_host,json=analyticsServiceHost,proto3" json:"analytics_service_host,omitempty"`
	AnalyticsServicePort     string `protobuf:"bytes,9,opt,name=analytics_service_port,json=analyticsServicePort,proto3" json:"analytics_service_port,omitempty"`
}

func (x *UpdateProjectRequest) Reset() {
	*x = UpdateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectRequest) ProtoMessage() {}

func (x *UpdateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProjectRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *UpdateProjectRequest) GetObjectBuilderServiceHost() string {
	if x != nil {
		return x.ObjectBuilderServiceHost
	}
	return ""
}

func (x *UpdateProjectRequest) GetObjectBuilderServicePort() string {
	if x != nil {
		return x.ObjectBuilderServicePort
	}
	return ""
}

func (x *UpdateProjectRequest) GetAuthServiceHost() string {
	if x != nil {
		return x.AuthServiceHost
	}
	return ""
}

func (x *UpdateProjectRequest) GetAuthServicePort() string {
	if x != nil {
		return x.AuthServicePort
	}
	return ""
}

func (x *UpdateProjectRequest) GetAnalyticsServiceHost() string {
	if x != nil {
		return x.AnalyticsServiceHost
	}
	return ""
}

func (x *UpdateProjectRequest) GetAnalyticsServicePort() string {
	if x != nil {
		return x.AnalyticsServicePort
	}
	return ""
}

type GetAllProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	Count    int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllProjectsResponse) Reset() {
	*x = GetAllProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProjectsResponse) ProtoMessage() {}

func (x *GetAllProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProjectsResponse.ProtoReflect.Descriptor instead.
func (*GetAllProjectsResponse) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllProjectsResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *GetAllProjectsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ProjectPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectPrimaryKey) Reset() {
	*x = ProjectPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectPrimaryKey) ProtoMessage() {}

func (x *ProjectPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectPrimaryKey.ProtoReflect.Descriptor instead.
func (*ProjectPrimaryKey) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllProjectsRequest) Reset() {
	*x = GetAllProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProjectsRequest) ProtoMessage() {}

func (x *GetAllProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProjectsRequest.ProtoReflect.Descriptor instead.
func (*GetAllProjectsRequest) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllProjectsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllProjectsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllProjectsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

var File_projects_proto protoreflect.FileDescriptor

var file_projects_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x03, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x34, 0x0a, 0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x9b, 0x03, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a,
	0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x8d, 0x03, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a,
	0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x9a, 0x03, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x64, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x11,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x32, 0xa6, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_projects_proto_rawDescOnce sync.Once
	file_projects_proto_rawDescData = file_projects_proto_rawDesc
)

func file_projects_proto_rawDescGZIP() []byte {
	file_projects_proto_rawDescOnce.Do(func() {
		file_projects_proto_rawDescData = protoimpl.X.CompressGZIP(file_projects_proto_rawDescData)
	})
	return file_projects_proto_rawDescData
}

var file_projects_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_projects_proto_goTypes = []interface{}{
	(*CreateProjectRequest)(nil),   // 0: project_service.CreateProjectRequest
	(*CreateProjectResponse)(nil),  // 1: project_service.CreateProjectResponse
	(*Project)(nil),                // 2: project_service.Project
	(*UpdateProjectRequest)(nil),   // 3: project_service.UpdateProjectRequest
	(*GetAllProjectsResponse)(nil), // 4: project_service.GetAllProjectsResponse
	(*ProjectPrimaryKey)(nil),      // 5: project_service.ProjectPrimaryKey
	(*GetAllProjectsRequest)(nil),  // 6: project_service.GetAllProjectsRequest
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_projects_proto_depIdxs = []int32{
	2, // 0: project_service.GetAllProjectsResponse.projects:type_name -> project_service.Project
	0, // 1: project_service.ProjectService.Create:input_type -> project_service.CreateProjectRequest
	5, // 2: project_service.ProjectService.GetByID:input_type -> project_service.ProjectPrimaryKey
	6, // 3: project_service.ProjectService.GetAll:input_type -> project_service.GetAllProjectsRequest
	3, // 4: project_service.ProjectService.Update:input_type -> project_service.UpdateProjectRequest
	5, // 5: project_service.ProjectService.Delete:input_type -> project_service.ProjectPrimaryKey
	1, // 6: project_service.ProjectService.Create:output_type -> project_service.CreateProjectResponse
	2, // 7: project_service.ProjectService.GetByID:output_type -> project_service.Project
	4, // 8: project_service.ProjectService.GetAll:output_type -> project_service.GetAllProjectsResponse
	7, // 9: project_service.ProjectService.Update:output_type -> google.protobuf.Empty
	7, // 10: project_service.ProjectService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_projects_proto_init() }
func file_projects_proto_init() {
	if File_projects_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_projects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
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
		file_projects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectResponse); i {
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
		file_projects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_projects_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProjectRequest); i {
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
		file_projects_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProjectsResponse); i {
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
		file_projects_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectPrimaryKey); i {
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
		file_projects_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProjectsRequest); i {
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
			RawDescriptor: file_projects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_projects_proto_goTypes,
		DependencyIndexes: file_projects_proto_depIdxs,
		MessageInfos:      file_projects_proto_msgTypes,
	}.Build()
	File_projects_proto = out.File
	file_projects_proto_rawDesc = nil
	file_projects_proto_goTypes = nil
	file_projects_proto_depIdxs = nil
}