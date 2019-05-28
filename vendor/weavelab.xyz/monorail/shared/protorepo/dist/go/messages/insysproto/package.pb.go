// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/messages/insys/package.proto

package insysproto

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

type SetFeaturesFromSalesforceProductsRequest struct {
	LocationId           string   `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Products             []string `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Vertical             int32    `protobuf:"varint,3,opt,name=vertical,proto3" json:"vertical,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFeaturesFromSalesforceProductsRequest) Reset() {
	*m = SetFeaturesFromSalesforceProductsRequest{}
}
func (m *SetFeaturesFromSalesforceProductsRequest) String() string { return proto.CompactTextString(m) }
func (*SetFeaturesFromSalesforceProductsRequest) ProtoMessage()    {}
func (*SetFeaturesFromSalesforceProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b757e2aec27b58d2, []int{0}
}

func (m *SetFeaturesFromSalesforceProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFeaturesFromSalesforceProductsRequest.Unmarshal(m, b)
}
func (m *SetFeaturesFromSalesforceProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFeaturesFromSalesforceProductsRequest.Marshal(b, m, deterministic)
}
func (m *SetFeaturesFromSalesforceProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFeaturesFromSalesforceProductsRequest.Merge(m, src)
}
func (m *SetFeaturesFromSalesforceProductsRequest) XXX_Size() int {
	return xxx_messageInfo_SetFeaturesFromSalesforceProductsRequest.Size(m)
}
func (m *SetFeaturesFromSalesforceProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFeaturesFromSalesforceProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetFeaturesFromSalesforceProductsRequest proto.InternalMessageInfo

func (m *SetFeaturesFromSalesforceProductsRequest) GetLocationId() string {
	if m != nil {
		return m.LocationId
	}
	return ""
}

func (m *SetFeaturesFromSalesforceProductsRequest) GetProducts() []string {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *SetFeaturesFromSalesforceProductsRequest) GetVertical() int32 {
	if m != nil {
		return m.Vertical
	}
	return 0
}

type SetFeaturesFromSalesforceProductsResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Products             []string `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Features             []string `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFeaturesFromSalesforceProductsResponse) Reset() {
	*m = SetFeaturesFromSalesforceProductsResponse{}
}
func (m *SetFeaturesFromSalesforceProductsResponse) String() string { return proto.CompactTextString(m) }
func (*SetFeaturesFromSalesforceProductsResponse) ProtoMessage()    {}
func (*SetFeaturesFromSalesforceProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b757e2aec27b58d2, []int{1}
}

func (m *SetFeaturesFromSalesforceProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFeaturesFromSalesforceProductsResponse.Unmarshal(m, b)
}
func (m *SetFeaturesFromSalesforceProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFeaturesFromSalesforceProductsResponse.Marshal(b, m, deterministic)
}
func (m *SetFeaturesFromSalesforceProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFeaturesFromSalesforceProductsResponse.Merge(m, src)
}
func (m *SetFeaturesFromSalesforceProductsResponse) XXX_Size() int {
	return xxx_messageInfo_SetFeaturesFromSalesforceProductsResponse.Size(m)
}
func (m *SetFeaturesFromSalesforceProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFeaturesFromSalesforceProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetFeaturesFromSalesforceProductsResponse proto.InternalMessageInfo

func (m *SetFeaturesFromSalesforceProductsResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SetFeaturesFromSalesforceProductsResponse) GetProducts() []string {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *SetFeaturesFromSalesforceProductsResponse) GetFeatures() []string {
	if m != nil {
		return m.Features
	}
	return nil
}

func init() {
	proto.RegisterType((*SetFeaturesFromSalesforceProductsRequest)(nil), "packageproto.SetFeaturesFromSalesforceProductsRequest")
	proto.RegisterType((*SetFeaturesFromSalesforceProductsResponse)(nil), "packageproto.SetFeaturesFromSalesforceProductsResponse")
}

func init() {
	proto.RegisterFile("protorepo/messages/insys/package.proto", fileDescriptor_b757e2aec27b58d2)
}

var fileDescriptor_b757e2aec27b58d2 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x15, 0x22, 0xa0, 0x35, 0x4c, 0x99, 0x22, 0x16, 0xa2, 0x0e, 0x28, 0x2c, 0xf5, 0xc0,
	0x37, 0xa8, 0x50, 0x25, 0x36, 0xe4, 0x6e, 0x2c, 0xe8, 0x6a, 0x5f, 0x83, 0x85, 0x93, 0x33, 0x3e,
	0xa7, 0x50, 0x24, 0x26, 0xbe, 0x38, 0xc2, 0x24, 0x45, 0x62, 0xa1, 0xd3, 0xe9, 0x77, 0xef, 0xfe,
	0x3c, 0x3d, 0x71, 0xe5, 0x03, 0x45, 0x0a, 0xe8, 0x49, 0xb6, 0xc8, 0x0c, 0x0d, 0xb2, 0xb4, 0x1d,
	0xef, 0x58, 0x7a, 0xd0, 0xcf, 0xd0, 0xe0, 0x3c, 0x0d, 0x14, 0xe7, 0x03, 0x26, 0x9a, 0x7d, 0x66,
	0xa2, 0x5e, 0x61, 0x5c, 0x22, 0xc4, 0x3e, 0x20, 0x2f, 0x03, 0xb5, 0x2b, 0x70, 0xc8, 0x1b, 0x0a,
	0x1a, 0xef, 0x03, 0x99, 0x5e, 0x47, 0x56, 0xf8, 0xd2, 0x23, 0xc7, 0xe2, 0x52, 0x9c, 0x39, 0xd2,
	0x10, 0x2d, 0x75, 0x8f, 0xd6, 0x94, 0x59, 0x95, 0xd5, 0x53, 0x25, 0xc6, 0xd6, 0x9d, 0x29, 0x2e,
	0xc4, 0xc4, 0x0f, 0x3b, 0xe5, 0x51, 0x95, 0xd7, 0x53, 0xb5, 0xe7, 0x6f, 0x6d, 0x8b, 0x21, 0x5a,
	0x0d, 0xae, 0xcc, 0xab, 0xac, 0x3e, 0x56, 0x7b, 0x9e, 0x7d, 0x88, 0xeb, 0x03, 0x4c, 0xb0, 0xa7,
	0x8e, 0xb1, 0x28, 0xc5, 0x29, 0xf7, 0x5a, 0x23, 0x73, 0x72, 0x30, 0x51, 0x23, 0xfe, 0xf7, 0x7e,
	0x33, 0xdc, 0x2f, 0xf3, 0x1f, 0x6d, 0xe4, 0xc5, 0xed, 0xc3, 0xe2, 0x15, 0x61, 0x8b, 0x0e, 0xd6,
	0xf3, 0xb7, 0xdd, 0xbb, 0x6c, 0xa9, 0xa3, 0x00, 0xd6, 0x49, 0x7e, 0x82, 0x80, 0x46, 0xfe, 0x26,
	0x6b, 0x2c, 0x47, 0xd9, 0xfc, 0x4d, 0x38, 0xe9, 0xeb, 0x93, 0x54, 0x6e, 0xbe, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x2d, 0x72, 0xfa, 0x89, 0x01, 0x00, 0x00,
}