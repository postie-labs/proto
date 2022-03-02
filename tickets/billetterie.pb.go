// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tickets/billetterie.proto

package routeguide

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

// Message
type IssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IssueRequest) Reset() {
	*x = IssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_billetterie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueRequest) ProtoMessage() {}

func (x *IssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_billetterie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueRequest.ProtoReflect.Descriptor instead.
func (*IssueRequest) Descriptor() ([]byte, []int) {
	return file_tickets_billetterie_proto_rawDescGZIP(), []int{0}
}

func (x *IssueRequest) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *IssueRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type IssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketHash string `protobuf:"bytes,1,opt,name=ticket_hash,json=ticketHash,proto3" json:"ticket_hash,omitempty"`
	Error      *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *IssueResponse) Reset() {
	*x = IssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_billetterie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueResponse) ProtoMessage() {}

func (x *IssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_billetterie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueResponse.ProtoReflect.Descriptor instead.
func (*IssueResponse) Descriptor() ([]byte, []int) {
	return file_tickets_billetterie_proto_rawDescGZIP(), []int{1}
}

func (x *IssueResponse) GetTicketHash() string {
	if x != nil {
		return x.TicketHash
	}
	return ""
}

func (x *IssueResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerFrom  string `protobuf:"bytes,1,opt,name=owner_from,json=ownerFrom,proto3" json:"owner_from,omitempty"`
	OwnerTo    string `protobuf:"bytes,2,opt,name=owner_to,json=ownerTo,proto3" json:"owner_to,omitempty"`
	TicketHash string `protobuf:"bytes,3,opt,name=ticket_hash,json=ticketHash,proto3" json:"ticket_hash,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_billetterie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_billetterie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_tickets_billetterie_proto_rawDescGZIP(), []int{2}
}

func (x *TransferRequest) GetOwnerFrom() string {
	if x != nil {
		return x.OwnerFrom
	}
	return ""
}

func (x *TransferRequest) GetOwnerTo() string {
	if x != nil {
		return x.OwnerTo
	}
	return ""
}

func (x *TransferRequest) GetTicketHash() string {
	if x != nil {
		return x.TicketHash
	}
	return ""
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transfered bool   `protobuf:"varint,1,opt,name=transfered,proto3" json:"transfered,omitempty"`
	Error      *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_billetterie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_billetterie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_tickets_billetterie_proto_rawDescGZIP(), []int{3}
}

func (x *TransferResponse) GetTransfered() bool {
	if x != nil {
		return x.Transfered
	}
	return false
}

func (x *TransferResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner      string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	TicketHash string `protobuf:"bytes,2,opt,name=ticket_hash,json=ticketHash,proto3" json:"ticket_hash,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_billetterie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_billetterie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_tickets_billetterie_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *VerifyRequest) GetTicketHash() string {
	if x != nil {
		return x.TicketHash
	}
	return ""
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verified bool   `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
	Error    *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_billetterie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_billetterie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_tickets_billetterie_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyResponse) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *VerifyResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_billetterie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_billetterie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_tickets_billetterie_proto_rawDescGZIP(), []int{6}
}

func (x *Error) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_tickets_billetterie_proto protoreflect.FileDescriptor

var file_tickets_billetterie_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0c, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x54, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x50, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x46, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22,
	0x4a, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2f, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x97, 0x01, 0x0a,
	0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x69, 0x65, 0x12, 0x28, 0x0a, 0x05,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x0d, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x12, 0x0e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickets_billetterie_proto_rawDescOnce sync.Once
	file_tickets_billetterie_proto_rawDescData = file_tickets_billetterie_proto_rawDesc
)

func file_tickets_billetterie_proto_rawDescGZIP() []byte {
	file_tickets_billetterie_proto_rawDescOnce.Do(func() {
		file_tickets_billetterie_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickets_billetterie_proto_rawDescData)
	})
	return file_tickets_billetterie_proto_rawDescData
}

var file_tickets_billetterie_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tickets_billetterie_proto_goTypes = []interface{}{
	(*IssueRequest)(nil),     // 0: IssueRequest
	(*IssueResponse)(nil),    // 1: IssueResponse
	(*TransferRequest)(nil),  // 2: TransferRequest
	(*TransferResponse)(nil), // 3: TransferResponse
	(*VerifyRequest)(nil),    // 4: VerifyRequest
	(*VerifyResponse)(nil),   // 5: VerifyResponse
	(*Error)(nil),            // 6: Error
}
var file_tickets_billetterie_proto_depIdxs = []int32{
	6, // 0: IssueResponse.error:type_name -> Error
	6, // 1: TransferResponse.error:type_name -> Error
	6, // 2: VerifyResponse.error:type_name -> Error
	0, // 3: Billetterie.Issue:input_type -> IssueRequest
	2, // 4: Billetterie.Transfer:input_type -> TransferRequest
	4, // 5: Billetterie.Verify:input_type -> VerifyRequest
	1, // 6: Billetterie.Issue:output_type -> IssueResponse
	3, // 7: Billetterie.Transfer:output_type -> TransferResponse
	5, // 8: Billetterie.Verify:output_type -> VerifyResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tickets_billetterie_proto_init() }
func file_tickets_billetterie_proto_init() {
	if File_tickets_billetterie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tickets_billetterie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueRequest); i {
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
		file_tickets_billetterie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueResponse); i {
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
		file_tickets_billetterie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
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
		file_tickets_billetterie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResponse); i {
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
		file_tickets_billetterie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_tickets_billetterie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
		file_tickets_billetterie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_tickets_billetterie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tickets_billetterie_proto_goTypes,
		DependencyIndexes: file_tickets_billetterie_proto_depIdxs,
		MessageInfos:      file_tickets_billetterie_proto_msgTypes,
	}.Build()
	File_tickets_billetterie_proto = out.File
	file_tickets_billetterie_proto_rawDesc = nil
	file_tickets_billetterie_proto_goTypes = nil
	file_tickets_billetterie_proto_depIdxs = nil
}