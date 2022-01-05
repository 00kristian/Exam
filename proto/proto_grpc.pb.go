// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/proto.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HashtableClient is the client API for Hashtable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashtableClient interface {
	Put(ctx context.Context, in *Keyvalue, opts ...grpc.CallOption) (*Result, error)
	Get(ctx context.Context, in *GetValue, opts ...grpc.CallOption) (*GetValue, error)
}

type hashtableClient struct {
	cc grpc.ClientConnInterface
}

func NewHashtableClient(cc grpc.ClientConnInterface) HashtableClient {
	return &hashtableClient{cc}
}

func (c *hashtableClient) Put(ctx context.Context, in *Keyvalue, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Proto.hashtable/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtableClient) Get(ctx context.Context, in *GetValue, opts ...grpc.CallOption) (*GetValue, error) {
	out := new(GetValue)
	err := c.cc.Invoke(ctx, "/Proto.hashtable/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashtableServer is the server API for Hashtable service.
// All implementations must embed UnimplementedHashtableServer
// for forward compatibility
type HashtableServer interface {
	Put(context.Context, *Keyvalue) (*Result, error)
	Get(context.Context, *GetValue) (*GetValue, error)
	mustEmbedUnimplementedHashtableServer()
}

// UnimplementedHashtableServer must be embedded to have forward compatible implementations.
type UnimplementedHashtableServer struct {
}

func (UnimplementedHashtableServer) Put(context.Context, *Keyvalue) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedHashtableServer) Get(context.Context, *GetValue) (*GetValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHashtableServer) mustEmbedUnimplementedHashtableServer() {}

// UnsafeHashtableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashtableServer will
// result in compilation errors.
type UnsafeHashtableServer interface {
	mustEmbedUnimplementedHashtableServer()
}

func RegisterHashtableServer(s grpc.ServiceRegistrar, srv HashtableServer) {
	s.RegisterService(&Hashtable_ServiceDesc, srv)
}

func _Hashtable_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keyvalue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtableServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.hashtable/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtableServer).Put(ctx, req.(*Keyvalue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hashtable_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtableServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.hashtable/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtableServer).Get(ctx, req.(*GetValue))
	}
	return interceptor(ctx, in, info, handler)
}

// Hashtable_ServiceDesc is the grpc.ServiceDesc for Hashtable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hashtable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.hashtable",
	HandlerType: (*HashtableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Hashtable_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Hashtable_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}
