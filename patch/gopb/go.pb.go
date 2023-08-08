// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: patch/go.proto

package gopb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Options represent Go-specific options for Protobuf messages, fields, oneofs, enums, or enum values.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name option renames the generated Go identifier and related identifiers.
	// For a message, this renames the generated Go struct and nested messages or enums, if any.
	// For a message field, this renames the generated Go struct field and getter method.
	// For a oneof field, this renames the generated Go struct field, getter method, interface type, and wrapper types.
	// For an enum, this renames the generated Go type.
	// For an enum value, this renames the generated Go const.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The embed option indicates the field should be embedded in the generated Go struct.
	// Only message types can be embedded. Oneof fields cannot be embedded.
	// See https://golang.org/ref/spec#Struct_types.
	Embed *bool `protobuf:"varint,2,opt,name=embed" json:"embed,omitempty"`
	// The type option changes the generated field type.
	// All generated code assumes that this type is castable to the protocol buffer field type.
	Type *string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	// The getter option renames the generated getter method (default: Get<Field>)
	// so a custom getter can be implemented in its place.
	Getter *string `protobuf:"bytes,10,opt,name=getter" json:"getter,omitempty"` // TODO: implement this
	// The tags option specifies additional struct tags which are appended a generated Go struct field.
	// This option may be specified on a message field or a oneof field.
	// The value should omit the enclosing backticks.
	Tags *string `protobuf:"bytes,20,opt,name=tags" json:"tags,omitempty"`
	// The stringer option renames a generated String() method (if any)
	// so a custom String() method can be implemented in its place.
	Stringer *string `protobuf:"bytes,30,opt,name=stringer" json:"stringer,omitempty"` // TODO: implement for messages
	// The stringer_name option is a deprecated alias for stringer.
	// It will be removed in a future version of this package.
	StringerName *string `protobuf:"bytes,31,opt,name=stringer_name,json=stringerName" json:"stringer_name,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patch_go_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_patch_go_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_patch_go_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Options) GetEmbed() bool {
	if x != nil && x.Embed != nil {
		return *x.Embed
	}
	return false
}

func (x *Options) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *Options) GetGetter() string {
	if x != nil && x.Getter != nil {
		return *x.Getter
	}
	return ""
}

func (x *Options) GetTags() string {
	if x != nil && x.Tags != nil {
		return *x.Tags
	}
	return ""
}

func (x *Options) GetStringer() string {
	if x != nil && x.Stringer != nil {
		return *x.Stringer
	}
	return ""
}

func (x *Options) GetStringerName() string {
	if x != nil && x.StringerName != nil {
		return *x.StringerName
	}
	return ""
}

// LintOptions represent options for linting a generated Go file.
type LintOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set all to true if all generated Go symbols should be linted.
	// This option affects generated structs, struct fields, enum types, and value constants.
	All *bool `protobuf:"varint,1,opt,name=all" json:"all,omitempty"`
	// Set messages to true if message names should be linted.
	// This does not affect message fields.
	Messages *bool `protobuf:"varint,2,opt,name=messages" json:"messages,omitempty"`
	// Set messages to true if message field names should be linted.
	// This does not affect message fields.
	Fields *bool `protobuf:"varint,3,opt,name=fields" json:"fields,omitempty"`
	// Set enums to true if generated enum names should be linted.
	// This does not affect enum values.
	Enums *bool `protobuf:"varint,4,opt,name=enums" json:"enums,omitempty"`
	// Set values to true if generated enum value constants should be linted.
	Values *bool `protobuf:"varint,5,opt,name=values" json:"values,omitempty"`
	// Set extensions to true if generated extension names should be linted.
	Extensions *bool `protobuf:"varint,6,opt,name=extensions" json:"extensions,omitempty"`
	// The initialisms option lets you specify strings that should not be generated as mixed-case,
	// Examples: ID, URL, HTTP, etc.
	Initialisms []string `protobuf:"bytes,10,rep,name=initialisms" json:"initialisms,omitempty"`
}

func (x *LintOptions) Reset() {
	*x = LintOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patch_go_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LintOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintOptions) ProtoMessage() {}

func (x *LintOptions) ProtoReflect() protoreflect.Message {
	mi := &file_patch_go_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintOptions.ProtoReflect.Descriptor instead.
func (*LintOptions) Descriptor() ([]byte, []int) {
	return file_patch_go_proto_rawDescGZIP(), []int{1}
}

func (x *LintOptions) GetAll() bool {
	if x != nil && x.All != nil {
		return *x.All
	}
	return false
}

func (x *LintOptions) GetMessages() bool {
	if x != nil && x.Messages != nil {
		return *x.Messages
	}
	return false
}

func (x *LintOptions) GetFields() bool {
	if x != nil && x.Fields != nil {
		return *x.Fields
	}
	return false
}

func (x *LintOptions) GetEnums() bool {
	if x != nil && x.Enums != nil {
		return *x.Enums
	}
	return false
}

func (x *LintOptions) GetValues() bool {
	if x != nil && x.Values != nil {
		return *x.Values
	}
	return false
}

func (x *LintOptions) GetExtensions() bool {
	if x != nil && x.Extensions != nil {
		return *x.Extensions
	}
	return false
}

func (x *LintOptions) GetInitialisms() []string {
	if x != nil {
		return x.Initialisms
	}
	return nil
}

var file_patch_go_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         7001,
		Name:          "go.message",
		Tag:           "bytes,7001,opt,name=message",
		Filename:      "patch/go.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         7001,
		Name:          "go.field",
		Tag:           "bytes,7001,opt,name=field",
		Filename:      "patch/go.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         7001,
		Name:          "go.oneof",
		Tag:           "bytes,7001,opt,name=oneof",
		Filename:      "patch/go.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         7001,
		Name:          "go.enum",
		Tag:           "bytes,7001,opt,name=enum",
		Filename:      "patch/go.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         7001,
		Name:          "go.value",
		Tag:           "bytes,7001,opt,name=value",
		Filename:      "patch/go.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*LintOptions)(nil),
		Field:         7001,
		Name:          "go.lint",
		Tag:           "bytes,7001,opt,name=lint",
		Filename:      "patch/go.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional go.Options message = 7001;
	E_Message = &file_patch_go_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional go.Options field = 7001;
	E_Field = &file_patch_go_proto_extTypes[1]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional go.Options oneof = 7001;
	E_Oneof = &file_patch_go_proto_extTypes[2]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional go.Options enum = 7001;
	E_Enum = &file_patch_go_proto_extTypes[3]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional go.Options value = 7001;
	E_Value = &file_patch_go_proto_extTypes[4]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional go.LintOptions lint = 7001;
	E_Lint = &file_patch_go_proto_extTypes[5]
)

var File_patch_go_proto protoreflect.FileDescriptor

var file_patch_go_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x67, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc3, 0x01,
	0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x73, 0x6d, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x73, 0x6d, 0x73, 0x3a, 0x47, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd9, 0x36, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x41, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x36, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x6f,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a,
	0x41, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x36, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x67, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x3a, 0x3e, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x36, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x67, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x65, 0x6e,
	0x75, 0x6d, 0x3a, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9,
	0x36, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x42, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd9, 0x36, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x74, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x2f, 0x67, 0x6f, 0x70, 0x62,
}

var (
	file_patch_go_proto_rawDescOnce sync.Once
	file_patch_go_proto_rawDescData = file_patch_go_proto_rawDesc
)

func file_patch_go_proto_rawDescGZIP() []byte {
	file_patch_go_proto_rawDescOnce.Do(func() {
		file_patch_go_proto_rawDescData = protoimpl.X.CompressGZIP(file_patch_go_proto_rawDescData)
	})
	return file_patch_go_proto_rawDescData
}

var file_patch_go_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_patch_go_proto_goTypes = []interface{}{
	(*Options)(nil),                       // 0: go.Options
	(*LintOptions)(nil),                   // 1: go.LintOptions
	(*descriptorpb.MessageOptions)(nil),   // 2: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),     // 3: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil),     // 4: google.protobuf.OneofOptions
	(*descriptorpb.EnumOptions)(nil),      // 5: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
	(*descriptorpb.FileOptions)(nil),      // 7: google.protobuf.FileOptions
}
var file_patch_go_proto_depIdxs = []int32{
	2,  // 0: go.message:extendee -> google.protobuf.MessageOptions
	3,  // 1: go.field:extendee -> google.protobuf.FieldOptions
	4,  // 2: go.oneof:extendee -> google.protobuf.OneofOptions
	5,  // 3: go.enum:extendee -> google.protobuf.EnumOptions
	6,  // 4: go.value:extendee -> google.protobuf.EnumValueOptions
	7,  // 5: go.lint:extendee -> google.protobuf.FileOptions
	0,  // 6: go.message:type_name -> go.Options
	0,  // 7: go.field:type_name -> go.Options
	0,  // 8: go.oneof:type_name -> go.Options
	0,  // 9: go.enum:type_name -> go.Options
	0,  // 10: go.value:type_name -> go.Options
	1,  // 11: go.lint:type_name -> go.LintOptions
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	6,  // [6:12] is the sub-list for extension type_name
	0,  // [0:6] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_patch_go_proto_init() }
func file_patch_go_proto_init() {
	if File_patch_go_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_patch_go_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
		file_patch_go_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LintOptions); i {
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
			RawDescriptor: file_patch_go_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_patch_go_proto_goTypes,
		DependencyIndexes: file_patch_go_proto_depIdxs,
		MessageInfos:      file_patch_go_proto_msgTypes,
		ExtensionInfos:    file_patch_go_proto_extTypes,
	}.Build()
	File_patch_go_proto = out.File
	file_patch_go_proto_rawDesc = nil
	file_patch_go_proto_goTypes = nil
	file_patch_go_proto_depIdxs = nil
}
