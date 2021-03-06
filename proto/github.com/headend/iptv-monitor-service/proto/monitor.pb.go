// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: monitor.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MonitorResponseStatus int32

const (
	MonitorResponseStatus_FAIL    MonitorResponseStatus = 0 /// Success
	MonitorResponseStatus_SUCCESS MonitorResponseStatus = 1 /// Failed
)

// Enum value maps for MonitorResponseStatus.
var (
	MonitorResponseStatus_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
	}
	MonitorResponseStatus_value = map[string]int32{
		"FAIL":    0,
		"SUCCESS": 1,
	}
)

func (x MonitorResponseStatus) Enum() *MonitorResponseStatus {
	p := new(MonitorResponseStatus)
	*p = x
	return p
}

func (x MonitorResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MonitorResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_monitor_proto_enumTypes[0].Descriptor()
}

func (MonitorResponseStatus) Type() protoreflect.EnumType {
	return &file_monitor_proto_enumTypes[0]
}

func (x MonitorResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MonitorResponseStatus.Descriptor instead.
func (MonitorResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{0}
}

type Monitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StatusSignal  int64  `protobuf:"varint,2,opt,name=status_signal,proto3" json:"status_signal,omitempty"`
	StatusVideo   int64  `protobuf:"varint,3,opt,name=status_video,proto3" json:"status_video,omitempty"`
	StatusAudio   int64  `protobuf:"varint,4,opt,name=status_audio,proto3" json:"status_audio,omitempty"`
	SignalMonitor bool   `protobuf:"varint,5,opt,name=signal_monitor,proto3" json:"signal_monitor,omitempty"`
	VideoMonitor  bool   `protobuf:"varint,6,opt,name=video_monitor,proto3" json:"video_monitor,omitempty"`
	AudioMonitor  bool   `protobuf:"varint,7,opt,name=audio_monitor,proto3" json:"audio_monitor,omitempty"`
	IsEnable      bool   `protobuf:"varint,8,opt,name=is_enable,proto3" json:"is_enable,omitempty"`
	DateUpdate    string `protobuf:"bytes,9,opt,name=date_update,proto3" json:"date_update,omitempty"`
	AgentId       int64  `protobuf:"varint,10,opt,name=agent_id,proto3" json:"agent_id,omitempty"`
	ProfileId     int64  `protobuf:"varint,11,opt,name=profile_id,proto3" json:"profile_id,omitempty"`
	StatusId      int64  `protobuf:"varint,12,opt,name=status_id,proto3" json:"status_id,omitempty"`
}

func (x *Monitor) Reset() {
	*x = Monitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitor) ProtoMessage() {}

func (x *Monitor) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitor.ProtoReflect.Descriptor instead.
func (*Monitor) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{0}
}

func (x *Monitor) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Monitor) GetStatusSignal() int64 {
	if x != nil {
		return x.StatusSignal
	}
	return 0
}

func (x *Monitor) GetStatusVideo() int64 {
	if x != nil {
		return x.StatusVideo
	}
	return 0
}

func (x *Monitor) GetStatusAudio() int64 {
	if x != nil {
		return x.StatusAudio
	}
	return 0
}

func (x *Monitor) GetSignalMonitor() bool {
	if x != nil {
		return x.SignalMonitor
	}
	return false
}

func (x *Monitor) GetVideoMonitor() bool {
	if x != nil {
		return x.VideoMonitor
	}
	return false
}

func (x *Monitor) GetAudioMonitor() bool {
	if x != nil {
		return x.AudioMonitor
	}
	return false
}

func (x *Monitor) GetIsEnable() bool {
	if x != nil {
		return x.IsEnable
	}
	return false
}

func (x *Monitor) GetDateUpdate() string {
	if x != nil {
		return x.DateUpdate
	}
	return ""
}

func (x *Monitor) GetAgentId() int64 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *Monitor) GetProfileId() int64 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

func (x *Monitor) GetStatusId() int64 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

//*
// Represents the params to identify monitor.
type MonitorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StatusSignal  int64  `protobuf:"varint,2,opt,name=status_signal,proto3" json:"status_signal,omitempty"`
	StatusVideo   int64  `protobuf:"varint,3,opt,name=status_video,proto3" json:"status_video,omitempty"`
	StatusAudio   int64  `protobuf:"varint,4,opt,name=status_audio,proto3" json:"status_audio,omitempty"`
	SignalMonitor bool   `protobuf:"varint,5,opt,name=signal_monitor,proto3" json:"signal_monitor,omitempty"`
	VideoMonitor  bool   `protobuf:"varint,6,opt,name=video_monitor,proto3" json:"video_monitor,omitempty"`
	AudioMonitor  bool   `protobuf:"varint,7,opt,name=audio_monitor,proto3" json:"audio_monitor,omitempty"`
	IsEnable      bool   `protobuf:"varint,8,opt,name=is_enable,proto3" json:"is_enable,omitempty"`
	DateUpdate    string `protobuf:"bytes,9,opt,name=date_update,proto3" json:"date_update,omitempty"`
	AgentId       int64  `protobuf:"varint,10,opt,name=agent_id,proto3" json:"agent_id,omitempty"`
	ProfileId     int64  `protobuf:"varint,11,opt,name=profile_id,proto3" json:"profile_id,omitempty"`
	StatusId      int64  `protobuf:"varint,12,opt,name=status_id,proto3" json:"status_id,omitempty"`
}

func (x *MonitorRequest) Reset() {
	*x = MonitorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorRequest) ProtoMessage() {}

func (x *MonitorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorRequest.ProtoReflect.Descriptor instead.
func (*MonitorRequest) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *MonitorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MonitorRequest) GetStatusSignal() int64 {
	if x != nil {
		return x.StatusSignal
	}
	return 0
}

func (x *MonitorRequest) GetStatusVideo() int64 {
	if x != nil {
		return x.StatusVideo
	}
	return 0
}

func (x *MonitorRequest) GetStatusAudio() int64 {
	if x != nil {
		return x.StatusAudio
	}
	return 0
}

func (x *MonitorRequest) GetSignalMonitor() bool {
	if x != nil {
		return x.SignalMonitor
	}
	return false
}

func (x *MonitorRequest) GetVideoMonitor() bool {
	if x != nil {
		return x.VideoMonitor
	}
	return false
}

func (x *MonitorRequest) GetAudioMonitor() bool {
	if x != nil {
		return x.AudioMonitor
	}
	return false
}

func (x *MonitorRequest) GetIsEnable() bool {
	if x != nil {
		return x.IsEnable
	}
	return false
}

func (x *MonitorRequest) GetDateUpdate() string {
	if x != nil {
		return x.DateUpdate
	}
	return ""
}

func (x *MonitorRequest) GetAgentId() int64 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *MonitorRequest) GetProfileId() int64 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

func (x *MonitorRequest) GetStatusId() int64 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

type MonitorFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AgentId   int64 `protobuf:"varint,2,opt,name=agent_id,proto3" json:"agent_id,omitempty"`
	ProfileId int64 `protobuf:"varint,3,opt,name=profile_id,proto3" json:"profile_id,omitempty"`
}

func (x *MonitorFilter) Reset() {
	*x = MonitorFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorFilter) ProtoMessage() {}

func (x *MonitorFilter) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorFilter.ProtoReflect.Descriptor instead.
func (*MonitorFilter) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *MonitorFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MonitorFilter) GetAgentId() int64 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *MonitorFilter) GetProfileId() int64 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

type MonitorDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MonitorDelete) Reset() {
	*x = MonitorDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorDelete) ProtoMessage() {}

func (x *MonitorDelete) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorDelete.ProtoReflect.Descriptor instead.
func (*MonitorDelete) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{3}
}

func (x *MonitorDelete) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MonitorGetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MonitorGetAll) Reset() {
	*x = MonitorGetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorGetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorGetAll) ProtoMessage() {}

func (x *MonitorGetAll) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorGetAll.ProtoReflect.Descriptor instead.
func (*MonitorGetAll) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{4}
}

type MonitorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Status
	Status MonitorResponseStatus `protobuf:"varint,1,opt,name=Status,json=status,proto3,enum=proto.MonitorResponseStatus" json:"Status,omitempty"`
	//*
	// Slice of agent object
	Monitors []*Monitor `protobuf:"bytes,2,rep,name=Monitors,json=data,proto3" json:"Monitors,omitempty"`
}

func (x *MonitorResponse) Reset() {
	*x = MonitorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorResponse) ProtoMessage() {}

func (x *MonitorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorResponse.ProtoReflect.Descriptor instead.
func (*MonitorResponse) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{5}
}

func (x *MonitorResponse) GetStatus() MonitorResponseStatus {
	if x != nil {
		return x.Status
	}
	return MonitorResponseStatus_FAIL
}

func (x *MonitorResponse) GetMonitors() []*Monitor {
	if x != nil {
		return x.Monitors
	}
	return nil
}

var File_monitor_proto protoreflect.FileDescriptor

var file_monitor_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x03, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x12, 0x26, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x22, 0x9c,
	0x03, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12,
	0x26, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x22, 0x5b, 0x0a,
	0x0d, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x22, 0x6f, 0x0a, 0x0f,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2e, 0x0a,
	0x15, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x32, 0xac, 0x02,
	0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x04, 0x47, 0x65, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x1a, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x6e, 0x64, 0x2f, 0x69, 0x70, 0x74, 0x76, 0x2d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_proto_rawDescOnce sync.Once
	file_monitor_proto_rawDescData = file_monitor_proto_rawDesc
)

func file_monitor_proto_rawDescGZIP() []byte {
	file_monitor_proto_rawDescOnce.Do(func() {
		file_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_proto_rawDescData)
	})
	return file_monitor_proto_rawDescData
}

var file_monitor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_monitor_proto_goTypes = []interface{}{
	(MonitorResponseStatus)(0), // 0: proto.MonitorResponseStatus
	(*Monitor)(nil),            // 1: proto.Monitor
	(*MonitorRequest)(nil),     // 2: proto.MonitorRequest
	(*MonitorFilter)(nil),      // 3: proto.MonitorFilter
	(*MonitorDelete)(nil),      // 4: proto.MonitorDelete
	(*MonitorGetAll)(nil),      // 5: proto.MonitorGetAll
	(*MonitorResponse)(nil),    // 6: proto.MonitorResponse
}
var file_monitor_proto_depIdxs = []int32{
	0, // 0: proto.MonitorResponse.Status:type_name -> proto.MonitorResponseStatus
	1, // 1: proto.MonitorResponse.Monitors:type_name -> proto.Monitor
	5, // 2: proto.MonitorService.Gets:input_type -> proto.MonitorGetAll
	3, // 3: proto.MonitorService.Get:input_type -> proto.MonitorFilter
	2, // 4: proto.MonitorService.Add:input_type -> proto.MonitorRequest
	2, // 5: proto.MonitorService.Update:input_type -> proto.MonitorRequest
	4, // 6: proto.MonitorService.Delete:input_type -> proto.MonitorDelete
	6, // 7: proto.MonitorService.Gets:output_type -> proto.MonitorResponse
	6, // 8: proto.MonitorService.Get:output_type -> proto.MonitorResponse
	6, // 9: proto.MonitorService.Add:output_type -> proto.MonitorResponse
	6, // 10: proto.MonitorService.Update:output_type -> proto.MonitorResponse
	6, // 11: proto.MonitorService.Delete:output_type -> proto.MonitorResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_monitor_proto_init() }
func file_monitor_proto_init() {
	if File_monitor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorDelete); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorGetAll); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitor_proto_goTypes,
		DependencyIndexes: file_monitor_proto_depIdxs,
		EnumInfos:         file_monitor_proto_enumTypes,
		MessageInfos:      file_monitor_proto_msgTypes,
	}.Build()
	File_monitor_proto = out.File
	file_monitor_proto_rawDesc = nil
	file_monitor_proto_goTypes = nil
	file_monitor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MonitorServiceClient is the client API for MonitorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorServiceClient interface {
	//*
	// Get All Monitor
	Gets(ctx context.Context, in *MonitorGetAll, opts ...grpc.CallOption) (*MonitorResponse, error)
	//*
	// Get Monitor
	Get(ctx context.Context, in *MonitorFilter, opts ...grpc.CallOption) (*MonitorResponse, error)
	//*
	// Add Monitor
	Add(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorResponse, error)
	//*
	// Update Monitor
	Update(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorResponse, error)
	//*
	// Delete Monitor
	Delete(ctx context.Context, in *MonitorDelete, opts ...grpc.CallOption) (*MonitorResponse, error)
}

type monitorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorServiceClient(cc grpc.ClientConnInterface) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) Gets(ctx context.Context, in *MonitorGetAll, opts ...grpc.CallOption) (*MonitorResponse, error) {
	out := new(MonitorResponse)
	err := c.cc.Invoke(ctx, "/proto.MonitorService/Gets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) Get(ctx context.Context, in *MonitorFilter, opts ...grpc.CallOption) (*MonitorResponse, error) {
	out := new(MonitorResponse)
	err := c.cc.Invoke(ctx, "/proto.MonitorService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) Add(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorResponse, error) {
	out := new(MonitorResponse)
	err := c.cc.Invoke(ctx, "/proto.MonitorService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) Update(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorResponse, error) {
	out := new(MonitorResponse)
	err := c.cc.Invoke(ctx, "/proto.MonitorService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) Delete(ctx context.Context, in *MonitorDelete, opts ...grpc.CallOption) (*MonitorResponse, error) {
	out := new(MonitorResponse)
	err := c.cc.Invoke(ctx, "/proto.MonitorService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServiceServer is the server API for MonitorService service.
type MonitorServiceServer interface {
	//*
	// Get All Monitor
	Gets(context.Context, *MonitorGetAll) (*MonitorResponse, error)
	//*
	// Get Monitor
	Get(context.Context, *MonitorFilter) (*MonitorResponse, error)
	//*
	// Add Monitor
	Add(context.Context, *MonitorRequest) (*MonitorResponse, error)
	//*
	// Update Monitor
	Update(context.Context, *MonitorRequest) (*MonitorResponse, error)
	//*
	// Delete Monitor
	Delete(context.Context, *MonitorDelete) (*MonitorResponse, error)
}

// UnimplementedMonitorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorServiceServer struct {
}

func (*UnimplementedMonitorServiceServer) Gets(context.Context, *MonitorGetAll) (*MonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gets not implemented")
}
func (*UnimplementedMonitorServiceServer) Get(context.Context, *MonitorFilter) (*MonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMonitorServiceServer) Add(context.Context, *MonitorRequest) (*MonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedMonitorServiceServer) Update(context.Context, *MonitorRequest) (*MonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedMonitorServiceServer) Delete(context.Context, *MonitorDelete) (*MonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterMonitorServiceServer(s *grpc.Server, srv MonitorServiceServer) {
	s.RegisterService(&_MonitorService_serviceDesc, srv)
}

func _MonitorService_Gets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorGetAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).Gets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MonitorService/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).Gets(ctx, req.(*MonitorGetAll))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MonitorService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).Get(ctx, req.(*MonitorFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MonitorService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).Add(ctx, req.(*MonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MonitorService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).Update(ctx, req.(*MonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MonitorService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).Delete(ctx, req.(*MonitorDelete))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MonitorService",
	HandlerType: (*MonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Gets",
			Handler:    _MonitorService_Gets_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MonitorService_Get_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _MonitorService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MonitorService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MonitorService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor.proto",
}
