// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: servicegrpc/quantity.proto

package servicegrpc

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

type InitProcessQuantityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidProcess string `protobuf:"bytes,1,opt,name=uuidProcess,proto3" json:"uuidProcess,omitempty"`
}

func (x *InitProcessQuantityRequest) Reset() {
	*x = InitProcessQuantityRequest{}
	mi := &file_servicegrpc_quantity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitProcessQuantityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitProcessQuantityRequest) ProtoMessage() {}

func (x *InitProcessQuantityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quantity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitProcessQuantityRequest.ProtoReflect.Descriptor instead.
func (*InitProcessQuantityRequest) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quantity_proto_rawDescGZIP(), []int{0}
}

func (x *InitProcessQuantityRequest) GetUuidProcess() string {
	if x != nil {
		return x.UuidProcess
	}
	return ""
}

type InitProcessQuantityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *InitProcessQuantityResponse) Reset() {
	*x = InitProcessQuantityResponse{}
	mi := &file_servicegrpc_quantity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitProcessQuantityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitProcessQuantityResponse) ProtoMessage() {}

func (x *InitProcessQuantityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quantity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitProcessQuantityResponse.ProtoReflect.Descriptor instead.
func (*InitProcessQuantityResponse) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quantity_proto_rawDescGZIP(), []int{1}
}

func (x *InitProcessQuantityResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type SendBlobQuantityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Blob []byte `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (x *SendBlobQuantityRequest) Reset() {
	*x = SendBlobQuantityRequest{}
	mi := &file_servicegrpc_quantity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendBlobQuantityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBlobQuantityRequest) ProtoMessage() {}

func (x *SendBlobQuantityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quantity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBlobQuantityRequest.ProtoReflect.Descriptor instead.
func (*SendBlobQuantityRequest) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quantity_proto_rawDescGZIP(), []int{2}
}

func (x *SendBlobQuantityRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SendBlobQuantityRequest) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

type SendBlobQuantityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendBlobQuantityResponse) Reset() {
	*x = SendBlobQuantityResponse{}
	mi := &file_servicegrpc_quantity_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendBlobQuantityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBlobQuantityResponse) ProtoMessage() {}

func (x *SendBlobQuantityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpc_quantity_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBlobQuantityResponse.ProtoReflect.Descriptor instead.
func (*SendBlobQuantityResponse) Descriptor() ([]byte, []int) {
	return file_servicegrpc_quantity_proto_rawDescGZIP(), []int{3}
}

var File_servicegrpc_quantity_proto protoreflect.FileDescriptor

var file_servicegrpc_quantity_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70,
	0x70, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3e, 0x0a, 0x1a, 0x49, 0x6e,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x75, 0x69, 0x64,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x75, 0x69, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x1b, 0x49, 0x6e,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x41, 0x0a, 0x17, 0x53, 0x65, 0x6e,
	0x64, 0x42, 0x6c, 0x6f, 0x62, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x22, 0x1a, 0x0a, 0x18,
	0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe6, 0x01, 0x0a, 0x0f, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x13,
	0x49, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x10, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servicegrpc_quantity_proto_rawDescOnce sync.Once
	file_servicegrpc_quantity_proto_rawDescData = file_servicegrpc_quantity_proto_rawDesc
)

func file_servicegrpc_quantity_proto_rawDescGZIP() []byte {
	file_servicegrpc_quantity_proto_rawDescOnce.Do(func() {
		file_servicegrpc_quantity_proto_rawDescData = protoimpl.X.CompressGZIP(file_servicegrpc_quantity_proto_rawDescData)
	})
	return file_servicegrpc_quantity_proto_rawDescData
}

var file_servicegrpc_quantity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_servicegrpc_quantity_proto_goTypes = []any{
	(*InitProcessQuantityRequest)(nil),  // 0: app.quantity.InitProcessQuantityRequest
	(*InitProcessQuantityResponse)(nil), // 1: app.quantity.InitProcessQuantityResponse
	(*SendBlobQuantityRequest)(nil),     // 2: app.quantity.SendBlobQuantityRequest
	(*SendBlobQuantityResponse)(nil),    // 3: app.quantity.SendBlobQuantityResponse
}
var file_servicegrpc_quantity_proto_depIdxs = []int32{
	0, // 0: app.quantity.QuantityService.InitProcessQuantity:input_type -> app.quantity.InitProcessQuantityRequest
	2, // 1: app.quantity.QuantityService.SendBlobQuantity:input_type -> app.quantity.SendBlobQuantityRequest
	1, // 2: app.quantity.QuantityService.InitProcessQuantity:output_type -> app.quantity.InitProcessQuantityResponse
	3, // 3: app.quantity.QuantityService.SendBlobQuantity:output_type -> app.quantity.SendBlobQuantityResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_servicegrpc_quantity_proto_init() }
func file_servicegrpc_quantity_proto_init() {
	if File_servicegrpc_quantity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_servicegrpc_quantity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servicegrpc_quantity_proto_goTypes,
		DependencyIndexes: file_servicegrpc_quantity_proto_depIdxs,
		MessageInfos:      file_servicegrpc_quantity_proto_msgTypes,
	}.Build()
	File_servicegrpc_quantity_proto = out.File
	file_servicegrpc_quantity_proto_rawDesc = nil
	file_servicegrpc_quantity_proto_goTypes = nil
	file_servicegrpc_quantity_proto_depIdxs = nil
}
