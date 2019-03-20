// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/messages/shared/uuid.proto

package sharedproto

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

type UUID struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b541901160df4f8c, []int{0}
}

func (m *UUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUID.Unmarshal(m, b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return xxx_messageInfo_UUID.Size(m)
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type ObjectKey struct {
	ID                   *UUID    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	LocationID           *UUID    `protobuf:"bytes,2,opt,name=LocationID,proto3" json:"LocationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectKey) Reset()         { *m = ObjectKey{} }
func (m *ObjectKey) String() string { return proto.CompactTextString(m) }
func (*ObjectKey) ProtoMessage()    {}
func (*ObjectKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b541901160df4f8c, []int{1}
}

func (m *ObjectKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectKey.Unmarshal(m, b)
}
func (m *ObjectKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectKey.Marshal(b, m, deterministic)
}
func (m *ObjectKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectKey.Merge(m, src)
}
func (m *ObjectKey) XXX_Size() int {
	return xxx_messageInfo_ObjectKey.Size(m)
}
func (m *ObjectKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectKey.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectKey proto.InternalMessageInfo

func (m *ObjectKey) GetID() *UUID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *ObjectKey) GetLocationID() *UUID {
	if m != nil {
		return m.LocationID
	}
	return nil
}

func init() {
	proto.RegisterType((*UUID)(nil), "shared.UUID")
	proto.RegisterType((*ObjectKey)(nil), "shared.ObjectKey")
}

func init() {
	proto.RegisterFile("protorepo/messages/shared/uuid.proto", fileDescriptor_b541901160df4f8c)
}

var fileDescriptor_b541901160df4f8c = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0x4a, 0x2d, 0xc8, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2f,
	0xce, 0x48, 0x2c, 0x4a, 0x4d, 0xd1, 0x2f, 0x2d, 0xcd, 0x4c, 0xd1, 0x03, 0x4b, 0x0b, 0xb1, 0x41,
	0x84, 0x94, 0x64, 0xb8, 0x58, 0x42, 0x43, 0x3d, 0x5d, 0x84, 0x44, 0xb8, 0x58, 0x9d, 0x2a, 0x4b,
	0x52, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0xa5, 0x70, 0x2e, 0x4e, 0xff,
	0xa4, 0xac, 0xd4, 0xe4, 0x12, 0xef, 0xd4, 0x4a, 0x21, 0x19, 0x2e, 0x26, 0x4f, 0x17, 0xb0, 0x3c,
	0xb7, 0x11, 0x8f, 0x1e, 0x44, 0xbf, 0x1e, 0x48, 0x73, 0x10, 0x93, 0xa7, 0x8b, 0x90, 0x0e, 0x17,
	0x97, 0x4f, 0x7e, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0xa7, 0x8b, 0x04, 0x13, 0x16, 0x55, 0x48,
	0xf2, 0x4e, 0xae, 0x51, 0xce, 0xe5, 0xa9, 0x89, 0x65, 0xa9, 0x39, 0x89, 0x49, 0x7a, 0x15, 0x95,
	0x55, 0xfa, 0xb9, 0xf9, 0x79, 0xf9, 0x45, 0x89, 0x99, 0x39, 0x30, 0x97, 0x22, 0xfc, 0x90, 0x92,
	0x59, 0x5c, 0xa2, 0x9f, 0x8e, 0xe1, 0x17, 0xb0, 0x82, 0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0x9e, 0xcd, 0x94, 0xf4, 0x00, 0x00, 0x00,
}
