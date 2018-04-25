// Code generated by protoc-gen-go. DO NOT EDIT.
// source: onboarding.proto

/*
Package onboardingservice is a generated protocol buffer package.

It is generated from these files:
	onboarding.proto

It has these top-level messages:
*/
package onboardingservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import onboardingproto "weavelab.xyz/protorepo/dist/go/messages/client/onboardingproto"

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

// Client API for Onboarding service

type OnboardingClient interface {
	Categories(ctx context.Context, in *onboardingproto.CategoriesRequest, opts ...grpc.CallOption) (*onboardingproto.CategoriesResponse, error)
	Tasks(ctx context.Context, in *onboardingproto.TasksRequest, opts ...grpc.CallOption) (*onboardingproto.TasksResponse, error)
	UpdateTask(ctx context.Context, in *onboardingproto.UpdateTaskRequest, opts ...grpc.CallOption) (*onboardingproto.UpdateTaskResponse, error)
}

type onboardingClient struct {
	cc *grpc.ClientConn
}

func NewOnboardingClient(cc *grpc.ClientConn) OnboardingClient {
	return &onboardingClient{cc}
}

func (c *onboardingClient) Categories(ctx context.Context, in *onboardingproto.CategoriesRequest, opts ...grpc.CallOption) (*onboardingproto.CategoriesResponse, error) {
	out := new(onboardingproto.CategoriesResponse)
	err := grpc.Invoke(ctx, "/onboardingproto.Onboarding/Categories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingClient) Tasks(ctx context.Context, in *onboardingproto.TasksRequest, opts ...grpc.CallOption) (*onboardingproto.TasksResponse, error) {
	out := new(onboardingproto.TasksResponse)
	err := grpc.Invoke(ctx, "/onboardingproto.Onboarding/Tasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingClient) UpdateTask(ctx context.Context, in *onboardingproto.UpdateTaskRequest, opts ...grpc.CallOption) (*onboardingproto.UpdateTaskResponse, error) {
	out := new(onboardingproto.UpdateTaskResponse)
	err := grpc.Invoke(ctx, "/onboardingproto.Onboarding/UpdateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Onboarding service

type OnboardingServer interface {
	Categories(context.Context, *onboardingproto.CategoriesRequest) (*onboardingproto.CategoriesResponse, error)
	Tasks(context.Context, *onboardingproto.TasksRequest) (*onboardingproto.TasksResponse, error)
	UpdateTask(context.Context, *onboardingproto.UpdateTaskRequest) (*onboardingproto.UpdateTaskResponse, error)
}

func RegisterOnboardingServer(s *grpc.Server, srv OnboardingServer) {
	s.RegisterService(&_Onboarding_serviceDesc, srv)
}

func _Onboarding_Categories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(onboardingproto.CategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).Categories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/Categories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).Categories(ctx, req.(*onboardingproto.CategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarding_Tasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(onboardingproto.TasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).Tasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/Tasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).Tasks(ctx, req.(*onboardingproto.TasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarding_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(onboardingproto.UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).UpdateTask(ctx, req.(*onboardingproto.UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Onboarding_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingproto.Onboarding",
	HandlerType: (*OnboardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Categories",
			Handler:    _Onboarding_Categories_Handler,
		},
		{
			MethodName: "Tasks",
			Handler:    _Onboarding_Tasks_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Onboarding_UpdateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onboarding.proto",
}

func init() { proto.RegisterFile("onboarding.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcf, 0x4b, 0xca,
	0x4f, 0x2c, 0x4a, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x47, 0x88,
	0x80, 0x05, 0xa4, 0xe4, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x8b, 0xf5, 0x33, 0xf3, 0x8a,
	0x2b, 0x8b, 0xf5, 0xd1, 0x75, 0x18, 0x75, 0x30, 0x71, 0x71, 0xf9, 0xc3, 0x05, 0x85, 0xc2, 0xb9,
	0xb8, 0x9c, 0x13, 0x4b, 0x52, 0xd3, 0xf3, 0x8b, 0x32, 0x53, 0x8b, 0x85, 0x94, 0xf4, 0xd0, 0xcc,
	0xd3, 0x43, 0x48, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x29, 0xe3, 0x55, 0x53, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xc4, 0x20, 0xe4, 0xc1, 0xc5, 0x1a, 0x92, 0x58, 0x9c, 0x5d, 0x2c,
	0x24, 0x8b, 0xa1, 0x1e, 0x2c, 0x0e, 0x33, 0x4e, 0x0e, 0x97, 0x34, 0xdc, 0xa4, 0x70, 0x2e, 0xae,
	0xd0, 0x82, 0x94, 0xc4, 0x92, 0x54, 0x90, 0x04, 0x16, 0x27, 0x22, 0x24, 0x71, 0x3b, 0x11, 0x59,
	0x0d, 0xcc, 0x60, 0x27, 0xeb, 0x28, 0xcb, 0xf2, 0xd4, 0xc4, 0xb2, 0xd4, 0x9c, 0xc4, 0x24, 0xbd,
	0x8a, 0xca, 0x2a, 0x7d, 0xb0, 0xd2, 0xa2, 0xd4, 0x82, 0x7c, 0xfd, 0x94, 0xcc, 0xe2, 0x12, 0xfd,
	0xf4, 0x7c, 0xfd, 0xe2, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xe4, 0x60, 0x84, 0x0a, 0x25, 0xb1,
	0x81, 0x55, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe4, 0x74, 0x34, 0x94, 0x01, 0x00,
	0x00,
}
