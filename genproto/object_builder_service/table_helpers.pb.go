// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: table_helpers.proto

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

type ImportFromJSONRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	AppId     string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ImportFromJSONRequest) Reset() {
	*x = ImportFromJSONRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_helpers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportFromJSONRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportFromJSONRequest) ProtoMessage() {}

func (x *ImportFromJSONRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_helpers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportFromJSONRequest.ProtoReflect.Descriptor instead.
func (*ImportFromJSONRequest) Descriptor() ([]byte, []int) {
	return file_table_helpers_proto_rawDescGZIP(), []int{0}
}

func (x *ImportFromJSONRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ImportFromJSONRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ImportFromJSONRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ExportToJSONRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	VersionId string `protobuf:"bytes,3,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
}

func (x *ExportToJSONRequest) Reset() {
	*x = ExportToJSONRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_helpers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportToJSONRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportToJSONRequest) ProtoMessage() {}

func (x *ExportToJSONRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_helpers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportToJSONRequest.ProtoReflect.Descriptor instead.
func (*ExportToJSONRequest) Descriptor() ([]byte, []int) {
	return file_table_helpers_proto_rawDescGZIP(), []int{1}
}

func (x *ExportToJSONRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ExportToJSONRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ExportToJSONRequest) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

type ExportToJSONReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *ExportToJSONReponse) Reset() {
	*x = ExportToJSONReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_helpers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportToJSONReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportToJSONReponse) ProtoMessage() {}

func (x *ExportToJSONReponse) ProtoReflect() protoreflect.Message {
	mi := &file_table_helpers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportToJSONReponse.ProtoReflect.Descriptor instead.
func (*ExportToJSONReponse) Descriptor() ([]byte, []int) {
	return file_table_helpers_proto_rawDescGZIP(), []int{2}
}

func (x *ExportToJSONReponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_table_helpers_proto protoreflect.FileDescriptor

var file_table_helpers_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x15, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x53, 0x4f, 0x4e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x6f, 0x4a, 0x53, 0x4f, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x4a, 0x53,
	0x4f, 0x4e, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0xdc, 0x01,
	0x0a, 0x13, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x6f, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x4a, 0x53, 0x4f, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x6f, 0x4a, 0x53, 0x4f, 0x4e, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4a,
	0x53, 0x4f, 0x4e, 0x12, 0x2d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x53, 0x4f, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_table_helpers_proto_rawDescOnce sync.Once
	file_table_helpers_proto_rawDescData = file_table_helpers_proto_rawDesc
)

func file_table_helpers_proto_rawDescGZIP() []byte {
	file_table_helpers_proto_rawDescOnce.Do(func() {
		file_table_helpers_proto_rawDescData = protoimpl.X.CompressGZIP(file_table_helpers_proto_rawDescData)
	})
	return file_table_helpers_proto_rawDescData
}

var file_table_helpers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_table_helpers_proto_goTypes = []interface{}{
	(*ImportFromJSONRequest)(nil), // 0: object_builder_service.ImportFromJSONRequest
	(*ExportToJSONRequest)(nil),   // 1: object_builder_service.ExportToJSONRequest
	(*ExportToJSONReponse)(nil),   // 2: object_builder_service.ExportToJSONReponse
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_table_helpers_proto_depIdxs = []int32{
	1, // 0: object_builder_service.TableHelpersService.ExportToJSON:input_type -> object_builder_service.ExportToJSONRequest
	0, // 1: object_builder_service.TableHelpersService.ImportFromJSON:input_type -> object_builder_service.ImportFromJSONRequest
	2, // 2: object_builder_service.TableHelpersService.ExportToJSON:output_type -> object_builder_service.ExportToJSONReponse
	3, // 3: object_builder_service.TableHelpersService.ImportFromJSON:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_table_helpers_proto_init() }
func file_table_helpers_proto_init() {
	if File_table_helpers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_table_helpers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportFromJSONRequest); i {
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
		file_table_helpers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportToJSONRequest); i {
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
		file_table_helpers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportToJSONReponse); i {
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
			RawDescriptor: file_table_helpers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_table_helpers_proto_goTypes,
		DependencyIndexes: file_table_helpers_proto_depIdxs,
		MessageInfos:      file_table_helpers_proto_msgTypes,
	}.Build()
	File_table_helpers_proto = out.File
	file_table_helpers_proto_rawDesc = nil
	file_table_helpers_proto_goTypes = nil
	file_table_helpers_proto_depIdxs = nil
}
