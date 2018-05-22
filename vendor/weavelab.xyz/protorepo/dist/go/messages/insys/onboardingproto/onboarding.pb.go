// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/messages/insys/onboarding.proto

package onboardingproto // import "weavelab.xyz/protorepo/dist/go/messages/insys/onboardingproto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import insysenums "weavelab.xyz/protorepo/dist/go/enums/insysenums"
import sharedproto "weavelab.xyz/protorepo/dist/go/messages/sharedproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Category struct {
	ID                   *sharedproto.UUID    `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	DisplayText          string               `protobuf:"bytes,2,opt,name=DisplayText" json:"DisplayText,omitempty"`
	DisplayOrder         int32                `protobuf:"varint,3,opt,name=DisplayOrder" json:"DisplayOrder,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt" json:"UpdatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{0}
}
func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (dst *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(dst, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetID() *sharedproto.UUID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Category) GetDisplayText() string {
	if m != nil {
		return m.DisplayText
	}
	return ""
}

func (m *Category) GetDisplayOrder() int32 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

func (m *Category) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Category) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type TaskInstance struct {
	ID                   *sharedproto.UUID               `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	LocationID           *sharedproto.UUID               `protobuf:"bytes,2,opt,name=LocationID" json:"LocationID,omitempty"`
	CategoryID           *sharedproto.UUID               `protobuf:"bytes,3,opt,name=CategoryID" json:"CategoryID,omitempty"`
	TaskID               *sharedproto.UUID               `protobuf:"bytes,4,opt,name=TaskID" json:"TaskID,omitempty"`
	Title                string                          `protobuf:"bytes,5,opt,name=Title" json:"Title,omitempty"`
	DisplayOrder         int32                           `protobuf:"varint,6,opt,name=DisplayOrder" json:"DisplayOrder,omitempty"`
	CompletedAt          *timestamp.Timestamp            `protobuf:"bytes,7,opt,name=CompletedAt" json:"CompletedAt,omitempty"`
	CompletedBy          string                          `protobuf:"bytes,8,opt,name=CompletedBy" json:"CompletedBy,omitempty"`
	VerifiedAt           *timestamp.Timestamp            `protobuf:"bytes,9,opt,name=VerifiedAt" json:"VerifiedAt,omitempty"`
	VerifiedBy           string                          `protobuf:"bytes,10,opt,name=VerifiedBy" json:"VerifiedBy,omitempty"`
	Status               insysenums.OnboardingTaskStatus `protobuf:"varint,11,opt,name=Status,enum=insysenums.OnboardingTaskStatus" json:"Status,omitempty"`
	StatusUpdatedAt      *timestamp.Timestamp            `protobuf:"bytes,12,opt,name=StatusUpdatedAt" json:"StatusUpdatedAt,omitempty"`
	StatusUpdatedBy      string                          `protobuf:"bytes,13,opt,name=StatusUpdatedBy" json:"StatusUpdatedBy,omitempty"`
	CreatedAt            *timestamp.Timestamp            `protobuf:"bytes,14,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp            `protobuf:"bytes,15,opt,name=UpdatedAt" json:"UpdatedAt,omitempty"`
	Content              string                          `protobuf:"bytes,16,opt,name=Content" json:"Content,omitempty"`
	ButtonContent        string                          `protobuf:"bytes,17,opt,name=ButtonContent" json:"ButtonContent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TaskInstance) Reset()         { *m = TaskInstance{} }
func (m *TaskInstance) String() string { return proto.CompactTextString(m) }
func (*TaskInstance) ProtoMessage()    {}
func (*TaskInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{1}
}
func (m *TaskInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInstance.Unmarshal(m, b)
}
func (m *TaskInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInstance.Marshal(b, m, deterministic)
}
func (dst *TaskInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInstance.Merge(dst, src)
}
func (m *TaskInstance) XXX_Size() int {
	return xxx_messageInfo_TaskInstance.Size(m)
}
func (m *TaskInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInstance.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInstance proto.InternalMessageInfo

func (m *TaskInstance) GetID() *sharedproto.UUID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *TaskInstance) GetLocationID() *sharedproto.UUID {
	if m != nil {
		return m.LocationID
	}
	return nil
}

func (m *TaskInstance) GetCategoryID() *sharedproto.UUID {
	if m != nil {
		return m.CategoryID
	}
	return nil
}

func (m *TaskInstance) GetTaskID() *sharedproto.UUID {
	if m != nil {
		return m.TaskID
	}
	return nil
}

func (m *TaskInstance) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskInstance) GetDisplayOrder() int32 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

func (m *TaskInstance) GetCompletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CompletedAt
	}
	return nil
}

func (m *TaskInstance) GetCompletedBy() string {
	if m != nil {
		return m.CompletedBy
	}
	return ""
}

func (m *TaskInstance) GetVerifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.VerifiedAt
	}
	return nil
}

func (m *TaskInstance) GetVerifiedBy() string {
	if m != nil {
		return m.VerifiedBy
	}
	return ""
}

func (m *TaskInstance) GetStatus() insysenums.OnboardingTaskStatus {
	if m != nil {
		return m.Status
	}
	return insysenums.OnboardingTaskStatus_WaitingOnCustomer
}

func (m *TaskInstance) GetStatusUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StatusUpdatedAt
	}
	return nil
}

func (m *TaskInstance) GetStatusUpdatedBy() string {
	if m != nil {
		return m.StatusUpdatedBy
	}
	return ""
}

func (m *TaskInstance) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TaskInstance) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TaskInstance) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskInstance) GetButtonContent() string {
	if m != nil {
		return m.ButtonContent
	}
	return ""
}

type CategoryRequest struct {
	ID                   *sharedproto.UUID `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CategoryRequest) Reset()         { *m = CategoryRequest{} }
func (m *CategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryRequest) ProtoMessage()    {}
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{2}
}
func (m *CategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryRequest.Unmarshal(m, b)
}
func (m *CategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryRequest.Marshal(b, m, deterministic)
}
func (dst *CategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryRequest.Merge(dst, src)
}
func (m *CategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryRequest.Size(m)
}
func (m *CategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryRequest proto.InternalMessageInfo

func (m *CategoryRequest) GetID() *sharedproto.UUID {
	if m != nil {
		return m.ID
	}
	return nil
}

type CategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=Category" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CategoryResponse) Reset()         { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()    {}
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{3}
}
func (m *CategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryResponse.Unmarshal(m, b)
}
func (m *CategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryResponse.Marshal(b, m, deterministic)
}
func (dst *CategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryResponse.Merge(dst, src)
}
func (m *CategoryResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryResponse.Size(m)
}
func (m *CategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryResponse proto.InternalMessageInfo

func (m *CategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type CreateTaskInstancesFromTasksRequest struct {
	LocationID           *sharedproto.UUID `protobuf:"bytes,1,opt,name=LocationID" json:"LocationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateTaskInstancesFromTasksRequest) Reset()         { *m = CreateTaskInstancesFromTasksRequest{} }
func (m *CreateTaskInstancesFromTasksRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskInstancesFromTasksRequest) ProtoMessage()    {}
func (*CreateTaskInstancesFromTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{4}
}
func (m *CreateTaskInstancesFromTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskInstancesFromTasksRequest.Unmarshal(m, b)
}
func (m *CreateTaskInstancesFromTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskInstancesFromTasksRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTaskInstancesFromTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskInstancesFromTasksRequest.Merge(dst, src)
}
func (m *CreateTaskInstancesFromTasksRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskInstancesFromTasksRequest.Size(m)
}
func (m *CreateTaskInstancesFromTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskInstancesFromTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskInstancesFromTasksRequest proto.InternalMessageInfo

func (m *CreateTaskInstancesFromTasksRequest) GetLocationID() *sharedproto.UUID {
	if m != nil {
		return m.LocationID
	}
	return nil
}

type TaskInstancesRequest struct {
	LocationID           *sharedproto.UUID `protobuf:"bytes,1,opt,name=LocationID" json:"LocationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskInstancesRequest) Reset()         { *m = TaskInstancesRequest{} }
func (m *TaskInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*TaskInstancesRequest) ProtoMessage()    {}
func (*TaskInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{5}
}
func (m *TaskInstancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInstancesRequest.Unmarshal(m, b)
}
func (m *TaskInstancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInstancesRequest.Marshal(b, m, deterministic)
}
func (dst *TaskInstancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInstancesRequest.Merge(dst, src)
}
func (m *TaskInstancesRequest) XXX_Size() int {
	return xxx_messageInfo_TaskInstancesRequest.Size(m)
}
func (m *TaskInstancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInstancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInstancesRequest proto.InternalMessageInfo

func (m *TaskInstancesRequest) GetLocationID() *sharedproto.UUID {
	if m != nil {
		return m.LocationID
	}
	return nil
}

type TaskInstancesResponse struct {
	TaskInstances        []*TaskInstance `protobuf:"bytes,1,rep,name=TaskInstances" json:"TaskInstances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TaskInstancesResponse) Reset()         { *m = TaskInstancesResponse{} }
func (m *TaskInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*TaskInstancesResponse) ProtoMessage()    {}
func (*TaskInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{6}
}
func (m *TaskInstancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInstancesResponse.Unmarshal(m, b)
}
func (m *TaskInstancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInstancesResponse.Marshal(b, m, deterministic)
}
func (dst *TaskInstancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInstancesResponse.Merge(dst, src)
}
func (m *TaskInstancesResponse) XXX_Size() int {
	return xxx_messageInfo_TaskInstancesResponse.Size(m)
}
func (m *TaskInstancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInstancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInstancesResponse proto.InternalMessageInfo

func (m *TaskInstancesResponse) GetTaskInstances() []*TaskInstance {
	if m != nil {
		return m.TaskInstances
	}
	return nil
}

type UpdateTaskInstanceRequest struct {
	ID                   *sharedproto.UUID               `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Status               insysenums.OnboardingTaskStatus `protobuf:"varint,2,opt,name=Status,enum=insysenums.OnboardingTaskStatus" json:"Status,omitempty"`
	StatusUpdatedBy      string                          `protobuf:"bytes,3,opt,name=StatusUpdatedBy" json:"StatusUpdatedBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UpdateTaskInstanceRequest) Reset()         { *m = UpdateTaskInstanceRequest{} }
func (m *UpdateTaskInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskInstanceRequest) ProtoMessage()    {}
func (*UpdateTaskInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{7}
}
func (m *UpdateTaskInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskInstanceRequest.Unmarshal(m, b)
}
func (m *UpdateTaskInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskInstanceRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTaskInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskInstanceRequest.Merge(dst, src)
}
func (m *UpdateTaskInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskInstanceRequest.Size(m)
}
func (m *UpdateTaskInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskInstanceRequest proto.InternalMessageInfo

func (m *UpdateTaskInstanceRequest) GetID() *sharedproto.UUID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *UpdateTaskInstanceRequest) GetStatus() insysenums.OnboardingTaskStatus {
	if m != nil {
		return m.Status
	}
	return insysenums.OnboardingTaskStatus_WaitingOnCustomer
}

func (m *UpdateTaskInstanceRequest) GetStatusUpdatedBy() string {
	if m != nil {
		return m.StatusUpdatedBy
	}
	return ""
}

type UpdateTaskInstanceResponse struct {
	TaskInstance         *TaskInstance `protobuf:"bytes,1,opt,name=TaskInstance" json:"TaskInstance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateTaskInstanceResponse) Reset()         { *m = UpdateTaskInstanceResponse{} }
func (m *UpdateTaskInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskInstanceResponse) ProtoMessage()    {}
func (*UpdateTaskInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_onboarding_0461c9a6cfb9f164, []int{8}
}
func (m *UpdateTaskInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskInstanceResponse.Unmarshal(m, b)
}
func (m *UpdateTaskInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskInstanceResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateTaskInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskInstanceResponse.Merge(dst, src)
}
func (m *UpdateTaskInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskInstanceResponse.Size(m)
}
func (m *UpdateTaskInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskInstanceResponse proto.InternalMessageInfo

func (m *UpdateTaskInstanceResponse) GetTaskInstance() *TaskInstance {
	if m != nil {
		return m.TaskInstance
	}
	return nil
}

func init() {
	proto.RegisterType((*Category)(nil), "onboardingproto.Category")
	proto.RegisterType((*TaskInstance)(nil), "onboardingproto.TaskInstance")
	proto.RegisterType((*CategoryRequest)(nil), "onboardingproto.CategoryRequest")
	proto.RegisterType((*CategoryResponse)(nil), "onboardingproto.CategoryResponse")
	proto.RegisterType((*CreateTaskInstancesFromTasksRequest)(nil), "onboardingproto.CreateTaskInstancesFromTasksRequest")
	proto.RegisterType((*TaskInstancesRequest)(nil), "onboardingproto.TaskInstancesRequest")
	proto.RegisterType((*TaskInstancesResponse)(nil), "onboardingproto.TaskInstancesResponse")
	proto.RegisterType((*UpdateTaskInstanceRequest)(nil), "onboardingproto.UpdateTaskInstanceRequest")
	proto.RegisterType((*UpdateTaskInstanceResponse)(nil), "onboardingproto.UpdateTaskInstanceResponse")
}

func init() {
	proto.RegisterFile("protorepo/messages/insys/onboarding.proto", fileDescriptor_onboarding_0461c9a6cfb9f164)
}

var fileDescriptor_onboarding_0461c9a6cfb9f164 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x6f, 0x6b, 0xd3, 0x5e,
	0x14, 0x26, 0xdd, 0x6f, 0xdd, 0x7a, 0xda, 0xae, 0xfb, 0x5d, 0x26, 0x64, 0xc5, 0x3f, 0x21, 0x16,
	0x89, 0x20, 0x09, 0x54, 0x84, 0x21, 0x8a, 0xac, 0x0d, 0x42, 0x40, 0x18, 0x64, 0xad, 0x2f, 0x44,
	0x90, 0xdb, 0xe5, 0x2e, 0x06, 0x9b, 0xdc, 0x98, 0x7b, 0xa3, 0x8b, 0x2f, 0xfd, 0x1e, 0x7e, 0x39,
	0x3f, 0x89, 0xf4, 0x26, 0x69, 0xfe, 0xb4, 0x36, 0x6e, 0xaf, 0xda, 0x7b, 0xce, 0xf3, 0x9c, 0x9c,
	0xf3, 0xe4, 0x39, 0x37, 0xf0, 0x34, 0x8c, 0x28, 0xa7, 0x11, 0x09, 0xa9, 0xe1, 0x13, 0xc6, 0xb0,
	0x4b, 0x98, 0xe1, 0x05, 0x2c, 0x61, 0x06, 0x0d, 0x16, 0x14, 0x47, 0x8e, 0x17, 0xb8, 0xba, 0xc0,
	0xa0, 0x41, 0x11, 0x11, 0x81, 0xe1, 0x23, 0x97, 0x52, 0x77, 0x49, 0x0c, 0x71, 0x5a, 0xc4, 0xd7,
	0x06, 0xf7, 0x7c, 0xc2, 0x38, 0xf6, 0xc3, 0x94, 0x31, 0x1c, 0x6d, 0x29, 0xce, 0x3e, 0xe3, 0x88,
	0x38, 0x46, 0x1c, 0x7b, 0x4e, 0x86, 0x7a, 0x52, 0xa0, 0x48, 0x10, 0xfb, 0x7f, 0x7b, 0xbe, 0xfa,
	0x5b, 0x82, 0xc3, 0x29, 0xe6, 0xc4, 0xa5, 0x51, 0x82, 0xee, 0x43, 0xcb, 0x32, 0x65, 0x49, 0x91,
	0xb4, 0xee, 0xb8, 0xa7, 0xa7, 0x45, 0xf5, 0xf9, 0xdc, 0x32, 0xed, 0x96, 0x65, 0x22, 0x05, 0xba,
	0xa6, 0xc7, 0xc2, 0x25, 0x4e, 0x66, 0xe4, 0x86, 0xcb, 0x2d, 0x45, 0xd2, 0x3a, 0x76, 0x39, 0x84,
	0x54, 0xe8, 0x65, 0xc7, 0x8b, 0xc8, 0x21, 0x91, 0xbc, 0xa7, 0x48, 0xda, 0xbe, 0x5d, 0x89, 0xa1,
	0x33, 0xe8, 0x4c, 0x23, 0x82, 0x39, 0x71, 0xce, 0xb9, 0xfc, 0x9f, 0x78, 0xd4, 0x50, 0x4f, 0x67,
	0xd6, 0xf3, 0x99, 0xf5, 0x59, 0x3e, 0xb3, 0x5d, 0x80, 0x57, 0xcc, 0x79, 0xe8, 0x64, 0xcc, 0xfd,
	0x66, 0xe6, 0x1a, 0xac, 0xfe, 0x6c, 0x43, 0x6f, 0x86, 0xd9, 0x17, 0x2b, 0x60, 0x1c, 0x07, 0x57,
	0xa4, 0x61, 0xd0, 0x67, 0x00, 0xef, 0xe8, 0x15, 0xe6, 0x1e, 0x0d, 0x2c, 0x53, 0xcc, 0x59, 0x47,
	0x95, 0xf2, 0x2b, 0x74, 0x2e, 0xa0, 0x65, 0x8a, 0x91, 0x37, 0xd0, 0x45, 0x1e, 0x8d, 0xa0, 0x2d,
	0x3a, 0x31, 0xb3, 0xd9, 0xab, 0xc8, 0x2c, 0x87, 0x4e, 0x60, 0x7f, 0xe6, 0xf1, 0x25, 0x11, 0x63,
	0x76, 0xec, 0xf4, 0xb0, 0x21, 0x6f, 0x7b, 0x8b, 0xbc, 0xaf, 0xa0, 0x3b, 0xa5, 0x7e, 0xb8, 0x24,
	0xa9, 0x4c, 0x07, 0x8d, 0x32, 0x95, 0xe1, 0xab, 0x57, 0xbc, 0x3e, 0x4e, 0x12, 0xf9, 0x30, 0x7d,
	0xc5, 0xa5, 0x10, 0x7a, 0x09, 0xf0, 0x9e, 0x44, 0xde, 0xb5, 0x27, 0xca, 0x77, 0x1a, 0xcb, 0x97,
	0xd0, 0xe8, 0x61, 0xc1, 0x9d, 0x24, 0x32, 0x88, 0xe2, 0xa5, 0x08, 0x3a, 0x83, 0xf6, 0x25, 0xc7,
	0x3c, 0x66, 0x72, 0x57, 0x91, 0xb4, 0xa3, 0xb1, 0xa2, 0x0b, 0xd3, 0x0a, 0xff, 0xea, 0x17, 0x6b,
	0xe7, 0xae, 0x34, 0x4a, 0x71, 0x76, 0x86, 0x47, 0x26, 0x0c, 0xd2, 0x7f, 0x85, 0x41, 0x7a, 0x8d,
	0xad, 0xd5, 0x29, 0x48, 0xab, 0x55, 0x99, 0x24, 0x72, 0x5f, 0x34, 0x59, 0x0f, 0x57, 0x4d, 0x7c,
	0x74, 0x67, 0x13, 0x0f, 0x6e, 0x61, 0x62, 0x24, 0xc3, 0xc1, 0x94, 0x06, 0x9c, 0x04, 0x5c, 0x3e,
	0x16, 0x5d, 0xe5, 0x47, 0x34, 0x82, 0xfe, 0x24, 0xe6, 0x9c, 0x06, 0x79, 0xfe, 0x7f, 0x91, 0xaf,
	0x06, 0x55, 0x03, 0x06, 0xb9, 0x0f, 0x6d, 0xf2, 0x35, 0x26, 0x8c, 0xef, 0x5e, 0x03, 0xd5, 0x82,
	0xe3, 0x82, 0xc0, 0x42, 0x1a, 0x30, 0x82, 0x5e, 0x14, 0xb7, 0x45, 0xc6, 0x3b, 0xd5, 0x6b, 0x37,
	0x98, 0xbe, 0x26, 0xad, 0xa1, 0xea, 0x25, 0x3c, 0x4e, 0x25, 0x28, 0x6f, 0x21, 0x7b, 0x1b, 0x51,
	0x7f, 0x15, 0x60, 0x79, 0x3f, 0xd5, 0xc5, 0x93, 0x76, 0x2f, 0x9e, 0x6a, 0xc2, 0x49, 0xa5, 0xdc,
	0xdd, 0xaa, 0x7c, 0x84, 0x7b, 0xb5, 0x2a, 0xd9, 0xa8, 0x53, 0xe8, 0x57, 0x12, 0xb2, 0xa4, 0xec,
	0x69, 0xdd, 0xf1, 0x83, 0x8d, 0x79, 0xcb, 0x28, 0xbb, 0xca, 0x51, 0x7f, 0x49, 0x70, 0x9a, 0xbe,
	0xc2, 0x0a, 0xea, 0x5f, 0xf4, 0x2f, 0xad, 0x43, 0xeb, 0x96, 0xeb, 0xb0, 0xc5, 0xc8, 0x7b, 0x5b,
	0x8d, 0xac, 0x7e, 0x82, 0xe1, 0xb6, 0xf6, 0x32, 0x09, 0xce, 0xab, 0xd7, 0x66, 0xd6, 0x69, 0x83,
	0x02, 0x15, 0xca, 0xe4, 0xcd, 0x87, 0xd7, 0xdf, 0x09, 0xfe, 0x46, 0x96, 0x78, 0xa1, 0xdf, 0x24,
	0x3f, 0x8c, 0xe2, 0xb3, 0xe4, 0x78, 0x8c, 0x1b, 0xee, 0x8e, 0x2f, 0x64, 0xba, 0x0e, 0x6d, 0xf1,
	0xf3, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xc0, 0x79, 0xa6, 0x54, 0x07, 0x00, 0x00,
}
