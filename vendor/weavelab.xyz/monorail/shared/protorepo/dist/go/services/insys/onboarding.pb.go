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
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x6f, 0x65, 0x29, 0x3a, 0xac, 0xbb, 0x32, 0x07, 0xc5, 0xa8, 0xa0, 0x05, 0x17, 0xff, 0x40,
	0x02, 0xf5, 0xcf, 0x45, 0x45, 0xd8, 0xb6, 0xcb, 0x56, 0x16, 0xaa, 0x5d, 0xf7, 0xb2, 0x78, 0x70,
	0xda, 0xbc, 0xa6, 0x83, 0xe9, 0x4c, 0x9c, 0x99, 0xd6, 0x8d, 0x07, 0xc1, 0x83, 0x27, 0xcf, 0x82,
	0x9f, 0xc5, 0xcf, 0x20, 0x7e, 0x26, 0x69, 0xc6, 0x4c, 0x9b, 0xcc, 0xa6, 0x69, 0x8b, 0xc7, 0xbc,
	0xf7, 0xfb, 0x97, 0x79, 0x6f, 0x18, 0x74, 0x3f, 0x12, 0x5c, 0x71, 0x01, 0x11, 0xf7, 0x24, 0x88,
	0x29, 0x1d, 0x80, 0xf4, 0x28, 0x93, 0xb1, 0xf4, 0x38, 0xeb, 0x73, 0x22, 0x7c, 0xca, 0x02, 0x37,
	0xc1, 0xe0, 0xdd, 0x79, 0x25, 0x29, 0x38, 0x37, 0x02, 0xce, 0x83, 0x10, 0xbc, 0xe4, 0xab, 0x3f,
	0x19, 0x7a, 0x30, 0x8e, 0x54, 0xac, 0xd1, 0xce, 0x82, 0xf0, 0x18, 0xa4, 0x24, 0x41, 0xa1, 0xb0,
	0x73, 0xaf, 0x0c, 0x0a, 0x42, 0x23, 0x1b, 0xbf, 0xb6, 0x10, 0xea, 0x1a, 0x3a, 0xfe, 0x82, 0x6e,
	0x36, 0x05, 0x10, 0x05, 0x6f, 0x89, 0xfc, 0xd0, 0x61, 0x52, 0x11, 0x36, 0x00, 0x79, 0x20, 0xf8,
	0x78, 0x56, 0x90, 0xf8, 0xb1, 0x9b, 0x8b, 0xec, 0x2e, 0x83, 0xf7, 0xe0, 0xe3, 0x04, 0xa4, 0x72,
	0xf6, 0x2c, 0x56, 0x06, 0xdf, 0x03, 0x19, 0x71, 0x26, 0xa1, 0x5e, 0xc1, 0x6f, 0xd0, 0xc5, 0x26,
	0x51, 0x10, 0x70, 0x11, 0xe3, 0xdb, 0xb6, 0xd7, 0xbf, 0x56, 0xaa, 0x7b, 0x67, 0x09, 0xc2, 0x48,
	0xbe, 0x47, 0x97, 0x33, 0x6e, 0xf8, 0x6e, 0x59, 0x9a, 0x75, 0x43, 0x73, 0x84, 0x4f, 0x22, 0x3f,
	0x77, 0x0a, 0xf8, 0x81, 0xc5, 0xb7, 0x41, 0xa9, 0xd7, 0xc3, 0x95, 0xb0, 0xc6, 0xf0, 0x5b, 0x15,
	0xdd, 0xb2, 0x01, 0xed, 0xb3, 0x28, 0x24, 0x8c, 0x28, 0xca, 0x19, 0x7e, 0xb2, 0x82, 0xe0, 0x02,
	0x7e, 0xb3, 0x1c, 0x8d, 0xdf, 0x17, 0xd0, 0xa5, 0x6e, 0xba, 0x50, 0xf8, 0x10, 0xed, 0xe8, 0x65,
	0xe8, 0x0a, 0xcd, 0xc2, 0xd7, 0x5d, 0xb3, 0x6e, 0x5a, 0xcd, 0x80, 0x9d, 0xe2, 0x56, 0xbd, 0x82,
	0x5f, 0xa1, 0x5a, 0x0b, 0x42, 0x50, 0x80, 0xf7, 0xf2, 0x30, 0x5d, 0x37, 0xe0, 0x34, 0xf8, 0x55,
	0x57, 0xdf, 0x1c, 0x37, 0xbd, 0x39, 0x6e, 0x7b, 0x76, 0x73, 0xea, 0x15, 0x4c, 0xd0, 0xce, 0x11,
	0x95, 0xca, 0x30, 0x16, 0xe6, 0x9f, 0x6a, 0x66, 0xfb, 0xd6, 0xfc, 0x8b, 0x60, 0x66, 0x1c, 0x07,
	0x68, 0xbb, 0x07, 0xc4, 0xdf, 0x8f, 0x4f, 0x24, 0x88, 0x4e, 0x6b, 0xd3, 0xdf, 0x6e, 0xfc, 0xa9,
	0x22, 0x3c, 0x37, 0x38, 0xe2, 0x03, 0x3d, 0xcb, 0x53, 0xeb, 0x5c, 0xeb, 0x85, 0x2a, 0x86, 0xe5,
	0xac, 0x80, 0xa9, 0x57, 0xf0, 0x3b, 0x74, 0x45, 0x47, 0x4f, 0x6b, 0x9d, 0xd6, 0xff, 0x53, 0x6f,
	0xfc, 0xa8, 0xa2, 0xdd, 0x43, 0xc2, 0x7c, 0x3e, 0x1c, 0x1e, 0x33, 0x12, 0xc9, 0x11, 0x57, 0xf8,
	0x6b, 0xd5, 0xfa, 0x1d, 0x7b, 0x59, 0x73, 0xa4, 0x2c, 0x3e, 0x1d, 0xd0, 0xd3, 0x75, 0x69, 0x66,
	0x6f, 0x7f, 0x6e, 0xa1, 0x6b, 0xcd, 0x11, 0x0d, 0xe9, 0x6b, 0x1a, 0x81, 0x38, 0x1e, 0x8c, 0xc0,
	0x9f, 0x84, 0xd0, 0x9e, 0x02, 0x53, 0x38, 0x46, 0xb5, 0xe6, 0x6c, 0xcd, 0xc3, 0x73, 0x62, 0xe9,
	0x46, 0x01, 0xb3, 0x38, 0x56, 0x09, 0xcd, 0xec, 0xd1, 0xcc, 0x3a, 0x89, 0x7c, 0x9e, 0x75, 0xd2,
	0x58, 0xdf, 0x7a, 0x39, 0xcd, 0x58, 0x7f, 0xaf, 0xa2, 0xed, 0xcc, 0x12, 0x3c, 0xb3, 0xa4, 0x16,
	0xdb, 0x25, 0x39, 0x9e, 0x6f, 0x46, 0x5e, 0x3c, 0x88, 0xc2, 0xd5, 0xd0, 0x8d, 0xb5, 0x0f, 0xa2,
	0x84, 0x96, 0x5a, 0xef, 0xbf, 0x3c, 0x7d, 0xf1, 0x09, 0xc8, 0x14, 0x42, 0xd2, 0x77, 0xcf, 0xe2,
	0xcf, 0xde, 0x98, 0x33, 0x2e, 0x08, 0x0d, 0x3d, 0x39, 0x22, 0x02, 0x7c, 0x6f, 0xfe, 0xb0, 0xfa,
	0x54, 0x2a, 0x2f, 0xc8, 0x3f, 0xf2, 0xfd, 0x5a, 0x82, 0x78, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xcf, 0x3a, 0x5e, 0x80, 0x07, 0x08, 0x00, 0x00,
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
	CreateOrUpdate(ctx context.Context, in *insysproto.HandoffSnapshotCreateOrUpdateRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotCreateOrUpdateResponse, error)
}

type handoffSnapshotClient struct {
	cc *grpc.ClientConn
}

func NewHandoffSnapshotClient(cc *grpc.ClientConn) HandoffSnapshotClient {
	return &handoffSnapshotClient{cc}
}

func (c *handoffSnapshotClient) CreateOrUpdate(ctx context.Context, in *insysproto.HandoffSnapshotCreateOrUpdateRequest, opts ...grpc.CallOption) (*insysproto.HandoffSnapshotCreateOrUpdateResponse, error) {
	out := new(insysproto.HandoffSnapshotCreateOrUpdateResponse)
	err := c.cc.Invoke(ctx, "/onboardingproto.HandoffSnapshot/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandoffSnapshotServer is the server API for HandoffSnapshot service.
type HandoffSnapshotServer interface {
	CreateOrUpdate(context.Context, *insysproto.HandoffSnapshotCreateOrUpdateRequest) (*insysproto.HandoffSnapshotCreateOrUpdateResponse, error)
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

var _HandoffSnapshot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboardingproto.HandoffSnapshot",
	HandlerType: (*HandoffSnapshotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _HandoffSnapshot_CreateOrUpdate_Handler,
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
