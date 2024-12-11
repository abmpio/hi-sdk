// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: hi.proto

package hi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthCheckResponse_ServiceStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServiceStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServiceStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServiceStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServiceStatus = 3
)

// Enum value maps for HealthCheckResponse_ServiceStatus.
var (
	HealthCheckResponse_ServiceStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
		3: "SERVICE_UNKNOWN",
	}
	HealthCheckResponse_ServiceStatus_value = map[string]int32{
		"UNKNOWN":         0,
		"SERVING":         1,
		"NOT_SERVING":     2,
		"SERVICE_UNKNOWN": 3,
	}
)

func (x HealthCheckResponse_ServiceStatus) Enum() *HealthCheckResponse_ServiceStatus {
	p := new(HealthCheckResponse_ServiceStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_hi_proto_enumTypes[0].Descriptor()
}

func (HealthCheckResponse_ServiceStatus) Type() protoreflect.EnumType {
	return &file_hi_proto_enumTypes[0]
}

func (x HealthCheckResponse_ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServiceStatus.Descriptor instead.
func (HealthCheckResponse_ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{2, 0}
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_hi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	mi := &file_hi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCheckResponse_ServiceStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.HealthCheckResponse_ServiceStatus" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	mi := &file_hi_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{2}
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServiceStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheckResponse_UNKNOWN
}

// 请求消息类型
type FindListByCodeTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App      string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`           // 应用标识
	CodeType string `protobuf:"bytes,2,opt,name=codeType,proto3" json:"codeType,omitempty"` // 代码类型
}

func (x *FindListByCodeTypeRequest) Reset() {
	*x = FindListByCodeTypeRequest{}
	mi := &file_hi_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindListByCodeTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindListByCodeTypeRequest) ProtoMessage() {}

func (x *FindListByCodeTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindListByCodeTypeRequest.ProtoReflect.Descriptor instead.
func (*FindListByCodeTypeRequest) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{3}
}

func (x *FindListByCodeTypeRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *FindListByCodeTypeRequest) GetCodeType() string {
	if x != nil {
		return x.CodeType
	}
	return ""
}

type FindOneByCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App      string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`           // 应用标识
	CodeType string `protobuf:"bytes,2,opt,name=codeType,proto3" json:"codeType,omitempty"` // 代码类型
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`         // 代码
}

func (x *FindOneByCodeRequest) Reset() {
	*x = FindOneByCodeRequest{}
	mi := &file_hi_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneByCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneByCodeRequest) ProtoMessage() {}

func (x *FindOneByCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneByCodeRequest.ProtoReflect.Descriptor instead.
func (*FindOneByCodeRequest) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneByCodeRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *FindOneByCodeRequest) GetCodeType() string {
	if x != nil {
		return x.CodeType
	}
	return ""
}

func (x *FindOneByCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// 响应消息类型
type CodeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`           // 代码
	ShortName string `protobuf:"bytes,2,opt,name=shortName,proto3" json:"shortName,omitempty"` // 简称
	FullName  string `protobuf:"bytes,3,opt,name=fullName,proto3" json:"fullName,omitempty"`   // 全称
	Memo      string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`           // 备注
}

func (x *CodeValue) Reset() {
	*x = CodeValue{}
	mi := &file_hi_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CodeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeValue) ProtoMessage() {}

func (x *CodeValue) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeValue.ProtoReflect.Descriptor instead.
func (*CodeValue) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{5}
}

func (x *CodeValue) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CodeValue) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *CodeValue) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CodeValue) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type FindListByCodeTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeValues []*CodeValue `protobuf:"bytes,1,rep,name=codeValues,proto3" json:"codeValues,omitempty"` // 返回的代码值列表
}

func (x *FindListByCodeTypeResponse) Reset() {
	*x = FindListByCodeTypeResponse{}
	mi := &file_hi_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindListByCodeTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindListByCodeTypeResponse) ProtoMessage() {}

func (x *FindListByCodeTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindListByCodeTypeResponse.ProtoReflect.Descriptor instead.
func (*FindListByCodeTypeResponse) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{6}
}

func (x *FindListByCodeTypeResponse) GetCodeValues() []*CodeValue {
	if x != nil {
		return x.CodeValues
	}
	return nil
}

type FindOneByCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeValue *CodeValue `protobuf:"bytes,1,opt,name=codeValue,proto3" json:"codeValue,omitempty"` // 单个代码值
}

func (x *FindOneByCodeResponse) Reset() {
	*x = FindOneByCodeResponse{}
	mi := &file_hi_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneByCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneByCodeResponse) ProtoMessage() {}

func (x *FindOneByCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hi_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneByCodeResponse.ProtoReflect.Descriptor instead.
func (*FindOneByCodeResponse) Descriptor() ([]byte, []int) {
	return file_hi_proto_rawDescGZIP(), []int{7}
}

func (x *FindOneByCodeResponse) GetCodeValue() *CodeValue {
	if x != nil {
		return x.CodeValue
	}
	return nil
}

var File_hi_proto protoreflect.FileDescriptor

var file_hi_proto_rawDesc = []byte{
	0x0a, 0x08, 0x68, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e,
	0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xa8,
	0x01, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x22, 0x49, 0x0a, 0x19, 0x46, 0x69, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x6d,
	0x0a, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x4e, 0x0a,
	0x1a, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x63,
	0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x47, 0x0a,
	0x15, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x64,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x4a, 0x0a, 0x02, 0x48, 0x69, 0x12, 0x44, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65,
	0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x68, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hi_proto_rawDescOnce sync.Once
	file_hi_proto_rawDescData = file_hi_proto_rawDesc
)

func file_hi_proto_rawDescGZIP() []byte {
	file_hi_proto_rawDescOnce.Do(func() {
		file_hi_proto_rawDescData = protoimpl.X.CompressGZIP(file_hi_proto_rawDescData)
	})
	return file_hi_proto_rawDescData
}

var file_hi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_hi_proto_goTypes = []any{
	(HealthCheckResponse_ServiceStatus)(0), // 0: proto.HealthCheckResponse.ServiceStatus
	(*BaseResponse)(nil),                   // 1: proto.BaseResponse
	(*HealthCheckRequest)(nil),             // 2: proto.HealthCheckRequest
	(*HealthCheckResponse)(nil),            // 3: proto.HealthCheckResponse
	(*FindListByCodeTypeRequest)(nil),      // 4: proto.FindListByCodeTypeRequest
	(*FindOneByCodeRequest)(nil),           // 5: proto.FindOneByCodeRequest
	(*CodeValue)(nil),                      // 6: proto.CodeValue
	(*FindListByCodeTypeResponse)(nil),     // 7: proto.FindListByCodeTypeResponse
	(*FindOneByCodeResponse)(nil),          // 8: proto.FindOneByCodeResponse
}
var file_hi_proto_depIdxs = []int32{
	0, // 0: proto.HealthCheckResponse.status:type_name -> proto.HealthCheckResponse.ServiceStatus
	6, // 1: proto.FindListByCodeTypeResponse.codeValues:type_name -> proto.CodeValue
	6, // 2: proto.FindOneByCodeResponse.codeValue:type_name -> proto.CodeValue
	2, // 3: proto.Hi.HealthCheck:input_type -> proto.HealthCheckRequest
	4, // 4: proto.CodeValueService.FindListByCodeType:input_type -> proto.FindListByCodeTypeRequest
	5, // 5: proto.CodeValueService.FindOneByCode:input_type -> proto.FindOneByCodeRequest
	3, // 6: proto.Hi.HealthCheck:output_type -> proto.HealthCheckResponse
	7, // 7: proto.CodeValueService.FindListByCodeType:output_type -> proto.FindListByCodeTypeResponse
	8, // 8: proto.CodeValueService.FindOneByCode:output_type -> proto.FindOneByCodeResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hi_proto_init() }
func file_hi_proto_init() {
	if File_hi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_hi_proto_goTypes,
		DependencyIndexes: file_hi_proto_depIdxs,
		EnumInfos:         file_hi_proto_enumTypes,
		MessageInfos:      file_hi_proto_msgTypes,
	}.Build()
	File_hi_proto = out.File
	file_hi_proto_rawDesc = nil
	file_hi_proto_goTypes = nil
	file_hi_proto_depIdxs = nil
}
