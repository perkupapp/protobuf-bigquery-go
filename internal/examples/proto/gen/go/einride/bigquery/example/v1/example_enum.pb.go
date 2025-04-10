// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: einride/bigquery/example/v1/example_enum.proto

package examplev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExampleEnum_Enum int32

const (
	ExampleEnum_ENUM_UNSPECIFIED ExampleEnum_Enum = 0
	ExampleEnum_ENUM_VALUE1      ExampleEnum_Enum = 1
	ExampleEnum_ENUM_VALUE2      ExampleEnum_Enum = 2
)

// Enum value maps for ExampleEnum_Enum.
var (
	ExampleEnum_Enum_name = map[int32]string{
		0: "ENUM_UNSPECIFIED",
		1: "ENUM_VALUE1",
		2: "ENUM_VALUE2",
	}
	ExampleEnum_Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_VALUE1":      1,
		"ENUM_VALUE2":      2,
	}
)

func (x ExampleEnum_Enum) Enum() *ExampleEnum_Enum {
	p := new(ExampleEnum_Enum)
	*p = x
	return p
}

func (x ExampleEnum_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExampleEnum_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_bigquery_example_v1_example_enum_proto_enumTypes[0].Descriptor()
}

func (ExampleEnum_Enum) Type() protoreflect.EnumType {
	return &file_einride_bigquery_example_v1_example_enum_proto_enumTypes[0]
}

func (x ExampleEnum_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExampleEnum_Enum.Descriptor instead.
func (ExampleEnum_Enum) EnumDescriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_enum_proto_rawDescGZIP(), []int{0, 0}
}

type ExampleEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EnumValue     ExampleEnum_Enum       `protobuf:"varint,1,opt,name=enum_value,json=enumValue,proto3,enum=einride.bigquery.example.v1.ExampleEnum_Enum" json:"enum_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExampleEnum) Reset() {
	*x = ExampleEnum{}
	mi := &file_einride_bigquery_example_v1_example_enum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleEnum) ProtoMessage() {}

func (x *ExampleEnum) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_enum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleEnum.ProtoReflect.Descriptor instead.
func (*ExampleEnum) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_enum_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleEnum) GetEnumValue() ExampleEnum_Enum {
	if x != nil {
		return x.EnumValue
	}
	return ExampleEnum_ENUM_UNSPECIFIED
}

var File_einride_bigquery_example_v1_example_enum_proto protoreflect.FileDescriptor

var file_einride_bigquery_example_v1_example_enum_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x9b, 0x01,
	0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x4c, 0x0a,
	0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x04, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e,
	0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x32, 0x10, 0x02, 0x42, 0xa7, 0x02, 0x0a, 0x1f,
	0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x10, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x42, 0x45, 0xaa, 0x02,
	0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x45,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a,
	0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_einride_bigquery_example_v1_example_enum_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_enum_proto_rawDescData []byte
)

func file_einride_bigquery_example_v1_example_enum_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_enum_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_enum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_einride_bigquery_example_v1_example_enum_proto_rawDesc), len(file_einride_bigquery_example_v1_example_enum_proto_rawDesc)))
	})
	return file_einride_bigquery_example_v1_example_enum_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_bigquery_example_v1_example_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_example_v1_example_enum_proto_goTypes = []any{
	(ExampleEnum_Enum)(0), // 0: einride.bigquery.example.v1.ExampleEnum.Enum
	(*ExampleEnum)(nil),   // 1: einride.bigquery.example.v1.ExampleEnum
}
var file_einride_bigquery_example_v1_example_enum_proto_depIdxs = []int32{
	0, // 0: einride.bigquery.example.v1.ExampleEnum.enum_value:type_name -> einride.bigquery.example.v1.ExampleEnum.Enum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_enum_proto_init() }
func file_einride_bigquery_example_v1_example_enum_proto_init() {
	if File_einride_bigquery_example_v1_example_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_einride_bigquery_example_v1_example_enum_proto_rawDesc), len(file_einride_bigquery_example_v1_example_enum_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_example_v1_example_enum_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_example_v1_example_enum_proto_depIdxs,
		EnumInfos:         file_einride_bigquery_example_v1_example_enum_proto_enumTypes,
		MessageInfos:      file_einride_bigquery_example_v1_example_enum_proto_msgTypes,
	}.Build()
	File_einride_bigquery_example_v1_example_enum_proto = out.File
	file_einride_bigquery_example_v1_example_enum_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_enum_proto_depIdxs = nil
}
