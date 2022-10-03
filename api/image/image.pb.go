// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: image/image.proto

package image

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Image_Status int32

const (
	Image_BUILDING  Image_Status = 0
	Image_PUBLISHED Image_Status = 1
	Image_EXPIRED   Image_Status = 2
	Image_FAILED    Image_Status = 3
)

// Enum value maps for Image_Status.
var (
	Image_Status_name = map[int32]string{
		0: "BUILDING",
		1: "PUBLISHED",
		2: "EXPIRED",
		3: "FAILED",
	}
	Image_Status_value = map[string]int32{
		"BUILDING":  0,
		"PUBLISHED": 1,
		"EXPIRED":   2,
		"FAILED":    3,
	}
)

func (x Image_Status) Enum() *Image_Status {
	p := new(Image_Status)
	*p = x
	return p
}

func (x Image_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Image_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_image_image_proto_enumTypes[0].Descriptor()
}

func (Image_Status) Type() protoreflect.EnumType {
	return &file_image_image_proto_enumTypes[0]
}

func (x Image_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Image_Status.Descriptor instead.
func (Image_Status) EnumDescriptor() ([]byte, []int) {
	return file_image_image_proto_rawDescGZIP(), []int{0, 0}
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"` // @gotags: bson:"_id,omitempty"
	Pid       string                 `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Image     string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Tag       string                 `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	Commit    string                 `protobuf:"bytes,6,opt,name=commit,proto3" json:"commit,omitempty"`
	Url       string                 `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Status    Image_Status           `protobuf:"varint,9,opt,name=status,proto3,enum=image.Image_Status" json:"status,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,103,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_image_image_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Image) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Image) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetStatus() Image_Status {
	if x != nil {
		return x.Status
	}
	return Image_BUILDING
}

func (x *Image) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Image) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ImageListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *ImageListRequest) Reset() {
	*x = ImageListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageListRequest) ProtoMessage() {}

func (x *ImageListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageListRequest.ProtoReflect.Descriptor instead.
func (*ImageListRequest) Descriptor() ([]byte, []int) {
	return file_image_image_proto_rawDescGZIP(), []int{1}
}

func (x *ImageListRequest) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type ImageListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Image `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ImageListResponse) Reset() {
	*x = ImageListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageListResponse) ProtoMessage() {}

func (x *ImageListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageListResponse.ProtoReflect.Descriptor instead.
func (*ImageListResponse) Descriptor() ([]byte, []int) {
	return file_image_image_proto_rawDescGZIP(), []int{2}
}

func (x *ImageListResponse) GetItems() []*Image {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ImageListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_image_image_proto protoreflect.FileDescriptor

var file_image_image_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6b, 0x69, 0x61, 0x65, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xfa,
	0x42, 0x02, 0x72, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xfa, 0x42, 0x02, 0x72, 0x00,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x22, 0x24, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x11,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x9c, 0x03, 0x0a, 0x0c,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12,
	0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x4e,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x69, 0x64, 0x7d, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7c,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x56, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x50, 0x1a, 0x22, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x3a, 0x01, 0x2a, 0x5a, 0x27, 0x32, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6b, 0x69, 0x61, 0x65, 0x2e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x64, 0x65, 0x76,
	0x2f, 0x6b, 0x69, 0x61, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_image_proto_rawDescOnce sync.Once
	file_image_image_proto_rawDescData = file_image_image_proto_rawDesc
)

func file_image_image_proto_rawDescGZIP() []byte {
	file_image_image_proto_rawDescOnce.Do(func() {
		file_image_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_image_proto_rawDescData)
	})
	return file_image_image_proto_rawDescData
}

var file_image_image_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_image_image_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_image_image_proto_goTypes = []interface{}{
	(Image_Status)(0),             // 0: image.Image.Status
	(*Image)(nil),                 // 1: image.Image
	(*ImageListRequest)(nil),      // 2: image.ImageListRequest
	(*ImageListResponse)(nil),     // 3: image.ImageListResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*kiae.IdRequest)(nil),        // 5: kiae.IdRequest
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_image_image_proto_depIdxs = []int32{
	0, // 0: image.Image.status:type_name -> image.Image.Status
	4, // 1: image.Image.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: image.Image.updated_at:type_name -> google.protobuf.Timestamp
	1, // 3: image.ImageListResponse.items:type_name -> image.Image
	2, // 4: image.ImageService.List:input_type -> image.ImageListRequest
	1, // 5: image.ImageService.Create:input_type -> image.Image
	1, // 6: image.ImageService.Update:input_type -> image.Image
	5, // 7: image.ImageService.Delete:input_type -> kiae.IdRequest
	3, // 8: image.ImageService.List:output_type -> image.ImageListResponse
	1, // 9: image.ImageService.Create:output_type -> image.Image
	1, // 10: image.ImageService.Update:output_type -> image.Image
	6, // 11: image.ImageService.Delete:output_type -> google.protobuf.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_image_image_proto_init() }
func file_image_image_proto_init() {
	if File_image_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_image_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageListRequest); i {
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
		file_image_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageListResponse); i {
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
			RawDescriptor: file_image_image_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_image_proto_goTypes,
		DependencyIndexes: file_image_image_proto_depIdxs,
		EnumInfos:         file_image_image_proto_enumTypes,
		MessageInfos:      file_image_image_proto_msgTypes,
	}.Build()
	File_image_image_proto = out.File
	file_image_image_proto_rawDesc = nil
	file_image_image_proto_goTypes = nil
	file_image_image_proto_depIdxs = nil
}