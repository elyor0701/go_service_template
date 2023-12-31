// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: function.proto

package object_builder_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreateFunctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string           `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Name        string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Body        *structpb.Struct `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	ProjectId   string           `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CommitId    int64            `protobuf:"varint,6,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	CommitGuid  string           `protobuf:"bytes,7,opt,name=commit_guid,json=commitGuid,proto3" json:"commit_guid,omitempty"`
}

func (x *CreateFunctionRequest) Reset() {
	*x = CreateFunctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFunctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFunctionRequest) ProtoMessage() {}

func (x *CreateFunctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFunctionRequest.ProtoReflect.Descriptor instead.
func (*CreateFunctionRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFunctionRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateFunctionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFunctionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateFunctionRequest) GetBody() *structpb.Struct {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *CreateFunctionRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateFunctionRequest) GetCommitId() int64 {
	if x != nil {
		return x.CommitId
	}
	return 0
}

func (x *CreateFunctionRequest) GetCommitGuid() string {
	if x != nil {
		return x.CommitGuid
	}
	return ""
}

type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path        string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Name        string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Body        *structpb.Struct `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	ProjectId   string           `protobuf:"bytes,6,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{1}
}

func (x *Function) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Function) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Function) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Function) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Function) GetBody() *structpb.Struct {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Function) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetAllFunctionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search    string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Limit     int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetAllFunctionsRequest) Reset() {
	*x = GetAllFunctionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFunctionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFunctionsRequest) ProtoMessage() {}

func (x *GetAllFunctionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFunctionsRequest.ProtoReflect.Descriptor instead.
func (*GetAllFunctionsRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllFunctionsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllFunctionsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllFunctionsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetAllFunctionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions []*Function `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty"`
	Count     int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllFunctionsResponse) Reset() {
	*x = GetAllFunctionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFunctionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFunctionsResponse) ProtoMessage() {}

func (x *GetAllFunctionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFunctionsResponse.ProtoReflect.Descriptor instead.
func (*GetAllFunctionsResponse) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllFunctionsResponse) GetFunctions() []*Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *GetAllFunctionsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type FunctionPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *FunctionPrimaryKey) Reset() {
	*x = FunctionPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionPrimaryKey) ProtoMessage() {}

func (x *FunctionPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionPrimaryKey.ProtoReflect.Descriptor instead.
func (*FunctionPrimaryKey) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{4}
}

func (x *FunctionPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FunctionPrimaryKey) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_function_proto protoreflect.FileDescriptor

var file_function_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x47, 0x75, 0x69,
	0x64, 0x22, 0xb0, 0x01, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x12,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x32, 0xcf, 0x03, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x2d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x6c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x2a, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x20, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2a, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_function_proto_rawDescOnce sync.Once
	file_function_proto_rawDescData = file_function_proto_rawDesc
)

func file_function_proto_rawDescGZIP() []byte {
	file_function_proto_rawDescOnce.Do(func() {
		file_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_function_proto_rawDescData)
	})
	return file_function_proto_rawDescData
}

var file_function_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_function_proto_goTypes = []interface{}{
	(*CreateFunctionRequest)(nil),   // 0: object_builder_service.CreateFunctionRequest
	(*Function)(nil),                // 1: object_builder_service.Function
	(*GetAllFunctionsRequest)(nil),  // 2: object_builder_service.GetAllFunctionsRequest
	(*GetAllFunctionsResponse)(nil), // 3: object_builder_service.GetAllFunctionsResponse
	(*FunctionPrimaryKey)(nil),      // 4: object_builder_service.FunctionPrimaryKey
	(*structpb.Struct)(nil),         // 5: google.protobuf.Struct
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
}
var file_function_proto_depIdxs = []int32{
	5, // 0: object_builder_service.CreateFunctionRequest.body:type_name -> google.protobuf.Struct
	5, // 1: object_builder_service.Function.body:type_name -> google.protobuf.Struct
	1, // 2: object_builder_service.GetAllFunctionsResponse.functions:type_name -> object_builder_service.Function
	0, // 3: object_builder_service.FunctionService.Create:input_type -> object_builder_service.CreateFunctionRequest
	2, // 4: object_builder_service.FunctionService.GetList:input_type -> object_builder_service.GetAllFunctionsRequest
	4, // 5: object_builder_service.FunctionService.GetSingle:input_type -> object_builder_service.FunctionPrimaryKey
	1, // 6: object_builder_service.FunctionService.Update:input_type -> object_builder_service.Function
	4, // 7: object_builder_service.FunctionService.Delete:input_type -> object_builder_service.FunctionPrimaryKey
	1, // 8: object_builder_service.FunctionService.Create:output_type -> object_builder_service.Function
	3, // 9: object_builder_service.FunctionService.GetList:output_type -> object_builder_service.GetAllFunctionsResponse
	1, // 10: object_builder_service.FunctionService.GetSingle:output_type -> object_builder_service.Function
	6, // 11: object_builder_service.FunctionService.Update:output_type -> google.protobuf.Empty
	6, // 12: object_builder_service.FunctionService.Delete:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_function_proto_init() }
func file_function_proto_init() {
	if File_function_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFunctionRequest); i {
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
		file_function_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
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
		file_function_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFunctionsRequest); i {
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
		file_function_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFunctionsResponse); i {
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
		file_function_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionPrimaryKey); i {
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
			RawDescriptor: file_function_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_function_proto_goTypes,
		DependencyIndexes: file_function_proto_depIdxs,
		MessageInfos:      file_function_proto_msgTypes,
	}.Build()
	File_function_proto = out.File
	file_function_proto_rawDesc = nil
	file_function_proto_goTypes = nil
	file_function_proto_depIdxs = nil
}
