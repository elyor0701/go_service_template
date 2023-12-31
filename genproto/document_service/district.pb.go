// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: district.proto

package document_service

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

type CreateDistrict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictId   string `protobuf:"bytes,1,opt,name=districtId,proto3" json:"districtId,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lang         string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	DistrictCode int32  `protobuf:"varint,4,opt,name=districtCode,proto3" json:"districtCode,omitempty"`
	RegionId     string `protobuf:"bytes,5,opt,name=regionId,proto3" json:"regionId,omitempty"`
}

func (x *CreateDistrict) Reset() {
	*x = CreateDistrict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDistrict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDistrict) ProtoMessage() {}

func (x *CreateDistrict) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDistrict.ProtoReflect.Descriptor instead.
func (*CreateDistrict) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDistrict) GetDistrictId() string {
	if x != nil {
		return x.DistrictId
	}
	return ""
}

func (x *CreateDistrict) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDistrict) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *CreateDistrict) GetDistrictCode() int32 {
	if x != nil {
		return x.DistrictCode
	}
	return 0
}

func (x *CreateDistrict) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type District struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictId   string `protobuf:"bytes,1,opt,name=districtId,proto3" json:"districtId,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lang         string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	DistrictCode int32  `protobuf:"varint,4,opt,name=districtCode,proto3" json:"districtCode,omitempty"`
	RegionId     int64  `protobuf:"varint,5,opt,name=regionId,proto3" json:"regionId,omitempty"`
}

func (x *District) Reset() {
	*x = District{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *District) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*District) ProtoMessage() {}

func (x *District) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use District.ProtoReflect.Descriptor instead.
func (*District) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{1}
}

func (x *District) GetDistrictId() string {
	if x != nil {
		return x.DistrictId
	}
	return ""
}

func (x *District) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *District) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *District) GetDistrictCode() int32 {
	if x != nil {
		return x.DistrictCode
	}
	return 0
}

func (x *District) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

type DistrictGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictId   string `protobuf:"bytes,1,opt,name=districtId,proto3" json:"districtId,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lang         string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	DistrictCode int32  `protobuf:"varint,4,opt,name=districtCode,proto3" json:"districtCode,omitempty"`
	RegionId     string `protobuf:"bytes,5,opt,name=regionId,proto3" json:"regionId,omitempty"`
}

func (x *DistrictGet) Reset() {
	*x = DistrictGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistrictGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistrictGet) ProtoMessage() {}

func (x *DistrictGet) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistrictGet.ProtoReflect.Descriptor instead.
func (*DistrictGet) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{2}
}

func (x *DistrictGet) GetDistrictId() string {
	if x != nil {
		return x.DistrictId
	}
	return ""
}

func (x *DistrictGet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DistrictGet) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *DistrictGet) GetDistrictCode() int32 {
	if x != nil {
		return x.DistrictCode
	}
	return 0
}

func (x *DistrictGet) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type GetListDistrictsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Lang   string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *GetListDistrictsRequest) Reset() {
	*x = GetListDistrictsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDistrictsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDistrictsRequest) ProtoMessage() {}

func (x *GetListDistrictsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDistrictsRequest.ProtoReflect.Descriptor instead.
func (*GetListDistrictsRequest) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{3}
}

func (x *GetListDistrictsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListDistrictsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListDistrictsRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type GetListDistrictsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Districts []*District `protobuf:"bytes,1,rep,name=districts,proto3" json:"districts,omitempty"`
	Total     int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetListDistrictsResponse) Reset() {
	*x = GetListDistrictsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDistrictsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDistrictsResponse) ProtoMessage() {}

func (x *GetListDistrictsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDistrictsResponse.ProtoReflect.Descriptor instead.
func (*GetListDistrictsResponse) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{4}
}

func (x *GetListDistrictsResponse) GetDistricts() []*District {
	if x != nil {
		return x.Districts
	}
	return nil
}

func (x *GetListDistrictsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DistrictId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictId string `protobuf:"bytes,1,opt,name=districtId,proto3" json:"districtId,omitempty"`
}

func (x *DistrictId) Reset() {
	*x = DistrictId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistrictId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistrictId) ProtoMessage() {}

func (x *DistrictId) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistrictId.ProtoReflect.Descriptor instead.
func (*DistrictId) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{5}
}

func (x *DistrictId) GetDistrictId() string {
	if x != nil {
		return x.DistrictId
	}
	return ""
}

type ReloadDistrictsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Districts []*District `protobuf:"bytes,1,rep,name=districts,proto3" json:"districts,omitempty"`
	Lang      string      `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *ReloadDistrictsListRequest) Reset() {
	*x = ReloadDistrictsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadDistrictsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadDistrictsListRequest) ProtoMessage() {}

func (x *ReloadDistrictsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadDistrictsListRequest.ProtoReflect.Descriptor instead.
func (*ReloadDistrictsListRequest) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{6}
}

func (x *ReloadDistrictsListRequest) GetDistricts() []*District {
	if x != nil {
		return x.Districts
	}
	return nil
}

func (x *ReloadDistrictsListRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type DeleteDistrictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictId string `protobuf:"bytes,1,opt,name=districtId,proto3" json:"districtId,omitempty"`
	Lang       string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *DeleteDistrictRequest) Reset() {
	*x = DeleteDistrictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDistrictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDistrictRequest) ProtoMessage() {}

func (x *DeleteDistrictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDistrictRequest.ProtoReflect.Descriptor instead.
func (*DeleteDistrictRequest) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDistrictRequest) GetDistrictId() string {
	if x != nil {
		return x.DistrictId
	}
	return ""
}

func (x *DeleteDistrictRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type DeleteDistrict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowAffect int64 `protobuf:"varint,1,opt,name=rowAffect,proto3" json:"rowAffect,omitempty"`
}

func (x *DeleteDistrict) Reset() {
	*x = DeleteDistrict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDistrict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDistrict) ProtoMessage() {}

func (x *DeleteDistrict) ProtoReflect() protoreflect.Message {
	mi := &file_district_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDistrict.ProtoReflect.Descriptor instead.
func (*DeleteDistrict) Descriptor() ([]byte, []int) {
	return file_district_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDistrict) GetRowAffect() int64 {
	if x != nil {
		return x.RowAffect
	}
	return 0
}

var File_district_proto protoreflect.FileDescriptor

var file_district_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x98, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x95, 0x01, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x47, 0x65, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x22, 0x6a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52,
	0x09, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x2c, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x22, 0x6a,
	0x0a, 0x1a, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x09, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x2e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x77,
	0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x6f,
	0x77, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x32, 0xb1, 0x03, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x47, 0x65, 0x74, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49,
	0x64, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x47, 0x65, 0x74,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x06, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x2e, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27,
	0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_district_proto_rawDescOnce sync.Once
	file_district_proto_rawDescData = file_district_proto_rawDesc
)

func file_district_proto_rawDescGZIP() []byte {
	file_district_proto_rawDescOnce.Do(func() {
		file_district_proto_rawDescData = protoimpl.X.CompressGZIP(file_district_proto_rawDescData)
	})
	return file_district_proto_rawDescData
}

var file_district_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_district_proto_goTypes = []interface{}{
	(*CreateDistrict)(nil),             // 0: document_service.CreateDistrict
	(*District)(nil),                   // 1: document_service.District
	(*DistrictGet)(nil),                // 2: document_service.DistrictGet
	(*GetListDistrictsRequest)(nil),    // 3: document_service.GetListDistrictsRequest
	(*GetListDistrictsResponse)(nil),   // 4: document_service.GetListDistrictsResponse
	(*DistrictId)(nil),                 // 5: document_service.DistrictId
	(*ReloadDistrictsListRequest)(nil), // 6: document_service.ReloadDistrictsListRequest
	(*DeleteDistrictRequest)(nil),      // 7: document_service.DeleteDistrictRequest
	(*DeleteDistrict)(nil),             // 8: document_service.DeleteDistrict
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_district_proto_depIdxs = []int32{
	1, // 0: document_service.GetListDistrictsResponse.districts:type_name -> document_service.District
	1, // 1: document_service.ReloadDistrictsListRequest.districts:type_name -> document_service.District
	0, // 2: document_service.DistrictService.Create:input_type -> document_service.CreateDistrict
	3, // 3: document_service.DistrictService.GetList:input_type -> document_service.GetListDistrictsRequest
	5, // 4: document_service.DistrictService.Get:input_type -> document_service.DistrictId
	6, // 5: document_service.DistrictService.Reload:input_type -> document_service.ReloadDistrictsListRequest
	7, // 6: document_service.DistrictService.Delete:input_type -> document_service.DeleteDistrictRequest
	2, // 7: document_service.DistrictService.Create:output_type -> document_service.DistrictGet
	4, // 8: document_service.DistrictService.GetList:output_type -> document_service.GetListDistrictsResponse
	2, // 9: document_service.DistrictService.Get:output_type -> document_service.DistrictGet
	9, // 10: document_service.DistrictService.Reload:output_type -> google.protobuf.Empty
	8, // 11: document_service.DistrictService.Delete:output_type -> document_service.DeleteDistrict
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_district_proto_init() }
func file_district_proto_init() {
	if File_district_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_district_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDistrict); i {
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
		file_district_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*District); i {
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
		file_district_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistrictGet); i {
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
		file_district_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDistrictsRequest); i {
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
		file_district_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDistrictsResponse); i {
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
		file_district_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistrictId); i {
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
		file_district_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadDistrictsListRequest); i {
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
		file_district_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDistrictRequest); i {
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
		file_district_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDistrict); i {
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
			RawDescriptor: file_district_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_district_proto_goTypes,
		DependencyIndexes: file_district_proto_depIdxs,
		MessageInfos:      file_district_proto_msgTypes,
	}.Build()
	File_district_proto = out.File
	file_district_proto_rawDesc = nil
	file_district_proto_goTypes = nil
	file_district_proto_depIdxs = nil
}
