// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acl.proto

package wiggum

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Permission int32

const (
	Permission_Unknown                         Permission = 0
	Permission_FeatureFlagRead                 Permission = 10
	Permission_FeatureFlagWrite                Permission = 11
	Permission_FeatureFlagCreate               Permission = 12
	Permission_UserDelete                      Permission = 20
	Permission_UserWrite                       Permission = 21
	Permission_SyncAppInstallAdvanced          Permission = 203
	Permission_AnalyticsAdvanced               Permission = 204
	Permission_CallRecordingSettingWrite       Permission = 400
	Permission_CallRecordingRead               Permission = 401
	Permission_CallFowardingSettingWrite       Permission = 402
	Permission_OfficeHoursSettingWrite         Permission = 403
	Permission_VoicemailOverrideWrite          Permission = 404
	Permission_DenyMobileAccess                Permission = 405
	Permission_WriteLocations                  Permission = 500
	Permission_AutomatedNotificationQueueWrite Permission = 600
	Permission_PaymentsRefunds                 Permission = 700
	Permission_PaymentsExpressDashboard        Permission = 701
	Permission_PaymentsExports                 Permission = 702
)

var Permission_name = map[int32]string{
	0:   "Unknown",
	10:  "FeatureFlagRead",
	11:  "FeatureFlagWrite",
	12:  "FeatureFlagCreate",
	20:  "UserDelete",
	21:  "UserWrite",
	203: "SyncAppInstallAdvanced",
	204: "AnalyticsAdvanced",
	400: "CallRecordingSettingWrite",
	401: "CallRecordingRead",
	402: "CallFowardingSettingWrite",
	403: "OfficeHoursSettingWrite",
	404: "VoicemailOverrideWrite",
	405: "DenyMobileAccess",
	500: "WriteLocations",
	600: "AutomatedNotificationQueueWrite",
	700: "PaymentsRefunds",
	701: "PaymentsExpressDashboard",
	702: "PaymentsExports",
}
var Permission_value = map[string]int32{
	"Unknown":                         0,
	"FeatureFlagRead":                 10,
	"FeatureFlagWrite":                11,
	"FeatureFlagCreate":               12,
	"UserDelete":                      20,
	"UserWrite":                       21,
	"SyncAppInstallAdvanced":          203,
	"AnalyticsAdvanced":               204,
	"CallRecordingSettingWrite":       400,
	"CallRecordingRead":               401,
	"CallFowardingSettingWrite":       402,
	"OfficeHoursSettingWrite":         403,
	"VoicemailOverrideWrite":          404,
	"DenyMobileAccess":                405,
	"WriteLocations":                  500,
	"AutomatedNotificationQueueWrite": 600,
	"PaymentsRefunds":                 700,
	"PaymentsExpressDashboard":        701,
	"PaymentsExports":                 702,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}
func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_acl_e77cd620683b7ffc, []int{0}
}

func init() {
	proto.RegisterEnum("aclproto.Permission", Permission_name, Permission_value)
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor_acl_e77cd620683b7ffc) }

var fileDescriptor_acl_e77cd620683b7ffc = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x5d, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0x59, 0x99, 0x94, 0x76, 0x0a, 0xad, 0x3b, 0x4d, 0x0a, 0x88, 0x2f, 0x09, 0xf1, 0xc4,
	0x03, 0x2f, 0x9c, 0x60, 0xd5, 0x34, 0x02, 0x09, 0x68, 0x49, 0x55, 0x78, 0x9e, 0xd8, 0xb3, 0xc5,
	0xc2, 0x6b, 0xaf, 0x6c, 0x6f, 0x93, 0xbd, 0x05, 0x9f, 0xb7, 0x01, 0x0e, 0x10, 0x38, 0x00, 0x37,
	0xe0, 0x02, 0x1c, 0x00, 0x65, 0x93, 0xa0, 0x44, 0x7d, 0xb3, 0xff, 0xbf, 0xf9, 0x8d, 0xfe, 0xd2,
	0xc0, 0x16, 0x29, 0xfb, 0xa4, 0x0a, 0x3e, 0x79, 0xdc, 0x24, 0x65, 0xdb, 0xd7, 0xe3, 0x3f, 0x02,
	0xe0, 0x84, 0x43, 0x69, 0x62, 0x34, 0xde, 0xe1, 0x36, 0x5c, 0x3b, 0x73, 0xef, 0x9d, 0x1f, 0x3b,
	0x79, 0x05, 0xf7, 0x61, 0x77, 0xc0, 0x94, 0xea, 0xc0, 0x03, 0x4b, 0xe7, 0x43, 0x26, 0x2d, 0x01,
	0xbb, 0x20, 0x57, 0xc2, 0xb7, 0xc1, 0x24, 0x96, 0xdb, 0xd8, 0x83, 0xbd, 0x95, 0xf4, 0x30, 0x30,
	0x25, 0x96, 0xd7, 0x71, 0x07, 0xe0, 0x2c, 0x72, 0xe8, 0xb3, 0xe5, 0xc4, 0xb2, 0x8b, 0x37, 0x60,
	0x6b, 0xf6, 0x9f, 0x5b, 0x3d, 0xbc, 0x03, 0x07, 0xa7, 0x8d, 0x53, 0x79, 0x55, 0x3d, 0x77, 0x31,
	0x91, 0xb5, 0xb9, 0xbe, 0x20, 0xa7, 0x58, 0xcb, 0x9f, 0x19, 0x1e, 0xc0, 0x5e, 0xee, 0xc8, 0x36,
	0xc9, 0xa8, 0xf8, 0x3f, 0xff, 0x95, 0xe1, 0x7d, 0xb8, 0x7d, 0x48, 0xd6, 0x0e, 0x59, 0xf9, 0xa0,
	0x8d, 0x3b, 0x3f, 0xe5, 0x94, 0x8c, 0x5b, 0x34, 0xf9, 0x20, 0x66, 0xde, 0x1a, 0x6f, 0x7b, 0x7f,
	0x14, 0x4b, 0x6f, 0xe0, 0xc7, 0x74, 0xd9, 0xfb, 0x24, 0xf0, 0x2e, 0xdc, 0x3c, 0x2e, 0x0a, 0xa3,
	0xf8, 0x99, 0xaf, 0x43, 0x5c, 0xa3, 0x9f, 0xc5, 0xac, 0xea, 0x1b, 0x6f, 0x14, 0x97, 0x64, 0xec,
	0xf1, 0x05, 0x87, 0x60, 0x34, 0xcf, 0xe1, 0x17, 0x81, 0x3d, 0x90, 0x7d, 0x76, 0xcd, 0x4b, 0x3f,
	0x32, 0x96, 0x73, 0xa5, 0x38, 0x46, 0xf9, 0x55, 0xe0, 0x3e, 0xec, 0xb4, 0x23, 0x2f, 0xbc, 0xa2,
	0x64, 0xbc, 0x8b, 0xf2, 0xaf, 0xc0, 0x47, 0xf0, 0x20, 0xaf, 0x93, 0x2f, 0x29, 0xb1, 0x7e, 0xe5,
	0x93, 0x29, 0xcc, 0x1c, 0xbe, 0xae, 0xb9, 0x5e, 0x6c, 0xfc, 0x7d, 0x15, 0xbb, 0xb0, 0x7b, 0x42,
	0x4d, 0xc9, 0x2e, 0xc5, 0x21, 0x17, 0xb5, 0xd3, 0x51, 0x7e, 0xeb, 0xe0, 0x3d, 0xb8, 0xb5, 0x4c,
	0x8f, 0x26, 0x55, 0xe0, 0x18, 0xfb, 0x14, 0xdf, 0x8d, 0x3c, 0x05, 0x2d, 0xbf, 0x77, 0x56, 0xa5,
	0xa3, 0x49, 0xe5, 0x43, 0x8a, 0xf2, 0x47, 0xe7, 0xe1, 0xc6, 0xe6, 0x34, 0x93, 0xd3, 0x6c, 0xb4,
	0xd1, 0x1e, 0xfc, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x5a, 0xf5, 0x0b, 0x07, 0x02,
	0x00, 0x00,
}
