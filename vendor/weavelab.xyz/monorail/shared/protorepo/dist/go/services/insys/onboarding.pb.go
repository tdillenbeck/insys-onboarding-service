// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/insys/onboarding.proto

package insys

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	proto.RegisterFile("protorepo/services/insys/onboarding.proto", fileDescriptor_49d6caaee6e07e7c)
}

var fileDescriptor_49d6caaee6e07e7c = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x4e, 0xff, 0xbf, 0x8a, 0x60, 0xd4, 0x0b, 0x9a, 0x05, 0x08, 0x03, 0x12, 0x44, 0xa2, 0x2a,
	0x54, 0xb2, 0xa5, 0x70, 0xd9, 0x00, 0x42, 0x34, 0x49, 0xd5, 0xa0, 0x4a, 0x81, 0xa4, 0xdd, 0x54,
	0x5d, 0x30, 0x8e, 0x4f, 0x1c, 0x0b, 0x67, 0xc6, 0xcc, 0x4c, 0x42, 0xcd, 0x82, 0x1d, 0x2b, 0x5e,
	0x80, 0x67, 0x61, 0xc3, 0x0b, 0x20, 0x9e, 0x09, 0xc5, 0x53, 0x4f, 0xec, 0x38, 0x8e, 0x13, 0x8b,
	0x65, 0xce, 0xf9, 0x2e, 0x67, 0xce, 0x25, 0x32, 0x7a, 0x14, 0x70, 0x26, 0x19, 0x87, 0x80, 0x59,
	0x02, 0xf8, 0xc4, 0xeb, 0x83, 0xb0, 0x3c, 0x2a, 0x42, 0x61, 0x31, 0x6a, 0x33, 0xc2, 0x1d, 0x8f,
	0xba, 0x66, 0x84, 0xc1, 0xbb, 0xb3, 0x48, 0x14, 0x30, 0xee, 0xb8, 0x8c, 0xb9, 0x3e, 0x58, 0xd1,
	0x2f, 0x7b, 0x3c, 0xb0, 0x60, 0x14, 0xc8, 0x50, 0xa1, 0x8d, 0x84, 0xf0, 0x08, 0x84, 0x20, 0x6e,
	0xae, 0xb0, 0xb1, 0x5f, 0x04, 0x05, 0xae, 0x90, 0xf5, 0x9f, 0x9b, 0x08, 0x75, 0x34, 0x1d, 0x7f,
	0x45, 0x77, 0x1b, 0x1c, 0x88, 0x84, 0x53, 0x22, 0x3e, 0xb6, 0xa9, 0x90, 0x84, 0xf6, 0x41, 0x1c,
	0x71, 0x36, 0x9a, 0x06, 0x04, 0x7e, 0x6a, 0xce, 0x95, 0x6c, 0x2e, 0x83, 0x77, 0xe1, 0xd3, 0x18,
	0x84, 0x34, 0xf6, 0x32, 0xac, 0x14, 0xbe, 0x0b, 0x22, 0x60, 0x54, 0x40, 0xad, 0x82, 0xdf, 0xa3,
	0x6b, 0x0d, 0x22, 0xc1, 0x65, 0x3c, 0xc4, 0xf7, 0xb3, 0x5e, 0x57, 0xa9, 0x58, 0xf7, 0xc1, 0x12,
	0x84, 0x96, 0xfc, 0x80, 0xb6, 0x53, 0x6e, 0xf8, 0x61, 0x51, 0x35, 0xeb, 0x16, 0xcd, 0x10, 0x3e,
	0x0b, 0x9c, 0xb9, 0x2e, 0xe0, 0xc7, 0x19, 0x7e, 0x16, 0x14, 0x7b, 0x1d, 0xac, 0x84, 0xd5, 0x86,
	0xdf, 0x36, 0xd0, 0xbd, 0x2c, 0xa0, 0x75, 0x19, 0xf8, 0x84, 0x12, 0xe9, 0x31, 0x8a, 0x9f, 0xad,
	0x20, 0x98, 0xc0, 0x97, 0xab, 0xa3, 0xfe, 0xfb, 0x3f, 0x74, 0xbd, 0x13, 0x2f, 0x14, 0x3e, 0x46,
	0x3b, 0x6a, 0x19, 0x3a, 0x5c, 0xb1, 0xf0, 0x6d, 0x53, 0xaf, 0x9b, 0x52, 0xd3, 0x60, 0x23, 0x3f,
	0x55, 0xab, 0xe0, 0xb7, 0xa8, 0xda, 0x04, 0x1f, 0x24, 0xe0, 0xbd, 0x79, 0x98, 0x8a, 0x6b, 0x70,
	0x5c, 0xf8, 0x4d, 0x53, 0x5d, 0x8e, 0x19, 0x5f, 0x8e, 0xd9, 0x9a, 0x5e, 0x4e, 0xad, 0x82, 0x09,
	0xda, 0x39, 0xf1, 0x84, 0xd4, 0x8c, 0xc4, 0xfc, 0x63, 0xcd, 0x74, 0x3e, 0x33, 0xff, 0x3c, 0x98,
	0x1e, 0xc7, 0x11, 0xda, 0xea, 0x02, 0x71, 0x0e, 0xc3, 0x33, 0x01, 0xbc, 0xdd, 0x2c, 0xfb, 0xec,
	0xfa, 0x9f, 0x0d, 0x84, 0x67, 0x06, 0x27, 0xac, 0xaf, 0x66, 0x79, 0x9e, 0xe9, 0x6b, 0x2d, 0x57,
	0x45, 0xb3, 0x8c, 0x15, 0x30, 0xb5, 0x0a, 0xbe, 0x40, 0x37, 0x54, 0xe9, 0x71, 0xac, 0xdd, 0xfc,
	0x77, 0xea, 0xf5, 0x5f, 0xff, 0xa3, 0xdd, 0x63, 0x42, 0x1d, 0x36, 0x18, 0xf4, 0x28, 0x09, 0xc4,
	0x90, 0x49, 0x2c, 0x32, 0xaf, 0xc9, 0xee, 0xea, 0x1c, 0x27, 0x8d, 0x8f, 0xe7, 0xb3, 0x5f, 0x44,
	0x4b, 0x4c, 0x68, 0x8c, 0x0c, 0xf5, 0xcc, 0x6c, 0x99, 0xed, 0x26, 0x3e, 0x28, 0x56, 0x22, 0x4e,
	0x19, 0xdb, 0x0b, 0x84, 0x7a, 0x63, 0x7b, 0xe4, 0xc9, 0x46, 0xef, 0xcd, 0xe9, 0xac, 0xaf, 0x9a,
	0x39, 0x4b, 0x96, 0x51, 0xb7, 0xd1, 0xb6, 0x12, 0xb8, 0x82, 0x2c, 0xf8, 0x63, 0x4b, 0xe5, 0x4b,
	0x78, 0xd4, 0x7f, 0x6c, 0xa2, 0x5b, 0x8d, 0xa1, 0xe7, 0x7b, 0xef, 0xbc, 0x00, 0x78, 0xaf, 0x3f,
	0x04, 0x67, 0xec, 0x43, 0x6b, 0x02, 0x54, 0xe2, 0x10, 0x55, 0x1b, 0xd3, 0x3f, 0x04, 0x7f, 0xc1,
	0x04, 0x55, 0x22, 0x87, 0x19, 0x17, 0xf2, 0x7c, 0x5d, 0x9a, 0x7e, 0xfa, 0xd4, 0x3a, 0x5a, 0x8a,
	0x45, 0xd6, 0x51, 0x62, 0x7d, 0xeb, 0xe5, 0x34, 0x6d, 0xfd, 0x7d, 0x03, 0x6d, 0xa5, 0xce, 0xe5,
	0x45, 0x46, 0x2a, 0x99, 0x2e, 0xa8, 0xe3, 0x65, 0x39, 0x72, 0xb2, 0x11, 0xb9, 0x57, 0xa4, 0x12,
	0x6b, 0x37, 0xa2, 0x80, 0x16, 0x5b, 0x1f, 0xbe, 0x3e, 0x7f, 0xf5, 0x19, 0xc8, 0x04, 0x7c, 0x62,
	0x9b, 0x97, 0xe1, 0x17, 0x6b, 0xc4, 0x28, 0xe3, 0xc4, 0xf3, 0x2d, 0x31, 0x24, 0x1c, 0x1c, 0x6b,
	0xf6, 0x09, 0xe2, 0x78, 0x42, 0x5a, 0xee, 0xfc, 0xe7, 0x90, 0x5d, 0x8d, 0x10, 0x4f, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x77, 0x9a, 0x92, 0x77, 0x31, 0x09, 0x00, 0x00,
}

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

// OnboardingServer is the server API for Onboarding service.
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

// OnboarderClient is the client API for Onboarder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnboarderClient interface {
	CreateOrUpdate(ctx context.Context, in *insysproto.Onboarder, opts ...grpc.CallOption) (*insysproto.Onboarder, error)
	Delete(ctx context.Context, in *insysproto.DeleteOnboarderRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListOnboarders(ctx context.Context, in *insysproto.ListOnboardersRequest, opts ...grpc.CallOption) (*insysproto.ListOnboardersResponse, error)
	ReadByUserID(ctx context.Context, in *insysproto.Onboarder, opts ...grpc.CallOption) (*insysproto.Onboarder, error)
}

type onboarderClient struct {
	cc *grpc.ClientConn
}

func NewOnboarderClient(cc *grpc.ClientConn) OnboarderClient {
	return &onboarderClient{cc}
}

func (c *onboarderClient) CreateOrUpdate(ctx context.Context, in *insysproto.Onboarder, opts ...grpc.CallOption) (*insysproto.Onboarder, error) {
	out := new(insysproto.Onboarder)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarder/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboarderClient) Delete(ctx context.Context, in *insysproto.DeleteOnboarderRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarder/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboarderClient) ListOnboarders(ctx context.Context, in *insysproto.ListOnboardersRequest, opts ...grpc.CallOption) (*insysproto.ListOnboardersResponse, error) {
	out := new(insysproto.ListOnboardersResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarder/ListOnboarders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboarderClient) ReadByUserID(ctx context.Context, in *insysproto.Onboarder, opts ...grpc.CallOption) (*insysproto.Onboarder, error) {
	out := new(insysproto.Onboarder)
	err := c.cc.Invoke(ctx, "/onboardingproto.Onboarder/ReadByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnboarderServer is the server API for Onboarder service.
type OnboarderServer interface {
	CreateOrUpdate(context.Context, *insysproto.Onboarder) (*insysproto.Onboarder, error)
	Delete(context.Context, *insysproto.DeleteOnboarderRequest) (*empty.Empty, error)
	ListOnboarders(context.Context, *insysproto.ListOnboardersRequest) (*insysproto.ListOnboardersResponse, error)
	ReadByUserID(context.Context, *insysproto.Onboarder) (*insysproto.Onboarder, error)
}

func RegisterOnboarderServer(s *grpc.Server, srv OnboarderServer) {
	s.RegisterService(&_Onboarder_serviceDesc, srv)
}

func _Onboarder_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.Onboarder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboarderServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarder/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboarderServer).CreateOrUpdate(ctx, req.(*insysproto.Onboarder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarder_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.DeleteOnboarderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboarderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarder/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboarderServer).Delete(ctx, req.(*insysproto.DeleteOnboarderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarder_ListOnboarders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.ListOnboardersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboarderServer).ListOnboarders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarder/ListOnboarders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboarderServer).ListOnboarders(ctx, req.(*insysproto.ListOnboardersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Onboarder_ReadByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.Onboarder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboarderServer).ReadByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.Onboarder/ReadByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboarderServer).ReadByUserID(ctx, req.(*insysproto.Onboarder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Onboarder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingproto.Onboarder",
	HandlerType: (*OnboarderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _Onboarder_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Onboarder_Delete_Handler,
		},
		{
			MethodName: "ListOnboarders",
			Handler:    _Onboarder_ListOnboarders_Handler,
		},
		{
			MethodName: "ReadByUserID",
			Handler:    _Onboarder_ReadByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/onboarding.proto",
}

// OnboardersLocationClient is the client API for OnboardersLocation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnboardersLocationClient interface {
	CreateOrUpdate(ctx context.Context, in *insysproto.OnboardersLocation, opts ...grpc.CallOption) (*insysproto.OnboardersLocation, error)
	ReadByLocationID(ctx context.Context, in *insysproto.OnboardersLocation, opts ...grpc.CallOption) (*insysproto.OnboardersLocation, error)
}

type onboardersLocationClient struct {
	cc *grpc.ClientConn
}

func NewOnboardersLocationClient(cc *grpc.ClientConn) OnboardersLocationClient {
	return &onboardersLocationClient{cc}
}

func (c *onboardersLocationClient) CreateOrUpdate(ctx context.Context, in *insysproto.OnboardersLocation, opts ...grpc.CallOption) (*insysproto.OnboardersLocation, error) {
	out := new(insysproto.OnboardersLocation)
	err := c.cc.Invoke(ctx, "/onboardingproto.OnboardersLocation/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onboardersLocationClient) ReadByLocationID(ctx context.Context, in *insysproto.OnboardersLocation, opts ...grpc.CallOption) (*insysproto.OnboardersLocation, error) {
	out := new(insysproto.OnboardersLocation)
	err := c.cc.Invoke(ctx, "/onboardingproto.OnboardersLocation/ReadByLocationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnboardersLocationServer is the server API for OnboardersLocation service.
type OnboardersLocationServer interface {
	CreateOrUpdate(context.Context, *insysproto.OnboardersLocation) (*insysproto.OnboardersLocation, error)
	ReadByLocationID(context.Context, *insysproto.OnboardersLocation) (*insysproto.OnboardersLocation, error)
}

func RegisterOnboardersLocationServer(s *grpc.Server, srv OnboardersLocationServer) {
	s.RegisterService(&_OnboardersLocation_serviceDesc, srv)
}

func _OnboardersLocation_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.OnboardersLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardersLocationServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.OnboardersLocation/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardersLocationServer).CreateOrUpdate(ctx, req.(*insysproto.OnboardersLocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnboardersLocation_ReadByLocationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.OnboardersLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardersLocationServer).ReadByLocationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.OnboardersLocation/ReadByLocationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardersLocationServer).ReadByLocationID(ctx, req.(*insysproto.OnboardersLocation))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnboardersLocation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingproto.OnboardersLocation",
	HandlerType: (*OnboardersLocationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _OnboardersLocation_CreateOrUpdate_Handler,
		},
		{
			MethodName: "ReadByLocationID",
			Handler:    _OnboardersLocation_ReadByLocationID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/onboarding.proto",
}

// HandoffSnapshotClient is the client API for HandoffSnapshot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandoffSnapshotClient interface {
	CreateOrUpdate(ctx context.Context, in *insysproto.HandoffSnapshotCreateOrUpdateRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error)
	ReadByOnboardersLocationID(ctx context.Context, in *insysproto.HandoffSnapshotReadRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error)
	SubmitCSAT(ctx context.Context, in *insysproto.SubmitCSATRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error)
	SubmitHandoff(ctx context.Context, in *insysproto.SubmitHandoffRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error)
}

type handoffSnapshotClient struct {
	cc *grpc.ClientConn
}

func NewHandoffSnapshotClient(cc *grpc.ClientConn) HandoffSnapshotClient {
	return &handoffSnapshotClient{cc}
}

func (c *handoffSnapshotClient) CreateOrUpdate(ctx context.Context, in *insysproto.HandoffSnapshotCreateOrUpdateRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error) {
	out := new(insysproto.HandoffSnapshotResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.HandoffSnapshot/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handoffSnapshotClient) ReadByOnboardersLocationID(ctx context.Context, in *insysproto.HandoffSnapshotReadRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error) {
	out := new(insysproto.HandoffSnapshotResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.HandoffSnapshot/ReadByOnboardersLocationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handoffSnapshotClient) SubmitCSAT(ctx context.Context, in *insysproto.SubmitCSATRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error) {
	out := new(insysproto.HandoffSnapshotResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.HandoffSnapshot/SubmitCSAT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handoffSnapshotClient) SubmitHandoff(ctx context.Context, in *insysproto.SubmitHandoffRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotResponse, error) {
	out := new(insysproto.HandoffSnapshotResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.HandoffSnapshot/SubmitHandoff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandoffSnapshotServer is the server API for HandoffSnapshot service.
type HandoffSnapshotServer interface {
	CreateOrUpdate(context.Context, *insysproto.HandoffSnapshotCreateOrUpdateRequest) (*insysproto.HandoffSnapshotResponse, error)
	ReadByOnboardersLocationID(context.Context, *insysproto.HandoffSnapshotReadRequest) (*insysproto.HandoffSnapshotResponse, error)
	SubmitCSAT(context.Context, *insysproto.SubmitCSATRequest) (*insysproto.HandoffSnapshotResponse, error)
	SubmitHandoff(context.Context, *insysproto.SubmitHandoffRequest) (*insysproto.HandoffSnapshotResponse, error)
}

func RegisterHandoffSnapshotServer(s *grpc.Server, srv HandoffSnapshotServer) {
	s.RegisterService(&_HandoffSnapshot_serviceDesc, srv)
}

func _HandoffSnapshot_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.HandoffSnapshotCreateOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandoffSnapshotServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.HandoffSnapshot/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandoffSnapshotServer).CreateOrUpdate(ctx, req.(*insysproto.HandoffSnapshotCreateOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandoffSnapshot_ReadByOnboardersLocationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.HandoffSnapshotReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandoffSnapshotServer).ReadByOnboardersLocationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.HandoffSnapshot/ReadByOnboardersLocationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandoffSnapshotServer).ReadByOnboardersLocationID(ctx, req.(*insysproto.HandoffSnapshotReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandoffSnapshot_SubmitCSAT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.SubmitCSATRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandoffSnapshotServer).SubmitCSAT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.HandoffSnapshot/SubmitCSAT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandoffSnapshotServer).SubmitCSAT(ctx, req.(*insysproto.SubmitCSATRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandoffSnapshot_SubmitHandoff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.SubmitHandoffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandoffSnapshotServer).SubmitHandoff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.HandoffSnapshot/SubmitHandoff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandoffSnapshotServer).SubmitHandoff(ctx, req.(*insysproto.SubmitHandoffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandoffSnapshot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingproto.HandoffSnapshot",
	HandlerType: (*HandoffSnapshotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _HandoffSnapshot_CreateOrUpdate_Handler,
		},
		{
			MethodName: "ReadByOnboardersLocationID",
			Handler:    _HandoffSnapshot_ReadByOnboardersLocationID_Handler,
		},
		{
			MethodName: "SubmitCSAT",
			Handler:    _HandoffSnapshot_SubmitCSAT_Handler,
		},
		{
			MethodName: "SubmitHandoff",
			Handler:    _HandoffSnapshot_SubmitHandoff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/onboarding.proto",
}

// ChiliPiperScheduleEventClient is the client API for ChiliPiperScheduleEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChiliPiperScheduleEventClient interface {
	Cancel(ctx context.Context, in *insysproto.CancelChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.CancelChiliPiperScheduleEventResponse, error)
	Create(ctx context.Context, in *insysproto.CreateChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.CreateChiliPiperScheduleEventResponse, error)
	ByLocationID(ctx context.Context, in *insysproto.ByLocationIDChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.ByLocationIDChiliPiperScheduleEventResponse, error)
	Update(ctx context.Context, in *insysproto.UpdateChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.UpdateChiliPiperScheduleEventResponse, error)
}

type chiliPiperScheduleEventClient struct {
	cc *grpc.ClientConn
}

func NewChiliPiperScheduleEventClient(cc *grpc.ClientConn) ChiliPiperScheduleEventClient {
	return &chiliPiperScheduleEventClient{cc}
}

func (c *chiliPiperScheduleEventClient) Cancel(ctx context.Context, in *insysproto.CancelChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.CancelChiliPiperScheduleEventResponse, error) {
	out := new(insysproto.CancelChiliPiperScheduleEventResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.ChiliPiperScheduleEvent/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chiliPiperScheduleEventClient) Create(ctx context.Context, in *insysproto.CreateChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.CreateChiliPiperScheduleEventResponse, error) {
	out := new(insysproto.CreateChiliPiperScheduleEventResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.ChiliPiperScheduleEvent/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chiliPiperScheduleEventClient) ByLocationID(ctx context.Context, in *insysproto.ByLocationIDChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.ByLocationIDChiliPiperScheduleEventResponse, error) {
	out := new(insysproto.ByLocationIDChiliPiperScheduleEventResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.ChiliPiperScheduleEvent/ByLocationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chiliPiperScheduleEventClient) Update(ctx context.Context, in *insysproto.UpdateChiliPiperScheduleEventRequest, opts ...grpc.CallOption) (*insysproto.UpdateChiliPiperScheduleEventResponse, error) {
	out := new(insysproto.UpdateChiliPiperScheduleEventResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.ChiliPiperScheduleEvent/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChiliPiperScheduleEventServer is the server API for ChiliPiperScheduleEvent service.
type ChiliPiperScheduleEventServer interface {
	Cancel(context.Context, *insysproto.CancelChiliPiperScheduleEventRequest) (*insysproto.CancelChiliPiperScheduleEventResponse, error)
	Create(context.Context, *insysproto.CreateChiliPiperScheduleEventRequest) (*insysproto.CreateChiliPiperScheduleEventResponse, error)
	ByLocationID(context.Context, *insysproto.ByLocationIDChiliPiperScheduleEventRequest) (*insysproto.ByLocationIDChiliPiperScheduleEventResponse, error)
	Update(context.Context, *insysproto.UpdateChiliPiperScheduleEventRequest) (*insysproto.UpdateChiliPiperScheduleEventResponse, error)
}

func RegisterChiliPiperScheduleEventServer(s *grpc.Server, srv ChiliPiperScheduleEventServer) {
	s.RegisterService(&_ChiliPiperScheduleEvent_serviceDesc, srv)
}

func _ChiliPiperScheduleEvent_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.CancelChiliPiperScheduleEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiliPiperScheduleEventServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.ChiliPiperScheduleEvent/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiliPiperScheduleEventServer).Cancel(ctx, req.(*insysproto.CancelChiliPiperScheduleEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChiliPiperScheduleEvent_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.CreateChiliPiperScheduleEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiliPiperScheduleEventServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.ChiliPiperScheduleEvent/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiliPiperScheduleEventServer).Create(ctx, req.(*insysproto.CreateChiliPiperScheduleEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChiliPiperScheduleEvent_ByLocationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.ByLocationIDChiliPiperScheduleEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiliPiperScheduleEventServer).ByLocationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.ChiliPiperScheduleEvent/ByLocationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiliPiperScheduleEventServer).ByLocationID(ctx, req.(*insysproto.ByLocationIDChiliPiperScheduleEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChiliPiperScheduleEvent_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(insysproto.UpdateChiliPiperScheduleEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiliPiperScheduleEventServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboardingproto.ChiliPiperScheduleEvent/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiliPiperScheduleEventServer).Update(ctx, req.(*insysproto.UpdateChiliPiperScheduleEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChiliPiperScheduleEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingproto.ChiliPiperScheduleEvent",
	HandlerType: (*ChiliPiperScheduleEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Cancel",
			Handler:    _ChiliPiperScheduleEvent_Cancel_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ChiliPiperScheduleEvent_Create_Handler,
		},
		{
			MethodName: "ByLocationID",
			Handler:    _ChiliPiperScheduleEvent_ByLocationID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ChiliPiperScheduleEvent_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/insys/onboarding.proto",
}
