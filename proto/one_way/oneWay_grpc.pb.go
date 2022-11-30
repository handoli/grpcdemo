// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: oneWay.proto

package oneWay

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

// OneWayServiceClient is the client API for OneWayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OneWayServiceClient interface {
	//rpc服务，服务名：“add” 请求参数（结构体）：AddReq  响应参数（结构体）：AddRes
	OneWay(ctx context.Context, in *OneWayReq, opts ...grpc.CallOption) (*OneWayRes, error)
}

type oneWayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOneWayServiceClient(cc grpc.ClientConnInterface) OneWayServiceClient {
	return &oneWayServiceClient{cc}
}

func (c *oneWayServiceClient) OneWay(ctx context.Context, in *OneWayReq, opts ...grpc.CallOption) (*OneWayRes, error) {
	out := new(OneWayRes)
	err := c.cc.Invoke(ctx, "/oneWay.OneWayService/OneWay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OneWayServiceServer is the server API for OneWayService service.
// All implementations must embed UnimplementedOneWayServiceServer
// for forward compatibility
type OneWayServiceServer interface {
	//rpc服务，服务名：“add” 请求参数（结构体）：AddReq  响应参数（结构体）：AddRes
	OneWay(context.Context, *OneWayReq) (*OneWayRes, error)
	mustEmbedUnimplementedOneWayServiceServer()
}

// UnimplementedOneWayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOneWayServiceServer struct {
}

func (UnimplementedOneWayServiceServer) OneWay(context.Context, *OneWayReq) (*OneWayRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneWay not implemented")
}
func (UnimplementedOneWayServiceServer) mustEmbedUnimplementedOneWayServiceServer() {}

// UnsafeOneWayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OneWayServiceServer will
// result in compilation errors.
type UnsafeOneWayServiceServer interface {
	mustEmbedUnimplementedOneWayServiceServer()
}

func RegisterOneWayServiceServer(s grpc.ServiceRegistrar, srv OneWayServiceServer) {
	s.RegisterService(&OneWayService_ServiceDesc, srv)
}

func _OneWayService_OneWay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneWayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OneWayServiceServer).OneWay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oneWay.OneWayService/OneWay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OneWayServiceServer).OneWay(ctx, req.(*OneWayReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OneWayService_ServiceDesc is the grpc.ServiceDesc for OneWayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OneWayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oneWay.OneWayService",
	HandlerType: (*OneWayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OneWay",
			Handler:    _OneWayService_OneWay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oneWay.proto",
}
