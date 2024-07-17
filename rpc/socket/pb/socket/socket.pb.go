// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.11
// source: socket.proto

package socket

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string     `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	ErrMsg string     `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	Code   string     `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string     `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Data   *anypb.Any `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *Result) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *Result) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Result) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Result) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type ReqPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *ReqPing) Reset() {
	*x = ReqPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPing) ProtoMessage() {}

func (x *ReqPing) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPing.ProtoReflect.Descriptor instead.
func (*ReqPing) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1}
}

func (x *ReqPing) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type ResPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ResPing) Reset() {
	*x = ResPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPing) ProtoMessage() {}

func (x *ResPing) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPing.ProtoReflect.Descriptor instead.
func (*ResPing) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{2}
}

func (x *ResPing) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// *****************广播消息处理的结构体******************
type ReqBroadcastMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      int32      `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Operate      int32      `protobuf:"varint,2,opt,name=operate,proto3" json:"operate,omitempty"`
	Method       string     `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	AuthToken    string     `protobuf:"bytes,4,opt,name=authToken,proto3" json:"authToken,omitempty"`
	RoomId       int64      `protobuf:"varint,5,opt,name=roomId,proto3" json:"roomId,omitempty"`
	FromUserId   int64      `protobuf:"varint,6,opt,name=fromUserId,proto3" json:"fromUserId,omitempty"`
	FromUserName string     `protobuf:"bytes,7,opt,name=fromUserName,proto3" json:"fromUserName,omitempty"`
	ToClientId   int64      `protobuf:"varint,8,opt,name=toClientId,proto3" json:"toClientId,omitempty"`
	ToUserId     int64      `protobuf:"varint,9,opt,name=toUserId,proto3" json:"toUserId,omitempty"`
	ToUserName   string     `protobuf:"bytes,10,opt,name=ToUserName,proto3" json:"ToUserName,omitempty"`
	Event        *BodyEvent `protobuf:"bytes,11,opt,name=event,proto3" json:"event,omitempty"`
	Extend       string     `protobuf:"bytes,12,opt,name=extend,proto3" json:"extend,omitempty"`
}

func (x *ReqBroadcastMsg) Reset() {
	*x = ReqBroadcastMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqBroadcastMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqBroadcastMsg) ProtoMessage() {}

func (x *ReqBroadcastMsg) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqBroadcastMsg.ProtoReflect.Descriptor instead.
func (*ReqBroadcastMsg) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{3}
}

func (x *ReqBroadcastMsg) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ReqBroadcastMsg) GetOperate() int32 {
	if x != nil {
		return x.Operate
	}
	return 0
}

func (x *ReqBroadcastMsg) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ReqBroadcastMsg) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *ReqBroadcastMsg) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ReqBroadcastMsg) GetFromUserId() int64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *ReqBroadcastMsg) GetFromUserName() string {
	if x != nil {
		return x.FromUserName
	}
	return ""
}

func (x *ReqBroadcastMsg) GetToClientId() int64 {
	if x != nil {
		return x.ToClientId
	}
	return 0
}

func (x *ReqBroadcastMsg) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *ReqBroadcastMsg) GetToUserName() string {
	if x != nil {
		return x.ToUserName
	}
	return ""
}

func (x *ReqBroadcastMsg) GetEvent() *BodyEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ReqBroadcastMsg) GetExtend() string {
	if x != nil {
		return x.Extend
	}
	return ""
}

type BodyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *anypb.Any `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Data   *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BodyEvent) Reset() {
	*x = BodyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyEvent) ProtoMessage() {}

func (x *BodyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyEvent.ProtoReflect.Descriptor instead.
func (*BodyEvent) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{4}
}

func (x *BodyEvent) GetParams() *anypb.Any {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *BodyEvent) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// 加入聊天房间params内容
type EventParamsLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   int64  `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	ClientId int64  `protobuf:"varint,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	UserId   int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName string `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *EventParamsLogin) Reset() {
	*x = EventParamsLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventParamsLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventParamsLogin) ProtoMessage() {}

func (x *EventParamsLogin) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventParamsLogin.ProtoReflect.Descriptor instead.
func (*EventParamsLogin) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{5}
}

func (x *EventParamsLogin) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *EventParamsLogin) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *EventParamsLogin) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *EventParamsLogin) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// 普通消息params内容
type EventParamsNormal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *EventParamsNormal) Reset() {
	*x = EventParamsNormal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventParamsNormal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventParamsNormal) ProtoMessage() {}

func (x *EventParamsNormal) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventParamsNormal.ProtoReflect.Descriptor instead.
func (*EventParamsNormal) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{6}
}

func (x *EventParamsNormal) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 加入聊天房间返回data
type EventDataLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   int64  `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	ClientId int64  `protobuf:"varint,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	UserId   int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName string `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *EventDataLogin) Reset() {
	*x = EventDataLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDataLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDataLogin) ProtoMessage() {}

func (x *EventDataLogin) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDataLogin.ProtoReflect.Descriptor instead.
func (*EventDataLogin) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{7}
}

func (x *EventDataLogin) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *EventDataLogin) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *EventDataLogin) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *EventDataLogin) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// 普通消息data内容
type EventDataNormal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId       int64  `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	FromUserId   int64  `protobuf:"varint,2,opt,name=fromUserId,proto3" json:"fromUserId,omitempty"`
	FromUserName string `protobuf:"bytes,3,opt,name=fromUserName,proto3" json:"fromUserName,omitempty"`
	Message      string `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *EventDataNormal) Reset() {
	*x = EventDataNormal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDataNormal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDataNormal) ProtoMessage() {}

func (x *EventDataNormal) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDataNormal.ProtoReflect.Descriptor instead.
func (*EventDataNormal) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{8}
}

func (x *EventDataNormal) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *EventDataNormal) GetFromUserId() int64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *EventDataNormal) GetFromUserName() string {
	if x != nil {
		return x.FromUserName
	}
	return ""
}

func (x *EventDataNormal) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_socket_proto protoreflect.FileDescriptor

var file_socket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x22, 0xf4, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x42, 0x6f, 0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x22, 0x63, 0x0a, 0x09, 0x42, 0x6f, 0x64, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a,
	0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x78, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x30, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0f, 0x2e,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x69, 0x6e, 0x67, 0x32, 0xbb,
	0x01, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0e,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17,
	0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x0e, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3a, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x17, 0x2e, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x1a, 0x0e, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x4f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x0e, 0x2e, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_socket_proto_rawDescOnce sync.Once
	file_socket_proto_rawDescData = file_socket_proto_rawDesc
)

func file_socket_proto_rawDescGZIP() []byte {
	file_socket_proto_rawDescOnce.Do(func() {
		file_socket_proto_rawDescData = protoimpl.X.CompressGZIP(file_socket_proto_rawDescData)
	})
	return file_socket_proto_rawDescData
}

var file_socket_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_socket_proto_goTypes = []interface{}{
	(*Result)(nil),            // 0: socket.Result
	(*ReqPing)(nil),           // 1: socket.ReqPing
	(*ResPing)(nil),           // 2: socket.ResPing
	(*ReqBroadcastMsg)(nil),   // 3: socket.ReqBroadcastMsg
	(*BodyEvent)(nil),         // 4: socket.BodyEvent
	(*EventParamsLogin)(nil),  // 5: socket.EventParamsLogin
	(*EventParamsNormal)(nil), // 6: socket.EventParamsNormal
	(*EventDataLogin)(nil),    // 7: socket.EventDataLogin
	(*EventDataNormal)(nil),   // 8: socket.EventDataNormal
	(*anypb.Any)(nil),         // 9: google.protobuf.Any
}
var file_socket_proto_depIdxs = []int32{
	9, // 0: socket.Result.data:type_name -> google.protobuf.Any
	4, // 1: socket.ReqBroadcastMsg.event:type_name -> socket.BodyEvent
	9, // 2: socket.BodyEvent.params:type_name -> google.protobuf.Any
	9, // 3: socket.BodyEvent.data:type_name -> google.protobuf.Any
	1, // 4: socket.Ping.Ping:input_type -> socket.ReqPing
	3, // 5: socket.Broadcast.BroadcastLogin:input_type -> socket.ReqBroadcastMsg
	3, // 6: socket.Broadcast.BroadcastNormal:input_type -> socket.ReqBroadcastMsg
	3, // 7: socket.Broadcast.BroadcastOut:input_type -> socket.ReqBroadcastMsg
	2, // 8: socket.Ping.Ping:output_type -> socket.ResPing
	0, // 9: socket.Broadcast.BroadcastLogin:output_type -> socket.Result
	0, // 10: socket.Broadcast.BroadcastNormal:output_type -> socket.Result
	0, // 11: socket.Broadcast.BroadcastOut:output_type -> socket.Result
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_socket_proto_init() }
func file_socket_proto_init() {
	if File_socket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_socket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_socket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPing); i {
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
		file_socket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResPing); i {
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
		file_socket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqBroadcastMsg); i {
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
		file_socket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyEvent); i {
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
		file_socket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventParamsLogin); i {
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
		file_socket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventParamsNormal); i {
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
		file_socket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDataLogin); i {
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
		file_socket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDataNormal); i {
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
			RawDescriptor: file_socket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_socket_proto_goTypes,
		DependencyIndexes: file_socket_proto_depIdxs,
		MessageInfos:      file_socket_proto_msgTypes,
	}.Build()
	File_socket_proto = out.File
	file_socket_proto_rawDesc = nil
	file_socket_proto_goTypes = nil
	file_socket_proto_depIdxs = nil
}
