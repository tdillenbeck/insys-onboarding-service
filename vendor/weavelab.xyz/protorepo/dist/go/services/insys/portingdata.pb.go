// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/insys/portingdata.proto

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

// PortingDataServiceClient is the client API for PortingDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortingDataServiceClient interface {
	Create(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataCreateResponse, error)
	Update(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error)
	ByLocationID(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts ...grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error)
}

type portingDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewPortingDataServiceClient(cc *grpc.ClientConn) PortingDataServiceClient {
	return &portingDataServiceClient{cc}
}

func (c *portingDataServiceClient) Create(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataCreateResponse, error) {
	out := new(insysproto.PortingDataCreateResponse)
	err := c.cc.Invoke(ctx, "/portingdataproto.PortingDataService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portingDataServiceClient) Update(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error) {
	out := new(insysproto.PortingDataUpdateResponse)
	err := c.cc.Invoke(ctx, "/portingdataproto.PortingDataService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portingDataServiceClient) ByLocationID(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts ...grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error) {
	out := new(insysproto.PortingDataByLocationIDResponse)
	err := c.cc.Invoke(ctx, "/portingdataproto.PortingDataService/ByLocationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortingDataServiceServer is the server API for PortingDataService service.
type PortingDataServiceServer interface {
	Create(context.Context, *insysproto.PortingDataCreateRequest) (*insysproto.PortingDataCreateResponse, error)
	Update(context.Context, *insysproto.PortingDataUpdateRequest) (*insysproto.PortingDataUpdateResponse, error)
	ByLocationID(context.Context, *insysproto.PortingDataByLocationIDRequest) (*insysproto.PortingDataByLocationIDResponse, error)
}

func RegisterPortingDataServiceServer(s *grpc.Server, srv PortingDataServiceServer) {
	s.RegisterService(&_PortingDataService_serviceDesc, srv)
}

func _PortingDataService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.PortingDataCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortingDataServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portingdataproto.PortingDataService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortingDataServiceServer).Create(ctx, req.(*insysproto.PortingDataCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortingDataService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.PortingDataUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortingDataServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portingdataproto.PortingDataService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortingDataServiceServer).Update(ctx, req.(*insysproto.PortingDataUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortingDataService_ByLocationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.PortingDataByLocationIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortingDataServiceServer).ByLocationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portingdataproto.PortingDataService/ByLocationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortingDataServiceServer).ByLocationID(ctx, req.(*insysproto.PortingDataByLocationIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortingDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portingdataproto.PortingDataService",
	HandlerType: (*PortingDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PortingDataService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PortingDataService_Update_Handler,
		},
		{
			MethodName: "ByLocationID",
			Handler:    _PortingDataService_ByLocationID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/portingdata.proto",
}

func init() {
	proto.RegisterFile("protorepo/services/insys/portingdata.proto", fileDescriptor_portingdata_e6c94c0007710943)
}

var fileDescriptor_portingdata_e6c94c0007710943 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3d, 0x4e, 0xc4, 0x30,
	0x10, 0x46, 0x11, 0x45, 0x0a, 0x8b, 0x02, 0xb9, 0x4c, 0x49, 0x19, 0x84, 0xcd, 0xcf, 0x0d, 0x42,
	0x1a, 0x24, 0x0a, 0x04, 0xa2, 0xa1, 0x9b, 0x24, 0xa3, 0xc8, 0x12, 0xd8, 0xc6, 0x33, 0x09, 0x84,
	0x7b, 0x72, 0x1f, 0x44, 0xac, 0xec, 0x5a, 0x91, 0x76, 0x37, 0xdb, 0x7e, 0x7a, 0x6f, 0x5e, 0x31,
	0xa2, 0xf0, 0xc1, 0xb1, 0x0b, 0xe8, 0x9d, 0x26, 0x0c, 0x83, 0x69, 0x90, 0xb4, 0xb1, 0x34, 0x92,
	0xf6, 0x2e, 0xb0, 0xb1, 0x5d, 0x0b, 0x0c, 0x6a, 0x82, 0xe4, 0x79, 0x32, 0x4d, 0x4b, 0x9e, 0xd8,
	0x1f, 0x48, 0x04, 0xdd, 0x6e, 0xfb, 0xf6, 0xf7, 0x54, 0xc8, 0xa7, 0xb8, 0x56, 0xc0, 0xf0, 0x12,
	0x6b, 0xb2, 0x11, 0xd9, 0x7d, 0x40, 0x60, 0x94, 0x85, 0x5a, 0xde, 0x57, 0x09, 0x1f, 0xa1, 0x67,
	0xfc, 0xec, 0x91, 0x38, 0xbf, 0x5c, 0xc5, 0x92, 0x77, 0x96, 0xf0, 0xe2, 0xe4, 0x3f, 0xf2, 0xea,
	0xdb, 0xc3, 0x91, 0x08, 0xad, 0x8b, 0xcc, 0xec, 0x26, 0xd2, 0x8b, 0xb3, 0x72, 0x7c, 0x74, 0x0d,
	0xb0, 0x71, 0xf6, 0xa1, 0x92, 0xd7, 0x7b, 0xf5, 0x14, 0x9d, 0x83, 0x37, 0x47, 0x18, 0x73, 0xb6,
	0xd4, 0x6f, 0x57, 0x5f, 0x08, 0x03, 0xbe, 0x43, 0xad, 0xbe, 0xc7, 0x1f, 0xbd, 0x7d, 0x49, 0x6b,
	0x88, 0x75, 0xb7, 0x7c, 0x6c, 0x9d, 0x4d, 0xc4, 0xdd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0xd4, 0x89, 0x94, 0xfb, 0x01, 0x00, 0x00,
}
