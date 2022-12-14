// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: clientStream.proto

package clientStream

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

// ClientStreamServiceClient is the client API for ClientStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientStreamServiceClient interface {
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (ClientStreamService_ClientStreamClient, error)
}

type clientStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamServiceClient(cc grpc.ClientConnInterface) ClientStreamServiceClient {
	return &clientStreamServiceClient{cc}
}

func (c *clientStreamServiceClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (ClientStreamService_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientStreamService_ServiceDesc.Streams[0], "/clientStream.ClientStreamService/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamServiceClientStreamClient{stream}
	return x, nil
}

type ClientStreamService_ClientStreamClient interface {
	Send(*ClientStreamReq) error
	CloseAndRecv() (*ClientStreamRes, error)
	grpc.ClientStream
}

type clientStreamServiceClientStreamClient struct {
	grpc.ClientStream
}

func (x *clientStreamServiceClientStreamClient) Send(m *ClientStreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamServiceClientStreamClient) CloseAndRecv() (*ClientStreamRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamServiceServer is the server API for ClientStreamService service.
// All implementations must embed UnimplementedClientStreamServiceServer
// for forward compatibility
type ClientStreamServiceServer interface {
	ClientStream(ClientStreamService_ClientStreamServer) error
	mustEmbedUnimplementedClientStreamServiceServer()
}

// UnimplementedClientStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientStreamServiceServer struct {
}

func (UnimplementedClientStreamServiceServer) ClientStream(ClientStreamService_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedClientStreamServiceServer) mustEmbedUnimplementedClientStreamServiceServer() {}

// UnsafeClientStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStreamServiceServer will
// result in compilation errors.
type UnsafeClientStreamServiceServer interface {
	mustEmbedUnimplementedClientStreamServiceServer()
}

func RegisterClientStreamServiceServer(s grpc.ServiceRegistrar, srv ClientStreamServiceServer) {
	s.RegisterService(&ClientStreamService_ServiceDesc, srv)
}

func _ClientStreamService_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamServiceServer).ClientStream(&clientStreamServiceClientStreamServer{stream})
}

type ClientStreamService_ClientStreamServer interface {
	SendAndClose(*ClientStreamRes) error
	Recv() (*ClientStreamReq, error)
	grpc.ServerStream
}

type clientStreamServiceClientStreamServer struct {
	grpc.ServerStream
}

func (x *clientStreamServiceClientStreamServer) SendAndClose(m *ClientStreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamServiceClientStreamServer) Recv() (*ClientStreamReq, error) {
	m := new(ClientStreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamService_ServiceDesc is the grpc.ServiceDesc for ClientStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clientStream.ClientStreamService",
	HandlerType: (*ClientStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _ClientStreamService_ClientStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "clientStream.proto",
}
