// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: fields_and_relations.proto

package object_builder_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateFieldsAndRelationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string                   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	VersionId string                   `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	CommitId  string                   `protobuf:"bytes,3,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	TableId   string                   `protobuf:"bytes,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Fields    []*CreateFieldRequest    `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Relations []*CreateRelationRequest `protobuf:"bytes,6,rep,name=relations,proto3" json:"relations,omitempty"`
}

func (x *CreateFieldsAndRelationsRequest) Reset() {
	*x = CreateFieldsAndRelationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fields_and_relations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFieldsAndRelationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFieldsAndRelationsRequest) ProtoMessage() {}

func (x *CreateFieldsAndRelationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fields_and_relations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFieldsAndRelationsRequest.ProtoReflect.Descriptor instead.
func (*CreateFieldsAndRelationsRequest) Descriptor() ([]byte, []int) {
	return file_fields_and_relations_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFieldsAndRelationsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateFieldsAndRelationsRequest) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *CreateFieldsAndRelationsRequest) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *CreateFieldsAndRelationsRequest) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *CreateFieldsAndRelationsRequest) GetFields() []*CreateFieldRequest {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *CreateFieldsAndRelationsRequest) GetRelations() []*CreateRelationRequest {
	if x != nil {
		return x.Relations
	}
	return nil
}

type CreateFieldsAndRelationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFieldsAndRelationsResponse) Reset() {
	*x = CreateFieldsAndRelationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fields_and_relations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFieldsAndRelationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFieldsAndRelationsResponse) ProtoMessage() {}

func (x *CreateFieldsAndRelationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fields_and_relations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFieldsAndRelationsResponse.ProtoReflect.Descriptor instead.
func (*CreateFieldsAndRelationsResponse) Descriptor() ([]byte, []int) {
	return file_fields_and_relations_proto_rawDescGZIP(), []int{1}
}

var File_fields_and_relations_proto protoreflect.FileDescriptor

var file_fields_and_relations_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a,
	0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x41, 0x6e, 0x64,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x4b, 0x0a, 0x09, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x22, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xab, 0x01, 0x0a, 0x17,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fields_and_relations_proto_rawDescOnce sync.Once
	file_fields_and_relations_proto_rawDescData = file_fields_and_relations_proto_rawDesc
)

func file_fields_and_relations_proto_rawDescGZIP() []byte {
	file_fields_and_relations_proto_rawDescOnce.Do(func() {
		file_fields_and_relations_proto_rawDescData = protoimpl.X.CompressGZIP(file_fields_and_relations_proto_rawDescData)
	})
	return file_fields_and_relations_proto_rawDescData
}

var file_fields_and_relations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fields_and_relations_proto_goTypes = []interface{}{
	(*CreateFieldsAndRelationsRequest)(nil),  // 0: object_builder_service.CreateFieldsAndRelationsRequest
	(*CreateFieldsAndRelationsResponse)(nil), // 1: object_builder_service.CreateFieldsAndRelationsResponse
	(*CreateFieldRequest)(nil),               // 2: object_builder_service.CreateFieldRequest
	(*CreateRelationRequest)(nil),            // 3: object_builder_service.CreateRelationRequest
}
var file_fields_and_relations_proto_depIdxs = []int32{
	2, // 0: object_builder_service.CreateFieldsAndRelationsRequest.fields:type_name -> object_builder_service.CreateFieldRequest
	3, // 1: object_builder_service.CreateFieldsAndRelationsRequest.relations:type_name -> object_builder_service.CreateRelationRequest
	0, // 2: object_builder_service.FieldAndRelationService.CreateFieldsAndRelations:input_type -> object_builder_service.CreateFieldsAndRelationsRequest
	1, // 3: object_builder_service.FieldAndRelationService.CreateFieldsAndRelations:output_type -> object_builder_service.CreateFieldsAndRelationsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fields_and_relations_proto_init() }
func file_fields_and_relations_proto_init() {
	if File_fields_and_relations_proto != nil {
		return
	}
	file_field_proto_init()
	file_relation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fields_and_relations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFieldsAndRelationsRequest); i {
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
		file_fields_and_relations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFieldsAndRelationsResponse); i {
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
			RawDescriptor: file_fields_and_relations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fields_and_relations_proto_goTypes,
		DependencyIndexes: file_fields_and_relations_proto_depIdxs,
		MessageInfos:      file_fields_and_relations_proto_msgTypes,
	}.Build()
	File_fields_and_relations_proto = out.File
	file_fields_and_relations_proto_rawDesc = nil
	file_fields_and_relations_proto_goTypes = nil
	file_fields_and_relations_proto_depIdxs = nil
}
