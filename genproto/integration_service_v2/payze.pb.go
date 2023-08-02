// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: payze.proto

package integration_service_v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PayzeLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method      string                `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	ProjectId   string                `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Environment string                `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	Data        *PayzeLinkRequestData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PayzeLinkRequest) Reset() {
	*x = PayzeLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payze_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayzeLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayzeLinkRequest) ProtoMessage() {}

func (x *PayzeLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payze_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayzeLinkRequest.ProtoReflect.Descriptor instead.
func (*PayzeLinkRequest) Descriptor() ([]byte, []int) {
	return file_payze_proto_rawDescGZIP(), []int{0}
}

func (x *PayzeLinkRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PayzeLinkRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PayzeLinkRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *PayzeLinkRequest) GetData() *PayzeLinkRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PayzeLinkRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount        int32  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Callback      string `protobuf:"bytes,3,opt,name=callback,proto3" json:"callback,omitempty"`
	CallbackError string `protobuf:"bytes,4,opt,name=callbackError,proto3" json:"callbackError,omitempty"`
	Preauthorize  bool   `protobuf:"varint,5,opt,name=preauthorize,proto3" json:"preauthorize,omitempty"`
	Lang          string `protobuf:"bytes,6,opt,name=lang,proto3" json:"lang,omitempty"`
	HookUrl       string `protobuf:"bytes,7,opt,name=hookUrl,proto3" json:"hookUrl,omitempty"`
	HookRefund    bool   `protobuf:"varint,8,opt,name=hookRefund,proto3" json:"hookRefund,omitempty"`
}

func (x *PayzeLinkRequestData) Reset() {
	*x = PayzeLinkRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payze_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayzeLinkRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayzeLinkRequestData) ProtoMessage() {}

func (x *PayzeLinkRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_payze_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayzeLinkRequestData.ProtoReflect.Descriptor instead.
func (*PayzeLinkRequestData) Descriptor() ([]byte, []int) {
	return file_payze_proto_rawDescGZIP(), []int{1}
}

func (x *PayzeLinkRequestData) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayzeLinkRequestData) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PayzeLinkRequestData) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

func (x *PayzeLinkRequestData) GetCallbackError() string {
	if x != nil {
		return x.CallbackError
	}
	return ""
}

func (x *PayzeLinkRequestData) GetPreauthorize() bool {
	if x != nil {
		return x.Preauthorize
	}
	return false
}

func (x *PayzeLinkRequestData) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *PayzeLinkRequestData) GetHookUrl() string {
	if x != nil {
		return x.HookUrl
	}
	return ""
}

func (x *PayzeLinkRequestData) GetHookRefund() bool {
	if x != nil {
		return x.HookRefund
	}
	return false
}

type PayzeLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string           `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string           `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	TransactionId string           `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Action        int32            `protobuf:"varint,5,opt,name=action,proto3" json:"action,omitempty"`
	Currency      string           `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Price         float32          `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
	CreatedDate   string           `protobuf:"bytes,8,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	CreatedAt     string           `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string           `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     string           `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Data          *structpb.Struct `protobuf:"bytes,12,opt,name=data,proto3" json:"data,omitempty"`
	ProductId     string           `protobuf:"bytes,13,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Amount        int32            `protobuf:"varint,14,opt,name=amount,proto3" json:"amount,omitempty"`
	RespId        string           `protobuf:"bytes,15,opt,name=resp_id,json=respId,proto3" json:"resp_id,omitempty"`
}

func (x *PayzeLinkResponse) Reset() {
	*x = PayzeLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payze_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayzeLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayzeLinkResponse) ProtoMessage() {}

func (x *PayzeLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payze_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayzeLinkResponse.ProtoReflect.Descriptor instead.
func (*PayzeLinkResponse) Descriptor() ([]byte, []int) {
	return file_payze_proto_rawDescGZIP(), []int{2}
}

func (x *PayzeLinkResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PayzeLinkResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PayzeLinkResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayzeLinkResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PayzeLinkResponse) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *PayzeLinkResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PayzeLinkResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PayzeLinkResponse) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *PayzeLinkResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PayzeLinkResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PayzeLinkResponse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *PayzeLinkResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PayzeLinkResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *PayzeLinkResponse) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayzeLinkResponse) GetRespId() string {
	if x != nil {
		return x.RespId
	}
	return ""
}

type PayzeLinkResponseSaveCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string           `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string           `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	TransactionId string           `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Action        int32            `protobuf:"varint,5,opt,name=action,proto3" json:"action,omitempty"`
	Currency      string           `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Price         float32          `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
	CreatedDate   string           `protobuf:"bytes,8,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	CreatedAt     string           `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string           `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     string           `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Data          *structpb.Struct `protobuf:"bytes,12,opt,name=data,proto3" json:"data,omitempty"`
	ProductId     string           `protobuf:"bytes,13,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Amount        int32            `protobuf:"varint,14,opt,name=amount,proto3" json:"amount,omitempty"`
	RespId        string           `protobuf:"bytes,15,opt,name=resp_id,json=respId,proto3" json:"resp_id,omitempty"`
	CardId        string           `protobuf:"bytes,16,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
}

func (x *PayzeLinkResponseSaveCard) Reset() {
	*x = PayzeLinkResponseSaveCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payze_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayzeLinkResponseSaveCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayzeLinkResponseSaveCard) ProtoMessage() {}

func (x *PayzeLinkResponseSaveCard) ProtoReflect() protoreflect.Message {
	mi := &file_payze_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayzeLinkResponseSaveCard.ProtoReflect.Descriptor instead.
func (*PayzeLinkResponseSaveCard) Descriptor() ([]byte, []int) {
	return file_payze_proto_rawDescGZIP(), []int{3}
}

func (x *PayzeLinkResponseSaveCard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *PayzeLinkResponseSaveCard) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PayzeLinkResponseSaveCard) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PayzeLinkResponseSaveCard) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayzeLinkResponseSaveCard) GetRespId() string {
	if x != nil {
		return x.RespId
	}
	return ""
}

func (x *PayzeLinkResponseSaveCard) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

var File_payze_proto protoreflect.FileDescriptor

var file_payze_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x79, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xfe, 0x01, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f,
	0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x6f,
	0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x22, 0xc4, 0x03, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x70, 0x49, 0x64, 0x22, 0xe5, 0x03, 0x0a, 0x19,
	0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x32, 0xea, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x28, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x76, 0x32, 0x2e, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x79,
	0x7a, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6e, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x53, 0x61, 0x76, 0x65, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x28, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x79, 0x7a, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x79, 0x7a, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x61, 0x76, 0x65, 0x43, 0x61, 0x72, 0x64, 0x22, 0x00,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payze_proto_rawDescOnce sync.Once
	file_payze_proto_rawDescData = file_payze_proto_rawDesc
)

func file_payze_proto_rawDescGZIP() []byte {
	file_payze_proto_rawDescOnce.Do(func() {
		file_payze_proto_rawDescData = protoimpl.X.CompressGZIP(file_payze_proto_rawDescData)
	})
	return file_payze_proto_rawDescData
}

var file_payze_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_payze_proto_goTypes = []interface{}{
	(*PayzeLinkRequest)(nil),          // 0: integration_service_v2.PayzeLinkRequest
	(*PayzeLinkRequestData)(nil),      // 1: integration_service_v2.PayzeLinkRequestData
	(*PayzeLinkResponse)(nil),         // 2: integration_service_v2.PayzeLinkResponse
	(*PayzeLinkResponseSaveCard)(nil), // 3: integration_service_v2.PayzeLinkResponseSaveCard
	(*structpb.Struct)(nil),           // 4: google.protobuf.Struct
}
var file_payze_proto_depIdxs = []int32{
	1, // 0: integration_service_v2.PayzeLinkRequest.data:type_name -> integration_service_v2.PayzeLinkRequestData
	4, // 1: integration_service_v2.PayzeLinkResponse.data:type_name -> google.protobuf.Struct
	4, // 2: integration_service_v2.PayzeLinkResponseSaveCard.data:type_name -> google.protobuf.Struct
	0, // 3: integration_service_v2.PayzeService.GeneratePayzeLink:input_type -> integration_service_v2.PayzeLinkRequest
	0, // 4: integration_service_v2.PayzeService.PayzeSaveCard:input_type -> integration_service_v2.PayzeLinkRequest
	2, // 5: integration_service_v2.PayzeService.GeneratePayzeLink:output_type -> integration_service_v2.PayzeLinkResponse
	3, // 6: integration_service_v2.PayzeService.PayzeSaveCard:output_type -> integration_service_v2.PayzeLinkResponseSaveCard
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_payze_proto_init() }
func file_payze_proto_init() {
	if File_payze_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payze_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayzeLinkRequest); i {
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
		file_payze_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayzeLinkRequestData); i {
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
		file_payze_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayzeLinkResponse); i {
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
		file_payze_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayzeLinkResponseSaveCard); i {
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
			RawDescriptor: file_payze_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payze_proto_goTypes,
		DependencyIndexes: file_payze_proto_depIdxs,
		MessageInfos:      file_payze_proto_msgTypes,
	}.Build()
	File_payze_proto = out.File
	file_payze_proto_rawDesc = nil
	file_payze_proto_goTypes = nil
	file_payze_proto_depIdxs = nil
}
