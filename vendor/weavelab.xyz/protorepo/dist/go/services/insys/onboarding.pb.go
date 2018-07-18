// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/insys/onboarding.proto

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

// OnboardingClient is the client API for Onboarding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnboardingClient interface {
	CreateTaskInstancesFromTasks(ctx context.Context, in *insysproto.CreateTaskInstancesFromTasksRequest, opts ...grpc.CallOption) (*insysproto.TaskInstancesResponse, error)
	Category(ctx context.Context, in *insysproto.CategoryRequest, opts ...grpc.CallOption) (*insysproto.CategoryResponse, error)
	TaskInstances(ctx context.Context, in *insysproto.TaskInstancesRequest, opts ...grpc.CallOption) (*insysproto.TaskInstancesResponse, error)
	UpdateTaskInstance(ctx context.Context, in *insysproto.UpdateTaskInstanceRequest, opts ...grpc.CallOption) (*insysproto.UpdateTaskInstanceResponse, error)
	UpdateTaskInstanceExplanation(ctx context.Context, in *insysproto.UpdateTaskInstanceExplanationRequest, opts ...grpc.CallOption) (*insysproto.UpdateTaskInstanceResponse, error)
}

type onboardingClient struct {
	cc *grpc.ClientConn
}

func NewOnboardingClient(cc *grpc.ClientConn) OnboardingClient {
	return &onboardingClient{cc}
}

func (c *onboardingClient) CreateTaskInstancesFromTasks(ctx context.Context, in *insysproto.CreateTaskInstancesFromTasksRequest, opts ...grpc.CallOption) (*insysproto.TaskInstancesResponse, error) {
	out := new(insysproto.TaskInstancesResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarding/CreateTaskInstancesFromTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingClient) Category(ctx context.Context, in *insysproto.CategoryRequest, opts ...grpc.CallOption) (*insysproto.CategoryResponse, error) {
	out := new(insysproto.CategoryResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarding/Category", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingClient) TaskInstances(ctx context.Context, in *insysproto.TaskInstancesRequest, opts ...grpc.CallOption) (*insysproto.TaskInstancesResponse, error) {
	out := new(insysproto.TaskInstancesResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarding/TaskInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingClient) UpdateTaskInstance(ctx context.Context, in *insysproto.UpdateTaskInstanceRequest, opts ...grpc.CallOption) (*insysproto.UpdateTaskInstanceResponse, error) {
	out := new(insysproto.UpdateTaskInstanceResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarding/UpdateTaskInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardingClient) UpdateTaskInstanceExplanation(ctx context.Context, in *insysproto.UpdateTaskInstanceExplanationRequest, opts ...grpc.CallOption) (*insysproto.UpdateTaskInstanceResponse, error) {
	out := new(insysproto.UpdateTaskInstanceResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarding/UpdateTaskInstanceExplanation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Onboarding service

type OnboardingServer interface {
	CreateTaskInstancesFromTasks(context.Context, *insysproto.CreateTaskInstancesFromTasksRequest) (*insysproto.TaskInstancesResponse, error)
	Category(context.Context, *insysproto.CategoryRequest) (*insysproto.CategoryResponse, error)
	TaskInstances(context.Context, *insysproto.TaskInstancesRequest) (*insysproto.TaskInstancesResponse, error)
	UpdateTaskInstance(context.Context, *insysproto.UpdateTaskInstanceRequest) (*insysproto.UpdateTaskInstanceResponse, error)
	UpdateTaskInstanceExplanation(context.Context, *insysproto.UpdateTaskInstanceExplanationRequest) (*insysproto.UpdateTaskInstanceResponse, error)
}

func RegisterOnboardingServer(s *grpc.Server, srv OnboardingServer) {
	s.RegisterService(&_Onboarding_serviceDesc, srv)
}

func _Onboarding_CreateTaskInstancesFromTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.CreateTaskInstancesFromTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).CreateTaskInstancesFromTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/CreateTaskInstancesFromTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).CreateTaskInstancesFromTasks(ctx, req.(*insysproto.CreateTaskInstancesFromTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarding_Category_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).Category(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/Category",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).Category(ctx, req.(*insysproto.CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarding_TaskInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.TaskInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).TaskInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/TaskInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).TaskInstances(ctx, req.(*insysproto.TaskInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarding_UpdateTaskInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.UpdateTaskInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).UpdateTaskInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/UpdateTaskInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).UpdateTaskInstance(ctx, req.(*insysproto.UpdateTaskInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarding_UpdateTaskInstanceExplanation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.UpdateTaskInstanceExplanationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServer).UpdateTaskInstanceExplanation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarding/UpdateTaskInstanceExplanation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServer).UpdateTaskInstanceExplanation(ctx, req.(*insysproto.UpdateTaskInstanceExplanationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Onboarding_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingproto.Onboarding",
	HandlerType: (*OnboardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaskInstancesFromTasks",
			Handler:    _Onboarding_CreateTaskInstancesFromTasks_Handler,
		},
		{
			MethodName: "Category",
			Handler:    _Onboarding_Category_Handler,
		},
		{
			MethodName: "TaskInstances",
			Handler:    _Onboarding_TaskInstances_Handler,
		},
		{
			MethodName: "UpdateTaskInstance",
			Handler:    _Onboarding_UpdateTaskInstance_Handler,
		},
		{
			MethodName: "UpdateTaskInstanceExplanation",
			Handler:    _Onboarding_UpdateTaskInstanceExplanation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/onboarding.proto",
}

func init() {
	proto.RegisterFile("protorepo/services/insys/onboarding.proto", fileDescriptor_onboarding_36a319172f8eb948)
}

var fileDescriptor_onboarding_36a319172f8eb948 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0x15, 0x44, 0x24, 0x20, 0xc2, 0x1c, 0x8b, 0x82, 0x0a, 0x0a, 0x2a, 0x6e, 0xc0, 0x8f,
	0x17, 0xb0, 0x28, 0x78, 0x12, 0x8b, 0x5e, 0x3c, 0x39, 0xdb, 0x1d, 0x42, 0xb0, 0xcd, 0xc4, 0x4c,
	0xac, 0x5d, 0x0f, 0xde, 0x7c, 0x29, 0x9f, 0x4e, 0x5c, 0xec, 0xb6, 0xdd, 0x60, 0xad, 0xbd, 0xfe,
	0xf3, 0xfb, 0x7f, 0x10, 0x46, 0x1d, 0xf8, 0xc0, 0x91, 0x03, 0x79, 0xd6, 0x42, 0x61, 0x60, 0xbb,
	0x24, 0xda, 0x3a, 0x29, 0x45, 0xb3, 0xcb, 0x19, 0x43, 0x61, 0x9d, 0xc9, 0x2a, 0x06, 0x36, 0xc6,
	0x4a, 0x25, 0xb4, 0x26, 0xbc, 0x7d, 0x12, 0x41, 0xf3, 0xab, 0xf7, 0xe4, 0x73, 0x45, 0xa9, 0x9b,
	0x5a, 0x84, 0x77, 0xb5, 0xd9, 0x0e, 0x84, 0x91, 0xee, 0x50, 0x9e, 0xae, 0x9d, 0x44, 0x74, 0x5d,
	0x92, 0xab, 0xc0, 0xfd, 0x6f, 0x41, 0xe0, 0x2c, 0x6b, 0x74, 0x65, 0xb3, 0xf0, 0x0e, 0x3d, 0xbf,
	0x90, 0xc4, 0xd6, 0x7e, 0xe2, 0x9a, 0xe2, 0x3b, 0x24, 0x9e, 0x9d, 0xd0, 0xee, 0x12, 0xdc, 0xaa,
	0xb5, 0x36, 0x46, 0x32, 0x1c, 0x4a, 0xd8, 0x4e, 0xbb, 0x7e, 0x9e, 0x46, 0xb9, 0x3b, 0x33, 0x88,
	0x3a, 0xf2, 0x51, 0xad, 0x4f, 0xb5, 0xc1, 0xde, 0x5f, 0x6b, 0xfe, 0x3b, 0x9a, 0x15, 0xdc, 0xfb,
	0xa2, 0xf1, 0x0b, 0x70, 0x98, 0xf8, 0x53, 0x68, 0xd4, 0x75, 0x34, 0x17, 0x5b, 0x17, 0x7e, 0x2c,
	0xab, 0xad, 0x14, 0xb8, 0x1c, 0xfa, 0x1e, 0x3a, 0x8c, 0x96, 0x1d, 0x9c, 0xcf, 0x11, 0x38, 0xc1,
	0x2f, 0xb6, 0xe3, 0x42, 0x3f, 0x1c, 0xbf, 0x12, 0x0e, 0xa8, 0x87, 0x79, 0x36, 0x2c, 0xdf, 0xf4,
	0xf8, 0xec, 0x0a, 0x2b, 0x51, 0x9b, 0xe6, 0xe9, 0xe6, 0xab, 0x15, 0x71, 0xfa, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0xfd, 0x29, 0x03, 0xdd, 0x02, 0x00, 0x00,
}
