//protobuf 版本3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: oneWay.proto

//go包名，go的空间域

package oneWay

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

//服务请求参数（结构体）
type OneWayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//类型 参数名 = 参数的序号（从1开始）
	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *OneWayReq) Reset() {
	*x = OneWayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneWay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneWayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneWayReq) ProtoMessage() {}

func (x *OneWayReq) ProtoReflect() protoreflect.Message {
	mi := &file_oneWay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneWayReq.ProtoReflect.Descriptor instead.
func (*OneWayReq) Descriptor() ([]byte, []int) {
	return file_oneWay_proto_rawDescGZIP(), []int{0}
}

func (x *OneWayReq) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *OneWayReq) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

//服务响应参数（结构体）
type OneWayRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int64 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *OneWayRes) Reset() {
	*x = OneWayRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneWay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneWayRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneWayRes) ProtoMessage() {}

func (x *OneWayRes) ProtoReflect() protoreflect.Message {
	mi := &file_oneWay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneWayRes.ProtoReflect.Descriptor instead.
func (*OneWayRes) Descriptor() ([]byte, []int) {
	return file_oneWay_proto_rawDescGZIP(), []int{1}
}

func (x *OneWayRes) GetRes() int64 {
	if x != nil {
		return x.Res
	}
	return 0
}

var File_oneWay_proto protoreflect.FileDescriptor

var file_oneWay_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x22, 0x27, 0x0a, 0x09, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x22,
	0x1d, 0x0a, 0x09, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x41,
	0x0a, 0x0d, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x06, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x12, 0x11, 0x2e, 0x6f, 0x6e, 0x65, 0x57,
	0x61, 0x79, 0x2e, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6f,
	0x6e, 0x65, 0x57, 0x61, 0x79, 0x2e, 0x4f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x6f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oneWay_proto_rawDescOnce sync.Once
	file_oneWay_proto_rawDescData = file_oneWay_proto_rawDesc
)

func file_oneWay_proto_rawDescGZIP() []byte {
	file_oneWay_proto_rawDescOnce.Do(func() {
		file_oneWay_proto_rawDescData = protoimpl.X.CompressGZIP(file_oneWay_proto_rawDescData)
	})
	return file_oneWay_proto_rawDescData
}

var file_oneWay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oneWay_proto_goTypes = []interface{}{
	(*OneWayReq)(nil), // 0: oneWay.OneWayReq
	(*OneWayRes)(nil), // 1: oneWay.OneWayRes
}
var file_oneWay_proto_depIdxs = []int32{
	0, // 0: oneWay.OneWayService.OneWay:input_type -> oneWay.OneWayReq
	1, // 1: oneWay.OneWayService.OneWay:output_type -> oneWay.OneWayRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oneWay_proto_init() }
func file_oneWay_proto_init() {
	if File_oneWay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oneWay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneWayReq); i {
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
		file_oneWay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneWayRes); i {
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
			RawDescriptor: file_oneWay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oneWay_proto_goTypes,
		DependencyIndexes: file_oneWay_proto_depIdxs,
		MessageInfos:      file_oneWay_proto_msgTypes,
	}.Build()
	File_oneWay_proto = out.File
	file_oneWay_proto_rawDesc = nil
	file_oneWay_proto_goTypes = nil
	file_oneWay_proto_depIdxs = nil
}
