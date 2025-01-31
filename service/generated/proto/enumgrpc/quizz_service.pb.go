// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: proto/enumgrpc/quizz_service.proto

package enumgrpc

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

type EntityType int32

const (
	EntityType_QUIZZ_VIDEO_LESSON EntityType = 0
	EntityType_QUIZZ_LESSON       EntityType = 1
)

// Enum value maps for EntityType.
var (
	EntityType_name = map[int32]string{
		0: "QUIZZ_VIDEO_LESSON",
		1: "QUIZZ_LESSON",
	}
	EntityType_value = map[string]int32{
		"QUIZZ_VIDEO_LESSON": 0,
		"QUIZZ_LESSON":       1,
	}
)

func (x EntityType) Enum() *EntityType {
	p := new(EntityType)
	*p = x
	return p
}

func (x EntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enumgrpc_quizz_service_proto_enumTypes[0].Descriptor()
}

func (EntityType) Type() protoreflect.EnumType {
	return &file_proto_enumgrpc_quizz_service_proto_enumTypes[0]
}

func (x EntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityType.Descriptor instead.
func (EntityType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enumgrpc_quizz_service_proto_rawDescGZIP(), []int{0}
}

type ResultType int32

const (
	ResultType_QUIZZ_SINGLE_RESULT ResultType = 0
	ResultType_QUIZZ_MULTI_RESULT  ResultType = 1
)

// Enum value maps for ResultType.
var (
	ResultType_name = map[int32]string{
		0: "QUIZZ_SINGLE_RESULT",
		1: "QUIZZ_MULTI_RESULT",
	}
	ResultType_value = map[string]int32{
		"QUIZZ_SINGLE_RESULT": 0,
		"QUIZZ_MULTI_RESULT":  1,
	}
)

func (x ResultType) Enum() *ResultType {
	p := new(ResultType)
	*p = x
	return p
}

func (x ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enumgrpc_quizz_service_proto_enumTypes[1].Descriptor()
}

func (ResultType) Type() protoreflect.EnumType {
	return &file_proto_enumgrpc_quizz_service_proto_enumTypes[1]
}

func (x ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultType.Descriptor instead.
func (ResultType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enumgrpc_quizz_service_proto_rawDescGZIP(), []int{1}
}

var File_proto_enumgrpc_quizz_service_proto protoreflect.FileDescriptor

var file_proto_enumgrpc_quizz_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x71, 0x75, 0x69, 0x7a, 0x7a, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x70, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x67, 0x72,
	0x70, 0x63, 0x2a, 0x36, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x51, 0x55, 0x49, 0x5a, 0x5a, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f,
	0x4c, 0x45, 0x53, 0x53, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x51, 0x55, 0x49, 0x5a,
	0x5a, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x2a, 0x3d, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x49, 0x5a,
	0x5a, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x51, 0x55, 0x49, 0x5a, 0x5a, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49,
	0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x70, 0x70,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_enumgrpc_quizz_service_proto_rawDescOnce sync.Once
	file_proto_enumgrpc_quizz_service_proto_rawDescData = file_proto_enumgrpc_quizz_service_proto_rawDesc
)

func file_proto_enumgrpc_quizz_service_proto_rawDescGZIP() []byte {
	file_proto_enumgrpc_quizz_service_proto_rawDescOnce.Do(func() {
		file_proto_enumgrpc_quizz_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_enumgrpc_quizz_service_proto_rawDescData)
	})
	return file_proto_enumgrpc_quizz_service_proto_rawDescData
}

var file_proto_enumgrpc_quizz_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_enumgrpc_quizz_service_proto_goTypes = []any{
	(EntityType)(0), // 0: app.enumgrpc.EntityType
	(ResultType)(0), // 1: app.enumgrpc.ResultType
}
var file_proto_enumgrpc_quizz_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_enumgrpc_quizz_service_proto_init() }
func file_proto_enumgrpc_quizz_service_proto_init() {
	if File_proto_enumgrpc_quizz_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_enumgrpc_quizz_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_enumgrpc_quizz_service_proto_goTypes,
		DependencyIndexes: file_proto_enumgrpc_quizz_service_proto_depIdxs,
		EnumInfos:         file_proto_enumgrpc_quizz_service_proto_enumTypes,
	}.Build()
	File_proto_enumgrpc_quizz_service_proto = out.File
	file_proto_enumgrpc_quizz_service_proto_rawDesc = nil
	file_proto_enumgrpc_quizz_service_proto_goTypes = nil
	file_proto_enumgrpc_quizz_service_proto_depIdxs = nil
}
