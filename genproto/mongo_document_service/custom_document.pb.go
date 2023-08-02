// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: custom_document.proto

package mongo_document_service

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

type DocumentMon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentDoc *DocumentDoc `protobuf:"bytes,1,opt,name=DocumentDoc,proto3" json:"DocumentDoc,omitempty"`
	Contract    *ContractDoc `protobuf:"bytes,2,opt,name=Contract,proto3" json:"Contract,omitempty"`
	DocumentId  string       `protobuf:"bytes,3,opt,name=DocumentId,proto3" json:"DocumentId,omitempty"`
	Owner       *Company     `protobuf:"bytes,4,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Partner     *Company     `protobuf:"bytes,5,opt,name=Partner,proto3" json:"Partner,omitempty"`
	Info        string       `protobuf:"bytes,6,opt,name=Info,proto3" json:"Info,omitempty"`
	File        string       `protobuf:"bytes,7,opt,name=File,proto3" json:"File,omitempty"`
}

func (x *DocumentMon) Reset() {
	*x = DocumentMon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentMon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentMon) ProtoMessage() {}

func (x *DocumentMon) ProtoReflect() protoreflect.Message {
	mi := &file_custom_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentMon.ProtoReflect.Descriptor instead.
func (*DocumentMon) Descriptor() ([]byte, []int) {
	return file_custom_document_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentMon) GetDocumentDoc() *DocumentDoc {
	if x != nil {
		return x.DocumentDoc
	}
	return nil
}

func (x *DocumentMon) GetContract() *ContractDoc {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *DocumentMon) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *DocumentMon) GetOwner() *Company {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *DocumentMon) GetPartner() *Company {
	if x != nil {
		return x.Partner
	}
	return nil
}

func (x *DocumentMon) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *DocumentMon) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document  *DocumentMon `protobuf:"bytes,1,opt,name=Document,proto3" json:"Document,omitempty"`
	Status    string       `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Sign      string       `protobuf:"bytes,3,opt,name=Sign,proto3" json:"Sign,omitempty"`
	UzPdfUrl  string       `protobuf:"bytes,4,opt,name=UzPdfUrl,proto3" json:"UzPdfUrl,omitempty"`
	RuPdfUrl  string       `protobuf:"bytes,5,opt,name=RuPdfUrl,proto3" json:"RuPdfUrl,omitempty"`
	CreatedAt string       `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt string       `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	HtmlUz    string       `protobuf:"bytes,8,opt,name=HtmlUz,proto3" json:"HtmlUz,omitempty"`
	HtmlRu    string       `protobuf:"bytes,9,opt,name=HtmlRu,proto3" json:"HtmlRu,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_custom_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_custom_document_proto_rawDescGZIP(), []int{1}
}

func (x *Document) GetDocument() *DocumentMon {
	if x != nil {
		return x.Document
	}
	return nil
}

func (x *Document) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Document) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *Document) GetUzPdfUrl() string {
	if x != nil {
		return x.UzPdfUrl
	}
	return ""
}

func (x *Document) GetRuPdfUrl() string {
	if x != nil {
		return x.RuPdfUrl
	}
	return ""
}

func (x *Document) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Document) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Document) GetHtmlUz() string {
	if x != nil {
		return x.HtmlUz
	}
	return ""
}

func (x *Document) GetHtmlRu() string {
	if x != nil {
		return x.HtmlRu
	}
	return ""
}

type DocumentDoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentNo   string `protobuf:"bytes,1,opt,name=DocumentNo,proto3" json:"DocumentNo,omitempty"`
	DocumentDate string `protobuf:"bytes,2,opt,name=DocumentDate,proto3" json:"DocumentDate,omitempty"`
	DocumentName string `protobuf:"bytes,3,opt,name=DocumentName,proto3" json:"DocumentName,omitempty"`
}

func (x *DocumentDoc) Reset() {
	*x = DocumentDoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_document_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentDoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentDoc) ProtoMessage() {}

func (x *DocumentDoc) ProtoReflect() protoreflect.Message {
	mi := &file_custom_document_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentDoc.ProtoReflect.Descriptor instead.
func (*DocumentDoc) Descriptor() ([]byte, []int) {
	return file_custom_document_proto_rawDescGZIP(), []int{2}
}

func (x *DocumentDoc) GetDocumentNo() string {
	if x != nil {
		return x.DocumentNo
	}
	return ""
}

func (x *DocumentDoc) GetDocumentDate() string {
	if x != nil {
		return x.DocumentDate
	}
	return ""
}

func (x *DocumentDoc) GetDocumentName() string {
	if x != nil {
		return x.DocumentName
	}
	return ""
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inn      string `protobuf:"bytes,1,opt,name=Inn,proto3" json:"Inn,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=Location,proto3" json:"Location,omitempty"`
	WithVat  string `protobuf:"bytes,4,opt,name=WithVat,proto3" json:"WithVat,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_document_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_custom_document_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_custom_document_proto_rawDescGZIP(), []int{3}
}

func (x *Company) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Company) GetWithVat() string {
	if x != nil {
		return x.WithVat
	}
	return ""
}

var File_custom_document_proto protoreflect.FileDescriptor

var file_custom_document_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02, 0x0a, 0x0b, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0b, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x6f, 0x63, 0x52, 0x0b, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63,
	0x12, 0x3f, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x35, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x07, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x9b, 0x02, 0x0a, 0x08,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x52,
	0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x7a, 0x50, 0x64, 0x66, 0x55, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x7a, 0x50, 0x64, 0x66, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x75, 0x50, 0x64, 0x66, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x75, 0x50, 0x64, 0x66, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x74, 0x6d,
	0x6c, 0x55, 0x7a, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x74, 0x6d, 0x6c, 0x55,
	0x7a, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x74, 0x6d, 0x6c, 0x52, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x48, 0x74, 0x6d, 0x6c, 0x52, 0x75, 0x22, 0x75, 0x0a, 0x0b, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x65, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x49,
	0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x6e, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x57, 0x69, 0x74, 0x68, 0x56, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x57, 0x69, 0x74, 0x68, 0x56, 0x61, 0x74, 0x32, 0xb8, 0x02, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x27, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_custom_document_proto_rawDescOnce sync.Once
	file_custom_document_proto_rawDescData = file_custom_document_proto_rawDesc
)

func file_custom_document_proto_rawDescGZIP() []byte {
	file_custom_document_proto_rawDescOnce.Do(func() {
		file_custom_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_custom_document_proto_rawDescData)
	})
	return file_custom_document_proto_rawDescData
}

var file_custom_document_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_custom_document_proto_goTypes = []interface{}{
	(*DocumentMon)(nil),     // 0: mongo_document_service.DocumentMon
	(*Document)(nil),        // 1: mongo_document_service.Document
	(*DocumentDoc)(nil),     // 2: mongo_document_service.DocumentDoc
	(*Company)(nil),         // 3: mongo_document_service.Company
	(*ContractDoc)(nil),     // 4: mongo_document_service.ContractDoc
	(*ById)(nil),            // 5: mongo_document_service.ById
	(*UpdateStatusReq)(nil), // 6: mongo_document_service.UpdateStatusReq
	(*emptypb.Empty)(nil),   // 7: google.protobuf.Empty
}
var file_custom_document_proto_depIdxs = []int32{
	2, // 0: mongo_document_service.DocumentMon.DocumentDoc:type_name -> mongo_document_service.DocumentDoc
	4, // 1: mongo_document_service.DocumentMon.Contract:type_name -> mongo_document_service.ContractDoc
	3, // 2: mongo_document_service.DocumentMon.Owner:type_name -> mongo_document_service.Company
	3, // 3: mongo_document_service.DocumentMon.Partner:type_name -> mongo_document_service.Company
	0, // 4: mongo_document_service.Document.Document:type_name -> mongo_document_service.DocumentMon
	1, // 5: mongo_document_service.CustomDocument.Upsert:input_type -> mongo_document_service.Document
	5, // 6: mongo_document_service.CustomDocument.GetById:input_type -> mongo_document_service.ById
	5, // 7: mongo_document_service.CustomDocument.Delete:input_type -> mongo_document_service.ById
	6, // 8: mongo_document_service.CustomDocument.UpdateStatus:input_type -> mongo_document_service.UpdateStatusReq
	7, // 9: mongo_document_service.CustomDocument.Upsert:output_type -> google.protobuf.Empty
	1, // 10: mongo_document_service.CustomDocument.GetById:output_type -> mongo_document_service.Document
	7, // 11: mongo_document_service.CustomDocument.Delete:output_type -> google.protobuf.Empty
	7, // 12: mongo_document_service.CustomDocument.UpdateStatus:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_custom_document_proto_init() }
func file_custom_document_proto_init() {
	if File_custom_document_proto != nil {
		return
	}
	file_mongo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_custom_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentMon); i {
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
		file_custom_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_custom_document_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentDoc); i {
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
		file_custom_document_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
			RawDescriptor: file_custom_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_custom_document_proto_goTypes,
		DependencyIndexes: file_custom_document_proto_depIdxs,
		MessageInfos:      file_custom_document_proto_msgTypes,
	}.Build()
	File_custom_document_proto = out.File
	file_custom_document_proto_rawDesc = nil
	file_custom_document_proto_goTypes = nil
	file_custom_document_proto_depIdxs = nil
}
