// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/enums/insys/agreements.proto

package insysenums

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Feature int32

const (
	Feature_PremiumFeatureSignUp Feature = 0
)

var Feature_name = map[int32]string{
	0: "PremiumFeatureSignUp",
}

var Feature_value = map[string]int32{
	"PremiumFeatureSignUp": 0,
}

func (x Feature) String() string {
	return proto.EnumName(Feature_name, int32(x))
}

func (Feature) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63e2eac2d238d9ba, []int{0}
}

func init() {
	proto.RegisterEnum("agreementsenums.Feature", Feature_name, Feature_value)
}

func init() {
	proto.RegisterFile("protorepo/enums/insys/agreements.proto", fileDescriptor_63e2eac2d238d9ba)
}

var fileDescriptor_63e2eac2d238d9ba = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0x4a, 0x2d, 0xc8, 0xd7, 0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0xd6, 0xcf, 0xcc, 0x2b, 0xae,
	0x2c, 0xd6, 0x4f, 0x4c, 0x2f, 0x4a, 0x4d, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0xd6, 0x03, 0x2b, 0x10,
	0xe2, 0x47, 0x88, 0x80, 0xd5, 0x69, 0x29, 0x73, 0xb1, 0xbb, 0xa5, 0x26, 0x96, 0x94, 0x16, 0xa5,
	0x0a, 0x49, 0x70, 0x89, 0x04, 0x14, 0xa5, 0xe6, 0x66, 0x96, 0xe6, 0x42, 0x45, 0x82, 0x33, 0xd3,
	0xf3, 0x42, 0x0b, 0x04, 0x18, 0x9c, 0x1c, 0xa3, 0xec, 0xcb, 0x53, 0x13, 0xcb, 0x52, 0x73, 0x12,
	0x93, 0xf4, 0x2a, 0x2a, 0xab, 0xf4, 0x73, 0xf3, 0xf3, 0xf2, 0x8b, 0x12, 0x33, 0x73, 0xf4, 0x8b,
	0x33, 0x12, 0x8b, 0x52, 0x53, 0xf4, 0x11, 0x96, 0xa7, 0x64, 0x16, 0x97, 0xe8, 0xa7, 0xa3, 0x38,
	0x02, 0xcc, 0x4c, 0x62, 0x03, 0xab, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x78, 0x3a, 0xee,
	0x18, 0xa9, 0x00, 0x00, 0x00,
}
