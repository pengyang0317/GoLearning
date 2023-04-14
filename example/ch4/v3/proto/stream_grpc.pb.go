// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: stream.proto

package streampb

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// 服务端流模式
	ServerStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Greeter_ServerStreamClient, error)
	// 客户端流模式
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_ClientStreamClient, error)
	// 双向流模式
	AllStreeam(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreeamClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) ServerStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Greeter_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/stream.v1.Greeter/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_ServerStreamClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type greeterServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterServerStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/stream.v1.Greeter/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterClientStreamClient{stream}
	return x, nil
}

type Greeter_ClientStreamClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type greeterClientStreamClient struct {
	grpc.ClientStream
}

func (x *greeterClientStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterClientStreamClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) AllStreeam(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreeamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/stream.v1.Greeter/AllStreeam", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterAllStreeamClient{stream}
	return x, nil
}

type Greeter_AllStreeamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type greeterAllStreeamClient struct {
	grpc.ClientStream
}

func (x *greeterAllStreeamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterAllStreeamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// 服务端流模式
	ServerStream(*StreamRequest, Greeter_ServerStreamServer) error
	// 客户端流模式
	ClientStream(Greeter_ClientStreamServer) error
	// 双向流模式
	AllStreeam(Greeter_AllStreeamServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) ServerStream(*StreamRequest, Greeter_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedGreeterServer) ClientStream(Greeter_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedGreeterServer) AllStreeam(Greeter_AllStreeamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStreeam not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).ServerStream(m, &greeterServerStreamServer{stream})
}

type Greeter_ServerStreamServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type greeterServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterServerStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).ClientStream(&greeterClientStreamServer{stream})
}

type Greeter_ClientStreamServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type greeterClientStreamServer struct {
	grpc.ServerStream
}

func (x *greeterClientStreamServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterClientStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_AllStreeam_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).AllStreeam(&greeterAllStreeamServer{stream})
}

type Greeter_AllStreeamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type greeterAllStreeamServer struct {
	grpc.ServerStream
}

func (x *greeterAllStreeamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterAllStreeamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Greeter_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _Greeter_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStreeam",
			Handler:       _Greeter_AllStreeam_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}