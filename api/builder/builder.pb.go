// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: builder/builder.proto

package builder

import (
	kiae "github.com/kiaedev/kiae/api/kiae"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Builder_Status int32

const (
	Builder_PENDING   Builder_Status = 0
	Builder_BUILDING  Builder_Status = 1
	Builder_PUBLISHED Builder_Status = 2
	Builder_EXPIRED   Builder_Status = 3
	Builder_FAILED    Builder_Status = 4
)

// Enum value maps for Builder_Status.
var (
	Builder_Status_name = map[int32]string{
		0: "PENDING",
		1: "BUILDING",
		2: "PUBLISHED",
		3: "EXPIRED",
		4: "FAILED",
	}
	Builder_Status_value = map[string]int32{
		"PENDING":   0,
		"BUILDING":  1,
		"PUBLISHED": 2,
		"EXPIRED":   3,
		"FAILED":    4,
	}
)

func (x Builder_Status) Enum() *Builder_Status {
	p := new(Builder_Status)
	*p = x
	return p
}

func (x Builder_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Builder_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_builder_builder_proto_enumTypes[0].Descriptor()
}

func (Builder_Status) Type() protoreflect.EnumType {
	return &file_builder_builder_proto_enumTypes[0]
}

func (x Builder_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Builder_Status.Descriptor instead.
func (Builder_Status) EnumDescriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{0, 0}
}

type Builder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"` // @gotags: bson:"_id,omitempty"
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Intro      string                 `protobuf:"bytes,4,opt,name=intro,proto3" json:"intro,omitempty"`
	RegistryId string                 `protobuf:"bytes,5,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty" bson:"registry_id,omitempty"` // @gotags: bson:"registry_id,omitempty"
	StackId    string                 `protobuf:"bytes,6,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty" bson:"stack_id,omitempty"`          // @gotags: bson:"stack_id,omitempty"
	BuildImage string                 `protobuf:"bytes,7,opt,name=build_image,json=buildImage,proto3" json:"build_image,omitempty" bson:"build_image,omitempty"` // @gotags: bson:"build_image,omitempty"
	RunImage   string                 `protobuf:"bytes,8,opt,name=run_image,json=runImage,proto3" json:"run_image,omitempty" bson:"run_image,omitempty"`       // @gotags: bson:"run_image,omitempty"
	Packs      []*Pack                `protobuf:"bytes,9,rep,name=packs,proto3" json:"packs,omitempty"`
	Status     Builder_Status         `protobuf:"varint,10,opt,name=status,proto3,enum=builder.Builder_Status" json:"status,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at,omitempty"` // @gotags: bson:"created_at,omitempty"
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,103,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bson:"updated_at,omitempty"` // @gotags: bson:"updated_at,omitempty"
}

func (x *Builder) Reset() {
	*x = Builder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Builder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Builder) ProtoMessage() {}

func (x *Builder) ProtoReflect() protoreflect.Message {
	mi := &file_builder_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Builder.ProtoReflect.Descriptor instead.
func (*Builder) Descriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{0}
}

func (x *Builder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Builder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Builder) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *Builder) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

func (x *Builder) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *Builder) GetBuildImage() string {
	if x != nil {
		return x.BuildImage
	}
	return ""
}

func (x *Builder) GetRunImage() string {
	if x != nil {
		return x.RunImage
	}
	return ""
}

func (x *Builder) GetPacks() []*Pack {
	if x != nil {
		return x.Packs
	}
	return nil
}

func (x *Builder) GetStatus() Builder_Status {
	if x != nil {
		return x.Status
	}
	return Builder_PENDING
}

func (x *Builder) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Builder) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Pack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang  string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Envs  []*Env `protobuf:"bytes,4,rep,name=envs,proto3" json:"envs,omitempty"`
}

func (x *Pack) Reset() {
	*x = Pack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pack) ProtoMessage() {}

func (x *Pack) ProtoReflect() protoreflect.Message {
	mi := &file_builder_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pack.ProtoReflect.Descriptor instead.
func (*Pack) Descriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{1}
}

func (x *Pack) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Pack) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Pack) GetEnvs() []*Env {
	if x != nil {
		return x.Envs
	}
	return nil
}

type Env struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Intro        string `protobuf:"bytes,2,opt,name=intro,proto3" json:"intro,omitempty"`
	DefaultValue string `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (x *Env) Reset() {
	*x = Env{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_builder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Env) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Env) ProtoMessage() {}

func (x *Env) ProtoReflect() protoreflect.Message {
	mi := &file_builder_builder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Env.ProtoReflect.Descriptor instead.
func (*Env) Descriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{2}
}

func (x *Env) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Env) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *Env) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

type BuilderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuilderListRequest) Reset() {
	*x = BuilderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_builder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuilderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuilderListRequest) ProtoMessage() {}

func (x *BuilderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builder_builder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuilderListRequest.ProtoReflect.Descriptor instead.
func (*BuilderListRequest) Descriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{3}
}

type BuilderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Builder `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *BuilderListResponse) Reset() {
	*x = BuilderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_builder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuilderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuilderListResponse) ProtoMessage() {}

func (x *BuilderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_builder_builder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuilderListResponse.ProtoReflect.Descriptor instead.
func (*BuilderListResponse) Descriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{4}
}

func (x *BuilderListResponse) GetItems() []*Builder {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *BuilderListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SuggestedStack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Intro      string `protobuf:"bytes,2,opt,name=intro,proto3" json:"intro,omitempty"`
	StackId    string `protobuf:"bytes,6,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty" bson:"stack_id,omitempty"`          // @gotags: bson:"stack_id,omitempty"
	BuildImage string `protobuf:"bytes,7,opt,name=build_image,json=buildImage,proto3" json:"build_image,omitempty" bson:"build_image,omitempty"` // @gotags: bson:"build_image,omitempty"
	RunImage   string `protobuf:"bytes,8,opt,name=run_image,json=runImage,proto3" json:"run_image,omitempty" bson:"run_image,omitempty"`       // @gotags: bson:"run_image,omitempty"
}

func (x *SuggestedStack) Reset() {
	*x = SuggestedStack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_builder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestedStack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestedStack) ProtoMessage() {}

func (x *SuggestedStack) ProtoReflect() protoreflect.Message {
	mi := &file_builder_builder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestedStack.ProtoReflect.Descriptor instead.
func (*SuggestedStack) Descriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{5}
}

func (x *SuggestedStack) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SuggestedStack) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *SuggestedStack) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *SuggestedStack) GetBuildImage() string {
	if x != nil {
		return x.BuildImage
	}
	return ""
}

func (x *SuggestedStack) GetRunImage() string {
	if x != nil {
		return x.RunImage
	}
	return ""
}

type SuggestedStackListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SuggestedStack `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SuggestedStackListResponse) Reset() {
	*x = SuggestedStackListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_builder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestedStackListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestedStackListResponse) ProtoMessage() {}

func (x *SuggestedStackListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_builder_builder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestedStackListResponse.ProtoReflect.Descriptor instead.
func (*SuggestedStackListResponse) Descriptor() ([]byte, []int) {
	return file_builder_builder_proto_rawDescGZIP(), []int{6}
}

func (x *SuggestedStackListResponse) GetItems() []*SuggestedStack {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_builder_builder_proto protoreflect.FileDescriptor

var file_builder_builder_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6b, 0x69,
	0x61, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd6, 0x03, 0x0a, 0x07, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x05,
	0x70, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4b, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x52, 0x0a, 0x04, 0x50, 0x61, 0x63,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x65,
	0x6e, 0x76, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x22, 0x54, 0x0a,
	0x03, 0x45, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x72,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x93,
	0x01, 0x0a, 0x0e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x1a, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x32, 0xf0, 0x03, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0f, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x23, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x5b,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x49, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x1a, 0x10, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x1a, 0x15, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x1a, 0x32, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x50, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6b,
	0x69, 0x61, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x64, 0x65, 0x76, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_builder_builder_proto_rawDescOnce sync.Once
	file_builder_builder_proto_rawDescData = file_builder_builder_proto_rawDesc
)

func file_builder_builder_proto_rawDescGZIP() []byte {
	file_builder_builder_proto_rawDescOnce.Do(func() {
		file_builder_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_builder_builder_proto_rawDescData)
	})
	return file_builder_builder_proto_rawDescData
}

var file_builder_builder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_builder_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_builder_builder_proto_goTypes = []interface{}{
	(Builder_Status)(0),                // 0: builder.Builder.Status
	(*Builder)(nil),                    // 1: builder.Builder
	(*Pack)(nil),                       // 2: builder.Pack
	(*Env)(nil),                        // 3: builder.Env
	(*BuilderListRequest)(nil),         // 4: builder.BuilderListRequest
	(*BuilderListResponse)(nil),        // 5: builder.BuilderListResponse
	(*SuggestedStack)(nil),             // 6: builder.SuggestedStack
	(*SuggestedStackListResponse)(nil), // 7: builder.SuggestedStackListResponse
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
	(*kiae.IdRequest)(nil),             // 10: kiae.IdRequest
}
var file_builder_builder_proto_depIdxs = []int32{
	2,  // 0: builder.Builder.packs:type_name -> builder.Pack
	0,  // 1: builder.Builder.status:type_name -> builder.Builder.Status
	8,  // 2: builder.Builder.created_at:type_name -> google.protobuf.Timestamp
	8,  // 3: builder.Builder.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 4: builder.Pack.envs:type_name -> builder.Env
	1,  // 5: builder.BuilderListResponse.items:type_name -> builder.Builder
	6,  // 6: builder.SuggestedStackListResponse.items:type_name -> builder.SuggestedStack
	9,  // 7: builder.BuilderService.SuggestedStacks:input_type -> google.protobuf.Empty
	4,  // 8: builder.BuilderService.List:input_type -> builder.BuilderListRequest
	1,  // 9: builder.BuilderService.Create:input_type -> builder.Builder
	1,  // 10: builder.BuilderService.Update:input_type -> builder.Builder
	10, // 11: builder.BuilderService.Delete:input_type -> kiae.IdRequest
	7,  // 12: builder.BuilderService.SuggestedStacks:output_type -> builder.SuggestedStackListResponse
	5,  // 13: builder.BuilderService.List:output_type -> builder.BuilderListResponse
	1,  // 14: builder.BuilderService.Create:output_type -> builder.Builder
	1,  // 15: builder.BuilderService.Update:output_type -> builder.Builder
	9,  // 16: builder.BuilderService.Delete:output_type -> google.protobuf.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_builder_builder_proto_init() }
func file_builder_builder_proto_init() {
	if File_builder_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_builder_builder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Builder); i {
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
		file_builder_builder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pack); i {
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
		file_builder_builder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Env); i {
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
		file_builder_builder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuilderListRequest); i {
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
		file_builder_builder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuilderListResponse); i {
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
		file_builder_builder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestedStack); i {
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
		file_builder_builder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestedStackListResponse); i {
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
			RawDescriptor: file_builder_builder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_builder_builder_proto_goTypes,
		DependencyIndexes: file_builder_builder_proto_depIdxs,
		EnumInfos:         file_builder_builder_proto_enumTypes,
		MessageInfos:      file_builder_builder_proto_msgTypes,
	}.Build()
	File_builder_builder_proto = out.File
	file_builder_builder_proto_rawDesc = nil
	file_builder_builder_proto_goTypes = nil
	file_builder_builder_proto_depIdxs = nil
}