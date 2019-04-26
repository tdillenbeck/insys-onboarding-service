// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/insys/netstats.proto

package insys

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
	insysproto "weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
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
	proto.RegisterFile("protorepo/services/insys/netstats.proto", fileDescriptor_f6b77bb8a720739c)
}

var fileDescriptor_f6b77bb8a720739c = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x2d, 0xe4, 0x38, 0xb7, 0xdc, 0x46, 0xb8, 0x42, 0xc5, 0xc6, 0x6e, 0x17, 0xce, 0x5a,
	0x04, 0xaf, 0x14, 0x22, 0x24, 0x8d, 0xd8, 0x4d, 0xcc, 0x23, 0x2e, 0x64, 0x77, 0xe3, 0xce, 0x24,
	0x1a, 0x7f, 0xbd, 0x18, 0x59, 0x94, 0x70, 0x29, 0xe7, 0xcd, 0xf7, 0x3d, 0x98, 0x51, 0x37, 0x7d,
	0x8a, 0x12, 0x13, 0xfa, 0x68, 0x19, 0x69, 0x74, 0xaf, 0x60, 0xeb, 0x02, 0x4f, 0x6c, 0x03, 0x84,
	0x85, 0x84, 0xcd, 0x4c, 0xe8, 0x6d, 0x9e, 0x77, 0xff, 0x14, 0x0f, 0x66, 0x6a, 0x57, 0x94, 0xfd,
	0x93, 0xda, 0x16, 0x90, 0xea, 0x27, 0xd1, 0x07, 0xb5, 0x39, 0x24, 0x90, 0x40, 0x9f, 0x9b, 0x8c,
	0xcd, 0x94, 0x29, 0x72, 0xef, 0xc5, 0x62, 0x51, 0x0d, 0xde, 0x53, 0x9a, 0x4a, 0x70, 0x1f, 0x03,
	0xe3, 0xfa, 0x64, 0xff, 0xac, 0xce, 0x7e, 0x43, 0x07, 0xd6, 0x8f, 0xea, 0xb4, 0x04, 0x35, 0xfa,
	0xf2, 0xa8, 0xe6, 0xc0, 0x25, 0xde, 0x07, 0xb0, 0xec, 0xae, 0xd6, 0x81, 0xdc, 0xfc, 0x70, 0xff,
	0x72, 0xf7, 0x01, 0x1a, 0xd1, 0x51, 0x6d, 0x3e, 0xa7, 0x2f, 0xeb, 0x63, 0x88, 0x89, 0x5c, 0x67,
	0xf9, 0x8d, 0x12, 0x1a, 0xfb, 0x77, 0x72, 0xe3, 0x58, 0x6c, 0xbb, 0xfc, 0x56, 0xbd, 0x99, 0x89,
	0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x02, 0xb9, 0x77, 0x50, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetStatsClient is the client API for NetStats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetStatsClient interface {
	Create(ctx context.Context, in *insysproto.Netstats, opts ...grpc.CallOption) (*insysproto.SummaryResponse, error)
}

type netStatsClient struct {
	cc *grpc.ClientConn
}

func NewNetStatsClient(cc *grpc.ClientConn) NetStatsClient {
	return &netStatsClient{cc}
}

func (c *netStatsClient) Create(ctx context.Context, in *insysproto.Netstats, opts ...grpc.CallOption) (*insysproto.SummaryResponse, error) {
	out := new(insysproto.SummaryResponse)
	err := c.cc.Invoke(ctx, "/netstats.NetStats/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetStatsServer is the server API for NetStats service.
type NetStatsServer interface {
	Create(context.Context, *insysproto.Netstats) (*insysproto.SummaryResponse, error)
}

func RegisterNetStatsServer(s *grpc.Server, srv NetStatsServer) {
	s.RegisterService(&_NetStats_serviceDesc, srv)
}

func _NetStats_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.Netstats)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStatsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstats.NetStats/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStatsServer).Create(ctx, req.(*insysproto.Netstats))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetStats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "netstats.NetStats",
	HandlerType: (*NetStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetStats_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/netstats.proto",
}

// SummariesClient is the client API for Summaries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SummariesClient interface {
	Read(ctx context.Context, in *insysproto.SummariesRequest, opts ...grpc.CallOption) (*insysproto.SummariesResponse, error)
}

type summariesClient struct {
	cc *grpc.ClientConn
}

func NewSummariesClient(cc *grpc.ClientConn) SummariesClient {
	return &summariesClient{cc}
}

func (c *summariesClient) Read(ctx context.Context, in *insysproto.SummariesRequest, opts ...grpc.CallOption) (*insysproto.SummariesResponse, error) {
	out := new(insysproto.SummariesResponse)
	err := c.cc.Invoke(ctx, "/netstats.Summaries/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummariesServer is the server API for Summaries service.
type SummariesServer interface {
	Read(context.Context, *insysproto.SummariesRequest) (*insysproto.SummariesResponse, error)
}

func RegisterSummariesServer(s *grpc.Server, srv SummariesServer) {
	s.RegisterService(&_Summaries_serviceDesc, srv)
}

func _Summaries_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.SummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummariesServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstats.Summaries/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummariesServer).Read(ctx, req.(*insysproto.SummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Summaries_serviceDesc = grpc.ServiceDesc{
	ServiceName: "netstats.Summaries",
	HandlerType: (*SummariesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Summaries_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/netstats.proto",
}