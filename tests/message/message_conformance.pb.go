// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: tests/message/message_conformance.proto

package message

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

type Basic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Basic) Reset() {
	*x = Basic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basic) ProtoMessage() {}

func (x *Basic) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basic.ProtoReflect.Descriptor instead.
func (*Basic) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{0}
}

func (x *Basic) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Basic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Union struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Contents:
	//
	//	*Union_Id
	//	*Union_Name
	Contents isUnion_Contents `protobuf_oneof:"contents"`
}

func (x *Union) Reset() {
	*x = Union{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Union) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Union) ProtoMessage() {}

func (x *Union) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Union.ProtoReflect.Descriptor instead.
func (*Union) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{1}
}

func (m *Union) GetContents() isUnion_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (x *Union) GetId() int32 {
	if x, ok := x.GetContents().(*Union_Id); ok {
		return x.Id
	}
	return 0
}

func (x *Union) GetName() string {
	if x, ok := x.GetContents().(*Union_Name); ok {
		return x.Name
	}
	return ""
}

type isUnion_Contents interface {
	isUnion_Contents()
}

type Union_Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type Union_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*Union_Id) isUnion_Contents() {}

func (*Union_Name) isUnion_Contents() {}

type Outer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Middle *Outer_Middle       `protobuf:"bytes,1,opt,name=middle,proto3" json:"middle,omitempty"`
	Inner  *Outer_Middle_Inner `protobuf:"bytes,2,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *Outer) Reset() {
	*x = Outer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outer) ProtoMessage() {}

func (x *Outer) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outer.ProtoReflect.Descriptor instead.
func (*Outer) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{2}
}

func (x *Outer) GetMiddle() *Outer_Middle {
	if x != nil {
		return x.Middle
	}
	return nil
}

func (x *Outer) GetInner() *Outer_Middle_Inner {
	if x != nil {
		return x.Inner
	}
	return nil
}

type Outer_Middle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *Outer_Middle_Inner `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *Outer_Middle) Reset() {
	*x = Outer_Middle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outer_Middle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outer_Middle) ProtoMessage() {}

func (x *Outer_Middle) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outer_Middle.ProtoReflect.Descriptor instead.
func (*Outer_Middle) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Outer_Middle) GetInner() *Outer_Middle_Inner {
	if x != nil {
		return x.Inner
	}
	return nil
}

type Outer_Middle_Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Outer_Middle_Inner) Reset() {
	*x = Outer_Middle_Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outer_Middle_Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outer_Middle_Inner) ProtoMessage() {}

func (x *Outer_Middle_Inner) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outer_Middle_Inner.ProtoReflect.Descriptor instead.
func (*Outer_Middle_Inner) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *Outer_Middle_Inner) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_tests_message_message_conformance_proto protoreflect.FileDescriptor

var file_tests_message_message_conformance_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x05, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x05, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x06,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x2e, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x5a, 0x0a, 0x06, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x17, 0x0a,
	0x05, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_message_message_conformance_proto_rawDescOnce sync.Once
	file_tests_message_message_conformance_proto_rawDescData = file_tests_message_message_conformance_proto_rawDesc
)

func file_tests_message_message_conformance_proto_rawDescGZIP() []byte {
	file_tests_message_message_conformance_proto_rawDescOnce.Do(func() {
		file_tests_message_message_conformance_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_message_message_conformance_proto_rawDescData)
	})
	return file_tests_message_message_conformance_proto_rawDescData
}

var file_tests_message_message_conformance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tests_message_message_conformance_proto_goTypes = []any{
	(*Basic)(nil),              // 0: tests.message.Basic
	(*Union)(nil),              // 1: tests.message.Union
	(*Outer)(nil),              // 2: tests.message.Outer
	(*Outer_Middle)(nil),       // 3: tests.message.Outer.Middle
	(*Outer_Middle_Inner)(nil), // 4: tests.message.Outer.Middle.Inner
}
var file_tests_message_message_conformance_proto_depIdxs = []int32{
	3, // 0: tests.message.Outer.middle:type_name -> tests.message.Outer.Middle
	4, // 1: tests.message.Outer.inner:type_name -> tests.message.Outer.Middle.Inner
	4, // 2: tests.message.Outer.Middle.inner:type_name -> tests.message.Outer.Middle.Inner
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tests_message_message_conformance_proto_init() }
func file_tests_message_message_conformance_proto_init() {
	if File_tests_message_message_conformance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_message_message_conformance_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Basic); i {
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
		file_tests_message_message_conformance_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Union); i {
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
		file_tests_message_message_conformance_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Outer); i {
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
		file_tests_message_message_conformance_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Outer_Middle); i {
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
		file_tests_message_message_conformance_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Outer_Middle_Inner); i {
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
	file_tests_message_message_conformance_proto_msgTypes[1].OneofWrappers = []any{
		(*Union_Id)(nil),
		(*Union_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_message_message_conformance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_message_message_conformance_proto_goTypes,
		DependencyIndexes: file_tests_message_message_conformance_proto_depIdxs,
		MessageInfos:      file_tests_message_message_conformance_proto_msgTypes,
	}.Build()
	File_tests_message_message_conformance_proto = out.File
	file_tests_message_message_conformance_proto_rawDesc = nil
	file_tests_message_message_conformance_proto_goTypes = nil
	file_tests_message_message_conformance_proto_depIdxs = nil
}
