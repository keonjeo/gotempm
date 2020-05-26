// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: promotion/promotion.proto

package promotion

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Discount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value       float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Type        int32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Discount) Reset() {
	*x = Discount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discount) ProtoMessage() {}

func (x *Discount) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discount.ProtoReflect.Descriptor instead.
func (*Discount) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{0}
}

func (x *Discount) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Discount) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Discount) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Discount) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UpcCode  string      `protobuf:"bytes,2,opt,name=upcCode,proto3" json:"upcCode,omitempty"`
	Discount []*Discount `protobuf:"bytes,3,rep,name=discount,proto3" json:"discount,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetUpcCode() string {
	if x != nil {
		return x.UpcCode
	}
	return ""
}

func (x *Product) GetDiscount() []*Discount {
	if x != nil {
		return x.Discount
	}
	return nil
}

type Promotion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description        string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ValidFrom          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=validFrom,proto3" json:"validFrom,omitempty"`
	ValidThru          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=validThru,proto3" json:"validThru,omitempty"`
	Active             bool                 `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	CustomerId         int64                `protobuf:"varint,7,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Product            []*Product           `protobuf:"bytes,8,rep,name=product,proto3" json:"product,omitempty"`
	ApprovalStatus     int32                `protobuf:"varint,9,opt,name=approvalStatus,proto3" json:"approvalStatus,omitempty"`
	PrevApprovalStatus int32                `protobuf:"varint,10,opt,name=prevApprovalStatus,proto3" json:"prevApprovalStatus,omitempty"`
}

func (x *Promotion) Reset() {
	*x = Promotion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promotion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promotion) ProtoMessage() {}

func (x *Promotion) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promotion.ProtoReflect.Descriptor instead.
func (*Promotion) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{2}
}

func (x *Promotion) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Promotion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Promotion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Promotion) GetValidFrom() *timestamp.Timestamp {
	if x != nil {
		return x.ValidFrom
	}
	return nil
}

func (x *Promotion) GetValidThru() *timestamp.Timestamp {
	if x != nil {
		return x.ValidThru
	}
	return nil
}

func (x *Promotion) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Promotion) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Promotion) GetProduct() []*Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *Promotion) GetApprovalStatus() int32 {
	if x != nil {
		return x.ApprovalStatus
	}
	return 0
}

func (x *Promotion) GetPrevApprovalStatus() int32 {
	if x != nil {
		return x.PrevApprovalStatus
	}
	return 0
}

type Promotions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Promotion []*Promotion `protobuf:"bytes,1,rep,name=promotion,proto3" json:"promotion,omitempty"`
}

func (x *Promotions) Reset() {
	*x = Promotions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promotions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promotions) ProtoMessage() {}

func (x *Promotions) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promotions.ProtoReflect.Descriptor instead.
func (*Promotions) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{3}
}

func (x *Promotions) GetPromotion() []*Promotion {
	if x != nil {
		return x.Promotion
	}
	return nil
}

type SearchParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProductId  int64                `protobuf:"varint,3,opt,name=productId,proto3" json:"productId,omitempty"`
	CustomerId int64                `protobuf:"varint,4,opt,name=customerId,proto3" json:"customerId,omitempty"`
	ValidDate  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=validDate,proto3" json:"validDate,omitempty"`
}

func (x *SearchParams) Reset() {
	*x = SearchParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchParams) ProtoMessage() {}

func (x *SearchParams) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchParams.ProtoReflect.Descriptor instead.
func (*SearchParams) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{4}
}

func (x *SearchParams) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchParams) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *SearchParams) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SearchParams) GetValidDate() *timestamp.Timestamp {
	if x != nil {
		return x.ValidDate
	}
	return nil
}

type SearchId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchId) Reset() {
	*x = SearchId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchId) ProtoMessage() {}

func (x *SearchId) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchId.ProtoReflect.Descriptor instead.
func (*SearchId) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{5}
}

func (x *SearchId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AffectedCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AffectedCount) Reset() {
	*x = AffectedCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffectedCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffectedCount) ProtoMessage() {}

func (x *AffectedCount) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffectedCount.ProtoReflect.Descriptor instead.
func (*AffectedCount) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{6}
}

func (x *AffectedCount) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ValidationErr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailureDesc []string `protobuf:"bytes,1,rep,name=failureDesc,proto3" json:"failureDesc,omitempty"`
}

func (x *ValidationErr) Reset() {
	*x = ValidationErr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationErr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationErr) ProtoMessage() {}

func (x *ValidationErr) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationErr.ProtoReflect.Descriptor instead.
func (*ValidationErr) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{7}
}

func (x *ValidationErr) GetFailureDesc() []string {
	if x != nil {
		return x.FailureDesc
	}
	return nil
}

type AfterFuncErr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailureDesc []string `protobuf:"bytes,1,rep,name=failureDesc,proto3" json:"failureDesc,omitempty"`
}

func (x *AfterFuncErr) Reset() {
	*x = AfterFuncErr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_promotion_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AfterFuncErr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AfterFuncErr) ProtoMessage() {}

func (x *AfterFuncErr) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_promotion_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AfterFuncErr.ProtoReflect.Descriptor instead.
func (*AfterFuncErr) Descriptor() ([]byte, []int) {
	return file_promotion_promotion_proto_rawDescGZIP(), []int{8}
}

func (x *AfterFuncErr) GetFailureDesc() []string {
	if x != nil {
		return x.FailureDesc
	}
	return nil
}

var File_promotion_promotion_proto protoreflect.FileDescriptor

var file_promotion_promotion_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x64, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70,
	0x63, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x63,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x83, 0x03, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x38, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x68, 0x72, 0x75,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x68, 0x72, 0x75, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x70,
	0x72, 0x65, 0x76, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x70, 0x72, 0x65, 0x76, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a, 0x0a, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xaa, 0x01,
	0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a,
	0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x22, 0x30, 0x0a, 0x0c, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x45, 0x72, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x32, 0x94, 0x06, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x72, 0x76, 0x12, 0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x64, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x15, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x15, 0x42, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72,
	0x72, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x15, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x14, 0x41, 0x66, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x75,
	0x6e, 0x63, 0x45, 0x72, 0x72, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x14, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x45, 0x72, 0x72, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x14, 0x41, 0x66, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x46, 0x75, 0x6e, 0x63, 0x45, 0x72, 0x72, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x6f, 0x54,
	0x65, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promotion_promotion_proto_rawDescOnce sync.Once
	file_promotion_promotion_proto_rawDescData = file_promotion_promotion_proto_rawDesc
)

func file_promotion_promotion_proto_rawDescGZIP() []byte {
	file_promotion_promotion_proto_rawDescOnce.Do(func() {
		file_promotion_promotion_proto_rawDescData = protoimpl.X.CompressGZIP(file_promotion_promotion_proto_rawDescData)
	})
	return file_promotion_promotion_proto_rawDescData
}

var file_promotion_promotion_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_promotion_promotion_proto_goTypes = []interface{}{
	(*Discount)(nil),            // 0: promotion.Discount
	(*Product)(nil),             // 1: promotion.Product
	(*Promotion)(nil),           // 2: promotion.Promotion
	(*Promotions)(nil),          // 3: promotion.Promotions
	(*SearchParams)(nil),        // 4: promotion.SearchParams
	(*SearchId)(nil),            // 5: promotion.SearchId
	(*AffectedCount)(nil),       // 6: promotion.affectedCount
	(*ValidationErr)(nil),       // 7: promotion.validationErr
	(*AfterFuncErr)(nil),        // 8: promotion.AfterFuncErr
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_promotion_promotion_proto_depIdxs = []int32{
	0,  // 0: promotion.Product.discount:type_name -> promotion.Discount
	9,  // 1: promotion.Promotion.validFrom:type_name -> google.protobuf.Timestamp
	9,  // 2: promotion.Promotion.validThru:type_name -> google.protobuf.Timestamp
	1,  // 3: promotion.Promotion.product:type_name -> promotion.Product
	2,  // 4: promotion.Promotions.promotion:type_name -> promotion.Promotion
	9,  // 5: promotion.SearchParams.validDate:type_name -> google.protobuf.Timestamp
	5,  // 6: promotion.PromotionSrv.GetPromotionById:input_type -> promotion.SearchId
	4,  // 7: promotion.PromotionSrv.GetPromotions:input_type -> promotion.SearchParams
	2,  // 8: promotion.PromotionSrv.CreatePromotion:input_type -> promotion.Promotion
	2,  // 9: promotion.PromotionSrv.UpdatePromotion:input_type -> promotion.Promotion
	5,  // 10: promotion.PromotionSrv.DeletePromotion:input_type -> promotion.SearchId
	2,  // 11: promotion.PromotionSrv.BeforeCreatePromotion:input_type -> promotion.Promotion
	2,  // 12: promotion.PromotionSrv.BeforeUpdatePromotion:input_type -> promotion.Promotion
	2,  // 13: promotion.PromotionSrv.BeforeDeletePromotion:input_type -> promotion.Promotion
	2,  // 14: promotion.PromotionSrv.AfterCreatePromotion:input_type -> promotion.Promotion
	2,  // 15: promotion.PromotionSrv.AfterUpdatePromotion:input_type -> promotion.Promotion
	2,  // 16: promotion.PromotionSrv.AfterDeletePromotion:input_type -> promotion.Promotion
	2,  // 17: promotion.PromotionSrv.GetPromotionById:output_type -> promotion.Promotion
	3,  // 18: promotion.PromotionSrv.GetPromotions:output_type -> promotion.Promotions
	2,  // 19: promotion.PromotionSrv.CreatePromotion:output_type -> promotion.Promotion
	2,  // 20: promotion.PromotionSrv.UpdatePromotion:output_type -> promotion.Promotion
	6,  // 21: promotion.PromotionSrv.DeletePromotion:output_type -> promotion.affectedCount
	7,  // 22: promotion.PromotionSrv.BeforeCreatePromotion:output_type -> promotion.validationErr
	7,  // 23: promotion.PromotionSrv.BeforeUpdatePromotion:output_type -> promotion.validationErr
	7,  // 24: promotion.PromotionSrv.BeforeDeletePromotion:output_type -> promotion.validationErr
	8,  // 25: promotion.PromotionSrv.AfterCreatePromotion:output_type -> promotion.AfterFuncErr
	8,  // 26: promotion.PromotionSrv.AfterUpdatePromotion:output_type -> promotion.AfterFuncErr
	8,  // 27: promotion.PromotionSrv.AfterDeletePromotion:output_type -> promotion.AfterFuncErr
	17, // [17:28] is the sub-list for method output_type
	6,  // [6:17] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_promotion_promotion_proto_init() }
func file_promotion_promotion_proto_init() {
	if File_promotion_promotion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_promotion_promotion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discount); i {
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
		file_promotion_promotion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_promotion_promotion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promotion); i {
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
		file_promotion_promotion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promotions); i {
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
		file_promotion_promotion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchParams); i {
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
		file_promotion_promotion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchId); i {
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
		file_promotion_promotion_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AffectedCount); i {
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
		file_promotion_promotion_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationErr); i {
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
		file_promotion_promotion_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AfterFuncErr); i {
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
			RawDescriptor: file_promotion_promotion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_promotion_promotion_proto_goTypes,
		DependencyIndexes: file_promotion_promotion_proto_depIdxs,
		MessageInfos:      file_promotion_promotion_proto_msgTypes,
	}.Build()
	File_promotion_promotion_proto = out.File
	file_promotion_promotion_proto_rawDesc = nil
	file_promotion_promotion_proto_goTypes = nil
	file_promotion_promotion_proto_depIdxs = nil
}
