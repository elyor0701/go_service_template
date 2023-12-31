// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: postgres_cascading.proto

package postgres_object_builder_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCascadingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableSlug string `protobuf:"bytes,1,opt,name=table_slug,json=tableSlug,proto3" json:"table_slug,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetCascadingRequest) Reset() {
	*x = GetCascadingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postgres_cascading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCascadingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCascadingRequest) ProtoMessage() {}

func (x *GetCascadingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postgres_cascading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCascadingRequest.ProtoReflect.Descriptor instead.
func (*GetCascadingRequest) Descriptor() ([]byte, []int) {
	return file_postgres_cascading_proto_rawDescGZIP(), []int{0}
}

func (x *GetCascadingRequest) GetTableSlug() string {
	if x != nil {
		return x.TableSlug
	}
	return ""
}

func (x *GetCascadingRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_postgres_cascading_proto protoreflect.FileDescriptor

var file_postgres_cascading_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x63, 0x61, 0x73, 0x63, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1d, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x73, 0x63, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x75, 0x67,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x32,
	0x8b, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x73, 0x63, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x63, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x34, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x63, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_postgres_cascading_proto_rawDescOnce sync.Once
	file_postgres_cascading_proto_rawDescData = file_postgres_cascading_proto_rawDesc
)

func file_postgres_cascading_proto_rawDescGZIP() []byte {
	file_postgres_cascading_proto_rawDescOnce.Do(func() {
		file_postgres_cascading_proto_rawDescData = protoimpl.X.CompressGZIP(file_postgres_cascading_proto_rawDescData)
	})
	return file_postgres_cascading_proto_rawDescData
}

var file_postgres_cascading_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_postgres_cascading_proto_goTypes = []interface{}{
	(*GetCascadingRequest)(nil), // 0: postgres_object_builder_service.GetCascadingRequest
	(*CommonMessage)(nil),       // 1: postgres_object_builder_service.CommonMessage
}
var file_postgres_cascading_proto_depIdxs = []int32{
	0, // 0: postgres_object_builder_service.CascadingService.GetCascadings:input_type -> postgres_object_builder_service.GetCascadingRequest
	1, // 1: postgres_object_builder_service.CascadingService.GetCascadings:output_type -> postgres_object_builder_service.CommonMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_postgres_cascading_proto_init() }
func file_postgres_cascading_proto_init() {
	if File_postgres_cascading_proto != nil {
		return
	}
	file_postgres_object_builder_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_postgres_cascading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCascadingRequest); i {
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
			RawDescriptor: file_postgres_cascading_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postgres_cascading_proto_goTypes,
		DependencyIndexes: file_postgres_cascading_proto_depIdxs,
		MessageInfos:      file_postgres_cascading_proto_msgTypes,
	}.Build()
	File_postgres_cascading_proto = out.File
	file_postgres_cascading_proto_rawDesc = nil
	file_postgres_cascading_proto_goTypes = nil
	file_postgres_cascading_proto_depIdxs = nil
}
