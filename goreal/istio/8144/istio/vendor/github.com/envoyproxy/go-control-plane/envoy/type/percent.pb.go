// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: envoy/type/percent.proto

package _type

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Fraction percentages support several fixed denominator values.
type FractionalPercent_DenominatorType int32

const (
	// 100.
	//
	// **Example**: 1/100 = 1%.
	FractionalPercent_HUNDRED FractionalPercent_DenominatorType = 0
	// 10,000.
	//
	// **Example**: 1/10000 = 0.01%.
	FractionalPercent_TEN_THOUSAND FractionalPercent_DenominatorType = 1
	// 1,000,000.
	//
	// **Example**: 1/1000000 = 0.0001%.
	FractionalPercent_MILLION FractionalPercent_DenominatorType = 2
)

// Enum value maps for FractionalPercent_DenominatorType.
var (
	FractionalPercent_DenominatorType_name = map[int32]string{
		0: "HUNDRED",
		1: "TEN_THOUSAND",
		2: "MILLION",
	}
	FractionalPercent_DenominatorType_value = map[string]int32{
		"HUNDRED":      0,
		"TEN_THOUSAND": 1,
		"MILLION":      2,
	}
)

func (x FractionalPercent_DenominatorType) Enum() *FractionalPercent_DenominatorType {
	p := new(FractionalPercent_DenominatorType)
	*p = x
	return p
}

func (x FractionalPercent_DenominatorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FractionalPercent_DenominatorType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_type_percent_proto_enumTypes[0].Descriptor()
}

func (FractionalPercent_DenominatorType) Type() protoreflect.EnumType {
	return &file_envoy_type_percent_proto_enumTypes[0]
}

func (x FractionalPercent_DenominatorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FractionalPercent_DenominatorType.Descriptor instead.
func (FractionalPercent_DenominatorType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_type_percent_proto_rawDescGZIP(), []int{1, 0}
}

// Identifies a percentage, in the range [0.0, 100.0].
type Percent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Percent) Reset() {
	*x = Percent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_percent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Percent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Percent) ProtoMessage() {}

func (x *Percent) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_percent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Percent.ProtoReflect.Descriptor instead.
func (*Percent) Descriptor() ([]byte, []int) {
	return file_envoy_type_percent_proto_rawDescGZIP(), []int{0}
}

func (x *Percent) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// A fractional percentage is used in cases in which for performance reasons performing floating
// point to integer conversions during randomness calculations is undesirable. The message includes
// both a numerator and denominator that together determine the final fractional value.
//
// * **Example**: 1/100 = 1%.
// * **Example**: 3/10000 = 0.03%.
type FractionalPercent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the numerator. Defaults to 0.
	Numerator uint32 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// Specifies the denominator. If the denominator specified is less than the numerator, the final
	// fractional percentage is capped at 1 (100%).
	Denominator FractionalPercent_DenominatorType `protobuf:"varint,2,opt,name=denominator,proto3,enum=envoy.type.FractionalPercent_DenominatorType" json:"denominator,omitempty"`
}

func (x *FractionalPercent) Reset() {
	*x = FractionalPercent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_percent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FractionalPercent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FractionalPercent) ProtoMessage() {}

func (x *FractionalPercent) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_percent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FractionalPercent.ProtoReflect.Descriptor instead.
func (*FractionalPercent) Descriptor() ([]byte, []int) {
	return file_envoy_type_percent_proto_rawDescGZIP(), []int{1}
}

func (x *FractionalPercent) GetNumerator() uint32 {
	if x != nil {
		return x.Numerator
	}
	return 0
}

func (x *FractionalPercent) GetDenominator() FractionalPercent_DenominatorType {
	if x != nil {
		return x.Denominator
	}
	return FractionalPercent_HUNDRED
}

var File_envoy_type_percent_proto protoreflect.FileDescriptor

var file_envoy_type_percent_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38,
	0x0a, 0x07, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x11, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x59, 0x0a, 0x0b,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x0f, 0x44, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x55,
	0x4e, 0x44, 0x52, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x45, 0x4e, 0x5f, 0x54,
	0x48, 0x4f, 0x55, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x4c,
	0x4c, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x42, 0x65, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x0c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_percent_proto_rawDescOnce sync.Once
	file_envoy_type_percent_proto_rawDescData = file_envoy_type_percent_proto_rawDesc
)

func file_envoy_type_percent_proto_rawDescGZIP() []byte {
	file_envoy_type_percent_proto_rawDescOnce.Do(func() {
		file_envoy_type_percent_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_percent_proto_rawDescData)
	})
	return file_envoy_type_percent_proto_rawDescData
}

var file_envoy_type_percent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_type_percent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_type_percent_proto_goTypes = []interface{}{
	(FractionalPercent_DenominatorType)(0), // 0: envoy.type.FractionalPercent.DenominatorType
	(*Percent)(nil),                        // 1: envoy.type.Percent
	(*FractionalPercent)(nil),              // 2: envoy.type.FractionalPercent
}
var file_envoy_type_percent_proto_depIdxs = []int32{
	0, // 0: envoy.type.FractionalPercent.denominator:type_name -> envoy.type.FractionalPercent.DenominatorType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_type_percent_proto_init() }
func file_envoy_type_percent_proto_init() {
	if File_envoy_type_percent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_percent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Percent); i {
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
		file_envoy_type_percent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FractionalPercent); i {
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
			RawDescriptor: file_envoy_type_percent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_percent_proto_goTypes,
		DependencyIndexes: file_envoy_type_percent_proto_depIdxs,
		EnumInfos:         file_envoy_type_percent_proto_enumTypes,
		MessageInfos:      file_envoy_type_percent_proto_msgTypes,
	}.Build()
	File_envoy_type_percent_proto = out.File
	file_envoy_type_percent_proto_rawDesc = nil
	file_envoy_type_percent_proto_goTypes = nil
	file_envoy_type_percent_proto_depIdxs = nil
}
