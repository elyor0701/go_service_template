// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: query_folder.proto

package object_builder_service

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

type CreateQueryFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ParentId  string `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *CreateQueryFolderRequest) Reset() {
	*x = CreateQueryFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_folder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQueryFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQueryFolderRequest) ProtoMessage() {}

func (x *CreateQueryFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_folder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQueryFolderRequest.ProtoReflect.Descriptor instead.
func (*CreateQueryFolderRequest) Descriptor() ([]byte, []int) {
	return file_query_folder_proto_rawDescGZIP(), []int{0}
}

func (x *CreateQueryFolderRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateQueryFolderRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CreateQueryFolderRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type QueryFolder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ParentId  string `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *QueryFolder) Reset() {
	*x = QueryFolder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_folder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFolder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFolder) ProtoMessage() {}

func (x *QueryFolder) ProtoReflect() protoreflect.Message {
	mi := &file_query_folder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFolder.ProtoReflect.Descriptor instead.
func (*QueryFolder) Descriptor() ([]byte, []int) {
	return file_query_folder_proto_rawDescGZIP(), []int{1}
}

func (x *QueryFolder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueryFolder) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *QueryFolder) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *QueryFolder) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type QueryFolderId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *QueryFolderId) Reset() {
	*x = QueryFolderId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_folder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFolderId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFolderId) ProtoMessage() {}

func (x *QueryFolderId) ProtoReflect() protoreflect.Message {
	mi := &file_query_folder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFolderId.ProtoReflect.Descriptor instead.
func (*QueryFolderId) Descriptor() ([]byte, []int) {
	return file_query_folder_proto_rawDescGZIP(), []int{2}
}

func (x *QueryFolderId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueryFolderId) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetAllQueryFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId  string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Limit     int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Search    string `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
	ProjectId string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetAllQueryFolderRequest) Reset() {
	*x = GetAllQueryFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_folder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllQueryFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllQueryFolderRequest) ProtoMessage() {}

func (x *GetAllQueryFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_folder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllQueryFolderRequest.ProtoReflect.Descriptor instead.
func (*GetAllQueryFolderRequest) Descriptor() ([]byte, []int) {
	return file_query_folder_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllQueryFolderRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetAllQueryFolderRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllQueryFolderRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllQueryFolderRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllQueryFolderRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetAllQueryFolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folders []*QueryFolder `protobuf:"bytes,1,rep,name=folders,proto3" json:"folders,omitempty"`
	Count   int32          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllQueryFolderResponse) Reset() {
	*x = GetAllQueryFolderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_folder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllQueryFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllQueryFolderResponse) ProtoMessage() {}

func (x *GetAllQueryFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_folder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllQueryFolderResponse.ProtoReflect.Descriptor instead.
func (*GetAllQueryFolderResponse) Descriptor() ([]byte, []int) {
	return file_query_folder_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllQueryFolderResponse) GetFolders() []*QueryFolder {
	if x != nil {
		return x.Folders
	}
	return nil
}

func (x *GetAllQueryFolderResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_query_folder_proto protoreflect.FileDescriptor

var file_query_folder_proto_rawDesc = []byte{
	0x0a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd5, 0x03, 0x0a, 0x12, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x61, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x30, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x25, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x25, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_folder_proto_rawDescOnce sync.Once
	file_query_folder_proto_rawDescData = file_query_folder_proto_rawDesc
)

func file_query_folder_proto_rawDescGZIP() []byte {
	file_query_folder_proto_rawDescOnce.Do(func() {
		file_query_folder_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_folder_proto_rawDescData)
	})
	return file_query_folder_proto_rawDescData
}

var file_query_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_query_folder_proto_goTypes = []interface{}{
	(*CreateQueryFolderRequest)(nil),  // 0: object_builder_service.CreateQueryFolderRequest
	(*QueryFolder)(nil),               // 1: object_builder_service.QueryFolder
	(*QueryFolderId)(nil),             // 2: object_builder_service.QueryFolderId
	(*GetAllQueryFolderRequest)(nil),  // 3: object_builder_service.GetAllQueryFolderRequest
	(*GetAllQueryFolderResponse)(nil), // 4: object_builder_service.GetAllQueryFolderResponse
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_query_folder_proto_depIdxs = []int32{
	1, // 0: object_builder_service.GetAllQueryFolderResponse.folders:type_name -> object_builder_service.QueryFolder
	0, // 1: object_builder_service.QueryFolderService.Create:input_type -> object_builder_service.CreateQueryFolderRequest
	3, // 2: object_builder_service.QueryFolderService.GetAll:input_type -> object_builder_service.GetAllQueryFolderRequest
	2, // 3: object_builder_service.QueryFolderService.GetById:input_type -> object_builder_service.QueryFolderId
	1, // 4: object_builder_service.QueryFolderService.Update:input_type -> object_builder_service.QueryFolder
	2, // 5: object_builder_service.QueryFolderService.Delete:input_type -> object_builder_service.QueryFolderId
	1, // 6: object_builder_service.QueryFolderService.Create:output_type -> object_builder_service.QueryFolder
	4, // 7: object_builder_service.QueryFolderService.GetAll:output_type -> object_builder_service.GetAllQueryFolderResponse
	1, // 8: object_builder_service.QueryFolderService.GetById:output_type -> object_builder_service.QueryFolder
	5, // 9: object_builder_service.QueryFolderService.Update:output_type -> google.protobuf.Empty
	5, // 10: object_builder_service.QueryFolderService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_query_folder_proto_init() }
func file_query_folder_proto_init() {
	if File_query_folder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_query_folder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQueryFolderRequest); i {
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
		file_query_folder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFolder); i {
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
		file_query_folder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFolderId); i {
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
		file_query_folder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllQueryFolderRequest); i {
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
		file_query_folder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllQueryFolderResponse); i {
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
			RawDescriptor: file_query_folder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_folder_proto_goTypes,
		DependencyIndexes: file_query_folder_proto_depIdxs,
		MessageInfos:      file_query_folder_proto_msgTypes,
	}.Build()
	File_query_folder_proto = out.File
	file_query_folder_proto_rawDesc = nil
	file_query_folder_proto_goTypes = nil
	file_query_folder_proto_depIdxs = nil
}
