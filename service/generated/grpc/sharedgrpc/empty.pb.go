// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: sharedgrpc/empty.proto

package sharedgrpc

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_sharedgrpc_empty_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_sharedgrpc_empty_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_sharedgrpc_empty_proto_rawDescGZIP(), []int{0}
}

var File_sharedgrpc_empty_proto protoreflect.FileDescriptor

var file_sharedgrpc_empty_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sharedgrpc_empty_proto_rawDescOnce sync.Once
	file_sharedgrpc_empty_proto_rawDescData = file_sharedgrpc_empty_proto_rawDesc
)

func file_sharedgrpc_empty_proto_rawDescGZIP() []byte {
	file_sharedgrpc_empty_proto_rawDescOnce.Do(func() {
		file_sharedgrpc_empty_proto_rawDescData = protoimpl.X.CompressGZIP(file_sharedgrpc_empty_proto_rawDescData)
	})
	return file_sharedgrpc_empty_proto_rawDescData
}

var file_sharedgrpc_empty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sharedgrpc_empty_proto_goTypes = []any{
	(*Empty)(nil), // 0: app.sharedgrpc.Empty
}
var file_sharedgrpc_empty_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sharedgrpc_empty_proto_init() }
func file_sharedgrpc_empty_proto_init() {
	if File_sharedgrpc_empty_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sharedgrpc_empty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sharedgrpc_empty_proto_goTypes,
		DependencyIndexes: file_sharedgrpc_empty_proto_depIdxs,
		MessageInfos:      file_sharedgrpc_empty_proto_msgTypes,
	}.Build()
	File_sharedgrpc_empty_proto = out.File
	file_sharedgrpc_empty_proto_rawDesc = nil
	file_sharedgrpc_empty_proto_goTypes = nil
	file_sharedgrpc_empty_proto_depIdxs = nil
}
