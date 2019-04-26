// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protorepo/messages/insys/onboarding.proto

package insysproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	insysenums "weavelab.xyz/monorail/shared/protorepo/dist/go/enums/insysenums"
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

type Category struct {
	ID                   *sharedproto.UUID    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	DisplayText          string               `protobuf:"bytes,2,opt,name=DisplayText,proto3" json:"DisplayText,omitempty"`
	DisplayOrder         int32                `protobuf:"varint,3,opt,name=DisplayOrder,proto3" json:"DisplayOrder,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
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
	ID                   *sharedproto.UUID               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	LocationID           *sharedproto.UUID               `protobuf:"bytes,2,opt,name=LocationID,proto3" json:"LocationID,omitempty"`
	CategoryID           *sharedproto.UUID               `protobuf:"bytes,3,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	TaskID               *sharedproto.UUID               `protobuf:"bytes,4,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	Title                string                          `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	DisplayOrder         int32                           `protobuf:"varint,6,opt,name=DisplayOrder,proto3" json:"DisplayOrder,omitempty"`
	CompletedAt          *timestamp.Timestamp            `protobuf:"bytes,7,opt,name=CompletedAt,proto3" json:"CompletedAt,omitempty"`
	CompletedBy          string                          `protobuf:"bytes,8,opt,name=CompletedBy,proto3" json:"CompletedBy,omitempty"`
	VerifiedAt           *timestamp.Timestamp            `protobuf:"bytes,9,opt,name=VerifiedAt,proto3" json:"VerifiedAt,omitempty"`
	VerifiedBy           string                          `protobuf:"bytes,10,opt,name=VerifiedBy,proto3" json:"VerifiedBy,omitempty"`
	Status               insysenums.OnboardingTaskStatus `protobuf:"varint,11,opt,name=Status,proto3,enum=insysenums.OnboardingTaskStatus" json:"Status,omitempty"`
	StatusUpdatedAt      *timestamp.Timestamp            `protobuf:"bytes,12,opt,name=StatusUpdatedAt,proto3" json:"StatusUpdatedAt,omitempty"`
	StatusUpdatedBy      string                          `protobuf:"bytes,13,opt,name=StatusUpdatedBy,proto3" json:"StatusUpdatedBy,omitempty"`
	CreatedAt            *timestamp.Timestamp            `protobuf:"bytes,14,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp            `protobuf:"bytes,15,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Content              string                          `protobuf:"bytes,16,opt,name=Content,proto3" json:"Content,omitempty"`
	ButtonContent        string                          `protobuf:"bytes,17,opt,name=ButtonContent,proto3" json:"ButtonContent,omitempty"`
	ButtonExternalURL    string                          `protobuf:"bytes,18,opt,name=ButtonExternalURL,proto3" json:"ButtonExternalURL,omitempty"`
	Explanation          string                          `protobuf:"bytes,19,opt,name=Explanation,proto3" json:"Explanation,omitempty"`
	ButtonInternalURL    string                          `protobuf:"bytes,20,opt,name=ButtonInternalURL,proto3" json:"ButtonInternalURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TaskInstance) Reset()         { *m = TaskInstance{} }
func (m *TaskInstance) String() string { return proto.CompactTextString(m) }
func (*TaskInstance) ProtoMessage()    {}
func (*TaskInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{1}
}

func (m *TaskInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInstance.Unmarshal(m, b)
}
func (m *TaskInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInstance.Marshal(b, m, deterministic)
}
func (m *TaskInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInstance.Merge(m, src)
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

func (m *TaskInstance) GetButtonExternalURL() string {
	if m != nil {
		return m.ButtonExternalURL
	}
	return ""
}

func (m *TaskInstance) GetExplanation() string {
	if m != nil {
		return m.Explanation
	}
	return ""
}

func (m *TaskInstance) GetButtonInternalURL() string {
	if m != nil {
		return m.ButtonInternalURL
	}
	return ""
}

type CategoryRequest struct {
	ID                   *sharedproto.UUID `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CategoryRequest) Reset()         { *m = CategoryRequest{} }
func (m *CategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryRequest) ProtoMessage()    {}
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{2}
}

func (m *CategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryRequest.Unmarshal(m, b)
}
func (m *CategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryRequest.Marshal(b, m, deterministic)
}
func (m *CategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryRequest.Merge(m, src)
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
	Category             *Category `protobuf:"bytes,1,opt,name=Category,proto3" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CategoryResponse) Reset()         { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()    {}
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{3}
}

func (m *CategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryResponse.Unmarshal(m, b)
}
func (m *CategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryResponse.Marshal(b, m, deterministic)
}
func (m *CategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryResponse.Merge(m, src)
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
	LocationID           *sharedproto.UUID `protobuf:"bytes,1,opt,name=LocationID,proto3" json:"LocationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateTaskInstancesFromTasksRequest) Reset()         { *m = CreateTaskInstancesFromTasksRequest{} }
func (m *CreateTaskInstancesFromTasksRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskInstancesFromTasksRequest) ProtoMessage()    {}
func (*CreateTaskInstancesFromTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{4}
}

func (m *CreateTaskInstancesFromTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskInstancesFromTasksRequest.Unmarshal(m, b)
}
func (m *CreateTaskInstancesFromTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskInstancesFromTasksRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskInstancesFromTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskInstancesFromTasksRequest.Merge(m, src)
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
	LocationID           *sharedproto.UUID `protobuf:"bytes,1,opt,name=LocationID,proto3" json:"LocationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskInstancesRequest) Reset()         { *m = TaskInstancesRequest{} }
func (m *TaskInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*TaskInstancesRequest) ProtoMessage()    {}
func (*TaskInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{5}
}

func (m *TaskInstancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInstancesRequest.Unmarshal(m, b)
}
func (m *TaskInstancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInstancesRequest.Marshal(b, m, deterministic)
}
func (m *TaskInstancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInstancesRequest.Merge(m, src)
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
	TaskInstances        []*TaskInstance `protobuf:"bytes,1,rep,name=TaskInstances,proto3" json:"TaskInstances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TaskInstancesResponse) Reset()         { *m = TaskInstancesResponse{} }
func (m *TaskInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*TaskInstancesResponse) ProtoMessage()    {}
func (*TaskInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{6}
}

func (m *TaskInstancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInstancesResponse.Unmarshal(m, b)
}
func (m *TaskInstancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInstancesResponse.Marshal(b, m, deterministic)
}
func (m *TaskInstancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInstancesResponse.Merge(m, src)
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
	ID                   *sharedproto.UUID               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status               insysenums.OnboardingTaskStatus `protobuf:"varint,2,opt,name=Status,proto3,enum=insysenums.OnboardingTaskStatus" json:"Status,omitempty"`
	StatusUpdatedBy      string                          `protobuf:"bytes,3,opt,name=StatusUpdatedBy,proto3" json:"StatusUpdatedBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UpdateTaskInstanceRequest) Reset()         { *m = UpdateTaskInstanceRequest{} }
func (m *UpdateTaskInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskInstanceRequest) ProtoMessage()    {}
func (*UpdateTaskInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{7}
}

func (m *UpdateTaskInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskInstanceRequest.Unmarshal(m, b)
}
func (m *UpdateTaskInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskInstanceRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTaskInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskInstanceRequest.Merge(m, src)
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

type UpdateTaskInstanceExplanationRequest struct {
	ID                   *sharedproto.UUID `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Explanation          string            `protobuf:"bytes,2,opt,name=Explanation,proto3" json:"Explanation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateTaskInstanceExplanationRequest) Reset()         { *m = UpdateTaskInstanceExplanationRequest{} }
func (m *UpdateTaskInstanceExplanationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskInstanceExplanationRequest) ProtoMessage()    {}
func (*UpdateTaskInstanceExplanationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{8}
}

func (m *UpdateTaskInstanceExplanationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskInstanceExplanationRequest.Unmarshal(m, b)
}
func (m *UpdateTaskInstanceExplanationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskInstanceExplanationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTaskInstanceExplanationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskInstanceExplanationRequest.Merge(m, src)
}
func (m *UpdateTaskInstanceExplanationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskInstanceExplanationRequest.Size(m)
}
func (m *UpdateTaskInstanceExplanationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskInstanceExplanationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskInstanceExplanationRequest proto.InternalMessageInfo

func (m *UpdateTaskInstanceExplanationRequest) GetID() *sharedproto.UUID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *UpdateTaskInstanceExplanationRequest) GetExplanation() string {
	if m != nil {
		return m.Explanation
	}
	return ""
}

type UpdateTaskInstanceResponse struct {
	TaskInstance         *TaskInstance `protobuf:"bytes,1,opt,name=TaskInstance,proto3" json:"TaskInstance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateTaskInstanceResponse) Reset()         { *m = UpdateTaskInstanceResponse{} }
func (m *UpdateTaskInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskInstanceResponse) ProtoMessage()    {}
func (*UpdateTaskInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d88e10ea11481c7, []int{9}
}

func (m *UpdateTaskInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskInstanceResponse.Unmarshal(m, b)
}
func (m *UpdateTaskInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskInstanceResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTaskInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskInstanceResponse.Merge(m, src)
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
	proto.RegisterType((*UpdateTaskInstanceExplanationRequest)(nil), "onboardingproto.UpdateTaskInstanceExplanationRequest")
	proto.RegisterType((*UpdateTaskInstanceResponse)(nil), "onboardingproto.UpdateTaskInstanceResponse")
}

func init() {
	proto.RegisterFile("protorepo/messages/insys/onboarding.proto", fileDescriptor_2d88e10ea11481c7)
}

var fileDescriptor_2d88e10ea11481c7 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x6d, 0x6b, 0xdb, 0x30,
	0x10, 0xc6, 0xe9, 0x9a, 0x36, 0x97, 0xb4, 0x69, 0xb5, 0x0e, 0xdc, 0xb0, 0x17, 0xe3, 0x85, 0x91,
	0x41, 0xb1, 0xa1, 0x63, 0x50, 0xc6, 0xbe, 0x34, 0x71, 0x07, 0x86, 0x42, 0xc1, 0x4d, 0xf6, 0x61,
	0x0c, 0x86, 0x5a, 0xab, 0x99, 0x99, 0x2d, 0x79, 0x96, 0xbc, 0xc5, 0xfb, 0xb4, 0x3f, 0xb2, 0x3f,
	0xb7, 0x5f, 0x32, 0x22, 0xdb, 0xf1, 0x4b, 0xb2, 0xa6, 0xed, 0xa7, 0x44, 0x77, 0xcf, 0x3d, 0xba,
	0x3b, 0x3d, 0x27, 0x19, 0x5e, 0x87, 0x11, 0x13, 0x2c, 0x22, 0x21, 0x33, 0x03, 0xc2, 0x39, 0x9e,
	0x12, 0x6e, 0x7a, 0x94, 0x27, 0xdc, 0x64, 0xf4, 0x8a, 0xe1, 0xc8, 0xf5, 0xe8, 0xd4, 0x90, 0x18,
	0xd4, 0x2d, 0x2c, 0xd2, 0xd0, 0x7b, 0x31, 0x65, 0x6c, 0xea, 0x13, 0x53, 0xae, 0xae, 0xe2, 0x1b,
	0x53, 0x78, 0x01, 0xe1, 0x02, 0x07, 0x61, 0x1a, 0xd1, 0xeb, 0xaf, 0x20, 0xe7, 0x5f, 0x71, 0x44,
	0x5c, 0x33, 0x8e, 0x3d, 0x37, 0x43, 0xbd, 0x2a, 0x50, 0x84, 0xc6, 0xc1, 0xff, 0xf6, 0xd7, 0xff,
	0x2a, 0xb0, 0x3d, 0xc2, 0x82, 0x4c, 0x59, 0x94, 0xa0, 0xa7, 0xd0, 0xb0, 0x2d, 0x55, 0xd1, 0x94,
	0x41, 0xfb, 0xb8, 0x63, 0xa4, 0xa4, 0xc6, 0x64, 0x62, 0x5b, 0x4e, 0xc3, 0xb6, 0x90, 0x06, 0x6d,
	0xcb, 0xe3, 0xa1, 0x8f, 0x93, 0x31, 0x99, 0x09, 0xb5, 0xa1, 0x29, 0x83, 0x96, 0x53, 0x36, 0x21,
	0x1d, 0x3a, 0xd9, 0xf2, 0x22, 0x72, 0x49, 0xa4, 0x6e, 0x68, 0xca, 0x60, 0xd3, 0xa9, 0xd8, 0xd0,
	0x09, 0xb4, 0x46, 0x11, 0xc1, 0x82, 0xb8, 0xa7, 0x42, 0x7d, 0x24, 0xb7, 0xea, 0x19, 0x69, 0xcd,
	0x46, 0x5e, 0xb3, 0x31, 0xce, 0x6b, 0x76, 0x0a, 0xf0, 0x3c, 0x72, 0x12, 0xba, 0x59, 0xe4, 0xe6,
	0xfa, 0xc8, 0x05, 0x58, 0xff, 0xbd, 0x05, 0x9d, 0x31, 0xe6, 0xdf, 0x6c, 0xca, 0x05, 0xa6, 0xd7,
	0x64, 0x4d, 0xa1, 0x47, 0x00, 0xe7, 0xec, 0x1a, 0x0b, 0x8f, 0x51, 0xdb, 0x92, 0x75, 0xd6, 0x51,
	0x25, 0xff, 0x1c, 0x9d, 0x37, 0xd0, 0xb6, 0x64, 0xc9, 0x4b, 0xe8, 0xc2, 0x8f, 0xfa, 0xd0, 0x94,
	0x99, 0x58, 0x59, 0xed, 0x55, 0x64, 0xe6, 0x43, 0x07, 0xb0, 0x39, 0xf6, 0x84, 0x4f, 0x64, 0x99,
	0x2d, 0x27, 0x5d, 0x2c, 0xb5, 0xb7, 0xb9, 0xa2, 0xbd, 0xef, 0xa1, 0x3d, 0x62, 0x41, 0xe8, 0x93,
	0xb4, 0x4d, 0x5b, 0x6b, 0xdb, 0x54, 0x86, 0xcf, 0x8f, 0x78, 0xb1, 0x1c, 0x26, 0xea, 0x76, 0x7a,
	0xc4, 0x25, 0x13, 0x7a, 0x07, 0xf0, 0x91, 0x44, 0xde, 0x8d, 0x27, 0xe9, 0x5b, 0x6b, 0xe9, 0x4b,
	0x68, 0xf4, 0xbc, 0x88, 0x1d, 0x26, 0x2a, 0x48, 0xf2, 0x92, 0x05, 0x9d, 0x40, 0xf3, 0x52, 0x60,
	0x11, 0x73, 0xb5, 0xad, 0x29, 0x83, 0xdd, 0x63, 0xcd, 0x90, 0xa2, 0x95, 0xfa, 0x35, 0x2e, 0x16,
	0xca, 0x9d, 0xf7, 0x28, 0xc5, 0x39, 0x19, 0x1e, 0x59, 0xd0, 0x4d, 0xff, 0x15, 0x02, 0xe9, 0xac,
	0x4d, 0xad, 0x1e, 0x82, 0x06, 0x35, 0x96, 0x61, 0xa2, 0xee, 0xc8, 0x24, 0xeb, 0xe6, 0xaa, 0x88,
	0x77, 0x1f, 0x2c, 0xe2, 0xee, 0x3d, 0x44, 0x8c, 0x54, 0xd8, 0x1a, 0x31, 0x2a, 0x08, 0x15, 0xea,
	0x9e, 0xcc, 0x2a, 0x5f, 0xa2, 0x3e, 0xec, 0x0c, 0x63, 0x21, 0x18, 0xcd, 0xfd, 0xfb, 0xd2, 0x5f,
	0x35, 0xa2, 0x23, 0xd8, 0x4f, 0x0d, 0x67, 0x33, 0x41, 0x22, 0x8a, 0xfd, 0x89, 0x73, 0xae, 0x22,
	0x89, 0x5c, 0x76, 0xcc, 0x95, 0x70, 0x36, 0x0b, 0x7d, 0x4c, 0xa5, 0xcc, 0xd5, 0xc7, 0xa9, 0x12,
	0x4a, 0xa6, 0x82, 0xcf, 0xa6, 0x05, 0xdf, 0x41, 0x99, 0xaf, 0xe4, 0xd0, 0x4d, 0xe8, 0xe6, 0x53,
	0xe0, 0x90, 0xef, 0x31, 0xe1, 0xe2, 0xf6, 0x21, 0xd4, 0x6d, 0xd8, 0x2b, 0x02, 0x78, 0xc8, 0x28,
	0x27, 0xe8, 0x6d, 0x71, 0x57, 0x65, 0x71, 0x87, 0x46, 0xed, 0xfe, 0x34, 0x16, 0x41, 0x0b, 0xa8,
	0x7e, 0x09, 0x2f, 0xd3, 0x03, 0x28, 0xdf, 0x01, 0xfc, 0x43, 0xc4, 0x82, 0xb9, 0x81, 0xe7, 0xf9,
	0x54, 0xc7, 0x5e, 0xb9, 0x7d, 0xec, 0x75, 0x0b, 0x0e, 0x2a, 0x74, 0x0f, 0x63, 0xf9, 0x0c, 0x4f,
	0x6a, 0x2c, 0x59, 0xa9, 0x23, 0xd8, 0xa9, 0x38, 0x54, 0x45, 0xdb, 0x18, 0xb4, 0x8f, 0x9f, 0x2d,
	0xd5, 0x5b, 0x46, 0x39, 0xd5, 0x18, 0xfd, 0x8f, 0x02, 0x87, 0xa9, 0x80, 0x2a, 0xa8, 0xbb, 0xf4,
	0xbf, 0x34, 0x8c, 0x8d, 0x7b, 0x0e, 0xe3, 0x8a, 0x31, 0xda, 0x58, 0x39, 0x46, 0xfa, 0x0d, 0xf4,
	0x97, 0xd3, 0x2b, 0x69, 0xec, 0x6e, 0x99, 0xd6, 0xa4, 0xda, 0x58, 0x92, 0xaa, 0xfe, 0x05, 0x7a,
	0xab, 0xda, 0x90, 0xb5, 0xfa, 0xb4, 0xfa, 0x38, 0x64, 0xfb, 0xac, 0xe9, 0x74, 0x25, 0x64, 0x68,
	0x7d, 0x1a, 0xfe, 0x24, 0xf8, 0x07, 0xf1, 0xf1, 0x95, 0x31, 0x4b, 0x7e, 0x99, 0x01, 0xa3, 0x2c,
	0xc2, 0x9e, 0x9f, 0x3f, 0xcc, 0xc5, 0x63, 0xec, 0x7a, 0x5c, 0x98, 0xd3, 0xfa, 0x77, 0x41, 0x3a,
	0xf9, 0x4d, 0xf9, 0xf3, 0xe6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xad, 0xb5, 0x1c, 0x3f,
	0x08, 0x00, 0x00,
}