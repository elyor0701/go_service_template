// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: postgres_variable.proto

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

type CreateVariableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug          string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Type          string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Label         string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	DashboardId   string   `protobuf:"bytes,4,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	FieldSlug     string   `protobuf:"bytes,5,opt,name=field_slug,json=fieldSlug,proto3" json:"field_slug,omitempty"`
	Options       []string `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`
	ViewFieldSlug string   `protobuf:"bytes,7,opt,name=view_field_slug,json=viewFieldSlug,proto3" json:"view_field_slug,omitempty"`
	Query         string   `protobuf:"bytes,8,opt,name=query,proto3" json:"query,omitempty"`
	ProjectId     string   `protobuf:"bytes,9,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CommitId      int64    `protobuf:"varint,10,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	CommitGuid    string   `protobuf:"bytes,11,opt,name=commit_guid,json=commitGuid,proto3" json:"commit_guid,omitempty"`
}

func (x *CreateVariableRequest) Reset() {
	*x = CreateVariableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_variable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVariableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVariableRequest) ProtoMessage() {}

func (x *CreateVariableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_variable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVariableRequest.ProtoReflect.Descriptor instead.
func (*CreateVariableRequest) Descriptor() ([]byte, []int) {
	return file_postgres_variable_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVariableRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CreateVariableRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateVariableRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CreateVariableRequest) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

func (x *CreateVariableRequest) GetFieldSlug() string {
	if x != nil {
		return x.FieldSlug
	}
	return ""
}

func (x *CreateVariableRequest) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CreateVariableRequest) GetViewFieldSlug() string {
	if x != nil {
		return x.ViewFieldSlug
	}
	return ""
}

func (x *CreateVariableRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *CreateVariableRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateVariableRequest) GetCommitId() int64 {
	if x != nil {
		return x.CommitId
	}
	return 0
}

func (x *CreateVariableRequest) GetCommitGuid() string {
	if x != nil {
		return x.CommitGuid
	}
	return ""
}

type Variable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug          string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Type          string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Label         string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	DashboardId   string   `protobuf:"bytes,5,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	FieldSlug     string   `protobuf:"bytes,6,opt,name=field_slug,json=fieldSlug,proto3" json:"field_slug,omitempty"`
	Options       []string `protobuf:"bytes,7,rep,name=options,proto3" json:"options,omitempty"`
	ViewFieldSlug string   `protobuf:"bytes,8,opt,name=view_field_slug,json=viewFieldSlug,proto3" json:"view_field_slug,omitempty"`
	Query         string   `protobuf:"bytes,9,opt,name=query,proto3" json:"query,omitempty"`
	ProjectId     string   `protobuf:"bytes,10,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *Variable) Reset() {
	*x = Variable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_variable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variable) ProtoMessage() {}

func (x *Variable) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_variable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variable.ProtoReflect.Descriptor instead.
func (*Variable) Descriptor() ([]byte, []int) {
	return file_postgres_variable_proto_rawDescGZIP(), []int{1}
}

func (x *Variable) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Variable) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Variable) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Variable) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Variable) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

func (x *Variable) GetFieldSlug() string {
	if x != nil {
		return x.FieldSlug
	}
	return ""
}

func (x *Variable) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Variable) GetViewFieldSlug() string {
	if x != nil {
		return x.ViewFieldSlug
	}
	return ""
}

func (x *Variable) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Variable) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetAllVariablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variables []*Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty"`
	Count     int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllVariablesResponse) Reset() {
	*x = GetAllVariablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_variable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllVariablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllVariablesResponse) ProtoMessage() {}

func (x *GetAllVariablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_variable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllVariablesResponse.ProtoReflect.Descriptor instead.
func (*GetAllVariablesResponse) Descriptor() ([]byte, []int) {
	return file_postgres_variable_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllVariablesResponse) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *GetAllVariablesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetAllVariablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug        string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	DashboardId string `protobuf:"bytes,2,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	ProjectId   string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetAllVariablesRequest) Reset() {
	*x = GetAllVariablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_variable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllVariablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllVariablesRequest) ProtoMessage() {}

func (x *GetAllVariablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_variable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllVariablesRequest.ProtoReflect.Descriptor instead.
func (*GetAllVariablesRequest) Descriptor() ([]byte, []int) {
	return file_postgres_variable_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllVariablesRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetAllVariablesRequest) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

func (x *GetAllVariablesRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type VariablePrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *VariablePrimaryKey) Reset() {
	*x = VariablePrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_variable_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariablePrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariablePrimaryKey) ProtoMessage() {}

func (x *VariablePrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_variable_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariablePrimaryKey.ProtoReflect.Descriptor instead.
func (*VariablePrimaryKey) Descriptor() ([]byte, []int) {
	return file_postgres_variable_proto_rawDescGZIP(), []int{4}
}

func (x *VariablePrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VariablePrimaryKey) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_postgres_variable_proto protoreflect.FileDescriptor

var file_postgres_variable_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x6c, 0x75, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x6c, 0x75,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x69, 0x65, 0x77, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53,
	0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x47, 0x75, 0x69, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x08, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53,
	0x6c, 0x75, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x6c, 0x75, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x69, 0x65, 0x77, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x32, 0x97, 0x04, 0x0a, 0x0f, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x33, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x29,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x33, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postgres_variable_proto_rawDescOnce sync.Once
	file_postgres_variable_proto_rawDescData = file_postgres_variable_proto_rawDesc
)

func file_postgres_variable_proto_rawDescGZIP() []byte {
	file_postgres_variable_proto_rawDescOnce.Do(func() {
		file_postgres_variable_proto_rawDescData = protoimpl.X.CompressGZIP(file_postgres_variable_proto_rawDescData)
	})
	return file_postgres_variable_proto_rawDescData
}

var file_postgres_variable_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_postgres_variable_proto_goTypes = []interface{}{
	(*CreateVariableRequest)(nil),   // 0: postgres_object_builder_service.CreateVariableRequest
	(*Variable)(nil),                // 1: postgres_object_builder_service.Variable
	(*GetAllVariablesResponse)(nil), // 2: postgres_object_builder_service.GetAllVariablesResponse
	(*GetAllVariablesRequest)(nil),  // 3: postgres_object_builder_service.GetAllVariablesRequest
	(*VariablePrimaryKey)(nil),      // 4: postgres_object_builder_service.VariablePrimaryKey
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_postgres_variable_proto_depIdxs = []int32{
	1, // 0: postgres_object_builder_service.GetAllVariablesResponse.variables:type_name -> postgres_object_builder_service.Variable
	0, // 1: postgres_object_builder_service.VariableService.Create:input_type -> postgres_object_builder_service.CreateVariableRequest
	3, // 2: postgres_object_builder_service.VariableService.GetList:input_type -> postgres_object_builder_service.GetAllVariablesRequest
	4, // 3: postgres_object_builder_service.VariableService.GetSingle:input_type -> postgres_object_builder_service.VariablePrimaryKey
	1, // 4: postgres_object_builder_service.VariableService.Update:input_type -> postgres_object_builder_service.Variable
	4, // 5: postgres_object_builder_service.VariableService.Delete:input_type -> postgres_object_builder_service.VariablePrimaryKey
	1, // 6: postgres_object_builder_service.VariableService.Create:output_type -> postgres_object_builder_service.Variable
	2, // 7: postgres_object_builder_service.VariableService.GetList:output_type -> postgres_object_builder_service.GetAllVariablesResponse
	1, // 8: postgres_object_builder_service.VariableService.GetSingle:output_type -> postgres_object_builder_service.Variable
	5, // 9: postgres_object_builder_service.VariableService.Update:output_type -> google.protobuf.Empty
	5, // 10: postgres_object_builder_service.VariableService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_postgres_variable_proto_init() }
func file_postgres_variable_proto_init() {
	if File_postgres_variable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_postgres_variable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVariableRequest); i {
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
		file_postgres_variable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variable); i {
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
		file_postgres_variable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllVariablesResponse); i {
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
		file_postgres_variable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllVariablesRequest); i {
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
		file_postgres_variable_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariablePrimaryKey); i {
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
			RawDescriptor: file_postgres_variable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postgres_variable_proto_goTypes,
		DependencyIndexes: file_postgres_variable_proto_depIdxs,
		MessageInfos:      file_postgres_variable_proto_msgTypes,
	}.Build()
	File_postgres_variable_proto = out.File
	file_postgres_variable_proto_rawDesc = nil
	file_postgres_variable_proto_goTypes = nil
	file_postgres_variable_proto_depIdxs = nil
}
