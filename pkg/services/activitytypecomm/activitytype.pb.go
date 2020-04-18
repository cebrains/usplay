// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activitytypecomm/activitytype.proto

package activitytypecomm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the ActivityType details
type CreateActivityTypeRequest struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateActivityTypeRequest) Reset()         { *m = CreateActivityTypeRequest{} }
func (m *CreateActivityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateActivityTypeRequest) ProtoMessage()    {}
func (*CreateActivityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{0}
}

func (m *CreateActivityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActivityTypeRequest.Unmarshal(m, b)
}
func (m *CreateActivityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActivityTypeRequest.Marshal(b, m, deterministic)
}
func (m *CreateActivityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActivityTypeRequest.Merge(m, src)
}
func (m *CreateActivityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateActivityTypeRequest.Size(m)
}
func (m *CreateActivityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActivityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActivityTypeRequest proto.InternalMessageInfo

func (m *CreateActivityTypeRequest) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateActivityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the id of the ActivityType
type CreateActivityTypeReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateActivityTypeReply) Reset()         { *m = CreateActivityTypeReply{} }
func (m *CreateActivityTypeReply) String() string { return proto.CompactTextString(m) }
func (*CreateActivityTypeReply) ProtoMessage()    {}
func (*CreateActivityTypeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{1}
}

func (m *CreateActivityTypeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActivityTypeReply.Unmarshal(m, b)
}
func (m *CreateActivityTypeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActivityTypeReply.Marshal(b, m, deterministic)
}
func (m *CreateActivityTypeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActivityTypeReply.Merge(m, src)
}
func (m *CreateActivityTypeReply) XXX_Size() int {
	return xxx_messageInfo_CreateActivityTypeReply.Size(m)
}
func (m *CreateActivityTypeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActivityTypeReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActivityTypeReply proto.InternalMessageInfo

func (m *CreateActivityTypeReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The request message for the read request
type ExistActivityTypeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistActivityTypeRequest) Reset()         { *m = ExistActivityTypeRequest{} }
func (m *ExistActivityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*ExistActivityTypeRequest) ProtoMessage()    {}
func (*ExistActivityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{2}
}

func (m *ExistActivityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistActivityTypeRequest.Unmarshal(m, b)
}
func (m *ExistActivityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistActivityTypeRequest.Marshal(b, m, deterministic)
}
func (m *ExistActivityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistActivityTypeRequest.Merge(m, src)
}
func (m *ExistActivityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_ExistActivityTypeRequest.Size(m)
}
func (m *ExistActivityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistActivityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExistActivityTypeRequest proto.InternalMessageInfo

func (m *ExistActivityTypeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message to the read request
type ExistActivityTypeReply struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistActivityTypeReply) Reset()         { *m = ExistActivityTypeReply{} }
func (m *ExistActivityTypeReply) String() string { return proto.CompactTextString(m) }
func (*ExistActivityTypeReply) ProtoMessage()    {}
func (*ExistActivityTypeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{3}
}

func (m *ExistActivityTypeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistActivityTypeReply.Unmarshal(m, b)
}
func (m *ExistActivityTypeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistActivityTypeReply.Marshal(b, m, deterministic)
}
func (m *ExistActivityTypeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistActivityTypeReply.Merge(m, src)
}
func (m *ExistActivityTypeReply) XXX_Size() int {
	return xxx_messageInfo_ExistActivityTypeReply.Size(m)
}
func (m *ExistActivityTypeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistActivityTypeReply.DiscardUnknown(m)
}

var xxx_messageInfo_ExistActivityTypeReply proto.InternalMessageInfo

func (m *ExistActivityTypeReply) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

// The request message for the read request
type ReadActivityTypeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadActivityTypeRequest) Reset()         { *m = ReadActivityTypeRequest{} }
func (m *ReadActivityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*ReadActivityTypeRequest) ProtoMessage()    {}
func (*ReadActivityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{4}
}

func (m *ReadActivityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadActivityTypeRequest.Unmarshal(m, b)
}
func (m *ReadActivityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadActivityTypeRequest.Marshal(b, m, deterministic)
}
func (m *ReadActivityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadActivityTypeRequest.Merge(m, src)
}
func (m *ReadActivityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_ReadActivityTypeRequest.Size(m)
}
func (m *ReadActivityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadActivityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadActivityTypeRequest proto.InternalMessageInfo

func (m *ReadActivityTypeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message to the read request
type ReadActivityTypeReply struct {
	ActivityType         *ActivityType `protobuf:"bytes,1,opt,name=ActivityType,proto3" json:"ActivityType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadActivityTypeReply) Reset()         { *m = ReadActivityTypeReply{} }
func (m *ReadActivityTypeReply) String() string { return proto.CompactTextString(m) }
func (*ReadActivityTypeReply) ProtoMessage()    {}
func (*ReadActivityTypeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{5}
}

func (m *ReadActivityTypeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadActivityTypeReply.Unmarshal(m, b)
}
func (m *ReadActivityTypeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadActivityTypeReply.Marshal(b, m, deterministic)
}
func (m *ReadActivityTypeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadActivityTypeReply.Merge(m, src)
}
func (m *ReadActivityTypeReply) XXX_Size() int {
	return xxx_messageInfo_ReadActivityTypeReply.Size(m)
}
func (m *ReadActivityTypeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadActivityTypeReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReadActivityTypeReply proto.InternalMessageInfo

func (m *ReadActivityTypeReply) GetActivityType() *ActivityType {
	if m != nil {
		return m.ActivityType
	}
	return nil
}

// The request message for the delete request
type DeleteActivityTypeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteActivityTypeRequest) Reset()         { *m = DeleteActivityTypeRequest{} }
func (m *DeleteActivityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteActivityTypeRequest) ProtoMessage()    {}
func (*DeleteActivityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{6}
}

func (m *DeleteActivityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteActivityTypeRequest.Unmarshal(m, b)
}
func (m *DeleteActivityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteActivityTypeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteActivityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteActivityTypeRequest.Merge(m, src)
}
func (m *DeleteActivityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteActivityTypeRequest.Size(m)
}
func (m *DeleteActivityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteActivityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteActivityTypeRequest proto.InternalMessageInfo

func (m *DeleteActivityTypeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message to the delete request
type DeleteActivityTypeReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteActivityTypeReply) Reset()         { *m = DeleteActivityTypeReply{} }
func (m *DeleteActivityTypeReply) String() string { return proto.CompactTextString(m) }
func (*DeleteActivityTypeReply) ProtoMessage()    {}
func (*DeleteActivityTypeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{7}
}

func (m *DeleteActivityTypeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteActivityTypeReply.Unmarshal(m, b)
}
func (m *DeleteActivityTypeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteActivityTypeReply.Marshal(b, m, deterministic)
}
func (m *DeleteActivityTypeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteActivityTypeReply.Merge(m, src)
}
func (m *DeleteActivityTypeReply) XXX_Size() int {
	return xxx_messageInfo_DeleteActivityTypeReply.Size(m)
}
func (m *DeleteActivityTypeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteActivityTypeReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteActivityTypeReply proto.InternalMessageInfo

// The request message for the update request
type UpdateActivityTypeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateActivityTypeRequest) Reset()         { *m = UpdateActivityTypeRequest{} }
func (m *UpdateActivityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateActivityTypeRequest) ProtoMessage()    {}
func (*UpdateActivityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{8}
}

func (m *UpdateActivityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActivityTypeRequest.Unmarshal(m, b)
}
func (m *UpdateActivityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActivityTypeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateActivityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActivityTypeRequest.Merge(m, src)
}
func (m *UpdateActivityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateActivityTypeRequest.Size(m)
}
func (m *UpdateActivityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActivityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActivityTypeRequest proto.InternalMessageInfo

func (m *UpdateActivityTypeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateActivityTypeRequest) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateActivityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message to the update requests
type UpdateActivityTypeReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateActivityTypeReply) Reset()         { *m = UpdateActivityTypeReply{} }
func (m *UpdateActivityTypeReply) String() string { return proto.CompactTextString(m) }
func (*UpdateActivityTypeReply) ProtoMessage()    {}
func (*UpdateActivityTypeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{9}
}

func (m *UpdateActivityTypeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActivityTypeReply.Unmarshal(m, b)
}
func (m *UpdateActivityTypeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActivityTypeReply.Marshal(b, m, deterministic)
}
func (m *UpdateActivityTypeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActivityTypeReply.Merge(m, src)
}
func (m *UpdateActivityTypeReply) XXX_Size() int {
	return xxx_messageInfo_UpdateActivityTypeReply.Size(m)
}
func (m *UpdateActivityTypeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActivityTypeReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActivityTypeReply proto.InternalMessageInfo

// The request message for the list ActivityTypes request
type ListActivityTypesRequest struct {
	FilterIds            []string `protobuf:"bytes,1,rep,name=FilterIds,proto3" json:"FilterIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListActivityTypesRequest) Reset()         { *m = ListActivityTypesRequest{} }
func (m *ListActivityTypesRequest) String() string { return proto.CompactTextString(m) }
func (*ListActivityTypesRequest) ProtoMessage()    {}
func (*ListActivityTypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{10}
}

func (m *ListActivityTypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActivityTypesRequest.Unmarshal(m, b)
}
func (m *ListActivityTypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActivityTypesRequest.Marshal(b, m, deterministic)
}
func (m *ListActivityTypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActivityTypesRequest.Merge(m, src)
}
func (m *ListActivityTypesRequest) XXX_Size() int {
	return xxx_messageInfo_ListActivityTypesRequest.Size(m)
}
func (m *ListActivityTypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActivityTypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListActivityTypesRequest proto.InternalMessageInfo

func (m *ListActivityTypesRequest) GetFilterIds() []string {
	if m != nil {
		return m.FilterIds
	}
	return nil
}

// The response message to the list ActivityTypes request.
type ListActivityTypesReply struct {
	ActivityTypes        []*ActivityType `protobuf:"bytes,1,rep,name=ActivityTypes,proto3" json:"ActivityTypes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListActivityTypesReply) Reset()         { *m = ListActivityTypesReply{} }
func (m *ListActivityTypesReply) String() string { return proto.CompactTextString(m) }
func (*ListActivityTypesReply) ProtoMessage()    {}
func (*ListActivityTypesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{11}
}

func (m *ListActivityTypesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActivityTypesReply.Unmarshal(m, b)
}
func (m *ListActivityTypesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActivityTypesReply.Marshal(b, m, deterministic)
}
func (m *ListActivityTypesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActivityTypesReply.Merge(m, src)
}
func (m *ListActivityTypesReply) XXX_Size() int {
	return xxx_messageInfo_ListActivityTypesReply.Size(m)
}
func (m *ListActivityTypesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActivityTypesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListActivityTypesReply proto.InternalMessageInfo

func (m *ListActivityTypesReply) GetActivityTypes() []*ActivityType {
	if m != nil {
		return m.ActivityTypes
	}
	return nil
}

type ActivityType struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityType) Reset()         { *m = ActivityType{} }
func (m *ActivityType) String() string { return proto.CompactTextString(m) }
func (*ActivityType) ProtoMessage()    {}
func (*ActivityType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa4a553331aef38, []int{12}
}

func (m *ActivityType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityType.Unmarshal(m, b)
}
func (m *ActivityType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityType.Marshal(b, m, deterministic)
}
func (m *ActivityType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityType.Merge(m, src)
}
func (m *ActivityType) XXX_Size() int {
	return xxx_messageInfo_ActivityType.Size(m)
}
func (m *ActivityType) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityType.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityType proto.InternalMessageInfo

func (m *ActivityType) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActivityType) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ActivityType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateActivityTypeRequest)(nil), "activitytypecomm.CreateActivityTypeRequest")
	proto.RegisterType((*CreateActivityTypeReply)(nil), "activitytypecomm.CreateActivityTypeReply")
	proto.RegisterType((*ExistActivityTypeRequest)(nil), "activitytypecomm.ExistActivityTypeRequest")
	proto.RegisterType((*ExistActivityTypeReply)(nil), "activitytypecomm.ExistActivityTypeReply")
	proto.RegisterType((*ReadActivityTypeRequest)(nil), "activitytypecomm.ReadActivityTypeRequest")
	proto.RegisterType((*ReadActivityTypeReply)(nil), "activitytypecomm.ReadActivityTypeReply")
	proto.RegisterType((*DeleteActivityTypeRequest)(nil), "activitytypecomm.DeleteActivityTypeRequest")
	proto.RegisterType((*DeleteActivityTypeReply)(nil), "activitytypecomm.DeleteActivityTypeReply")
	proto.RegisterType((*UpdateActivityTypeRequest)(nil), "activitytypecomm.UpdateActivityTypeRequest")
	proto.RegisterType((*UpdateActivityTypeReply)(nil), "activitytypecomm.UpdateActivityTypeReply")
	proto.RegisterType((*ListActivityTypesRequest)(nil), "activitytypecomm.ListActivityTypesRequest")
	proto.RegisterType((*ListActivityTypesReply)(nil), "activitytypecomm.ListActivityTypesReply")
	proto.RegisterType((*ActivityType)(nil), "activitytypecomm.ActivityType")
}

func init() {
	proto.RegisterFile("activitytypecomm/activitytype.proto", fileDescriptor_8aa4a553331aef38)
}

var fileDescriptor_8aa4a553331aef38 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0x57, 0x23, 0x32, 0x7c, 0x6a, 0x25, 0x92, 0x38, 0x42, 0xa8, 0x5a, 0x0e, 0xa4, 0xad,
	0x64, 0xa3, 0x72, 0xe1, 0x86, 0xa8, 0x4b, 0x84, 0xa5, 0x8a, 0x83, 0x0b, 0x17, 0x40, 0x45, 0x6b,
	0x7b, 0x54, 0x2c, 0x6c, 0xbc, 0x78, 0x37, 0x11, 0xfe, 0xaf, 0xfc, 0x18, 0xb4, 0xeb, 0x02, 0xb6,
	0xb3, 0x2b, 0x2c, 0x6e, 0xce, 0xf8, 0xcd, 0x7b, 0x93, 0x37, 0x6f, 0x0c, 0x4f, 0x58, 0x2c, 0xd3,
	0x5d, 0x2a, 0x2b, 0x59, 0x71, 0x8c, 0x8b, 0x3c, 0xf7, 0x9a, 0x05, 0x97, 0x97, 0x85, 0x2c, 0xc8,
	0x83, 0x2e, 0x88, 0xfa, 0xe0, 0xf8, 0x25, 0x32, 0x89, 0xaf, 0x6e, 0xde, 0xbc, 0xab, 0x38, 0x86,
	0xf8, 0x7d, 0x8b, 0x42, 0x12, 0x02, 0x13, 0xbf, 0x48, 0x70, 0x39, 0x3c, 0x1c, 0xae, 0x0f, 0x42,
	0xfd, 0xac, 0x6a, 0x6f, 0x59, 0x8e, 0xcb, 0xd1, 0xe1, 0x70, 0x3d, 0x0b, 0xf5, 0x33, 0x3d, 0x82,
	0x85, 0x89, 0x84, 0x67, 0x15, 0xb9, 0x07, 0xa3, 0x20, 0xd1, 0x04, 0xb3, 0x70, 0x14, 0x24, 0xf4,
	0x18, 0x96, 0xaf, 0x7f, 0xa4, 0x42, 0x9a, 0xe4, 0xba, 0xd8, 0x67, 0x30, 0x37, 0x60, 0x15, 0xeb,
	0x1c, 0xa6, 0xa8, 0xde, 0x08, 0x8d, 0xbe, 0x15, 0xde, 0xfc, 0x52, 0x83, 0x84, 0xc8, 0x92, 0x3e,
	0xe4, 0x1f, 0xe1, 0xe1, 0x3e, 0x54, 0x71, 0x9f, 0xc1, 0x9d, 0x66, 0x51, 0xb7, 0xdc, 0x3e, 0x7d,
	0xec, 0x76, 0xad, 0x73, 0x5b, 0xad, 0xad, 0x1e, 0x7a, 0x02, 0xce, 0x39, 0x66, 0x68, 0x76, 0xb5,
	0x3b, 0x89, 0x03, 0x0b, 0x13, 0x98, 0x67, 0x15, 0xbd, 0x04, 0xe7, 0x3d, 0x4f, 0x58, 0x2f, 0x9e,
	0x3f, 0xdb, 0x1a, 0x19, 0xb6, 0x35, 0x6e, 0x6c, 0xcb, 0x81, 0x85, 0x89, 0x54, 0xe9, 0xbd, 0x80,
	0xe5, 0x45, 0xc7, 0x70, 0xf1, 0x5b, 0xee, 0x11, 0xcc, 0x36, 0x69, 0x26, 0xb1, 0x0c, 0x12, 0x65,
	0xfb, 0x78, 0x3d, 0x0b, 0xff, 0x16, 0xe8, 0x15, 0xcc, 0x0d, 0x9d, 0xca, 0xcf, 0x73, 0xb8, 0xdb,
	0xaa, 0xea, 0xde, 0x7f, 0x1b, 0xda, 0x6e, 0xa2, 0x9b, 0xf6, 0x56, 0xfe, 0xf7, 0xcf, 0x9f, 0xfe,
	0x9c, 0xc0, 0xfd, 0x26, 0xd1, 0xe5, 0x2e, 0x26, 0x11, 0x4c, 0xeb, 0xf8, 0x92, 0x93, 0xfd, 0xa1,
	0xac, 0xd7, 0xb1, 0x3a, 0xea, 0x07, 0x56, 0xbe, 0x0e, 0xc8, 0x67, 0x38, 0xd0, 0x59, 0x26, 0xc7,
	0xfb, 0x5d, 0xb6, 0x83, 0x58, 0xad, 0x7b, 0x61, 0x6b, 0x81, 0x4f, 0x30, 0x51, 0x79, 0x26, 0x86,
	0xa9, 0x2c, 0x27, 0xb1, 0x7a, 0xda, 0x07, 0x5a, 0xb3, 0x47, 0x30, 0xad, 0x33, 0x6a, 0xb2, 0xc8,
	0x1a, 0x75, 0x93, 0x45, 0xb6, 0xa8, 0x6b, 0x8d, 0x3a, 0x97, 0x26, 0x0d, 0xeb, 0x19, 0x98, 0x34,
	0x6c, 0xf1, 0x1e, 0x90, 0x2b, 0x98, 0x5c, 0x58, 0xb6, 0x60, 0x0b, 0xbe, 0x69, 0x0b, 0xe6, 0xa8,
	0xd3, 0xc1, 0x99, 0xff, 0x66, 0xfc, 0xe1, 0xe5, 0x75, 0x2a, 0xbf, 0x6c, 0x23, 0x37, 0x2e, 0x72,
	0x6f, 0x53, 0xb2, 0x6f, 0x31, 0x8a, 0xb8, 0x08, 0x32, 0x56, 0xa6, 0x85, 0xb7, 0x15, 0x3c, 0x63,
	0x95, 0xc7, 0xbf, 0x5e, 0x7b, 0x02, 0xcb, 0x5d, 0x1a, 0xa3, 0xf0, 0xba, 0xec, 0xd1, 0x54, 0x7f,
	0xac, 0x9f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x26, 0xe6, 0x25, 0xd3, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActivityTypeSvcClient is the client API for ActivityTypeSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityTypeSvcClient interface {
	// Creates a new ActivityType
	Create(ctx context.Context, in *CreateActivityTypeRequest, opts ...grpc.CallOption) (*CreateActivityTypeReply, error)
	// Exists an activity
	Exist(ctx context.Context, in *ExistActivityTypeRequest, opts ...grpc.CallOption) (*ExistActivityTypeReply, error)
	// Reads an ActivityType
	Read(ctx context.Context, in *ReadActivityTypeRequest, opts ...grpc.CallOption) (*ReadActivityTypeReply, error)
	// Delete an ActivityType
	Delete(ctx context.Context, in *DeleteActivityTypeRequest, opts ...grpc.CallOption) (*DeleteActivityTypeReply, error)
	// Update an ActivityType
	Update(ctx context.Context, in *UpdateActivityTypeRequest, opts ...grpc.CallOption) (*UpdateActivityTypeReply, error)
	// List an ActivityType
	List(ctx context.Context, in *ListActivityTypesRequest, opts ...grpc.CallOption) (*ListActivityTypesReply, error)
}

type activityTypeSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityTypeSvcClient(cc grpc.ClientConnInterface) ActivityTypeSvcClient {
	return &activityTypeSvcClient{cc}
}

func (c *activityTypeSvcClient) Create(ctx context.Context, in *CreateActivityTypeRequest, opts ...grpc.CallOption) (*CreateActivityTypeReply, error) {
	out := new(CreateActivityTypeReply)
	err := c.cc.Invoke(ctx, "/activitytypecomm.ActivityTypeSvc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityTypeSvcClient) Exist(ctx context.Context, in *ExistActivityTypeRequest, opts ...grpc.CallOption) (*ExistActivityTypeReply, error) {
	out := new(ExistActivityTypeReply)
	err := c.cc.Invoke(ctx, "/activitytypecomm.ActivityTypeSvc/Exist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityTypeSvcClient) Read(ctx context.Context, in *ReadActivityTypeRequest, opts ...grpc.CallOption) (*ReadActivityTypeReply, error) {
	out := new(ReadActivityTypeReply)
	err := c.cc.Invoke(ctx, "/activitytypecomm.ActivityTypeSvc/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityTypeSvcClient) Delete(ctx context.Context, in *DeleteActivityTypeRequest, opts ...grpc.CallOption) (*DeleteActivityTypeReply, error) {
	out := new(DeleteActivityTypeReply)
	err := c.cc.Invoke(ctx, "/activitytypecomm.ActivityTypeSvc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityTypeSvcClient) Update(ctx context.Context, in *UpdateActivityTypeRequest, opts ...grpc.CallOption) (*UpdateActivityTypeReply, error) {
	out := new(UpdateActivityTypeReply)
	err := c.cc.Invoke(ctx, "/activitytypecomm.ActivityTypeSvc/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityTypeSvcClient) List(ctx context.Context, in *ListActivityTypesRequest, opts ...grpc.CallOption) (*ListActivityTypesReply, error) {
	out := new(ListActivityTypesReply)
	err := c.cc.Invoke(ctx, "/activitytypecomm.ActivityTypeSvc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityTypeSvcServer is the server API for ActivityTypeSvc service.
type ActivityTypeSvcServer interface {
	// Creates a new ActivityType
	Create(context.Context, *CreateActivityTypeRequest) (*CreateActivityTypeReply, error)
	// Exists an activity
	Exist(context.Context, *ExistActivityTypeRequest) (*ExistActivityTypeReply, error)
	// Reads an ActivityType
	Read(context.Context, *ReadActivityTypeRequest) (*ReadActivityTypeReply, error)
	// Delete an ActivityType
	Delete(context.Context, *DeleteActivityTypeRequest) (*DeleteActivityTypeReply, error)
	// Update an ActivityType
	Update(context.Context, *UpdateActivityTypeRequest) (*UpdateActivityTypeReply, error)
	// List an ActivityType
	List(context.Context, *ListActivityTypesRequest) (*ListActivityTypesReply, error)
}

// UnimplementedActivityTypeSvcServer can be embedded to have forward compatible implementations.
type UnimplementedActivityTypeSvcServer struct {
}

func (*UnimplementedActivityTypeSvcServer) Create(ctx context.Context, req *CreateActivityTypeRequest) (*CreateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedActivityTypeSvcServer) Exist(ctx context.Context, req *ExistActivityTypeRequest) (*ExistActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exist not implemented")
}
func (*UnimplementedActivityTypeSvcServer) Read(ctx context.Context, req *ReadActivityTypeRequest) (*ReadActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedActivityTypeSvcServer) Delete(ctx context.Context, req *DeleteActivityTypeRequest) (*DeleteActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedActivityTypeSvcServer) Update(ctx context.Context, req *UpdateActivityTypeRequest) (*UpdateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedActivityTypeSvcServer) List(ctx context.Context, req *ListActivityTypesRequest) (*ListActivityTypesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterActivityTypeSvcServer(s *grpc.Server, srv ActivityTypeSvcServer) {
	s.RegisterService(&_ActivityTypeSvc_serviceDesc, srv)
}

func _ActivityTypeSvc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityTypeSvcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitytypecomm.ActivityTypeSvc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityTypeSvcServer).Create(ctx, req.(*CreateActivityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityTypeSvc_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistActivityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityTypeSvcServer).Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitytypecomm.ActivityTypeSvc/Exist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityTypeSvcServer).Exist(ctx, req.(*ExistActivityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityTypeSvc_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadActivityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityTypeSvcServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitytypecomm.ActivityTypeSvc/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityTypeSvcServer).Read(ctx, req.(*ReadActivityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityTypeSvc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActivityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityTypeSvcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitytypecomm.ActivityTypeSvc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityTypeSvcServer).Delete(ctx, req.(*DeleteActivityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityTypeSvc_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActivityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityTypeSvcServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitytypecomm.ActivityTypeSvc/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityTypeSvcServer).Update(ctx, req.(*UpdateActivityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityTypeSvc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivityTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityTypeSvcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activitytypecomm.ActivityTypeSvc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityTypeSvcServer).List(ctx, req.(*ListActivityTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivityTypeSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "activitytypecomm.ActivityTypeSvc",
	HandlerType: (*ActivityTypeSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ActivityTypeSvc_Create_Handler,
		},
		{
			MethodName: "Exist",
			Handler:    _ActivityTypeSvc_Exist_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ActivityTypeSvc_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ActivityTypeSvc_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ActivityTypeSvc_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ActivityTypeSvc_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activitytypecomm/activitytype.proto",
}
