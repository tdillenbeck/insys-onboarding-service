// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/services/platform/usersettings.proto

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
	proto.RegisterFile("protorepo/services/platform/usersettings.proto", fileDescriptor_fbcfb0531c538669)
}

var fileDescriptor_fbcfb0531c538669 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x77, 0xd0, 0x81, 0x61, 0x07, 0x09, 0xb8, 0x43, 0x77, 0x91, 0x21, 0xea, 0x41, 0x12,
	0xd4, 0x3f, 0xa0, 0xdb, 0x44, 0xa6, 0x43, 0x64, 0x5a, 0x04, 0xbd, 0x98, 0xad, 0xdf, 0x6a, 0xa0,
	0x6d, 0x62, 0xbe, 0x74, 0x5a, 0xff, 0xa1, 0xff, 0x4a, 0xda, 0xac, 0x3a, 0xea, 0x86, 0x13, 0x3d,
	0x96, 0x3c, 0xef, 0xd3, 0x7c, 0x6f, 0x12, 0xc2, 0xb4, 0x51, 0x56, 0x19, 0xd0, 0x8a, 0x23, 0x98,
	0xa9, 0x1c, 0x03, 0x72, 0x1d, 0x09, 0x3b, 0x51, 0x26, 0xe6, 0x29, 0x82, 0x41, 0xb0, 0x56, 0x26,
	0x21, 0x3a, 0xd0, 0x6b, 0x85, 0x4a, 0x85, 0x11, 0xf0, 0xe2, 0x6b, 0x94, 0x4e, 0x38, 0xc4, 0xda,
	0x66, 0xb3, 0xc5, 0x39, 0x59, 0x0c, 0x88, 0x22, 0xfc, 0x41, 0xb6, 0xb3, 0x80, 0xc7, 0x27, 0x61,
	0x20, 0xe0, 0x69, 0x2a, 0x03, 0x47, 0x1d, 0xbd, 0xaf, 0x93, 0x96, 0x8f, 0x60, 0xae, 0x94, 0x95,
	0x13, 0x39, 0x16, 0x56, 0xaa, 0xe4, 0x66, 0x26, 0x3a, 0xbd, 0xee, 0xd3, 0x80, 0xac, 0x0d, 0x24,
	0x5a, 0x7a, 0xc8, 0xca, 0x7f, 0x15, 0x39, 0xb6, 0x24, 0x93, 0xb3, 0x43, 0x78, 0x4e, 0x01, 0xad,
	0xb7, 0xb7, 0x5a, 0x04, 0xdb, 0x35, 0x7a, 0x41, 0xea, 0xbe, 0x46, 0x30, 0x96, 0xee, 0xae, 0x16,
	0xf2, 0x9a, 0xcc, 0x75, 0xc5, 0xca, 0xae, 0xd8, 0x59, 0xde, 0x95, 0x73, 0xf5, 0x20, 0x02, 0x0b,
	0xff, 0xe0, 0xba, 0x24, 0x9b, 0xf9, 0x44, 0x9d, 0x6c, 0xa0, 0x5c, 0xa4, 0xdf, 0xa3, 0x0d, 0xe6,
	0x5a, 0x64, 0xbe, 0xdf, 0xef, 0xfd, 0x66, 0xc8, 0x73, 0xd2, 0x70, 0xb2, 0x9c, 0xf9, 0x8b, 0xe8,
	0x8e, 0x6c, 0x74, 0x45, 0x52, 0x2c, 0x66, 0x2b, 0x0f, 0xb9, 0x5d, 0xe1, 0x3e, 0x0d, 0x43, 0x40,
	0xad, 0x12, 0x84, 0x76, 0x8d, 0x3e, 0x90, 0xad, 0x7c, 0x87, 0xf3, 0xf1, 0xdb, 0x4c, 0x03, 0xd2,
	0x25, 0x0d, 0x79, 0xfb, 0x15, 0xe9, 0xb7, 0xe4, 0x9c, 0xfc, 0x91, 0x34, 0xbb, 0x06, 0x84, 0x85,
	0x2a, 0x44, 0x0f, 0xaa, 0x5b, 0x5b, 0x88, 0x95, 0xd7, 0x6a, 0xe9, 0x69, 0x75, 0x3a, 0xf7, 0x27,
	0x2f, 0x20, 0xa6, 0x10, 0x89, 0x11, 0x7b, 0xcd, 0xde, 0x78, 0xac, 0x12, 0x65, 0x84, 0x8c, 0xca,
	0x6b, 0xff, 0xf5, 0x20, 0x02, 0x89, 0x96, 0x87, 0x0b, 0x5e, 0xe5, 0xa8, 0x5e, 0x40, 0xc7, 0x1f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0xc7, 0x24, 0x4c, 0xbb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserNotificationSettingsAPIClient is the client API for UserNotificationSettingsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserNotificationSettingsAPIClient interface {
	List(ctx context.Context, in *platformproto.UserNotificationSettingListRequest, opts ...grpc.CallOption) (*platformproto.UserNotificationSettings, error)
	Upsert(ctx context.Context, in *platformproto.UserNotificationSetting, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *platformproto.UserNotificationSetting, opts ...grpc.CallOption) (*empty.Empty, error)
	ListByLocationID(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*platformproto.UserNotificationSettings, error)
	ListByUserID(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*platformproto.UserNotificationSettings, error)
	CanNotify(ctx context.Context, in *platformproto.UserNotificationSetting, opts ...grpc.CallOption) (*platformproto.CanNotifyResponse, error)
	ListNotificationTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*platformproto.NotificationTypesResponse, error)
	CreateNotificationType(ctx context.Context, in *platformproto.CreateNotificationTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userNotificationSettingsAPIClient struct {
	cc *grpc.ClientConn
}

func NewUserNotificationSettingsAPIClient(cc *grpc.ClientConn) UserNotificationSettingsAPIClient {
	return &userNotificationSettingsAPIClient{cc}
}

func (c *userNotificationSettingsAPIClient) List(ctx context.Context, in *platformproto.UserNotificationSettingListRequest, opts ...grpc.CallOption) (*platformproto.UserNotificationSettings, error) {
	out := new(platformproto.UserNotificationSettings)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNotificationSettingsAPIClient) Upsert(ctx context.Context, in *platformproto.UserNotificationSetting, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNotificationSettingsAPIClient) Delete(ctx context.Context, in *platformproto.UserNotificationSetting, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNotificationSettingsAPIClient) ListByLocationID(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*platformproto.UserNotificationSettings, error) {
	out := new(platformproto.UserNotificationSettings)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/ListByLocationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNotificationSettingsAPIClient) ListByUserID(ctx context.Context, in *sharedproto.UUID, opts ...grpc.CallOption) (*platformproto.UserNotificationSettings, error) {
	out := new(platformproto.UserNotificationSettings)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/ListByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNotificationSettingsAPIClient) CanNotify(ctx context.Context, in *platformproto.UserNotificationSetting, opts ...grpc.CallOption) (*platformproto.CanNotifyResponse, error) {
	out := new(platformproto.CanNotifyResponse)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/CanNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNotificationSettingsAPIClient) ListNotificationTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*platformproto.NotificationTypesResponse, error) {
	out := new(platformproto.NotificationTypesResponse)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/ListNotificationTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNotificationSettingsAPIClient) CreateNotificationType(ctx context.Context, in *platformproto.CreateNotificationTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserNotificationSettingsAPI/CreateNotificationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserNotificationSettingsAPIServer is the server API for UserNotificationSettingsAPI service.
type UserNotificationSettingsAPIServer interface {
	List(context.Context, *platformproto.UserNotificationSettingListRequest) (*platformproto.UserNotificationSettings, error)
	Upsert(context.Context, *platformproto.UserNotificationSetting) (*empty.Empty, error)
	Delete(context.Context, *platformproto.UserNotificationSetting) (*empty.Empty, error)
	ListByLocationID(context.Context, *sharedproto.UUID) (*platformproto.UserNotificationSettings, error)
	ListByUserID(context.Context, *sharedproto.UUID) (*platformproto.UserNotificationSettings, error)
	CanNotify(context.Context, *platformproto.UserNotificationSetting) (*platformproto.CanNotifyResponse, error)
	ListNotificationTypes(context.Context, *empty.Empty) (*platformproto.NotificationTypesResponse, error)
	CreateNotificationType(context.Context, *platformproto.CreateNotificationTypeRequest) (*empty.Empty, error)
}

func RegisterUserNotificationSettingsAPIServer(s *grpc.Server, srv UserNotificationSettingsAPIServer) {
	s.RegisterService(&_UserNotificationSettingsAPI_serviceDesc, srv)
}

func _UserNotificationSettingsAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platformproto.UserNotificationSettingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).List(ctx, req.(*platformproto.UserNotificationSettingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNotificationSettingsAPI_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platformproto.UserNotificationSetting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).Upsert(ctx, req.(*platformproto.UserNotificationSetting))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNotificationSettingsAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platformproto.UserNotificationSetting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).Delete(ctx, req.(*platformproto.UserNotificationSetting))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNotificationSettingsAPI_ListByLocationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sharedproto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).ListByLocationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/ListByLocationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).ListByLocationID(ctx, req.(*sharedproto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNotificationSettingsAPI_ListByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sharedproto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).ListByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/ListByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).ListByUserID(ctx, req.(*sharedproto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNotificationSettingsAPI_CanNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platformproto.UserNotificationSetting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).CanNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/CanNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).CanNotify(ctx, req.(*platformproto.UserNotificationSetting))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNotificationSettingsAPI_ListNotificationTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).ListNotificationTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/ListNotificationTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).ListNotificationTypes(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNotificationSettingsAPI_CreateNotificationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platformproto.CreateNotificationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNotificationSettingsAPIServer).CreateNotificationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserNotificationSettingsAPI/CreateNotificationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNotificationSettingsAPIServer).CreateNotificationType(ctx, req.(*platformproto.CreateNotificationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserNotificationSettingsAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserNotificationSettingsAPI",
	HandlerType: (*UserNotificationSettingsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserNotificationSettingsAPI_List_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _UserNotificationSettingsAPI_Upsert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserNotificationSettingsAPI_Delete_Handler,
		},
		{
			MethodName: "ListByLocationID",
			Handler:    _UserNotificationSettingsAPI_ListByLocationID_Handler,
		},
		{
			MethodName: "ListByUserID",
			Handler:    _UserNotificationSettingsAPI_ListByUserID_Handler,
		},
		{
			MethodName: "CanNotify",
			Handler:    _UserNotificationSettingsAPI_CanNotify_Handler,
		},
		{
			MethodName: "ListNotificationTypes",
			Handler:    _UserNotificationSettingsAPI_ListNotificationTypes_Handler,
		},
		{
			MethodName: "CreateNotificationType",
			Handler:    _UserNotificationSettingsAPI_CreateNotificationType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protorepo/services/platform/usersettings.proto",
}
