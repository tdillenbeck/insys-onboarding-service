// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/messages/platform/notifications.proto

package platformproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	platformenums "weavelab.xyz/monorail/shared/protorepo/dist/go/enums/platformenums"
	sharedproto "weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
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

//Proto used to send a notification via NotificationAPI
type WeaveNotification struct {
	LocationID           *sharedproto.UUID                     `protobuf:"bytes,1,opt,name=LocationID,proto3" json:"LocationID,omitempty"`
	UserID               *sharedproto.UUID                     `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Type                 platformenums.NotificationType        `protobuf:"varint,3,opt,name=Type,proto3,enum=platformenums.NotificationType" json:"Type,omitempty"`
	TypeID               *sharedproto.UUID                     `protobuf:"bytes,4,opt,name=TypeID,proto3" json:"TypeID,omitempty"`
	Alert                *Alert                                `protobuf:"bytes,5,opt,name=Alert,proto3" json:"Alert,omitempty"`
	RetryMax             int64                                 `protobuf:"varint,6,opt,name=RetryMax,proto3" json:"RetryMax,omitempty"`
	RetryCount           int64                                 `protobuf:"varint,7,opt,name=RetryCount,proto3" json:"RetryCount,omitempty"`
	Destination          platformenums.NotificationDestination `protobuf:"varint,8,opt,name=Destination,proto3,enum=platformenums.NotificationDestination" json:"Destination,omitempty"`
	DestinationID        string                                `protobuf:"bytes,9,opt,name=DestinationID,proto3" json:"DestinationID,omitempty"`
	Event                string                                `protobuf:"bytes,10,opt,name=Event,proto3" json:"Event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *WeaveNotification) Reset()         { *m = WeaveNotification{} }
func (m *WeaveNotification) String() string { return proto.CompactTextString(m) }
func (*WeaveNotification) ProtoMessage()    {}
func (*WeaveNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9c7f35ba37178f6, []int{0}
}

func (m *WeaveNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeaveNotification.Unmarshal(m, b)
}
func (m *WeaveNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeaveNotification.Marshal(b, m, deterministic)
}
func (m *WeaveNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeaveNotification.Merge(m, src)
}
func (m *WeaveNotification) XXX_Size() int {
	return xxx_messageInfo_WeaveNotification.Size(m)
}
func (m *WeaveNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_WeaveNotification.DiscardUnknown(m)
}

var xxx_messageInfo_WeaveNotification proto.InternalMessageInfo

func (m *WeaveNotification) GetLocationID() *sharedproto.UUID {
	if m != nil {
		return m.LocationID
	}
	return nil
}

func (m *WeaveNotification) GetUserID() *sharedproto.UUID {
	if m != nil {
		return m.UserID
	}
	return nil
}

func (m *WeaveNotification) GetType() platformenums.NotificationType {
	if m != nil {
		return m.Type
	}
	return platformenums.NotificationType_Unknown
}

func (m *WeaveNotification) GetTypeID() *sharedproto.UUID {
	if m != nil {
		return m.TypeID
	}
	return nil
}

func (m *WeaveNotification) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *WeaveNotification) GetRetryMax() int64 {
	if m != nil {
		return m.RetryMax
	}
	return 0
}

func (m *WeaveNotification) GetRetryCount() int64 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *WeaveNotification) GetDestination() platformenums.NotificationDestination {
	if m != nil {
		return m.Destination
	}
	return platformenums.NotificationDestination_All
}

func (m *WeaveNotification) GetDestinationID() string {
	if m != nil {
		return m.DestinationID
	}
	return ""
}

func (m *WeaveNotification) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func init() {
	proto.RegisterType((*WeaveNotification)(nil), "platformproto.WeaveNotification")
}

func init() {
	proto.RegisterFile("protorepo/messages/platform/notifications.proto", fileDescriptor_b9c7f35ba37178f6)
}

var fileDescriptor_b9c7f35ba37178f6 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xe9, 0xbb, 0x3f, 0xef, 0x76, 0xf6, 0xee, 0x05, 0xc3, 0x2e, 0x42, 0x2f, 0xb4, 0xc8,
	0xd0, 0x22, 0x23, 0x85, 0xed, 0x13, 0xa8, 0x1d, 0x5a, 0x50, 0x2f, 0x8a, 0x43, 0xf0, 0x2e, 0x73,
	0xd9, 0x2c, 0xb4, 0x4d, 0x49, 0xd2, 0xb9, 0xf9, 0xa9, 0xfc, 0x88, 0xd2, 0xc4, 0xb9, 0x0c, 0xca,
	0xae, 0x92, 0x93, 0xe7, 0x77, 0x4e, 0xce, 0x93, 0x1c, 0x08, 0x0a, 0xc1, 0x15, 0x17, 0xac, 0xe0,
	0x41, 0xc6, 0xa4, 0xa4, 0x2b, 0x26, 0x83, 0x22, 0xa5, 0x6a, 0xc9, 0x45, 0x16, 0xe4, 0x5c, 0x25,
	0xcb, 0xe4, 0x8d, 0xaa, 0x84, 0xe7, 0x92, 0x68, 0x12, 0xf5, 0x77, 0xaa, 0x0e, 0xdd, 0x61, 0x4d,
	0xbe, 0x7c, 0xa7, 0x82, 0x2d, 0x82, 0xb2, 0x4c, 0x16, 0x26, 0xc9, 0x1d, 0xed, 0x29, 0x96, 0x97,
	0xd9, 0xf1, 0x2b, 0xdc, 0xcb, 0x63, 0x3d, 0xd1, 0x94, 0x09, 0x65, 0xc0, 0xf3, 0xaf, 0x06, 0x9c,
	0xbc, 0x30, 0xba, 0x66, 0x4f, 0x56, 0x15, 0x34, 0x02, 0x78, 0xe0, 0x66, 0x1f, 0x85, 0xd8, 0xf1,
	0x1c, 0xbf, 0x37, 0xfe, 0x47, 0x4c, 0x53, 0x64, 0x36, 0x8b, 0xc2, 0xd8, 0xd2, 0xd1, 0x10, 0xda,
	0x33, 0xc9, 0x44, 0x14, 0xe2, 0x3f, 0x35, 0xe4, 0x8f, 0x86, 0x26, 0xd0, 0x7c, 0xde, 0x16, 0x0c,
	0x37, 0x3c, 0xc7, 0xff, 0x3f, 0x3e, 0x23, 0xbb, 0x76, 0xb4, 0x1b, 0x62, 0x5f, 0x5f, 0x61, 0xb1,
	0x86, 0xab, 0xd2, 0xd5, 0x1a, 0x85, 0xb8, 0x59, 0x57, 0xda, 0x68, 0xe8, 0x0a, 0x5a, 0xd7, 0x95,
	0x27, 0xdc, 0xd2, 0xd0, 0x80, 0x1c, 0x3c, 0x30, 0xd1, 0x5a, 0x6c, 0x10, 0xe4, 0x42, 0x27, 0x66,
	0x4a, 0x6c, 0x1f, 0xe9, 0x06, 0xb7, 0x3d, 0xc7, 0x6f, 0xc4, 0xbf, 0x31, 0x3a, 0x05, 0xd0, 0xfb,
	0x5b, 0x5e, 0xe6, 0x0a, 0xff, 0xd5, 0xaa, 0x75, 0x82, 0xee, 0xa1, 0x17, 0x32, 0xa9, 0x92, 0x5c,
	0xb7, 0x89, 0x3b, 0xda, 0xc9, 0xc5, 0x11, 0x27, 0x16, 0x1d, 0xdb, 0xa9, 0x68, 0x08, 0x7d, 0x2b,
	0x8c, 0x42, 0xdc, 0xf5, 0x1c, 0xbf, 0x1b, 0x1f, 0x1e, 0xa2, 0x01, 0xb4, 0xa6, 0x6b, 0x96, 0x2b,
	0x0c, 0x5a, 0x35, 0xc1, 0xcd, 0xdd, 0xeb, 0xf4, 0xa3, 0xfa, 0xb1, 0x94, 0xce, 0xc9, 0x66, 0xfb,
	0x19, 0x64, 0x3c, 0xe7, 0x82, 0x26, 0xe9, 0x6e, 0x68, 0xf6, 0x5f, 0xbf, 0x48, 0xa4, 0x0a, 0x56,
	0x35, 0x23, 0xa0, 0x91, 0x79, 0x5b, 0x2f, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xe8,
	0x6d, 0xe9, 0xc1, 0x02, 0x00, 0x00,
}