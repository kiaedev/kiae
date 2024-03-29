// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: entry/entry.proto

package entry

import (
	context "context"
	kiae "github.com/kiaedev/kiae/api/kiae"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EntryServiceClient is the client API for EntryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Entry, error)
	Delete(ctx context.Context, in *kiae.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type entryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryServiceClient(cc grpc.ClientConnInterface) EntryServiceClient {
	return &entryServiceClient{cc}
}

func (c *entryServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/entry.EntryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) Create(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/entry.EntryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/entry.EntryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServiceClient) Delete(ctx context.Context, in *kiae.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/entry.EntryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntryServiceServer is the server API for EntryService service.
// All implementations should embed UnimplementedEntryServiceServer
// for forward compatibility
type EntryServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *Entry) (*Entry, error)
	Update(context.Context, *UpdateRequest) (*Entry, error)
	Delete(context.Context, *kiae.IdRequest) (*emptypb.Empty, error)
}

// UnimplementedEntryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntryServiceServer struct {
}

func (UnimplementedEntryServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEntryServiceServer) Create(context.Context, *Entry) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEntryServiceServer) Update(context.Context, *UpdateRequest) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEntryServiceServer) Delete(context.Context, *kiae.IdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeEntryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryServiceServer will
// result in compilation errors.
type UnsafeEntryServiceServer interface {
	mustEmbedUnimplementedEntryServiceServer()
}

func RegisterEntryServiceServer(s grpc.ServiceRegistrar, srv EntryServiceServer) {
	s.RegisterService(&EntryService_ServiceDesc, srv)
}

func _EntryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry.EntryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry.EntryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).Create(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry.EntryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kiae.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entry.EntryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).Delete(ctx, req.(*kiae.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntryService_ServiceDesc is the grpc.ServiceDesc for EntryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entry.EntryService",
	HandlerType: (*EntryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _EntryService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EntryService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EntryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EntryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entry/entry.proto",
}
