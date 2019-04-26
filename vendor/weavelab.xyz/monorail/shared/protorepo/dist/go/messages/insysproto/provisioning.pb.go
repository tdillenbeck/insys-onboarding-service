// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/messages/insys/provisioning.proto

package insysproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type InitialProvisionRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	OfficeEmail          string   `protobuf:"bytes,3,opt,name=office_email,json=officeEmail,proto3" json:"office_email,omitempty"`
	TimeZone             string   `protobuf:"bytes,4,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	VerticalId           int32    `protobuf:"varint,5,opt,name=vertical_id,json=verticalId,proto3" json:"vertical_id,omitempty"`
	CallerNumber         string   `protobuf:"bytes,6,opt,name=caller_number,json=callerNumber,proto3" json:"caller_number,omitempty"`
	PrimaryDataCenter    string   `protobuf:"bytes,7,opt,name=primary_data_center,json=primaryDataCenter,proto3" json:"primary_data_center,omitempty"`
	Street1              string   `protobuf:"bytes,8,opt,name=street1,proto3" json:"street1,omitempty"`
	Street2              string   `protobuf:"bytes,9,opt,name=street2,proto3" json:"street2,omitempty"`
	City                 string   `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,11,opt,name=state,proto3" json:"state,omitempty"`
	Zip                  string   `protobuf:"bytes,12,opt,name=zip,proto3" json:"zip,omitempty"`
	Country              string   `protobuf:"bytes,13,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitialProvisionRequest) Reset()         { *m = InitialProvisionRequest{} }
func (m *InitialProvisionRequest) String() string { return proto.CompactTextString(m) }
func (*InitialProvisionRequest) ProtoMessage()    {}
func (*InitialProvisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11524db7586d182a, []int{0}
}

func (m *InitialProvisionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialProvisionRequest.Unmarshal(m, b)
}
func (m *InitialProvisionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialProvisionRequest.Marshal(b, m, deterministic)
}
func (m *InitialProvisionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialProvisionRequest.Merge(m, src)
}
func (m *InitialProvisionRequest) XXX_Size() int {
	return xxx_messageInfo_InitialProvisionRequest.Size(m)
}
func (m *InitialProvisionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialProvisionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitialProvisionRequest proto.InternalMessageInfo

func (m *InitialProvisionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InitialProvisionRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *InitialProvisionRequest) GetOfficeEmail() string {
	if m != nil {
		return m.OfficeEmail
	}
	return ""
}

func (m *InitialProvisionRequest) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *InitialProvisionRequest) GetVerticalId() int32 {
	if m != nil {
		return m.VerticalId
	}
	return 0
}

func (m *InitialProvisionRequest) GetCallerNumber() string {
	if m != nil {
		return m.CallerNumber
	}
	return ""
}

func (m *InitialProvisionRequest) GetPrimaryDataCenter() string {
	if m != nil {
		return m.PrimaryDataCenter
	}
	return ""
}

func (m *InitialProvisionRequest) GetStreet1() string {
	if m != nil {
		return m.Street1
	}
	return ""
}

func (m *InitialProvisionRequest) GetStreet2() string {
	if m != nil {
		return m.Street2
	}
	return ""
}

func (m *InitialProvisionRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *InitialProvisionRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *InitialProvisionRequest) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *InitialProvisionRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type InitialProvisionResponse struct {
	LocationCreatedSuccessfully          bool     `protobuf:"varint,1,opt,name=location_created_successfully,json=locationCreatedSuccessfully,proto3" json:"location_created_successfully,omitempty"`
	LocationCreatedError                 string   `protobuf:"bytes,2,opt,name=location_created_error,json=locationCreatedError,proto3" json:"location_created_error,omitempty"`
	PhoneDataTenantCreatedSuccessfully   bool     `protobuf:"varint,3,opt,name=phone_data_tenant_created_successfully,json=phoneDataTenantCreatedSuccessfully,proto3" json:"phone_data_tenant_created_successfully,omitempty"`
	PhoneDataTenantCreatedError          string   `protobuf:"bytes,4,opt,name=phone_data_tenant_created_error,json=phoneDataTenantCreatedError,proto3" json:"phone_data_tenant_created_error,omitempty"`
	SoftwareIntegrationSetupSuccessfully bool     `protobuf:"varint,5,opt,name=software_integration_setup_successfully,json=softwareIntegrationSetupSuccessfully,proto3" json:"software_integration_setup_successfully,omitempty"`
	SoftwareIntegrationError             string   `protobuf:"bytes,6,opt,name=software_integration_error,json=softwareIntegrationError,proto3" json:"software_integration_error,omitempty"`
	TempNumberSetupSuccessfully          bool     `protobuf:"varint,7,opt,name=temp_number_setup_successfully,json=tempNumberSetupSuccessfully,proto3" json:"temp_number_setup_successfully,omitempty"`
	TempNumberSetupError                 string   `protobuf:"bytes,8,opt,name=temp_number_setup_error,json=tempNumberSetupError,proto3" json:"temp_number_setup_error,omitempty"`
	E911SetupSuccessfully                bool     `protobuf:"varint,9,opt,name=e911_setup_successfully,json=e911SetupSuccessfully,proto3" json:"e911_setup_successfully,omitempty"`
	E911SetupError                       string   `protobuf:"bytes,10,opt,name=e911_setup_error,json=e911SetupError,proto3" json:"e911_setup_error,omitempty"`
	CustomizationFlagsSetupSuccessfully  bool     `protobuf:"varint,11,opt,name=customization_flags_setup_successfully,json=customizationFlagsSetupSuccessfully,proto3" json:"customization_flags_setup_successfully,omitempty"`
	CustomizationFlagsError              string   `protobuf:"bytes,12,opt,name=customization_flags_error,json=customizationFlagsError,proto3" json:"customization_flags_error,omitempty"`
	LocationId                           string   `protobuf:"bytes,13,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{} `json:"-"`
	XXX_unrecognized                     []byte   `json:"-"`
	XXX_sizecache                        int32    `json:"-"`
}

func (m *InitialProvisionResponse) Reset()         { *m = InitialProvisionResponse{} }
func (m *InitialProvisionResponse) String() string { return proto.CompactTextString(m) }
func (*InitialProvisionResponse) ProtoMessage()    {}
func (*InitialProvisionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11524db7586d182a, []int{1}
}

func (m *InitialProvisionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialProvisionResponse.Unmarshal(m, b)
}
func (m *InitialProvisionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialProvisionResponse.Marshal(b, m, deterministic)
}
func (m *InitialProvisionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialProvisionResponse.Merge(m, src)
}
func (m *InitialProvisionResponse) XXX_Size() int {
	return xxx_messageInfo_InitialProvisionResponse.Size(m)
}
func (m *InitialProvisionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialProvisionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitialProvisionResponse proto.InternalMessageInfo

func (m *InitialProvisionResponse) GetLocationCreatedSuccessfully() bool {
	if m != nil {
		return m.LocationCreatedSuccessfully
	}
	return false
}

func (m *InitialProvisionResponse) GetLocationCreatedError() string {
	if m != nil {
		return m.LocationCreatedError
	}
	return ""
}

func (m *InitialProvisionResponse) GetPhoneDataTenantCreatedSuccessfully() bool {
	if m != nil {
		return m.PhoneDataTenantCreatedSuccessfully
	}
	return false
}

func (m *InitialProvisionResponse) GetPhoneDataTenantCreatedError() string {
	if m != nil {
		return m.PhoneDataTenantCreatedError
	}
	return ""
}

func (m *InitialProvisionResponse) GetSoftwareIntegrationSetupSuccessfully() bool {
	if m != nil {
		return m.SoftwareIntegrationSetupSuccessfully
	}
	return false
}

func (m *InitialProvisionResponse) GetSoftwareIntegrationError() string {
	if m != nil {
		return m.SoftwareIntegrationError
	}
	return ""
}

func (m *InitialProvisionResponse) GetTempNumberSetupSuccessfully() bool {
	if m != nil {
		return m.TempNumberSetupSuccessfully
	}
	return false
}

func (m *InitialProvisionResponse) GetTempNumberSetupError() string {
	if m != nil {
		return m.TempNumberSetupError
	}
	return ""
}

func (m *InitialProvisionResponse) GetE911SetupSuccessfully() bool {
	if m != nil {
		return m.E911SetupSuccessfully
	}
	return false
}

func (m *InitialProvisionResponse) GetE911SetupError() string {
	if m != nil {
		return m.E911SetupError
	}
	return ""
}

func (m *InitialProvisionResponse) GetCustomizationFlagsSetupSuccessfully() bool {
	if m != nil {
		return m.CustomizationFlagsSetupSuccessfully
	}
	return false
}

func (m *InitialProvisionResponse) GetCustomizationFlagsError() string {
	if m != nil {
		return m.CustomizationFlagsError
	}
	return ""
}

func (m *InitialProvisionResponse) GetLocationId() string {
	if m != nil {
		return m.LocationId
	}
	return ""
}

func init() {
	proto.RegisterType((*InitialProvisionRequest)(nil), "provisioningproto.InitialProvisionRequest")
	proto.RegisterType((*InitialProvisionResponse)(nil), "provisioningproto.InitialProvisionResponse")
}

func init() {
	proto.RegisterFile("protorepo/messages/insys/provisioning.proto", fileDescriptor_11524db7586d182a)
}

var fileDescriptor_11524db7586d182a = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5f, 0x4f, 0xdb, 0x3e,
	0x14, 0x15, 0xbf, 0x52, 0xa0, 0xb7, 0xf0, 0x13, 0x78, 0x6c, 0xf5, 0x86, 0x36, 0x18, 0x4c, 0x0c,
	0x69, 0x12, 0x11, 0xec, 0x8f, 0xb4, 0x69, 0x4f, 0xfc, 0x99, 0xd4, 0x97, 0x69, 0x2a, 0xdb, 0x0b,
	0x2f, 0x91, 0x49, 0x6e, 0x8b, 0xa5, 0xc4, 0xce, 0x6c, 0x07, 0xd6, 0xbe, 0xed, 0x33, 0xee, 0x0b,
	0x4d, 0xf6, 0x6d, 0xa0, 0x34, 0xd9, 0x53, 0xec, 0x73, 0x8f, 0xcf, 0x39, 0xb6, 0x6f, 0x0c, 0x6f,
	0x0a, 0xa3, 0x9d, 0x36, 0x58, 0xe8, 0x28, 0x47, 0x6b, 0xc5, 0x08, 0x6d, 0x24, 0x95, 0x1d, 0xdb,
	0xa8, 0x30, 0xfa, 0x46, 0x5a, 0xa9, 0x95, 0x54, 0xa3, 0xc3, 0xc0, 0x62, 0x1b, 0xb3, 0x58, 0x80,
	0x76, 0x7f, 0xb7, 0xa0, 0xd7, 0x57, 0xd2, 0x49, 0x91, 0x7d, 0xab, 0x8a, 0x03, 0xfc, 0x59, 0xa2,
	0x75, 0x8c, 0xc1, 0xa2, 0x12, 0x39, 0xf2, 0x85, 0x9d, 0x85, 0x83, 0xce, 0x20, 0x8c, 0x3d, 0x66,
	0xb3, 0x72, 0xc4, 0xff, 0x23, 0xcc, 0x8f, 0xd9, 0x4b, 0x58, 0xd5, 0xc3, 0xa1, 0x4c, 0x30, 0xc6,
	0x5c, 0xc8, 0x8c, 0xb7, 0x42, 0xad, 0x4b, 0xd8, 0xb9, 0x87, 0xd8, 0x16, 0x74, 0x9c, 0xcc, 0x31,
	0x9e, 0x68, 0x85, 0x7c, 0x31, 0xd4, 0x57, 0x3c, 0x70, 0xa9, 0x15, 0xb2, 0x6d, 0xe8, 0xde, 0xa0,
	0x71, 0x32, 0x11, 0x59, 0x2c, 0x53, 0xde, 0xde, 0x59, 0x38, 0x68, 0x0f, 0xa0, 0x82, 0xfa, 0x29,
	0xdb, 0x83, 0xb5, 0x44, 0x64, 0x19, 0x9a, 0x58, 0x95, 0xf9, 0x15, 0x1a, 0xbe, 0x14, 0x14, 0x56,
	0x09, 0xfc, 0x1a, 0x30, 0x76, 0x08, 0x8f, 0x0a, 0x23, 0x73, 0x61, 0xc6, 0x71, 0x2a, 0x9c, 0x88,
	0x13, 0x54, 0x0e, 0x0d, 0x5f, 0x0e, 0xd4, 0x8d, 0x69, 0xe9, 0x4c, 0x38, 0x71, 0x1a, 0x0a, 0x8c,
	0xc3, 0xb2, 0x75, 0x06, 0xd1, 0x1d, 0xf1, 0x95, 0xc0, 0xa9, 0xa6, 0xf7, 0x95, 0x63, 0xde, 0x99,
	0xad, 0x1c, 0xfb, 0xdd, 0x27, 0xd2, 0x8d, 0x39, 0xd0, 0xee, 0xfd, 0x98, 0x6d, 0x42, 0xdb, 0x3a,
	0xe1, 0x90, 0x77, 0x03, 0x48, 0x13, 0xb6, 0x0e, 0xad, 0x89, 0x2c, 0xf8, 0x6a, 0xc0, 0xfc, 0xd0,
	0xab, 0x26, 0xba, 0x54, 0xce, 0x8c, 0xf9, 0x1a, 0xa9, 0x4e, 0xa7, 0xbb, 0x7f, 0x96, 0x80, 0xd7,
	0xef, 0xc0, 0x16, 0x5a, 0x59, 0x64, 0x27, 0xf0, 0x3c, 0xd3, 0x89, 0x70, 0x52, 0xab, 0x38, 0x31,
	0x28, 0x1c, 0xa6, 0xb1, 0x2d, 0x93, 0x04, 0xad, 0x1d, 0x96, 0x59, 0x36, 0x0e, 0xb7, 0xb3, 0x32,
	0xd8, 0xaa, 0x48, 0xa7, 0xc4, 0xb9, 0x98, 0xa1, 0xb0, 0x77, 0xf0, 0xa4, 0xa6, 0x81, 0xc6, 0x68,
	0x33, 0xbd, 0xc6, 0xcd, 0xb9, 0xc5, 0xe7, 0xbe, 0xc6, 0x06, 0xb0, 0x5f, 0x5c, 0x6b, 0x85, 0x74,
	0x9c, 0x0e, 0x95, 0x50, 0xae, 0x39, 0x42, 0x2b, 0x44, 0xd8, 0x0d, 0x6c, 0x7f, 0xc2, 0xdf, 0x03,
	0xb7, 0x29, 0xc9, 0x19, 0x6c, 0xff, 0x5b, 0x93, 0x22, 0x51, 0x77, 0x6c, 0x35, 0x8b, 0x51, 0xb2,
	0x1f, 0xf0, 0xda, 0xea, 0xa1, 0xbb, 0x15, 0x06, 0x63, 0xa9, 0x1c, 0x8e, 0x0c, 0xed, 0xcd, 0xa2,
	0x2b, 0x8b, 0x87, 0xd1, 0xda, 0x21, 0xda, 0xab, 0x8a, 0xde, 0xbf, 0x67, 0x5f, 0x78, 0xf2, 0x83,
	0x70, 0x9f, 0xe1, 0x59, 0xa3, 0x2c, 0xe5, 0xa2, 0x9e, 0xe3, 0x0d, 0x4a, 0x14, 0xea, 0x14, 0x5e,
	0x38, 0xcc, 0x8b, 0x69, 0x8b, 0x36, 0x65, 0x59, 0xa6, 0x9b, 0xf2, 0x2c, 0xea, 0xd9, 0x7a, 0x84,
	0xf7, 0xd0, 0xab, 0x8b, 0x90, 0x3f, 0x35, 0xe9, 0xe6, 0xdc, 0x6a, 0xf2, 0xfe, 0x00, 0x3d, 0xfc,
	0x78, 0x74, 0xd4, 0x64, 0xda, 0x09, 0xa6, 0x8f, 0x7d, 0xb9, 0x6e, 0x77, 0x00, 0xeb, 0x33, 0xeb,
	0xc8, 0x87, 0x7a, 0xfb, 0xff, 0xbb, 0x05, 0xe4, 0x70, 0x01, 0xfb, 0x49, 0x69, 0x9d, 0xce, 0xe5,
	0x84, 0x0e, 0x65, 0x98, 0x89, 0x91, 0x6d, 0x32, 0xec, 0x06, 0xc3, 0xbd, 0x07, 0xec, 0x2f, 0x9e,
	0x5c, 0xb7, 0xff, 0x04, 0x4f, 0x9b, 0x44, 0x29, 0x07, 0xfd, 0x3a, 0xbd, 0xba, 0x0e, 0x05, 0xda,
	0x86, 0xee, 0x5d, 0x4f, 0xcb, 0x74, 0xfa, 0x4b, 0x41, 0x05, 0xf5, 0xd3, 0x93, 0xb3, 0xcb, 0x93,
	0x5b, 0x14, 0x37, 0x98, 0x89, 0xab, 0xc3, 0x5f, 0xe3, 0x49, 0x94, 0x6b, 0xa5, 0x8d, 0x90, 0x59,
	0x64, 0xaf, 0x85, 0xc1, 0x34, 0xba, 0x7f, 0x38, 0x53, 0x69, 0x5d, 0x34, 0x9a, 0x7f, 0x40, 0x43,
	0xfd, 0x6a, 0x29, 0x7c, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x60, 0x83, 0x8b, 0x96, 0x68,
	0x05, 0x00, 0x00,
}