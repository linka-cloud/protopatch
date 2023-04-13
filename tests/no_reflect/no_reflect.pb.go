// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: tests/no_reflect/no_reflect.proto

package no_reflect

import (
	_ "github.com/alta/protopatch/patch/gopb"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

type Enum int32

const (
	Enum_ENUM_VALUE Enum = 0
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ENUM_VALUE",
	}
	Enum_value = map[string]int32{
		"ENUM_VALUE": 0,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

type NoReflect struct {
	unknownFields protoimpl.UnknownFields

	StringField string  `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	Int32Field  int32   `protobuf:"varint,2,opt,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	Int64Field  int64   `protobuf:"varint,3,opt,name=int64_field,json=int64Field,proto3" json:"int64_field,omitempty"`
	Uint32Field uint32  `protobuf:"varint,4,opt,name=uint32_field,json=uint32Field,proto3" json:"uint32_field,omitempty"`
	Uint64Field uint64  `protobuf:"varint,5,opt,name=uint64_field,json=uint64Field,proto3" json:"uint64_field,omitempty"`
	BoolField   bool    `protobuf:"varint,6,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	FloatField  float32 `protobuf:"fixed32,7,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	DoubleField float64 `protobuf:"fixed64,8,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	BytesField  []byte  `protobuf:"bytes,9,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	NestedField *Nested `protobuf:"bytes,10,opt,name=nested_field,json=nestedField,proto3" json:"nested_field,omitempty"`
}

type Nested struct {
	unknownFields protoimpl.UnknownFields

	StringField string `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
}

// 0: tests.no_reflect.Enum
// 1: tests.no_reflect.NoReflect
// 2: tests.no_reflect.Nested

// 0: tests.no_reflect.NoReflect.nested_field:type_name -> tests.no_reflect.Nested
// [1:1] is the sub-list for method output_type
// [1:1] is the sub-list for method input_type
// [1:1] is the sub-list for extension type_name
// [1:1] is the sub-list for extension extendee
// [0:1] is the sub-list for field type_name
