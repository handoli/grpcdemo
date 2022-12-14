// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: twoWay.proto

package twoWay

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

// TwoWayServiceClient is the client API for TwoWayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwoWayServiceClient interface {
	//rpc服务，服务名：“add” 请求参数（结构体）：AddReq  响应参数（结构体）：AddRes
	TwoWay(ctx context.Context, opts ...grpc.CallOption) (TwoWayService_TwoWayClient, error)
}

type twoWayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTwoWayServiceClient(cc grpc.ClientConnInterface) TwoWayServiceClient {
	return &twoWayServiceClient{cc}
}

func (c *twoWayServiceClient) TwoWay(ctx context.Context, opts ...grpc.CallOption) (TwoWayService_TwoWayClient, error) {
	stream, err := c.cc.NewStream(ctx, &TwoWayService_ServiceDesc.Streams[0], "/twoWay.TwoWayService/TwoWay", opts...)
	if err != nil {
		return nil, err
	}
	x := &twoWayServiceTwoWayClient{stream}
	return x, nil
}

type TwoWayService_TwoWayClient interface {
	Send(*TwoWayReq) error
	Recv() (*TwoWayRes, error)
	grpc.ClientStream
}

type twoWayServiceTwoWayClient struct {
	grpc.ClientStream
}

func (x *twoWayServiceTwoWayClient) Send(m *TwoWayReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *twoWayServiceTwoWayClient) Recv() (*TwoWayRes, error) {
	m := new(TwoWayRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TwoWayServiceServer is the server API for TwoWayService service.
// All implementations must embed UnimplementedTwoWayServiceServer
// for forward compatibility
type TwoWayServiceServer interface {
	//rpc服务，服务名：“add” 请求参数（结构体）：AddReq  响应参数（结构体）：AddRes
	TwoWay(TwoWayService_TwoWayServer) error
	mustEmbedUnimplementedTwoWayServiceServer()
}

// UnimplementedTwoWayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTwoWayServiceServer struct {
}

func (UnimplementedTwoWayServiceServer) TwoWay(TwoWayService_TwoWayServer) error {
	return status.Errorf(codes.Unimplemented, "method TwoWay not implemented")
}
func (UnimplementedTwoWayServiceServer) mustEmbedUnimplementedTwoWayServiceServer() {}

// UnsafeTwoWayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwoWayServiceServer will
// result in compilation errors.
type UnsafeTwoWayServiceServer interface {
	mustEmbedUnimplementedTwoWayServiceServer()
}

func RegisterTwoWayServiceServer(s grpc.ServiceRegistrar, srv TwoWayServiceServer) {
	s.RegisterService(&TwoWayService_ServiceDesc, srv)
}

func _TwoWayService_TwoWay_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TwoWayServiceServer).TwoWay(&twoWayServiceTwoWayServer{stream})
}

type TwoWayService_TwoWayServer interface {
	Send(*TwoWayRes) error
	Recv() (*TwoWayReq, error)
	grpc.ServerStream
}

type twoWayServiceTwoWayServer struct {
	grpc.ServerStream
}

func (x *twoWayServiceTwoWayServer) Send(m *TwoWayRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *twoWayServiceTwoWayServer) Recv() (*TwoWayReq, error) {
	m := new(TwoWayReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TwoWayService_ServiceDesc is the grpc.ServiceDesc for TwoWayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TwoWayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "twoWay.TwoWayService",
	HandlerType: (*TwoWayServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TwoWay",
			Handler:       _TwoWayService_TwoWay_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "twoWay.proto",
}
