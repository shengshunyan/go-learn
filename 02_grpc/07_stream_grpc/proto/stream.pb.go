// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: stream.proto

package proto

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

type StreamReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamReq) Reset() {
	*x = StreamReq{}
	mi := &file_stream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReq) ProtoMessage() {}

func (x *StreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReq.ProtoReflect.Descriptor instead.
func (*StreamReq) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *StreamReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamResp) Reset() {
	*x = StreamResp{}
	mi := &file_stream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResp) ProtoMessage() {}

func (x *StreamResp) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResp.ProtoReflect.Descriptor instead.
func (*StreamResp) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x1f, 0x0a, 0x09, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcb, 0x01,
	0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x09, 0x41,
	0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x30,
	0x37, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData []byte
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stream_proto_rawDesc), len(file_stream_proto_rawDesc)))
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stream_proto_goTypes = []any{
	(*StreamReq)(nil),  // 0: helloworld.StreamReq
	(*StreamResp)(nil), // 1: helloworld.StreamResp
}
var file_stream_proto_depIdxs = []int32{
	0, // 0: helloworld.Greeter.ServerStream:input_type -> helloworld.StreamReq
	0, // 1: helloworld.Greeter.ClientStream:input_type -> helloworld.StreamReq
	0, // 2: helloworld.Greeter.AllStream:input_type -> helloworld.StreamReq
	1, // 3: helloworld.Greeter.ServerStream:output_type -> helloworld.StreamResp
	1, // 4: helloworld.Greeter.ClientStream:output_type -> helloworld.StreamResp
	1, // 5: helloworld.Greeter.AllStream:output_type -> helloworld.StreamResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stream_proto_rawDesc), len(file_stream_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}
