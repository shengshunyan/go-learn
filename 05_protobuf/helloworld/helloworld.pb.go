// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: helloworld.proto

package helloworld

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

// 定义枚举类型
type Color int32

const (
	// 第一个枚举值必须为 0
	Color_COLOR_UNSPECIFIED Color = 0
	Color_RED               Color = 1
	Color_GREEN             Color = 2
	Color_BLUE              Color = 3
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "COLOR_UNSPECIFIED",
		1: "RED",
		2: "GREEN",
		3: "BLUE",
	}
	Color_value = map[string]int32{
		"COLOR_UNSPECIFIED": 0,
		"RED":               1,
		"GREEN":             2,
		"BLUE":              3,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_proto_enumTypes[0].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_helloworld_proto_enumTypes[0]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 1是编号
	Color         Color                  `protobuf:"varint,2,opt,name=color,proto3,enum=hello_world.Color" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_helloworld_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetColor() Color {
	if x != nil {
		return x.Color
	}
	return Color_COLOR_UNSPECIFIED
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         string                 `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_helloworld_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22,
	0x4c, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x25, 0x0a,
	0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x2a, 0x3c, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x15, 0x0a,
	0x11, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45,
	0x10, 0x03, 0x32, 0x4c, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x12, 0x3e, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x18, 0x5a, 0x16, 0x30, 0x35, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_proto_rawDescData []byte
)

func file_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_helloworld_proto_rawDesc), len(file_helloworld_proto_rawDesc)))
	})
	return file_helloworld_proto_rawDescData
}

var file_helloworld_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_helloworld_proto_goTypes = []any{
	(Color)(0),            // 0: hello_world.Color
	(*HelloRequest)(nil),  // 1: hello_world.HelloRequest
	(*HelloResponse)(nil), // 2: hello_world.HelloResponse
}
var file_helloworld_proto_depIdxs = []int32{
	0, // 0: hello_world.HelloRequest.color:type_name -> hello_world.Color
	1, // 1: hello_world.HelloWorld.Hello:input_type -> hello_world.HelloRequest
	2, // 2: hello_world.HelloWorld.Hello:output_type -> hello_world.HelloResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_helloworld_proto_init() }
func file_helloworld_proto_init() {
	if File_helloworld_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_helloworld_proto_rawDesc), len(file_helloworld_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_proto_depIdxs,
		EnumInfos:         file_helloworld_proto_enumTypes,
		MessageInfos:      file_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_proto = out.File
	file_helloworld_proto_goTypes = nil
	file_helloworld_proto_depIdxs = nil
}
