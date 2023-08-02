// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: table_folder.proto

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

type TableFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ParentId  string `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

func (x *TableFolderRequest) Reset() {
	*x = TableFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_folder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableFolderRequest) ProtoMessage() {}

func (x *TableFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_folder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableFolderRequest.ProtoReflect.Descriptor instead.
func (*TableFolderRequest) Descriptor() ([]byte, []int) {
	return file_table_folder_proto_rawDescGZIP(), []int{0}
}

func (x *TableFolderRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TableFolderRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *TableFolderRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type CreateTableFolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *CreateTableFolderResponse) Reset() {
	*x = CreateTableFolderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_folder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableFolderResponse) ProtoMessage() {}

func (x *CreateTableFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_table_folder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableFolderResponse.ProtoReflect.Descriptor instead.
func (*CreateTableFolderResponse) Descriptor() ([]byte, []int) {
	return file_table_folder_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTableFolderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTableFolderResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTableFolderResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type TableFolder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ParentId  string `protobuf:"bytes,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

func (x *TableFolder) Reset() {
	*x = TableFolder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_folder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableFolder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableFolder) ProtoMessage() {}

func (x *TableFolder) ProtoReflect() protoreflect.Message {
	mi := &file_table_folder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableFolder.ProtoReflect.Descriptor instead.
func (*TableFolder) Descriptor() ([]byte, []int) {
	return file_table_folder_proto_rawDescGZIP(), []int{2}
}

func (x *TableFolder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TableFolder) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TableFolder) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *TableFolder) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type GetAllTableFoldersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folders []*TableFolder `protobuf:"bytes,1,rep,name=folders,proto3" json:"folders,omitempty"`
	Count   int32          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllTableFoldersResponse) Reset() {
	*x = GetAllTableFoldersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_folder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTableFoldersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTableFoldersResponse) ProtoMessage() {}

func (x *GetAllTableFoldersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_table_folder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTableFoldersResponse.ProtoReflect.Descriptor instead.
func (*GetAllTableFoldersResponse) Descriptor() ([]byte, []int) {
	return file_table_folder_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTableFoldersResponse) GetFolders() []*TableFolder {
	if x != nil {
		return x.Folders
	}
	return nil
}

func (x *GetAllTableFoldersResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TableFolderPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *TableFolderPrimaryKey) Reset() {
	*x = TableFolderPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_folder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableFolderPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableFolderPrimaryKey) ProtoMessage() {}

func (x *TableFolderPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_table_folder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableFolderPrimaryKey.ProtoReflect.Descriptor instead.
func (*TableFolderPrimaryKey) Descriptor() ([]byte, []int) {
	return file_table_folder_proto_rawDescGZIP(), []int{4}
}

func (x *TableFolderPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TableFolderPrimaryKey) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetAllTableFoldersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset    int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search    string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetAllTableFoldersRequest) Reset() {
	*x = GetAllTableFoldersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_folder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTableFoldersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTableFoldersRequest) ProtoMessage() {}

func (x *GetAllTableFoldersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_folder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTableFoldersRequest.ProtoReflect.Descriptor instead.
func (*GetAllTableFoldersRequest) Descriptor() ([]byte, []int) {
	return file_table_folder_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllTableFoldersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllTableFoldersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllTableFoldersRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllTableFoldersRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_table_folder_proto protoreflect.FileDescriptor

var file_table_folder_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x12, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x60, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x15, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x80, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x32, 0xef, 0x03, 0x0a, 0x12, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x2d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x23,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x31, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_table_folder_proto_rawDescOnce sync.Once
	file_table_folder_proto_rawDescData = file_table_folder_proto_rawDesc
)

func file_table_folder_proto_rawDescGZIP() []byte {
	file_table_folder_proto_rawDescOnce.Do(func() {
		file_table_folder_proto_rawDescData = protoimpl.X.CompressGZIP(file_table_folder_proto_rawDescData)
	})
	return file_table_folder_proto_rawDescData
}

var file_table_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_table_folder_proto_goTypes = []interface{}{
	(*TableFolderRequest)(nil),         // 0: object_builder_service.TableFolderRequest
	(*CreateTableFolderResponse)(nil),  // 1: object_builder_service.CreateTableFolderResponse
	(*TableFolder)(nil),                // 2: object_builder_service.TableFolder
	(*GetAllTableFoldersResponse)(nil), // 3: object_builder_service.GetAllTableFoldersResponse
	(*TableFolderPrimaryKey)(nil),      // 4: object_builder_service.TableFolderPrimaryKey
	(*GetAllTableFoldersRequest)(nil),  // 5: object_builder_service.GetAllTableFoldersRequest
	(*emptypb.Empty)(nil),              // 6: google.protobuf.Empty
}
var file_table_folder_proto_depIdxs = []int32{
	2, // 0: object_builder_service.GetAllTableFoldersResponse.folders:type_name -> object_builder_service.TableFolder
	0, // 1: object_builder_service.TableFolderService.Create:input_type -> object_builder_service.TableFolderRequest
	4, // 2: object_builder_service.TableFolderService.GetByID:input_type -> object_builder_service.TableFolderPrimaryKey
	5, // 3: object_builder_service.TableFolderService.GetAll:input_type -> object_builder_service.GetAllTableFoldersRequest
	2, // 4: object_builder_service.TableFolderService.Update:input_type -> object_builder_service.TableFolder
	4, // 5: object_builder_service.TableFolderService.Delete:input_type -> object_builder_service.TableFolderPrimaryKey
	1, // 6: object_builder_service.TableFolderService.Create:output_type -> object_builder_service.CreateTableFolderResponse
	2, // 7: object_builder_service.TableFolderService.GetByID:output_type -> object_builder_service.TableFolder
	3, // 8: object_builder_service.TableFolderService.GetAll:output_type -> object_builder_service.GetAllTableFoldersResponse
	6, // 9: object_builder_service.TableFolderService.Update:output_type -> google.protobuf.Empty
	6, // 10: object_builder_service.TableFolderService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_table_folder_proto_init() }
func file_table_folder_proto_init() {
	if File_table_folder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_table_folder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableFolderRequest); i {
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
		file_table_folder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableFolderResponse); i {
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
		file_table_folder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableFolder); i {
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
		file_table_folder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTableFoldersResponse); i {
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
		file_table_folder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableFolderPrimaryKey); i {
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
		file_table_folder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTableFoldersRequest); i {
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
			RawDescriptor: file_table_folder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_table_folder_proto_goTypes,
		DependencyIndexes: file_table_folder_proto_depIdxs,
		MessageInfos:      file_table_folder_proto_msgTypes,
	}.Build()
	File_table_folder_proto = out.File
	file_table_folder_proto_rawDesc = nil
	file_table_folder_proto_goTypes = nil
	file_table_folder_proto_depIdxs = nil
}