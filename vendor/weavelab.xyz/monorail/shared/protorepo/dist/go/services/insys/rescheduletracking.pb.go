// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/insys/rescheduletracking.proto

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
	proto.RegisterFile("protorepo/services/insys/rescheduletracking.proto", fileDescriptor_3a110eb68234e860)
}

var fileDescriptor_3a110eb68234e860 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0x4a, 0x2d, 0xc8, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0xcf,
	0xcc, 0x2b, 0xae, 0x2c, 0xd6, 0x2f, 0x4a, 0x2d, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x2d,
	0x29, 0x4a, 0x4c, 0xce, 0xce, 0xcc, 0x4b, 0xd7, 0x03, 0xab, 0x15, 0x12, 0xc7, 0x94, 0x01, 0x4b,
	0x48, 0x21, 0x99, 0x95, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x4e, 0xd0, 0x2c, 0xa3, 0xf9, 0x8c, 0x5c,
	0xe2, 0x41, 0x70, 0xc9, 0x10, 0xa8, 0xa4, 0x6b, 0x59, 0x6a, 0x5e, 0x89, 0x50, 0x33, 0x23, 0x97,
	0x58, 0x50, 0x6a, 0x62, 0x0a, 0xa6, 0xbc, 0x90, 0x91, 0x1e, 0x0e, 0x37, 0xe8, 0x61, 0x2a, 0x0e,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x32, 0x26, 0x49, 0x4f, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x12, 0x83, 0x93, 0x7d, 0x94, 0x6d, 0x79, 0x6a, 0x62, 0x59, 0x6a, 0x4e, 0x62, 0x92, 0x5e,
	0x45, 0x65, 0x95, 0x7e, 0x6e, 0x7e, 0x5e, 0x7e, 0x51, 0x62, 0x66, 0x8e, 0x7e, 0x71, 0x46, 0x62,
	0x51, 0x6a, 0x8a, 0x3e, 0xc2, 0xcf, 0x29, 0x99, 0xc5, 0x25, 0xfa, 0xe9, 0xe8, 0xe1, 0x98, 0xc4,
	0x06, 0x56, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x64, 0x81, 0x41, 0x6a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RescheduleTrackingEventClient is the client API for RescheduleTrackingEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RescheduleTrackingEventClient interface {
	ReadRescheduleTracking(ctx context.Context, in *insysproto.RescheduleTrackingRequest, opts ...grpc.CallOption) (*insysproto.RescheduleTrackingResponse, error)
}

type rescheduleTrackingEventClient struct {
	cc *grpc.ClientConn
}

func NewRescheduleTrackingEventClient(cc *grpc.ClientConn) RescheduleTrackingEventClient {
	return &rescheduleTrackingEventClient{cc}
}

func (c *rescheduleTrackingEventClient) ReadRescheduleTracking(ctx context.Context, in *insysproto.RescheduleTrackingRequest, opts ...grpc.CallOption) (*insysproto.RescheduleTrackingResponse, error) {
	out := new(insysproto.RescheduleTrackingResponse)
	err := c.cc.Invoke(ctx, "/rescheduletrackingproto.RescheduleTrackingEvent/ReadRescheduleTracking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RescheduleTrackingEventServer is the server API for RescheduleTrackingEvent service.
type RescheduleTrackingEventServer interface {
	ReadRescheduleTracking(context.Context, *insysproto.RescheduleTrackingRequest) (*insysproto.RescheduleTrackingResponse, error)
}

func RegisterRescheduleTrackingEventServer(s *grpc.Server, srv RescheduleTrackingEventServer) {
	s.RegisterService(&_RescheduleTrackingEvent_serviceDesc, srv)
}

func _RescheduleTrackingEvent_ReadRescheduleTracking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.RescheduleTrackingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RescheduleTrackingEventServer).ReadRescheduleTracking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rescheduletrackingproto.RescheduleTrackingEvent/ReadRescheduleTracking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RescheduleTrackingEventServer).ReadRescheduleTracking(ctx, req.(*insysproto.RescheduleTrackingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RescheduleTrackingEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rescheduletrackingproto.RescheduleTrackingEvent",
	HandlerType: (*RescheduleTrackingEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadRescheduleTracking",
			Handler:    _RescheduleTrackingEvent_ReadRescheduleTracking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/rescheduletracking.proto",
}
