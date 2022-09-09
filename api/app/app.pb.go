// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: app/app.proto

package app

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	kiae "github.com/kiaedev/kiae/api/kiae"
	project "github.com/kiaedev/kiae/api/project"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Size int32

const (
	Size_SIZE_NANO    Size = 0
	Size_SIZE_MIRCO   Size = 1
	Size_SIZE_MINI    Size = 2
	Size_SIZE_SMALL   Size = 3
	Size_SIZE_MEDIUM  Size = 4
	Size_SIZE_LARGE   Size = 5
	Size_SIZE_XLARGE  Size = 6
	Size_SIZE_XXLARGE Size = 7
)

// Enum value maps for Size.
var (
	Size_name = map[int32]string{
		0: "SIZE_NANO",
		1: "SIZE_MIRCO",
		2: "SIZE_MINI",
		3: "SIZE_SMALL",
		4: "SIZE_MEDIUM",
		5: "SIZE_LARGE",
		6: "SIZE_XLARGE",
		7: "SIZE_XXLARGE",
	}
	Size_value = map[string]int32{
		"SIZE_NANO":    0,
		"SIZE_MIRCO":   1,
		"SIZE_MINI":    2,
		"SIZE_SMALL":   3,
		"SIZE_MEDIUM":  4,
		"SIZE_LARGE":   5,
		"SIZE_XLARGE":  6,
		"SIZE_XXLARGE": 7,
	}
)

func (x Size) Enum() *Size {
	p := new(Size)
	*p = x
	return p
}

func (x Size) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Size) Descriptor() protoreflect.EnumDescriptor {
	return file_app_app_proto_enumTypes[0].Descriptor()
}

func (Size) Type() protoreflect.EnumType {
	return &file_app_app_proto_enumTypes[0]
}

func (x Size) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Size.Descriptor instead.
func (Size) EnumDescriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_STATUS_CREATED Status = 0
	Status_STATUS_RUNNING Status = 1
	Status_STATUS_STOPPED Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_CREATED",
		1: "STATUS_RUNNING",
		2: "STATUS_STOPPED",
	}
	Status_value = map[string]int32{
		"STATUS_CREATED": 0,
		"STATUS_RUNNING": 1,
		"STATUS_STOPPED": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_app_app_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_app_app_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{1}
}

type ActionPayload_Action int32

const (
	ActionPayload_START   ActionPayload_Action = 0
	ActionPayload_STOP    ActionPayload_Action = 1
	ActionPayload_RESTART ActionPayload_Action = 2
)

// Enum value maps for ActionPayload_Action.
var (
	ActionPayload_Action_name = map[int32]string{
		0: "START",
		1: "STOP",
		2: "RESTART",
	}
	ActionPayload_Action_value = map[string]int32{
		"START":   0,
		"STOP":    1,
		"RESTART": 2,
	}
)

func (x ActionPayload_Action) Enum() *ActionPayload_Action {
	p := new(ActionPayload_Action)
	*p = x
	return p
}

func (x ActionPayload_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionPayload_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_app_app_proto_enumTypes[2].Descriptor()
}

func (ActionPayload_Action) Type() protoreflect.EnumType {
	return &file_app_app_proto_enumTypes[2]
}

func (x ActionPayload_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionPayload_Action.Descriptor instead.
func (ActionPayload_Action) EnumDescriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{4, 0}
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"` // @gotags: bson:"_id,omitempty"
	Pid                   string                   `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Env                   string                   `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	Name                  string                   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Image                 string                   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Ports                 []*project.Port          `protobuf:"bytes,6,rep,name=ports,proto3" json:"ports,omitempty"`
	Status                Status                   `protobuf:"varint,7,opt,name=status,proto3,enum=app.Status" json:"status,omitempty"`
	Size                  Size                     `protobuf:"varint,8,opt,name=size,proto3,enum=app.Size" json:"size,omitempty"`
	Replicas              uint32                   `protobuf:"varint,9,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Annotations           map[string]string        `protobuf:"bytes,10,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LivenessProbeEnabled  bool                     `protobuf:"varint,15,opt,name=liveness_probe_enabled,json=livenessProbeEnabled,proto3" json:"liveness_probe_enabled,omitempty"`
	ReadinessProbeEnabled bool                     `protobuf:"varint,16,opt,name=readiness_probe_enabled,json=readinessProbeEnabled,proto3" json:"readiness_probe_enabled,omitempty"`
	Configs               []*project.Configuration `protobuf:"bytes,17,rep,name=configs,proto3" json:"configs,omitempty"`
	Middlewares           []*project.Middleware    `protobuf:"bytes,18,rep,name=middlewares,proto3" json:"middlewares,omitempty"`
	CreatedAt             *timestamppb.Timestamp   `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at,omitempty"` // @gotags: bson:"created_at,omitempty"
	CreatedBy             *timestamppb.Timestamp   `protobuf:"bytes,102,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty" bson:"created_by,omitempty"` // @gotags: bson:"created_by,omitempty"
	UpdatedAt             *timestamppb.Timestamp   `protobuf:"bytes,103,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bson:"updated_at,omitempty"` // @gotags: bson:"updated_at,omitempty"
	UpdatedBy             *timestamppb.Timestamp   `protobuf:"bytes,104,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty" bson:"updated_by,omitempty"` // @gotags: bson:"updated_by,omitempty"
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Application) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *Application) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Application) GetPorts() []*project.Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Application) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_CREATED
}

func (x *Application) GetSize() Size {
	if x != nil {
		return x.Size
	}
	return Size_SIZE_NANO
}

func (x *Application) GetReplicas() uint32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Application) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Application) GetLivenessProbeEnabled() bool {
	if x != nil {
		return x.LivenessProbeEnabled
	}
	return false
}

func (x *Application) GetReadinessProbeEnabled() bool {
	if x != nil {
		return x.ReadinessProbeEnabled
	}
	return false
}

func (x *Application) GetConfigs() []*project.Configuration {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *Application) GetMiddlewares() []*project.Middleware {
	if x != nil {
		return x.Middlewares
	}
	return nil
}

func (x *Application) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Application) GetCreatedBy() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *Application) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Application) GetUpdatedBy() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Application `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetItems() []*Application {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload    *Application           `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetPayload() *Application {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *UpdateRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type ActionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Action ActionPayload_Action `protobuf:"varint,2,opt,name=action,proto3,enum=app.ActionPayload_Action" json:"action,omitempty"`
}

func (x *ActionPayload) Reset() {
	*x = ActionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionPayload) ProtoMessage() {}

func (x *ActionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionPayload.ProtoReflect.Descriptor instead.
func (*ActionPayload) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{4}
}

func (x *ActionPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActionPayload) GetAction() ActionPayload_Action {
	if x != nil {
		return x.Action
	}
	return ActionPayload_START
}

var File_app_app_proto protoreflect.FileDescriptor

var file_app_app_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x6b, 0x69, 0x61, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x06, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x18, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x10, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x03, 0x18, 0x80, 0x01, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x34, 0x0a, 0x16, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x35, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x18, 0x12,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x1f, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x78, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x7e, 0x0a, 0x0d, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x2a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x2a, 0x88, 0x01, 0x0a, 0x04,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4e, 0x41, 0x4e,
	0x4f, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4d, 0x49, 0x52, 0x43,
	0x4f, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x49,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x53, 0x4d, 0x41, 0x4c, 0x4c,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55,
	0x4d, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4c, 0x41, 0x52, 0x47,
	0x45, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x58, 0x4c, 0x41, 0x52,
	0x47, 0x45, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x58, 0x58, 0x4c,
	0x41, 0x52, 0x47, 0x45, 0x10, 0x07, 0x2a, 0x44, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52,
	0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x32, 0x85, 0x04, 0x0a,
	0x0a, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x12, 0x45,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4a, 0x1a, 0x19,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5a, 0x24, 0x32, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x69, 0x64, 0x7d, 0x3a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6b, 0x69, 0x61, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x44, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0f,
	0x2e, 0x6b, 0x69, 0x61, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a, 0x08,
	0x44, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x10, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x3a, 0x01, 0x2a, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x64, 0x65, 0x76, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_app_proto_rawDescOnce sync.Once
	file_app_app_proto_rawDescData = file_app_app_proto_rawDesc
)

func file_app_app_proto_rawDescGZIP() []byte {
	file_app_app_proto_rawDescOnce.Do(func() {
		file_app_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_app_proto_rawDescData)
	})
	return file_app_app_proto_rawDescData
}

var file_app_app_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_app_app_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_app_app_proto_goTypes = []interface{}{
	(Size)(0),                     // 0: app.Size
	(Status)(0),                   // 1: app.Status
	(ActionPayload_Action)(0),     // 2: app.ActionPayload.Action
	(*Application)(nil),           // 3: app.Application
	(*ListRequest)(nil),           // 4: app.ListRequest
	(*ListResponse)(nil),          // 5: app.ListResponse
	(*UpdateRequest)(nil),         // 6: app.UpdateRequest
	(*ActionPayload)(nil),         // 7: app.ActionPayload
	nil,                           // 8: app.Application.AnnotationsEntry
	(*project.Port)(nil),          // 9: project.Port
	(*project.Configuration)(nil), // 10: project.Configuration
	(*project.Middleware)(nil),    // 11: project.Middleware
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 13: google.protobuf.FieldMask
	(*kiae.IdRequest)(nil),        // 14: kiae.IdRequest
	(*emptypb.Empty)(nil),         // 15: google.protobuf.Empty
}
var file_app_app_proto_depIdxs = []int32{
	9,  // 0: app.Application.ports:type_name -> project.Port
	1,  // 1: app.Application.status:type_name -> app.Status
	0,  // 2: app.Application.size:type_name -> app.Size
	8,  // 3: app.Application.annotations:type_name -> app.Application.AnnotationsEntry
	10, // 4: app.Application.configs:type_name -> project.Configuration
	11, // 5: app.Application.middlewares:type_name -> project.Middleware
	12, // 6: app.Application.created_at:type_name -> google.protobuf.Timestamp
	12, // 7: app.Application.created_by:type_name -> google.protobuf.Timestamp
	12, // 8: app.Application.updated_at:type_name -> google.protobuf.Timestamp
	12, // 9: app.Application.updated_by:type_name -> google.protobuf.Timestamp
	3,  // 10: app.ListResponse.items:type_name -> app.Application
	3,  // 11: app.UpdateRequest.payload:type_name -> app.Application
	13, // 12: app.UpdateRequest.update_mask:type_name -> google.protobuf.FieldMask
	2,  // 13: app.ActionPayload.action:type_name -> app.ActionPayload.Action
	4,  // 14: app.AppService.List:input_type -> app.ListRequest
	3,  // 15: app.AppService.Create:input_type -> app.Application
	6,  // 16: app.AppService.Update:input_type -> app.UpdateRequest
	14, // 17: app.AppService.Delete:input_type -> kiae.IdRequest
	14, // 18: app.AppService.Read:input_type -> kiae.IdRequest
	7,  // 19: app.AppService.DoAction:input_type -> app.ActionPayload
	5,  // 20: app.AppService.List:output_type -> app.ListResponse
	3,  // 21: app.AppService.Create:output_type -> app.Application
	3,  // 22: app.AppService.Update:output_type -> app.Application
	15, // 23: app.AppService.Delete:output_type -> google.protobuf.Empty
	3,  // 24: app.AppService.Read:output_type -> app.Application
	3,  // 25: app.AppService.DoAction:output_type -> app.Application
	20, // [20:26] is the sub-list for method output_type
	14, // [14:20] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_app_app_proto_init() }
func file_app_app_proto_init() {
	if File_app_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_app_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_app_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_app_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_app_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionPayload); i {
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
			RawDescriptor: file_app_app_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_app_proto_goTypes,
		DependencyIndexes: file_app_app_proto_depIdxs,
		EnumInfos:         file_app_app_proto_enumTypes,
		MessageInfos:      file_app_app_proto_msgTypes,
	}.Build()
	File_app_app_proto = out.File
	file_app_app_proto_rawDesc = nil
	file_app_app_proto_goTypes = nil
	file_app_app_proto_depIdxs = nil
}
