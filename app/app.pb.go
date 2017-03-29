// Code generated by protoc-gen-go.
// source: app/app.proto
// DO NOT EDIT!

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	app/app.proto

It has these top-level messages:
	KeyValMessage
	EmptyMessage
*/
package app

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyValMessage struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValMessage) Reset()                    { *m = KeyValMessage{} }
func (m *KeyValMessage) String() string            { return proto.CompactTextString(m) }
func (*KeyValMessage) ProtoMessage()               {}
func (*KeyValMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyValMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*KeyValMessage)(nil), "app.KeyValMessage")
	proto.RegisterType((*EmptyMessage)(nil), "app.EmptyMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for App service

type AppClient interface {
	KeyValCreate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
	KeyValRead(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
	KeyValUpdate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
	KeyValDelete(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) KeyValCreate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/app.App/KeyValCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) KeyValRead(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/app.App/KeyValRead", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) KeyValUpdate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/app.App/KeyValUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) KeyValDelete(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/app.App/KeyValDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppServer interface {
	KeyValCreate(context.Context, *KeyValMessage) (*EmptyMessage, error)
	KeyValRead(context.Context, *KeyValMessage) (*KeyValMessage, error)
	KeyValUpdate(context.Context, *KeyValMessage) (*EmptyMessage, error)
	KeyValDelete(context.Context, *KeyValMessage) (*EmptyMessage, error)
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_KeyValCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).KeyValCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/KeyValCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).KeyValCreate(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_KeyValRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).KeyValRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/KeyValRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).KeyValRead(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_KeyValUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).KeyValUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/KeyValUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).KeyValUpdate(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_KeyValDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).KeyValDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/KeyValDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).KeyValDelete(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeyValCreate",
			Handler:    _App_KeyValCreate_Handler,
		},
		{
			MethodName: "KeyValRead",
			Handler:    _App_KeyValRead_Handler,
		},
		{
			MethodName: "KeyValUpdate",
			Handler:    _App_KeyValUpdate_Handler,
		},
		{
			MethodName: "KeyValDelete",
			Handler:    _App_KeyValDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/app.proto",
}

func init() { proto.RegisterFile("app/app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0x28, 0xd0,
	0x4f, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0x28, 0x90, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x51, 0x32, 0xe7, 0xe2, 0xf5, 0x4e, 0xad, 0x0c,
	0x4b, 0xcc, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad,
	0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73,
	0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x12, 0x1f, 0x17, 0x8f, 0x6b, 0x6e, 0x41, 0x49,
	0x25, 0x54, 0x9f, 0xd1, 0x39, 0x26, 0x2e, 0x66, 0xc7, 0x82, 0x02, 0xa1, 0x20, 0x2e, 0x1e, 0x88,
	0x81, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x42, 0x42, 0x7a, 0x20, 0xf7, 0xa0, 0xd8, 0x21, 0x25,
	0x08, 0x16, 0x43, 0xd6, 0xae, 0x24, 0xdd, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x51, 0x29, 0x01, 0xfd,
	0x32, 0x43, 0xfd, 0xec, 0xd4, 0xca, 0xb2, 0xc4, 0x1c, 0xfd, 0xea, 0xec, 0xd4, 0xca, 0x5a, 0x2b,
	0x46, 0x2d, 0x21, 0x3f, 0x2e, 0x2e, 0x88, 0x01, 0x41, 0xa9, 0x89, 0x29, 0x58, 0x4d, 0xc4, 0x22,
	0xa6, 0x24, 0x01, 0x36, 0x52, 0x48, 0x08, 0xc3, 0x48, 0x84, 0x1b, 0x43, 0x0b, 0x52, 0x48, 0x77,
	0xa3, 0x12, 0x56, 0x37, 0xfa, 0xc3, 0xcc, 0x74, 0x49, 0xcd, 0x49, 0x25, 0xde, 0x4c, 0xa8, 0x23,
	0xb5, 0x30, 0xcc, 0x4c, 0x62, 0x03, 0x47, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x77, 0x22,
	0x3e, 0xa1, 0xd4, 0x01, 0x00, 0x00,
}
