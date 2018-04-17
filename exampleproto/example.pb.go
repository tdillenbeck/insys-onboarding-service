// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

/*
Package exampleproto is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	ExampleRequestMessage
	ExampleResponseMessage
*/
package exampleproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wgrpcprotouuid "weavelab.xyz/wlib/wgrpc/wgrpcproto/wgrpcprotouuid"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExampleRequestMessage struct {
	SomeID *wgrpcprotouuid.UUID `protobuf:"bytes,1,opt,name=SomeID" json:"SomeID,omitempty"`
}

func (m *ExampleRequestMessage) Reset()                    { *m = ExampleRequestMessage{} }
func (m *ExampleRequestMessage) String() string            { return proto.CompactTextString(m) }
func (*ExampleRequestMessage) ProtoMessage()               {}
func (*ExampleRequestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExampleRequestMessage) GetSomeID() *wgrpcprotouuid.UUID {
	if m != nil {
		return m.SomeID
	}
	return nil
}

type ExampleResponseMessage struct {
	Message string `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
}

func (m *ExampleResponseMessage) Reset()                    { *m = ExampleResponseMessage{} }
func (m *ExampleResponseMessage) String() string            { return proto.CompactTextString(m) }
func (*ExampleResponseMessage) ProtoMessage()               {}
func (*ExampleResponseMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExampleResponseMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ExampleRequestMessage)(nil), "exampleproto.ExampleRequestMessage")
	proto.RegisterType((*ExampleResponseMessage)(nil), "exampleproto.ExampleResponseMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ExampleAPI service

type ExampleAPIClient interface {
	ExampleRequest(ctx context.Context, in *ExampleRequestMessage, opts ...grpc.CallOption) (*ExampleResponseMessage, error)
}

type exampleAPIClient struct {
	cc *grpc.ClientConn
}

func NewExampleAPIClient(cc *grpc.ClientConn) ExampleAPIClient {
	return &exampleAPIClient{cc}
}

func (c *exampleAPIClient) ExampleRequest(ctx context.Context, in *ExampleRequestMessage, opts ...grpc.CallOption) (*ExampleResponseMessage, error) {
	out := new(ExampleResponseMessage)
	err := grpc.Invoke(ctx, "/exampleproto.ExampleAPI/ExampleRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExampleAPI service

type ExampleAPIServer interface {
	ExampleRequest(context.Context, *ExampleRequestMessage) (*ExampleResponseMessage, error)
}

func RegisterExampleAPIServer(s *grpc.Server, srv ExampleAPIServer) {
	s.RegisterService(&_ExampleAPI_serviceDesc, srv)
}

func _ExampleAPI_ExampleRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServer).ExampleRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exampleproto.ExampleAPI/ExampleRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServer).ExampleRequest(ctx, req.(*ExampleRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exampleproto.ExampleAPI",
	HandlerType: (*ExampleAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExampleRequest",
			Handler:    _ExampleAPI_ExampleRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0x72, 0xc1, 0x3c, 0x29,
	0x9b, 0xf2, 0xd4, 0xc4, 0xb2, 0xd4, 0x9c, 0xc4, 0x24, 0xbd, 0x8a, 0xca, 0x2a, 0xfd, 0xf2, 0x9c,
	0xcc, 0x24, 0xfd, 0xf2, 0xf4, 0xa2, 0x82, 0x64, 0x08, 0x09, 0x56, 0x83, 0xc4, 0x2c, 0x2d, 0xcd,
	0x4c, 0xd1, 0x07, 0x11, 0x10, 0xb3, 0x94, 0x5c, 0xb9, 0x44, 0x5d, 0x21, 0xa6, 0x05, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0xf8, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0xe9, 0x70, 0xb1, 0x05,
	0xe7, 0xe7, 0xa6, 0x7a, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xe8, 0xa1, 0x1a,
	0xa2, 0x17, 0x1a, 0xea, 0xe9, 0x12, 0x04, 0x55, 0xa3, 0x64, 0xc4, 0x25, 0x06, 0x37, 0xa6, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0x15, 0x66, 0x8e, 0x04, 0x17, 0x3b, 0x94, 0x09, 0x36, 0x88, 0x33, 0x08,
	0xc6, 0x35, 0xca, 0xe6, 0xe2, 0x82, 0xea, 0x71, 0x0c, 0xf0, 0x14, 0x8a, 0xe5, 0xe2, 0x43, 0x75,
	0x88, 0x90, 0xb2, 0x1e, 0xb2, 0x3f, 0xf5, 0xb0, 0x3a, 0x53, 0x4a, 0x05, 0x87, 0x22, 0x14, 0x47,
	0x28, 0x31, 0x24, 0xb1, 0x81, 0xe5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xf4, 0xbc,
	0x1d, 0x4b, 0x01, 0x00, 0x00,
}
