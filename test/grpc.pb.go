// Code generated by protoc-gen-gogo.
// source: grpc.proto
// DO NOT EDIT!

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	MyRequest
	MyResponse
	MyMsg
	MyMsg2
*/
package grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyRequest struct {
	Value  int64 `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	Value2 int32 `protobuf:"varint,2,opt,name=Value2,json=value2,proto3" json:"Value2,omitempty"`
}

func (m *MyRequest) Reset()                    { *m = MyRequest{} }
func (m *MyRequest) String() string            { return proto.CompactTextString(m) }
func (*MyRequest) ProtoMessage()               {}
func (*MyRequest) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{0} }

type MyResponse struct {
	Value int64 `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
}

func (m *MyResponse) Reset()                    { *m = MyResponse{} }
func (m *MyResponse) String() string            { return proto.CompactTextString(m) }
func (*MyResponse) ProtoMessage()               {}
func (*MyResponse) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{1} }

type MyMsg struct {
	Value int64 `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
}

func (m *MyMsg) Reset()                    { *m = MyMsg{} }
func (m *MyMsg) String() string            { return proto.CompactTextString(m) }
func (*MyMsg) ProtoMessage()               {}
func (*MyMsg) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{2} }

type MyMsg2 struct {
	Value int64 `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
}

func (m *MyMsg2) Reset()                    { *m = MyMsg2{} }
func (m *MyMsg2) String() string            { return proto.CompactTextString(m) }
func (*MyMsg2) ProtoMessage()               {}
func (*MyMsg2) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{3} }

func init() {
	proto.RegisterType((*MyRequest)(nil), "grpc.MyRequest")
	proto.RegisterType((*MyResponse)(nil), "grpc.MyResponse")
	proto.RegisterType((*MyMsg)(nil), "grpc.MyMsg")
	proto.RegisterType((*MyMsg2)(nil), "grpc.MyMsg2")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion3

// Client API for MyTest service

type MyTestClient interface {
	UnaryCall(ctx context.Context, in *MyRequest, opts ...grpc1.CallOption) (*MyResponse, error)
	// This RPC streams from the server only.
	Downstream(ctx context.Context, in *MyRequest, opts ...grpc1.CallOption) (MyTest_DownstreamClient, error)
	// This RPC streams from the client.
	Upstream(ctx context.Context, opts ...grpc1.CallOption) (MyTest_UpstreamClient, error)
	// This one streams in both directions.
	Bidi(ctx context.Context, opts ...grpc1.CallOption) (MyTest_BidiClient, error)
}

type myTestClient struct {
	cc *grpc1.ClientConn
}

func NewMyTestClient(cc *grpc1.ClientConn) MyTestClient {
	return &myTestClient{cc}
}

func (c *myTestClient) UnaryCall(ctx context.Context, in *MyRequest, opts ...grpc1.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := grpc1.Invoke(ctx, "/grpc.MyTest/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myTestClient) Downstream(ctx context.Context, in *MyRequest, opts ...grpc1.CallOption) (MyTest_DownstreamClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_MyTest_serviceDesc.Streams[0], c.cc, "/grpc.MyTest/Downstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &myTestDownstreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MyTest_DownstreamClient interface {
	Recv() (*MyMsg, error)
	grpc1.ClientStream
}

type myTestDownstreamClient struct {
	grpc1.ClientStream
}

func (x *myTestDownstreamClient) Recv() (*MyMsg, error) {
	m := new(MyMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *myTestClient) Upstream(ctx context.Context, opts ...grpc1.CallOption) (MyTest_UpstreamClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_MyTest_serviceDesc.Streams[1], c.cc, "/grpc.MyTest/Upstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &myTestUpstreamClient{stream}
	return x, nil
}

type MyTest_UpstreamClient interface {
	Send(*MyMsg) error
	CloseAndRecv() (*MyResponse, error)
	grpc1.ClientStream
}

type myTestUpstreamClient struct {
	grpc1.ClientStream
}

func (x *myTestUpstreamClient) Send(m *MyMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myTestUpstreamClient) CloseAndRecv() (*MyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *myTestClient) Bidi(ctx context.Context, opts ...grpc1.CallOption) (MyTest_BidiClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_MyTest_serviceDesc.Streams[2], c.cc, "/grpc.MyTest/Bidi", opts...)
	if err != nil {
		return nil, err
	}
	x := &myTestBidiClient{stream}
	return x, nil
}

type MyTest_BidiClient interface {
	Send(*MyMsg) error
	Recv() (*MyMsg2, error)
	grpc1.ClientStream
}

type myTestBidiClient struct {
	grpc1.ClientStream
}

func (x *myTestBidiClient) Send(m *MyMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myTestBidiClient) Recv() (*MyMsg2, error) {
	m := new(MyMsg2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MyTest service

type MyTestServer interface {
	UnaryCall(context.Context, *MyRequest) (*MyResponse, error)
	// This RPC streams from the server only.
	Downstream(*MyRequest, MyTest_DownstreamServer) error
	// This RPC streams from the client.
	Upstream(MyTest_UpstreamServer) error
	// This one streams in both directions.
	Bidi(MyTest_BidiServer) error
}

func RegisterMyTestServer(s *grpc1.Server, srv MyTestServer) {
	s.RegisterService(&_MyTest_serviceDesc, srv)
}

func _MyTest_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyTestServer).UnaryCall(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MyTest/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyTestServer).UnaryCall(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyTest_Downstream_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(MyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyTestServer).Downstream(m, &myTestDownstreamServer{stream})
}

type MyTest_DownstreamServer interface {
	Send(*MyMsg) error
	grpc1.ServerStream
}

type myTestDownstreamServer struct {
	grpc1.ServerStream
}

func (x *myTestDownstreamServer) Send(m *MyMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _MyTest_Upstream_Handler(srv interface{}, stream grpc1.ServerStream) error {
	return srv.(MyTestServer).Upstream(&myTestUpstreamServer{stream})
}

type MyTest_UpstreamServer interface {
	SendAndClose(*MyResponse) error
	Recv() (*MyMsg, error)
	grpc1.ServerStream
}

type myTestUpstreamServer struct {
	grpc1.ServerStream
}

func (x *myTestUpstreamServer) SendAndClose(m *MyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myTestUpstreamServer) Recv() (*MyMsg, error) {
	m := new(MyMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MyTest_Bidi_Handler(srv interface{}, stream grpc1.ServerStream) error {
	return srv.(MyTestServer).Bidi(&myTestBidiServer{stream})
}

type MyTest_BidiServer interface {
	Send(*MyMsg2) error
	Recv() (*MyMsg, error)
	grpc1.ServerStream
}

type myTestBidiServer struct {
	grpc1.ServerStream
}

func (x *myTestBidiServer) Send(m *MyMsg2) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myTestBidiServer) Recv() (*MyMsg, error) {
	m := new(MyMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MyTest_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.MyTest",
	HandlerType: (*MyTestServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _MyTest_UnaryCall_Handler,
		},
	},
	Streams: []grpc1.StreamDesc{
		{
			StreamName:    "Downstream",
			Handler:       _MyTest_Downstream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upstream",
			Handler:       _MyTest_Upstream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Bidi",
			Handler:       _MyTest_Bidi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptorGrpc,
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptorGrpc) }

var fileDescriptorGrpc = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x2c, 0xb9, 0x38, 0x7d, 0x2b,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0xc3, 0x12, 0x73, 0x4a, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x58, 0xcb, 0x40, 0x1c, 0x21, 0x31, 0x2e, 0x36, 0xb0,
	0xa8, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x1b, 0x58, 0xd8, 0x48, 0x49, 0x89, 0x8b,
	0x0b, 0xa4, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xbb, 0x5e, 0x25, 0x59, 0x2e, 0x56, 0xdf,
	0x4a, 0xdf, 0xe2, 0x74, 0x1c, 0xd2, 0x72, 0x5c, 0x6c, 0x60, 0x69, 0x23, 0xec, 0xf2, 0x46, 0xbb,
	0x18, 0x41, 0x0a, 0x42, 0x40, 0x6e, 0xd3, 0xe3, 0xe2, 0x0c, 0xcd, 0x4b, 0x2c, 0xaa, 0x74, 0x4e,
	0xcc, 0xc9, 0x11, 0xe2, 0xd7, 0x03, 0x7b, 0x04, 0xee, 0x72, 0x29, 0x01, 0x84, 0x00, 0xd4, 0x3d,
	0x3a, 0x5c, 0x5c, 0x2e, 0xf9, 0xe5, 0x79, 0xc5, 0x25, 0x45, 0xa9, 0x89, 0xb9, 0x98, 0x1a, 0xb8,
	0x61, 0x02, 0xbe, 0xc5, 0xe9, 0x06, 0x8c, 0x42, 0xda, 0x5c, 0x1c, 0xa1, 0x05, 0x50, 0xb5, 0xc8,
	0x52, 0x98, 0x06, 0x6b, 0x30, 0x0a, 0xa9, 0x72, 0xb1, 0x38, 0x65, 0xa6, 0x64, 0xa2, 0x2a, 0xe4,
	0x41, 0xe2, 0x18, 0x69, 0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0xc3, 0xd9, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0x1f, 0xe1, 0xe1, 0x75, 0x01, 0x00, 0x00,
}
