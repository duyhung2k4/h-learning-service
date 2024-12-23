// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: servicegrpc/quizz.proto

package servicegrpc

import (
	enumgrpc "app/generated/grpc/enumgrpc"
	sharedgrpc "app/generated/grpc/sharedgrpc"
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

type Quizz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ask        string              `protobuf:"bytes,2,opt,name=ask,proto3" json:"ask,omitempty"`
	Time       int32               `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Result     []string            `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
	Option     []string            `protobuf:"bytes,5,rep,name=option,proto3" json:"option,omitempty"`
	EntityType enumgrpc.EntityType `protobuf:"varint,6,opt,name=entityType,proto3,enum=app.enumgrpc.EntityType" json:"entityType,omitempty"`
	EntityId   uint64              `protobuf:"varint,7,opt,name=entityId,proto3" json:"entityId,omitempty"`
	CreatedAt  uint64              `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt  uint32              `protobuf:"varint,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt  uint64              `protobuf:"varint,10,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Quizz) Reset() {
	*x = Quizz{}
	mi := &file_servicegrpc_quizz_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Quizz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quizz) ProtoMessage() {}

func (x *Quizz) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quizz_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quizz.ProtoReflect.Descriptor instead.
func (*Quizz) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quizz_proto_rawDescGZIP(), []int{0}
}

func (x *Quizz) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Quizz) GetAsk() string {
	if x != nil {
		return x.Ask
	}
	return ""
}

func (x *Quizz) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Quizz) GetResult() []string {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Quizz) GetOption() []string {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *Quizz) GetEntityType() enumgrpc.EntityType {
	if x != nil {
		return x.EntityType
	}
	return enumgrpc.EntityType(0)
}

func (x *Quizz) GetEntityId() uint64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *Quizz) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Quizz) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Quizz) GetDeletedAt() uint64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type CreateQuizzRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quizz *Quizz `protobuf:"bytes,1,opt,name=quizz,proto3" json:"quizz,omitempty"`
}

func (x *CreateQuizzRequest) Reset() {
	*x = CreateQuizzRequest{}
	mi := &file_servicegrpc_quizz_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateQuizzRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuizzRequest) ProtoMessage() {}

func (x *CreateQuizzRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quizz_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuizzRequest.ProtoReflect.Descriptor instead.
func (*CreateQuizzRequest) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quizz_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuizzRequest) GetQuizz() *Quizz {
	if x != nil {
		return x.Quizz
	}
	return nil
}

type UpdateQuizzRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quizz *Quizz `protobuf:"bytes,2,opt,name=quizz,proto3" json:"quizz,omitempty"`
}

func (x *UpdateQuizzRequest) Reset() {
	*x = UpdateQuizzRequest{}
	mi := &file_servicegrpc_quizz_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateQuizzRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuizzRequest) ProtoMessage() {}

func (x *UpdateQuizzRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quizz_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuizzRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuizzRequest) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quizz_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateQuizzRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateQuizzRequest) GetQuizz() *Quizz {
	if x != nil {
		return x.Quizz
	}
	return nil
}

type GetListByEntityIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quizzs []*Quizz `protobuf:"bytes,1,rep,name=quizzs,proto3" json:"quizzs,omitempty"`
	Total  int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetListByEntityIdResponse) Reset() {
	*x = GetListByEntityIdResponse{}
	mi := &file_servicegrpc_quizz_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListByEntityIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListByEntityIdResponse) ProtoMessage() {}

func (x *GetListByEntityIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quizz_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListByEntityIdResponse.ProtoReflect.Descriptor instead.
func (*GetListByEntityIdResponse) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quizz_proto_rawDescGZIP(), []int{3}
}

func (x *GetListByEntityIdResponse) GetQuizzs() []*Quizz {
	if x != nil {
		return x.Quizzs
	}
	return nil
}

func (x *GetListByEntityIdResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_servicegrpc_quizz_proto protoreflect.FileDescriptor

var file_servicegrpc_quizz_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x71, 0x75,
	0x69, 0x7a, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x65, 0x6e, 0x75, 0x6d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x73,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x69, 0x7a, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x71,
	0x75, 0x69, 0x7a, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x69,
	0x7a, 0x7a, 0x52, 0x05, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x22, 0x52, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2c, 0x0a, 0x05, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x52, 0x05, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x22, 0x61, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x71, 0x75,
	0x69, 0x7a, 0x7a, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x69,
	0x7a, 0x7a, 0x52, 0x06, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x32, 0xf0, 0x02, 0x0a, 0x0c, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x44,
	0x1a, 0x16, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x49, 0x44, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x7a,
	0x12, 0x23, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x44, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servicegrpc_quizz_proto_rawDescOnce sync.Once
	file_servicegrpc_quizz_proto_rawDescData = file_servicegrpc_quizz_proto_rawDesc
)

func file_servicegrpc_quizz_proto_rawDescGZIP() []byte {
	file_servicegrpc_quizz_proto_rawDescOnce.Do(func() {
		file_servicegrpc_quizz_proto_rawDescData = protoimpl.X.CompressGZIP(file_servicegrpc_quizz_proto_rawDescData)
	})
	return file_servicegrpc_quizz_proto_rawDescData
}

var file_servicegrpc_quizz_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_servicegrpc_quizz_proto_goTypes = []any{
	(*Quizz)(nil),                     // 0: app.servicegrpc.Quizz
	(*CreateQuizzRequest)(nil),        // 1: app.servicegrpc.CreateQuizzRequest
	(*UpdateQuizzRequest)(nil),        // 2: app.servicegrpc.UpdateQuizzRequest
	(*GetListByEntityIdResponse)(nil), // 3: app.servicegrpc.GetListByEntityIdResponse
	(enumgrpc.EntityType)(0),          // 4: app.enumgrpc.EntityType
	(*sharedgrpc.ID)(nil),             // 5: app.sharedgrpc.ID
	(*sharedgrpc.Empty)(nil),          // 6: app.sharedgrpc.Empty
}
var file_servicegrpc_quizz_proto_depIdxs = []int32{
	4, // 0: app.servicegrpc.Quizz.entityType:type_name -> app.enumgrpc.EntityType
	0, // 1: app.servicegrpc.CreateQuizzRequest.quizz:type_name -> app.servicegrpc.Quizz
	0, // 2: app.servicegrpc.UpdateQuizzRequest.quizz:type_name -> app.servicegrpc.Quizz
	0, // 3: app.servicegrpc.GetListByEntityIdResponse.quizzs:type_name -> app.servicegrpc.Quizz
	5, // 4: app.servicegrpc.QuizzService.GetById:input_type -> app.sharedgrpc.ID
	5, // 5: app.servicegrpc.QuizzService.GetListByEntityId:input_type -> app.sharedgrpc.ID
	1, // 6: app.servicegrpc.QuizzService.CreateQuizz:input_type -> app.servicegrpc.CreateQuizzRequest
	2, // 7: app.servicegrpc.QuizzService.UpdateQuizz:input_type -> app.servicegrpc.UpdateQuizzRequest
	5, // 8: app.servicegrpc.QuizzService.DeleteById:input_type -> app.sharedgrpc.ID
	0, // 9: app.servicegrpc.QuizzService.GetById:output_type -> app.servicegrpc.Quizz
	3, // 10: app.servicegrpc.QuizzService.GetListByEntityId:output_type -> app.servicegrpc.GetListByEntityIdResponse
	5, // 11: app.servicegrpc.QuizzService.CreateQuizz:output_type -> app.sharedgrpc.ID
	6, // 12: app.servicegrpc.QuizzService.UpdateQuizz:output_type -> app.sharedgrpc.Empty
	6, // 13: app.servicegrpc.QuizzService.DeleteById:output_type -> app.sharedgrpc.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_servicegrpc_quizz_proto_init() }
func file_servicegrpc_quizz_proto_init() {
	if File_servicegrpc_quizz_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_servicegrpc_quizz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servicegrpc_quizz_proto_goTypes,
		DependencyIndexes: file_servicegrpc_quizz_proto_depIdxs,
		MessageInfos:      file_servicegrpc_quizz_proto_msgTypes,
	}.Build()
	File_servicegrpc_quizz_proto = out.File
	file_servicegrpc_quizz_proto_rawDesc = nil
	file_servicegrpc_quizz_proto_goTypes = nil
	file_servicegrpc_quizz_proto_depIdxs = nil
}
