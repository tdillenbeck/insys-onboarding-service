// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/insys/officeinfo.proto

package insys // import "weavelab.xyz/protorepo/dist/go/services/insys"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import insysproto "weavelab.xyz/protorepo/dist/go/messages/insysproto"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OfficeInfoServiceClient is the client API for OfficeInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OfficeInfoServiceClient interface {
	Submit(ctx context.Context, in *insysproto.OfficeInfoSubmitRequest, opts ...grpc.CallOption) (*insysproto.OfficeInfoSubmitResponse, error)
}

type officeInfoServiceClient struct {
	cc *grpc.ClientConn
}

func NewOfficeInfoServiceClient(cc *grpc.ClientConn) OfficeInfoServiceClient {
	return &officeInfoServiceClient{cc}
}

func (c *officeInfoServiceClient) Submit(ctx context.Context, in *insysproto.OfficeInfoSubmitRequest, opts ...grpc.CallOption) (*insysproto.OfficeInfoSubmitResponse, error) {
	out := new(insysproto.OfficeInfoSubmitResponse)
	err := c.cc.Invoke(ctx, "/officeinfoproto.OfficeInfoService/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OfficeInfoService service

type OfficeInfoServiceServer interface {
	Submit(context.Context, *insysproto.OfficeInfoSubmitRequest) (*insysproto.OfficeInfoSubmitResponse, error)
}

func RegisterOfficeInfoServiceServer(s *grpc.Server, srv OfficeInfoServiceServer) {
	s.RegisterService(&_OfficeInfoService_serviceDesc, srv)
}

func _OfficeInfoService_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.OfficeInfoSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeInfoServiceServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/officeinfoproto.OfficeInfoService/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeInfoServiceServer).Submit(ctx, req.(*insysproto.OfficeInfoSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OfficeInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "officeinfoproto.OfficeInfoService",
	HandlerType: (*OfficeInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _OfficeInfoService_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/officeinfo.proto",
}

func init() {
	proto.RegisterFile("protorepo/services/insys/officeinfo.proto", fileDescriptor_officeinfo_a7171f86884aac24)
}

var fileDescriptor_officeinfo_a7171f86884aac24 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0x4a, 0x2d, 0xc8, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0xcf,
	0xcc, 0x2b, 0xae, 0x2c, 0xd6, 0xcf, 0x4f, 0x4b, 0xcb, 0x4c, 0x4e, 0xcd, 0xcc, 0x4b, 0xcb, 0xd7,
	0x03, 0xab, 0x11, 0xe2, 0x47, 0x88, 0x80, 0x05, 0xa4, 0x90, 0xf4, 0xe6, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xe3, 0xd4, 0x6b, 0x54, 0xc2, 0x25, 0xe8, 0x0f, 0x16, 0xf3, 0xcc, 0x4b, 0xcb, 0x0f, 0x86,
	0x58, 0x24, 0x14, 0xcf, 0xc5, 0x16, 0x5c, 0x9a, 0x94, 0x9b, 0x59, 0x22, 0xa4, 0xa1, 0x87, 0x66,
	0xb6, 0x1e, 0x92, 0x6a, 0xb0, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x4d, 0x22,
	0x54, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x31, 0x38, 0xe9, 0x47, 0xe9, 0x96, 0xa7, 0x26,
	0x96, 0xa5, 0xe6, 0x24, 0x26, 0xe9, 0x55, 0x54, 0x56, 0xe9, 0x23, 0xdc, 0x9b, 0x92, 0x59, 0x5c,
	0xa2, 0x9f, 0x8e, 0xee, 0xe7, 0x24, 0x36, 0xb0, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0x66, 0x0d, 0xcc, 0x16, 0x01, 0x00, 0x00,
}