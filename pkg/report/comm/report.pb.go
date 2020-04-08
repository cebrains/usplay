// Code generated by protoc-gen-go. DO NOT EDIT.
// source: report.proto

package comm

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the report details
type GenerateReportRequest struct {
	Name                 string    `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Period               *TimeSpan `protobuf:"bytes,2,opt,name=Period,proto3" json:"Period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GenerateReportRequest) Reset()         { *m = GenerateReportRequest{} }
func (m *GenerateReportRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateReportRequest) ProtoMessage()    {}
func (*GenerateReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{0}
}

func (m *GenerateReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateReportRequest.Unmarshal(m, b)
}
func (m *GenerateReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateReportRequest.Marshal(b, m, deterministic)
}
func (m *GenerateReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateReportRequest.Merge(m, src)
}
func (m *GenerateReportRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateReportRequest.Size(m)
}
func (m *GenerateReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateReportRequest proto.InternalMessageInfo

func (m *GenerateReportRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenerateReportRequest) GetPeriod() *TimeSpan {
	if m != nil {
		return m.Period
	}
	return nil
}

// The response message containing the id of the report
type GenerateReportReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateReportReply) Reset()         { *m = GenerateReportReply{} }
func (m *GenerateReportReply) String() string { return proto.CompactTextString(m) }
func (*GenerateReportReply) ProtoMessage()    {}
func (*GenerateReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{1}
}

func (m *GenerateReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateReportReply.Unmarshal(m, b)
}
func (m *GenerateReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateReportReply.Marshal(b, m, deterministic)
}
func (m *GenerateReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateReportReply.Merge(m, src)
}
func (m *GenerateReportReply) XXX_Size() int {
	return xxx_messageInfo_GenerateReportReply.Size(m)
}
func (m *GenerateReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateReportReply proto.InternalMessageInfo

func (m *GenerateReportReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The request message for the read request
type ReadReportRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadReportRequest) Reset()         { *m = ReadReportRequest{} }
func (m *ReadReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReadReportRequest) ProtoMessage()    {}
func (*ReadReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{2}
}

func (m *ReadReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadReportRequest.Unmarshal(m, b)
}
func (m *ReadReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadReportRequest.Marshal(b, m, deterministic)
}
func (m *ReadReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadReportRequest.Merge(m, src)
}
func (m *ReadReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReadReportRequest.Size(m)
}
func (m *ReadReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadReportRequest proto.InternalMessageInfo

func (m *ReadReportRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message to the read request
type ReadReportReply struct {
	Report               *Report  `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadReportReply) Reset()         { *m = ReadReportReply{} }
func (m *ReadReportReply) String() string { return proto.CompactTextString(m) }
func (*ReadReportReply) ProtoMessage()    {}
func (*ReadReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{3}
}

func (m *ReadReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadReportReply.Unmarshal(m, b)
}
func (m *ReadReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadReportReply.Marshal(b, m, deterministic)
}
func (m *ReadReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadReportReply.Merge(m, src)
}
func (m *ReadReportReply) XXX_Size() int {
	return xxx_messageInfo_ReadReportReply.Size(m)
}
func (m *ReadReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReadReportReply proto.InternalMessageInfo

func (m *ReadReportReply) GetReport() *Report {
	if m != nil {
		return m.Report
	}
	return nil
}

// The request message for the delete request
type DeleteReportRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReportRequest) Reset()         { *m = DeleteReportRequest{} }
func (m *DeleteReportRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteReportRequest) ProtoMessage()    {}
func (*DeleteReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{4}
}

func (m *DeleteReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReportRequest.Unmarshal(m, b)
}
func (m *DeleteReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReportRequest.Marshal(b, m, deterministic)
}
func (m *DeleteReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReportRequest.Merge(m, src)
}
func (m *DeleteReportRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteReportRequest.Size(m)
}
func (m *DeleteReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReportRequest proto.InternalMessageInfo

func (m *DeleteReportRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message to the delete request
type DeleteReportReply struct {
	Report               *Report  `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReportReply) Reset()         { *m = DeleteReportReply{} }
func (m *DeleteReportReply) String() string { return proto.CompactTextString(m) }
func (*DeleteReportReply) ProtoMessage()    {}
func (*DeleteReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{5}
}

func (m *DeleteReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReportReply.Unmarshal(m, b)
}
func (m *DeleteReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReportReply.Marshal(b, m, deterministic)
}
func (m *DeleteReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReportReply.Merge(m, src)
}
func (m *DeleteReportReply) XXX_Size() int {
	return xxx_messageInfo_DeleteReportReply.Size(m)
}
func (m *DeleteReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReportReply proto.InternalMessageInfo

func (m *DeleteReportReply) GetReport() *Report {
	if m != nil {
		return m.Report
	}
	return nil
}

// The request message for the update request
type UpdateReportRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReportRequest) Reset()         { *m = UpdateReportRequest{} }
func (m *UpdateReportRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateReportRequest) ProtoMessage()    {}
func (*UpdateReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{6}
}

func (m *UpdateReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReportRequest.Unmarshal(m, b)
}
func (m *UpdateReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReportRequest.Marshal(b, m, deterministic)
}
func (m *UpdateReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReportRequest.Merge(m, src)
}
func (m *UpdateReportRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateReportRequest.Size(m)
}
func (m *UpdateReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReportRequest proto.InternalMessageInfo

func (m *UpdateReportRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateReportRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateReportRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *UpdateReportRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// The response message to the update request.
// Returns the old values for the Report
type UpdateReportReply struct {
	Report               *Report  `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReportReply) Reset()         { *m = UpdateReportReply{} }
func (m *UpdateReportReply) String() string { return proto.CompactTextString(m) }
func (*UpdateReportReply) ProtoMessage()    {}
func (*UpdateReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{7}
}

func (m *UpdateReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReportReply.Unmarshal(m, b)
}
func (m *UpdateReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReportReply.Marshal(b, m, deterministic)
}
func (m *UpdateReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReportReply.Merge(m, src)
}
func (m *UpdateReportReply) XXX_Size() int {
	return xxx_messageInfo_UpdateReportReply.Size(m)
}
func (m *UpdateReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReportReply proto.InternalMessageInfo

func (m *UpdateReportReply) GetReport() *Report {
	if m != nil {
		return m.Report
	}
	return nil
}

// The request message for the list reports request
type ListReportsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReportsRequest) Reset()         { *m = ListReportsRequest{} }
func (m *ListReportsRequest) String() string { return proto.CompactTextString(m) }
func (*ListReportsRequest) ProtoMessage()    {}
func (*ListReportsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{8}
}

func (m *ListReportsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReportsRequest.Unmarshal(m, b)
}
func (m *ListReportsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReportsRequest.Marshal(b, m, deterministic)
}
func (m *ListReportsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReportsRequest.Merge(m, src)
}
func (m *ListReportsRequest) XXX_Size() int {
	return xxx_messageInfo_ListReportsRequest.Size(m)
}
func (m *ListReportsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReportsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListReportsRequest proto.InternalMessageInfo

// The response message to the list reports request.
type ListReportsReply struct {
	Reports              []*Report `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListReportsReply) Reset()         { *m = ListReportsReply{} }
func (m *ListReportsReply) String() string { return proto.CompactTextString(m) }
func (*ListReportsReply) ProtoMessage()    {}
func (*ListReportsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{9}
}

func (m *ListReportsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReportsReply.Unmarshal(m, b)
}
func (m *ListReportsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReportsReply.Marshal(b, m, deterministic)
}
func (m *ListReportsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReportsReply.Merge(m, src)
}
func (m *ListReportsReply) XXX_Size() int {
	return xxx_messageInfo_ListReportsReply.Size(m)
}
func (m *ListReportsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReportsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListReportsReply proto.InternalMessageInfo

func (m *ListReportsReply) GetReports() []*Report {
	if m != nil {
		return m.Reports
	}
	return nil
}

type Report struct {
	Id                   string      `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	OrderId              string      `protobuf:"bytes,3,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	Period               *TimeSpan   `protobuf:"bytes,4,opt,name=Period,proto3" json:"Period,omitempty"`
	Activities           []*Activity `protobuf:"bytes,5,rep,name=Activities,proto3" json:"Activities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{10}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Report) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Report) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Report) GetPeriod() *TimeSpan {
	if m != nil {
		return m.Period
	}
	return nil
}

func (m *Report) GetActivities() []*Activity {
	if m != nil {
		return m.Activities
	}
	return nil
}

type TimeSpan struct {
	From                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeSpan) Reset()         { *m = TimeSpan{} }
func (m *TimeSpan) String() string { return proto.CompactTextString(m) }
func (*TimeSpan) ProtoMessage()    {}
func (*TimeSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{11}
}

func (m *TimeSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSpan.Unmarshal(m, b)
}
func (m *TimeSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSpan.Marshal(b, m, deterministic)
}
func (m *TimeSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSpan.Merge(m, src)
}
func (m *TimeSpan) XXX_Size() int {
	return xxx_messageInfo_TimeSpan.Size(m)
}
func (m *TimeSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSpan.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSpan proto.InternalMessageInfo

func (m *TimeSpan) GetFrom() *timestamp.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TimeSpan) GetTo() *timestamp.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

type Activity struct {
	Id                   string        `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Type                 *ActivityType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Period               *TimeSpan     `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{12}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Activity) GetType() *ActivityType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Activity) GetPeriod() *TimeSpan {
	if m != nil {
		return m.Period
	}
	return nil
}

type ActivityType struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityType) Reset()         { *m = ActivityType{} }
func (m *ActivityType) String() string { return proto.CompactTextString(m) }
func (*ActivityType) ProtoMessage()    {}
func (*ActivityType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{13}
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

func (m *ActivityType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ActivityType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateReportRequest)(nil), "gen.GenerateReportRequest")
	proto.RegisterType((*GenerateReportReply)(nil), "gen.GenerateReportReply")
	proto.RegisterType((*ReadReportRequest)(nil), "gen.ReadReportRequest")
	proto.RegisterType((*ReadReportReply)(nil), "gen.ReadReportReply")
	proto.RegisterType((*DeleteReportRequest)(nil), "gen.DeleteReportRequest")
	proto.RegisterType((*DeleteReportReply)(nil), "gen.DeleteReportReply")
	proto.RegisterType((*UpdateReportRequest)(nil), "gen.UpdateReportRequest")
	proto.RegisterType((*UpdateReportReply)(nil), "gen.UpdateReportReply")
	proto.RegisterType((*ListReportsRequest)(nil), "gen.ListReportsRequest")
	proto.RegisterType((*ListReportsReply)(nil), "gen.ListReportsReply")
	proto.RegisterType((*Report)(nil), "gen.Report")
	proto.RegisterType((*TimeSpan)(nil), "gen.TimeSpan")
	proto.RegisterType((*Activity)(nil), "gen.Activity")
	proto.RegisterType((*ActivityType)(nil), "gen.ActivityType")
}

func init() {
	proto.RegisterFile("report.proto", fileDescriptor_3eedb623aa6ca98c)
}

var fileDescriptor_3eedb623aa6ca98c = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x17, 0xd2, 0x74, 0x52, 0x2e, 0xd9, 0x5e, 0xb0, 0xfc, 0x42, 0xe4, 0x2a, 0x52, 0x85,
	0x84, 0x2b, 0x05, 0xa9, 0x2a, 0x12, 0x2f, 0x94, 0x0a, 0x88, 0x84, 0x00, 0x6d, 0xcb, 0x07, 0xb8,
	0xf1, 0x24, 0xac, 0x14, 0x7b, 0x97, 0xf5, 0xb6, 0x92, 0xff, 0x85, 0x7f, 0x05, 0x79, 0x77, 0x0d,
	0xb6, 0x63, 0xb5, 0x79, 0x73, 0xe6, 0x9c, 0x39, 0x33, 0x3b, 0x33, 0x27, 0xb0, 0x27, 0x51, 0x70,
	0xa9, 0x62, 0x21, 0xb9, 0xe2, 0xc4, 0x5b, 0x61, 0x1e, 0xbe, 0x5c, 0x71, 0xbe, 0x5a, 0xe3, 0xa9,
	0x0e, 0xdd, 0xdc, 0x2e, 0x4f, 0x15, 0xcb, 0xb0, 0x50, 0x49, 0x26, 0x0c, 0x2b, 0xa2, 0x70, 0xf8,
	0x09, 0x73, 0x94, 0x89, 0x42, 0xaa, 0xb3, 0x29, 0xfe, 0xba, 0xc5, 0x42, 0x11, 0x02, 0xfe, 0xd7,
	0x24, 0xc3, 0xc0, 0x99, 0x38, 0x27, 0xbb, 0x54, 0x7f, 0x93, 0x29, 0x0c, 0xbe, 0xa3, 0x64, 0x3c,
	0x0d, 0xdc, 0x89, 0x73, 0x32, 0x9a, 0x3d, 0x89, 0x57, 0x98, 0xc7, 0xd7, 0x2c, 0xc3, 0x2b, 0x91,
	0xe4, 0xd4, 0x82, 0xd1, 0x14, 0xf6, 0xbb, 0x9a, 0x62, 0x5d, 0x92, 0xa7, 0xe0, 0xce, 0x53, 0xab,
	0xe7, 0xce, 0xd3, 0xe8, 0x18, 0xc6, 0x14, 0x93, 0xb4, 0x5d, 0xb6, 0x4b, 0x3a, 0x83, 0x67, 0x4d,
	0x52, 0xa5, 0x73, 0x0c, 0x03, 0xf3, 0x50, 0x4d, 0x1b, 0xcd, 0x46, 0xba, 0x0b, 0xcb, 0xb0, 0x50,
	0xd5, 0xc3, 0x25, 0xae, 0xb1, 0xfb, 0xaa, 0xae, 0xfc, 0x39, 0x8c, 0xdb, 0xb4, 0xad, 0x0b, 0x70,
	0xd8, 0xff, 0x21, 0xd2, 0xe4, 0x81, 0x02, 0xff, 0xc6, 0xe8, 0x36, 0xc6, 0x48, 0xc0, 0xff, 0xc0,
	0x53, 0x0c, 0x3c, 0x13, 0xab, 0xbe, 0xc9, 0x04, 0x46, 0x97, 0x58, 0x2c, 0x24, 0x13, 0x8a, 0xf1,
	0x3c, 0xf0, 0x35, 0xd4, 0x0c, 0x55, 0xad, 0xb6, 0x0b, 0x6e, 0xdd, 0xea, 0x01, 0x90, 0x2f, 0xac,
	0x50, 0x26, 0x5a, 0xd8, 0x4e, 0xa3, 0xb7, 0xf0, 0xbc, 0x15, 0xad, 0xe4, 0xa6, 0xb0, 0x63, 0x72,
	0x8a, 0xc0, 0x99, 0x78, 0x5d, 0xbd, 0x1a, 0x8b, 0x7e, 0x3b, 0x30, 0x30, 0xb1, 0xad, 0xde, 0x1b,
	0xc0, 0xce, 0x37, 0x99, 0xa2, 0x9c, 0xa7, 0xf6, 0xc9, 0xf5, 0xcf, 0xc6, 0x41, 0xf9, 0xf7, 0x1c,
	0x14, 0x79, 0x0d, 0xf0, 0x7e, 0xa1, 0xd8, 0x1d, 0x53, 0x0c, 0x8b, 0xe0, 0xb1, 0xee, 0xcc, 0x50,
	0x6d, 0xb8, 0xa4, 0x0d, 0x42, 0xb4, 0x84, 0x61, 0x2d, 0x41, 0x62, 0xf0, 0x97, 0x92, 0x67, 0x76,
	0x3c, 0x61, 0x6c, 0xfc, 0x10, 0xd7, 0x7e, 0xd0, 0xb5, 0xb4, 0x1f, 0xa8, 0xe6, 0x91, 0x57, 0xe0,
	0x2a, 0x6e, 0xcf, 0xfb, 0x3e, 0xb6, 0xab, 0x78, 0xf4, 0x13, 0x86, 0x75, 0xfd, 0x8d, 0x39, 0x4c,
	0xc1, 0x57, 0xa5, 0x40, 0xab, 0x34, 0x6e, 0x35, 0x7b, 0x5d, 0x0a, 0xa4, 0x1a, 0xae, 0x06, 0x20,
	0xcc, 0x00, 0xbc, 0xde, 0x01, 0x18, 0x30, 0xfa, 0x08, 0x7b, 0xcd, 0xe4, 0xbe, 0xa9, 0xeb, 0x8b,
	0x72, 0x1b, 0x17, 0x55, 0x6f, 0xc2, 0xfb, 0xbf, 0x89, 0xd9, 0x1f, 0x07, 0x76, 0xcd, 0xe2, 0xae,
	0xee, 0x16, 0xe4, 0x02, 0x86, 0xb5, 0x4f, 0x49, 0xa8, 0x0b, 0xf7, 0xfe, 0x15, 0x84, 0x41, 0x2f,
	0x26, 0xd6, 0x65, 0xf4, 0x88, 0x9c, 0x81, 0x5f, 0xf9, 0x93, 0x1c, 0xd9, 0x43, 0xe9, 0xf8, 0x39,
	0x3c, 0xd8, 0x88, 0x9b, 0xbc, 0x77, 0x30, 0x30, 0xc6, 0x23, 0x46, 0xbd, 0xc7, 0xac, 0xe1, 0x51,
	0x0f, 0x62, 0xb2, 0xcf, 0xc1, 0xaf, 0x6e, 0x97, 0xbc, 0xd0, 0x8c, 0xcd, 0xe3, 0x0e, 0x0f, 0x37,
	0x01, 0x9d, 0x79, 0xe1, 0x7e, 0xf6, 0x6e, 0x06, 0x7a, 0x9f, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x3b, 0x5d, 0x24, 0x30, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportSvcClient is the client API for ReportSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportSvcClient interface {
	// Generates a new report
	Generate(ctx context.Context, in *GenerateReportRequest, opts ...grpc.CallOption) (*GenerateReportReply, error)
	// Reads an report
	Read(ctx context.Context, in *ReadReportRequest, opts ...grpc.CallOption) (*ReadReportReply, error)
	// Delete an report
	Delete(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*DeleteReportReply, error)
	// List an report
	List(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsReply, error)
}

type reportSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewReportSvcClient(cc grpc.ClientConnInterface) ReportSvcClient {
	return &reportSvcClient{cc}
}

func (c *reportSvcClient) Generate(ctx context.Context, in *GenerateReportRequest, opts ...grpc.CallOption) (*GenerateReportReply, error) {
	out := new(GenerateReportReply)
	err := c.cc.Invoke(ctx, "/gen.ReportSvc/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportSvcClient) Read(ctx context.Context, in *ReadReportRequest, opts ...grpc.CallOption) (*ReadReportReply, error) {
	out := new(ReadReportReply)
	err := c.cc.Invoke(ctx, "/gen.ReportSvc/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportSvcClient) Delete(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*DeleteReportReply, error) {
	out := new(DeleteReportReply)
	err := c.cc.Invoke(ctx, "/gen.ReportSvc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportSvcClient) List(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsReply, error) {
	out := new(ListReportsReply)
	err := c.cc.Invoke(ctx, "/gen.ReportSvc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportSvcServer is the server API for ReportSvc service.
type ReportSvcServer interface {
	// Generates a new report
	Generate(context.Context, *GenerateReportRequest) (*GenerateReportReply, error)
	// Reads an report
	Read(context.Context, *ReadReportRequest) (*ReadReportReply, error)
	// Delete an report
	Delete(context.Context, *DeleteReportRequest) (*DeleteReportReply, error)
	// List an report
	List(context.Context, *ListReportsRequest) (*ListReportsReply, error)
}

// UnimplementedReportSvcServer can be embedded to have forward compatible implementations.
type UnimplementedReportSvcServer struct {
}

func (*UnimplementedReportSvcServer) Generate(ctx context.Context, req *GenerateReportRequest) (*GenerateReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (*UnimplementedReportSvcServer) Read(ctx context.Context, req *ReadReportRequest) (*ReadReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedReportSvcServer) Delete(ctx context.Context, req *DeleteReportRequest) (*DeleteReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedReportSvcServer) List(ctx context.Context, req *ListReportsRequest) (*ListReportsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterReportSvcServer(s *grpc.Server, srv ReportSvcServer) {
	s.RegisterService(&_ReportSvc_serviceDesc, srv)
}

func _ReportSvc_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportSvcServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.ReportSvc/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportSvcServer).Generate(ctx, req.(*GenerateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportSvc_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportSvcServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.ReportSvc/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportSvcServer).Read(ctx, req.(*ReadReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportSvc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportSvcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.ReportSvc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportSvcServer).Delete(ctx, req.(*DeleteReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportSvc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportSvcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.ReportSvc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportSvcServer).List(ctx, req.(*ListReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gen.ReportSvc",
	HandlerType: (*ReportSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _ReportSvc_Generate_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ReportSvc_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReportSvc_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ReportSvc_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "report.proto",
}
