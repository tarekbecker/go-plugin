// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_broker.proto

/*
Package plugin is a generated protocol buffer package.

It is generated from these files:
	grpc_broker.proto

It has these top-level messages:
	ConnInfo
*/
package go_plugin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ConnInfo struct {
	ServiceId uint32 `protobuf:"varint,1,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	Network   string `protobuf:"bytes,2,opt,name=network" json:"network,omitempty"`
	Address   string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *ConnInfo) Reset()                    { *m = ConnInfo{} }
func (m *ConnInfo) String() string            { return proto.CompactTextString(m) }
func (*ConnInfo) ProtoMessage()               {}
func (*ConnInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnInfo) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *ConnInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ConnInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*ConnInfo)(nil), "plugin.ConnInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GRPCBroker service

type GRPCBrokerClient interface {
	StartStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBroker_StartStreamClient, error)
}

type gRPCBrokerClient struct {
	cc *grpc.ClientConn
}

func NewGRPCBrokerClient(cc *grpc.ClientConn) GRPCBrokerClient {
	return &gRPCBrokerClient{cc}
}

func (c *gRPCBrokerClient) StartStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBroker_StartStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCBroker_serviceDesc.Streams[0], c.cc, "/plugin.GRPCBroker/StartStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCBrokerStartStreamClient{stream}
	return x, nil
}

type GRPCBroker_StartStreamClient interface {
	Send(*ConnInfo) error
	Recv() (*ConnInfo, error)
	grpc.ClientStream
}

type gRPCBrokerStartStreamClient struct {
	grpc.ClientStream
}

func (x *gRPCBrokerStartStreamClient) Send(m *ConnInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCBrokerStartStreamClient) Recv() (*ConnInfo, error) {
	m := new(ConnInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GRPCBroker service

type GRPCBrokerServer interface {
	StartStream(GRPCBroker_StartStreamServer) error
}

func RegisterGRPCBrokerServer(s *grpc.Server, srv GRPCBrokerServer) {
	s.RegisterService(&_GRPCBroker_serviceDesc, srv)
}

func _GRPCBroker_StartStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCBrokerServer).StartStream(&gRPCBrokerStartStreamServer{stream})
}

type GRPCBroker_StartStreamServer interface {
	Send(*ConnInfo) error
	Recv() (*ConnInfo, error)
	grpc.ServerStream
}

type gRPCBrokerStartStreamServer struct {
	grpc.ServerStream
}

func (x *gRPCBrokerStartStreamServer) Send(m *ConnInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCBrokerStartStreamServer) Recv() (*ConnInfo, error) {
	m := new(ConnInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GRPCBroker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.GRPCBroker",
	HandlerType: (*GRPCBrokerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartStream",
			Handler:       _GRPCBroker_StartStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_broker.proto",
}

func init() { proto.RegisterFile("grpc_broker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2f, 0x2a, 0x48,
	0x8e, 0x4f, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b,
	0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x53, 0x8a, 0xe5, 0xe2, 0x70, 0xce, 0xcf, 0xcb, 0xf3, 0xcc, 0x4b,
	0xcb, 0x17, 0x92, 0xe5, 0xe2, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x8d, 0xcf, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0xe2, 0x84, 0x8a, 0x78, 0xa6, 0x08, 0x49, 0x70, 0xb1, 0xe7,
	0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x20,
	0x99, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x66, 0x88, 0x0c, 0x94, 0x6b, 0xe4, 0xcc,
	0xc5, 0xe5, 0x1e, 0x14, 0xe0, 0xec, 0x04, 0xb6, 0x5a, 0xc8, 0x94, 0x8b, 0x3b, 0xb8, 0x24, 0xb1,
	0xa8, 0x24, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x57, 0x48, 0x40, 0x0f, 0xe2, 0x08, 0x3d, 0x98, 0x0b,
	0xa4, 0x30, 0x44, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x4e, 0x36, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0x5d, 0xfb, 0xe1, 0xc7, 0x00, 0x00, 0x00,
}
