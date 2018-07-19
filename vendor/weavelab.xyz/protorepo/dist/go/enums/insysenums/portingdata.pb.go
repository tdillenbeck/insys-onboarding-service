// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/enums/insys/portingdata.proto

package insysenums // import "weavelab.xyz/protorepo/dist/go/enums/insysenums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PortType int32

const (
	PortType_PortRequest  PortType = 0
	PortType_NewNumber    PortType = 1
	PortType_InternalPort PortType = 2
)

var PortType_name = map[int32]string{
	0: "PortRequest",
	1: "NewNumber",
	2: "InternalPort",
}
var PortType_value = map[string]int32{
	"PortRequest":  0,
	"NewNumber":    1,
	"InternalPort": 2,
}

func (x PortType) String() string {
	return proto.EnumName(PortType_name, int32(x))
}
func (PortType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_portingdata_a186defbfd5df1e2, []int{0}
}

type PortStatus int32

const (
	PortStatus_Pending           PortStatus = 0
	PortStatus_Submitted         PortStatus = 1
	PortStatus_Rejected          PortStatus = 2
	PortStatus_RejectionResolved PortStatus = 3
	PortStatus_Resubmitted       PortStatus = 4
	PortStatus_SUPPending        PortStatus = 5
	PortStatus_SUPSubmitted      PortStatus = 6
	PortStatus_CancelPending     PortStatus = 7
	PortStatus_CancelSubmitted   PortStatus = 8
	PortStatus_Cancelled         PortStatus = 9
	PortStatus_Accepted          PortStatus = 10
	PortStatus_Completed         PortStatus = 11
)

var PortStatus_name = map[int32]string{
	0:  "Pending",
	1:  "Submitted",
	2:  "Rejected",
	3:  "RejectionResolved",
	4:  "Resubmitted",
	5:  "SUPPending",
	6:  "SUPSubmitted",
	7:  "CancelPending",
	8:  "CancelSubmitted",
	9:  "Cancelled",
	10: "Accepted",
	11: "Completed",
}
var PortStatus_value = map[string]int32{
	"Pending":           0,
	"Submitted":         1,
	"Rejected":          2,
	"RejectionResolved": 3,
	"Resubmitted":       4,
	"SUPPending":        5,
	"SUPSubmitted":      6,
	"CancelPending":     7,
	"CancelSubmitted":   8,
	"Cancelled":         9,
	"Accepted":          10,
	"Completed":         11,
}

func (x PortStatus) String() string {
	return proto.EnumName(PortStatus_name, int32(x))
}
func (PortStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_portingdata_a186defbfd5df1e2, []int{1}
}

type CurrentAccountType int32

const (
	CurrentAccountType_Business    CurrentAccountType = 0
	CurrentAccountType_Residential CurrentAccountType = 1
)

var CurrentAccountType_name = map[int32]string{
	0: "Business",
	1: "Residential",
}
var CurrentAccountType_value = map[string]int32{
	"Business":    0,
	"Residential": 1,
}

func (x CurrentAccountType) String() string {
	return proto.EnumName(CurrentAccountType_name, int32(x))
}
func (CurrentAccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_portingdata_a186defbfd5df1e2, []int{2}
}

func init() {
	proto.RegisterEnum("portingenums.PortType", PortType_name, PortType_value)
	proto.RegisterEnum("portingenums.PortStatus", PortStatus_name, PortStatus_value)
	proto.RegisterEnum("portingenums.CurrentAccountType", CurrentAccountType_name, CurrentAccountType_value)
}

func init() {
	proto.RegisterFile("protorepo/enums/insys/portingdata.proto", fileDescriptor_portingdata_a186defbfd5df1e2)
}

var fileDescriptor_portingdata_a186defbfd5df1e2 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x6e, 0xc2, 0x30,
	0x10, 0x84, 0x09, 0x6d, 0x21, 0x2c, 0xa1, 0x18, 0x57, 0x7d, 0x86, 0x4a, 0x1c, 0x88, 0x2a, 0xae,
	0xbd, 0x00, 0xa7, 0x5e, 0x50, 0x94, 0x94, 0x4b, 0x6f, 0x4e, 0xbc, 0x42, 0xae, 0x1c, 0x3b, 0xf5,
	0x0f, 0x94, 0x3e, 0x67, 0x1f, 0xa8, 0x72, 0x02, 0x45, 0xbd, 0xcd, 0x8c, 0xbf, 0x5d, 0x8d, 0xbc,
	0xf0, 0xd4, 0x18, 0xed, 0xb4, 0xc1, 0x46, 0xa7, 0xa8, 0x7c, 0x6d, 0x53, 0xa1, 0xec, 0xc9, 0xa6,
	0x8d, 0x36, 0x4e, 0xa8, 0x3d, 0x67, 0x8e, 0x2d, 0x5a, 0x82, 0x26, 0xe7, 0xa8, 0xa5, 0xe6, 0x2f,
	0x10, 0x67, 0xda, 0xb8, 0xb7, 0x53, 0x83, 0x74, 0x0a, 0xe3, 0xa0, 0x73, 0xfc, 0xf4, 0x68, 0x1d,
	0xe9, 0xd1, 0x09, 0x8c, 0xb6, 0x78, 0xdc, 0xfa, 0xba, 0x44, 0x43, 0x22, 0x4a, 0x20, 0x79, 0x55,
	0x0e, 0x8d, 0x62, 0x32, 0x70, 0xa4, 0x3f, 0xff, 0x89, 0x00, 0x82, 0x2c, 0x1c, 0x73, 0xde, 0xd2,
	0x31, 0x0c, 0x33, 0x54, 0x5c, 0xa8, 0x7d, 0x37, 0x5c, 0xf8, 0xb2, 0x16, 0xce, 0x21, 0x27, 0x11,
	0x4d, 0x20, 0xce, 0xf1, 0x03, 0xab, 0xe0, 0xfa, 0xf4, 0x11, 0x66, 0x9d, 0x13, 0x5a, 0xe5, 0x68,
	0xb5, 0x3c, 0x20, 0x27, 0x37, 0xa1, 0x41, 0x8e, 0xf6, 0x6f, 0xea, 0x96, 0xde, 0x03, 0x14, 0xbb,
	0xec, 0xb2, 0xf4, 0x2e, 0x54, 0x28, 0x76, 0xd9, 0x75, 0xef, 0x80, 0xce, 0x60, 0xb2, 0x61, 0xaa,
	0x42, 0x79, 0x81, 0x86, 0xf4, 0x01, 0xa6, 0x5d, 0x74, 0xe5, 0xe2, 0x50, 0xa7, 0x0b, 0x25, 0x72,
	0x32, 0x0a, 0x75, 0x56, 0x55, 0x85, 0x4d, 0x78, 0x84, 0xf6, 0x51, 0xd7, 0x8d, 0xc4, 0x60, 0xc7,
	0xf3, 0x25, 0xd0, 0x8d, 0x37, 0x06, 0x95, 0x5b, 0x55, 0x95, 0xf6, 0xaa, 0xfb, 0x9e, 0x04, 0xe2,
	0xb5, 0xb7, 0x42, 0xa1, 0xb5, 0xa4, 0x77, 0xae, 0x2a, 0x38, 0x2a, 0x27, 0x98, 0x24, 0xd1, 0xfa,
	0xf9, 0x3d, 0x3d, 0x22, 0x3b, 0xa0, 0x64, 0xe5, 0xe2, 0xeb, 0xf4, 0x9d, 0x5e, 0xef, 0xc1, 0x85,
	0x75, 0xe9, 0xfe, 0xdf, 0x5d, 0x5a, 0x59, 0x0e, 0x5a, 0x66, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0x3a, 0x2b, 0x95, 0xbc, 0x01, 0x00, 0x00,
}