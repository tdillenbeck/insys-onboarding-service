// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/enums/platform/useremails.proto

package platformenums

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

type EmailType int32

const (
	EmailType_unknown       EmailType = 0
	EmailType_resetPassword EmailType = 1
	EmailType_inviteUser    EmailType = 2
	EmailType_voicemail     EmailType = 3
	EmailType_provision     EmailType = 4
)

var EmailType_name = map[int32]string{
	0: "unknown",
	1: "resetPassword",
	2: "inviteUser",
	3: "voicemail",
	4: "provision",
}

var EmailType_value = map[string]int32{
	"unknown":       0,
	"resetPassword": 1,
	"inviteUser":    2,
	"voicemail":     3,
	"provision":     4,
}

func (x EmailType) String() string {
	return proto.EnumName(EmailType_name, int32(x))
}

func (EmailType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2b3a88ae3b91f714, []int{0}
}

func init() {
	proto.RegisterEnum("platformenums.EmailType", EmailType_name, EmailType_value)
}

func init() {
	proto.RegisterFile("protorepo/enums/platform/useremails.proto", fileDescriptor_2b3a88ae3b91f714)
}

var fileDescriptor_2b3a88ae3b91f714 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x46, 0xf9, 0x13, 0x28, 0x46, 0x41, 0xc6, 0xb7, 0x80, 0x45, 0xbc, 0xe0, 0x06, 0x11, 0xec,
	0x59, 0xc0, 0x02, 0x76, 0x0e, 0x19, 0xda, 0x51, 0x6d, 0x8f, 0x35, 0xe3, 0x38, 0x4d, 0x4f, 0x5f,
	0xc5, 0x52, 0x55, 0xa9, 0xcb, 0x4f, 0xef, 0xe9, 0xd3, 0x53, 0x2f, 0x89, 0x29, 0x13, 0x43, 0x22,
	0x0b, 0x71, 0x0a, 0x62, 0x93, 0x77, 0xf9, 0x9f, 0x38, 0xd8, 0x49, 0x80, 0x21, 0x38, 0xf4, 0xd2,
	0x55, 0xc7, 0xb4, 0x27, 0x54, 0xc5, 0xd7, 0x1f, 0xd5, 0x7c, 0xac, 0xf8, 0x6b, 0x49, 0x60, 0x1e,
	0xd5, 0xc3, 0x14, 0x77, 0x91, 0xe6, 0xa8, 0xaf, 0xcc, 0xb3, 0x6a, 0x19, 0x04, 0xf2, 0xa7, 0x13,
	0x99, 0x89, 0x47, 0x7d, 0x6d, 0x9e, 0x94, 0xc2, 0x58, 0x30, 0xc3, 0xb7, 0x00, 0xeb, 0x1b, 0xd3,
	0xaa, 0xa6, 0x10, 0xfe, 0xd5, 0x7f, 0x7d, 0xbb, 0xce, 0xc4, 0x54, 0x50, 0x90, 0xa2, 0xbe, 0xeb,
	0xdf, 0x7f, 0xfb, 0x19, 0x5c, 0x01, 0xef, 0x86, 0x6e, 0xbf, 0x1c, 0x6c, 0xa0, 0x48, 0xec, 0xd0,
	0x5b, 0xd9, 0x3a, 0x86, 0xd1, 0x9e, 0x9b, 0x47, 0x94, 0x6c, 0x37, 0x97, 0xed, 0x75, 0x0d, 0xf7,
	0x55, 0x7b, 0x3b, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x92, 0x63, 0x48, 0xe3, 0x00, 0x00, 0x00,
}