// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/insys/portingdata.proto

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
	proto.RegisterFile("protorepo/services/insys/portingdata.proto", fileDescriptor_837113dababed2fd)
}

var fileDescriptor_837113dababed2fd = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x05, 0xa1, 0x87, 0xa0, 0x20, 0x39, 0xf6, 0xe8, 0xb1, 0xc2, 0xc6, 0x3f, 0x67, 0x11,
	0x6a, 0x2f, 0x82, 0x07, 0xb1, 0x78, 0xf1, 0x36, 0xdd, 0x0c, 0x6b, 0xa4, 0xcd, 0xc4, 0xcc, 0x6c,
	0x75, 0xfd, 0x76, 0x7e, 0x33, 0x71, 0x63, 0x6a, 0x28, 0x68, 0xd7, 0xeb, 0xe3, 0xf7, 0xe6, 0xc7,
	0xc0, 0x53, 0x93, 0x10, 0x49, 0x28, 0x62, 0x20, 0xc3, 0x18, 0xd7, 0xae, 0x46, 0x36, 0xce, 0x73,
	0xc7, 0x26, 0x50, 0x14, 0xe7, 0x1b, 0x0b, 0x02, 0x55, 0x0f, 0xe9, 0xa3, 0x22, 0xea, 0x93, 0x71,
	0xd1, 0x5e, 0x21, 0x33, 0x34, 0xbf, 0xb7, 0xcf, 0x3f, 0xf6, 0x95, 0xbe, 0x4b, 0xe9, 0x0c, 0x04,
	0xe6, 0xc9, 0xa6, 0x6b, 0x35, 0xba, 0x8e, 0x08, 0x82, 0x7a, 0x52, 0x6d, 0xdf, 0xaf, 0x0a, 0x3e,
	0x41, 0xf7, 0xf8, 0xd2, 0x22, 0xcb, 0xf8, 0x64, 0x10, 0xcb, 0x81, 0x3c, 0xe3, 0xf1, 0xde, 0x97,
	0xe4, 0x21, 0xd8, 0xdd, 0x92, 0x04, 0x0d, 0x93, 0x64, 0x76, 0x23, 0x69, 0xd5, 0xc1, 0xb4, 0xbb,
	0xa5, 0x1a, 0xc4, 0x91, 0xbf, 0x99, 0xe9, 0xd3, 0x3f, 0xeb, 0x25, 0x9a, 0x85, 0x67, 0xff, 0x68,
	0x6c, 0xb4, 0xcf, 0xea, 0xf0, 0x1b, 0x9a, 0x0b, 0x48, 0xcb, 0x3b, 0x5e, 0x4c, 0xd0, 0xb0, 0x17,
	0x33, 0x9b, 0x5d, 0xd3, 0xab, 0xc7, 0xcb, 0x57, 0x84, 0x35, 0x2e, 0x61, 0x51, 0xbd, 0x75, 0xef,
	0x66, 0x45, 0x9e, 0x22, 0xb8, 0xa5, 0xe1, 0x27, 0x88, 0x68, 0xcd, 0xcf, 0x1c, 0xac, 0x63, 0x31,
	0xcd, 0xf6, 0xa8, 0x16, 0xa3, 0x9e, 0xb8, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xef, 0x21, 0x8c,
	0x5a, 0x77, 0x02, 0x00, 0x00,
}

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
	PortingStatus(ctx context.Context, in *insysproto.PortingDataStatusRequest, opts ...grpc.CallOption) (*insysproto.PortingDataStatusResponse, error)
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

func (c *portingDataServiceClient) PortingStatus(ctx context.Context, in *insysproto.PortingDataStatusRequest, opts ...grpc.CallOption) (*insysproto.PortingDataStatusResponse, error) {
	out := new(insysproto.PortingDataStatusResponse)
	err := c.cc.Invoke(ctx, "/portingdataproto.PortingDataService/PortingStatus", in, out, opts...)
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
	PortingStatus(context.Context, *insysproto.PortingDataStatusRequest) (*insysproto.PortingDataStatusResponse, error)
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

func _PortingDataService_PortingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.PortingDataStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortingDataServiceServer).PortingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portingdataproto.PortingDataService/PortingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortingDataServiceServer).PortingStatus(ctx, req.(*insysproto.PortingDataStatusRequest))
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
		{
			MethodName: "PortingStatus",
			Handler:    _PortingDataService_PortingStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/portingdata.proto",
}