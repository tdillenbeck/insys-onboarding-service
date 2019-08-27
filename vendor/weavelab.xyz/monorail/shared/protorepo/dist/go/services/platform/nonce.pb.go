// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/platform/nonce.proto

package platform

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	math "math"
	platformproto "weavelab.xyz/monorail/shared/protorepo/dist/go/messages/platformproto"
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

func init() {
	proto.RegisterFile("protorepo/services/platform/nonce.proto", fileDescriptor_beaf2b2357382704)
}

var fileDescriptor_beaf2b2357382704 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x63, 0xa1, 0xc8, 0x7a, 0x55, 0x0a, 0x91, 0xd5, 0x2a, 0x08, 0x76, 0xb3, 0xa8, 0x60,
	0x63, 0xa3, 0xf1, 0xae, 0xb8, 0x46, 0x44, 0xb8, 0xc6, 0x6e, 0x72, 0x99, 0x8b, 0x0b, 0x49, 0x66,
	0xdd, 0xdd, 0x9c, 0x17, 0x7f, 0x96, 0xbf, 0x50, 0xb2, 0x4b, 0x50, 0x8f, 0xbb, 0x6e, 0x86, 0xf9,
	0xde, 0x1b, 0xde, 0x13, 0x57, 0xc6, 0xb2, 0x67, 0x4b, 0x86, 0x95, 0x23, 0xbb, 0xd6, 0x4b, 0x72,
	0xca, 0xd4, 0xe8, 0x57, 0x6c, 0x1b, 0xd5, 0x72, 0xbb, 0x24, 0x08, 0x44, 0x2a, 0xc2, 0x12, 0x66,
	0x79, 0x5e, 0x31, 0x57, 0x35, 0xa9, 0xb0, 0x15, 0xdd, 0x4a, 0x51, 0x63, 0x7c, 0x1f, 0x41, 0xf9,
	0xc7, 0xb1, 0x21, 0xe7, 0xb0, 0xda, 0xe3, 0x28, 0x2f, 0x77, 0x80, 0xee, 0x1d, 0x2d, 0x95, 0xaa,
	0xeb, 0x74, 0x19, 0xa9, 0x9b, 0xef, 0x03, 0x71, 0xfc, 0x3c, 0xa8, 0x1e, 0x5f, 0xe6, 0xe9, 0xb5,
	0x38, 0x0c, 0x73, 0x3a, 0x81, 0x48, 0xc2, 0x62, 0x31, 0x9f, 0xca, 0x33, 0x18, 0x1f, 0x04, 0x0d,
	0x04, 0x66, 0x8a, 0x1e, 0xb3, 0x24, 0xbd, 0x17, 0x27, 0x4f, 0x96, 0xd0, 0x53, 0x14, 0x5e, 0x6c,
	0xa1, 0xf1, 0xf6, 0x4a, 0x1f, 0x1d, 0x39, 0x2f, 0xff, 0xd9, 0x66, 0x49, 0x7a, 0x27, 0x26, 0xb3,
	0x8d, 0xd1, 0x96, 0xf2, 0x7e, 0xb0, 0xdb, 0x7a, 0x7b, 0x0a, 0xb1, 0x07, 0x18, 0x7b, 0x80, 0xd9,
	0xd0, 0x43, 0x96, 0xe4, 0xf9, 0xdb, 0xc3, 0x27, 0xe1, 0x9a, 0x6a, 0x2c, 0x60, 0xd3, 0x7f, 0xa9,
	0x86, 0x5b, 0xb6, 0xa8, 0xeb, 0x31, 0xdf, 0x6f, 0xf2, 0x52, 0x3b, 0xaf, 0xaa, 0x1d, 0xe5, 0x17,
	0x47, 0x01, 0xba, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x29, 0x40, 0x34, 0xa2, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NonceAPIClient is the client API for NonceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NonceAPIClient interface {
	Nonce(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*platformproto.NonceData, error)
	CreateNonce(ctx context.Context, in *platformproto.CreateRequest, opts ...grpc.CallOption) (*sharedproto.UUID, error)
	ExpireByData(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*empty.Empty, error)
}

type nonceAPIClient struct {
	cc *grpc.ClientConn
}

func NewNonceAPIClient(cc *grpc.ClientConn) NonceAPIClient {
	return &nonceAPIClient{cc}
}

func (c *nonceAPIClient) Nonce(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*platformproto.NonceData, error) {
	out := new(platformproto.NonceData)
	err := c.cc.Invoke(ctx, "/nonceproto.NonceAPI/Nonce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nonceAPIClient) CreateNonce(ctx context.Context, in *platformproto.CreateRequest, opts ...grpc.CallOption) (*sharedproto.UUID, error) {
	out := new(sharedproto.UUID)
	err := c.cc.Invoke(ctx, "/nonceproto.NonceAPI/CreateNonce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nonceAPIClient) ExpireByData(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nonceproto.NonceAPI/ExpireByData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NonceAPIServer is the server API for NonceAPI service.
type NonceAPIServer interface {
	Nonce(context.Context, *sharedproto.UUID) (*platformproto.NonceData, error)
	CreateNonce(context.Context, *platformproto.CreateRequest) (*sharedproto.UUID, error)
	ExpireByData(context.Context, *sharedproto.UUID) (*empty.Empty, error)
}

func RegisterNonceAPIServer(s *grpc.Server, srv NonceAPIServer) {
	s.RegisterService(&_NonceAPI_serviceDesc, srv)
}

func _NonceAPI_Nonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sharedproto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonceAPIServer).Nonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nonceproto.NonceAPI/Nonce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonceAPIServer).Nonce(ctx, req.(*sharedproto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _NonceAPI_CreateNonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platformproto.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonceAPIServer).CreateNonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nonceproto.NonceAPI/CreateNonce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonceAPIServer).CreateNonce(ctx, req.(*platformproto.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NonceAPI_ExpireByData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sharedproto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonceAPIServer).ExpireByData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nonceproto.NonceAPI/ExpireByData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonceAPIServer).ExpireByData(ctx, req.(*sharedproto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

var _NonceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nonceproto.NonceAPI",
	HandlerType: (*NonceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nonce",
			Handler:    _NonceAPI_Nonce_Handler,
		},
		{
			MethodName: "CreateNonce",
			Handler:    _NonceAPI_CreateNonce_Handler,
		},
		{
			MethodName: "ExpireByData",
			Handler:    _NonceAPI_ExpireByData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/platform/nonce.proto",
}
