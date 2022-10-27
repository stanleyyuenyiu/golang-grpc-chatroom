// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: protos/event/event.proto

package eventpush

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

// EventPushClient is the client API for EventPush service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventPushClient interface {
	Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (EventPush_JoinClient, error)
	SendMsg(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendReqRes, error)
	BoardCast(ctx context.Context, in *BoardCastReq, opts ...grpc.CallOption) (*SendReqRes, error)
}

type eventPushClient struct {
	cc grpc.ClientConnInterface
}

func NewEventPushClient(cc grpc.ClientConnInterface) EventPushClient {
	return &eventPushClient{cc}
}

func (c *eventPushClient) Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (EventPush_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventPush_ServiceDesc.Streams[0], "/eventpush.EventPush/Join", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventPushJoinClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventPush_JoinClient interface {
	Recv() (*EventStream, error)
	grpc.ClientStream
}

type eventPushJoinClient struct {
	grpc.ClientStream
}

func (x *eventPushJoinClient) Recv() (*EventStream, error) {
	m := new(EventStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventPushClient) SendMsg(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendReqRes, error) {
	out := new(SendReqRes)
	err := c.cc.Invoke(ctx, "/eventpush.EventPush/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventPushClient) BoardCast(ctx context.Context, in *BoardCastReq, opts ...grpc.CallOption) (*SendReqRes, error) {
	out := new(SendReqRes)
	err := c.cc.Invoke(ctx, "/eventpush.EventPush/BoardCast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventPushServer is the server API for EventPush service.
// All implementations must embed UnimplementedEventPushServer
// for forward compatibility
type EventPushServer interface {
	Join(*JoinReq, EventPush_JoinServer) error
	SendMsg(context.Context, *SendReq) (*SendReqRes, error)
	BoardCast(context.Context, *BoardCastReq) (*SendReqRes, error)
	mustEmbedUnimplementedEventPushServer()
}

// UnimplementedEventPushServer must be embedded to have forward compatible implementations.
type UnimplementedEventPushServer struct {
}

func (UnimplementedEventPushServer) Join(*JoinReq, EventPush_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedEventPushServer) SendMsg(context.Context, *SendReq) (*SendReqRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedEventPushServer) BoardCast(context.Context, *BoardCastReq) (*SendReqRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BoardCast not implemented")
}
func (UnimplementedEventPushServer) mustEmbedUnimplementedEventPushServer() {}

// UnsafeEventPushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventPushServer will
// result in compilation errors.
type UnsafeEventPushServer interface {
	mustEmbedUnimplementedEventPushServer()
}

func RegisterEventPushServer(s grpc.ServiceRegistrar, srv EventPushServer) {
	s.RegisterService(&EventPush_ServiceDesc, srv)
}

func _EventPush_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventPushServer).Join(m, &eventPushJoinServer{stream})
}

type EventPush_JoinServer interface {
	Send(*EventStream) error
	grpc.ServerStream
}

type eventPushJoinServer struct {
	grpc.ServerStream
}

func (x *eventPushJoinServer) Send(m *EventStream) error {
	return x.ServerStream.SendMsg(m)
}

func _EventPush_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventPushServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpush.EventPush/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventPushServer).SendMsg(ctx, req.(*SendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventPush_BoardCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardCastReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventPushServer).BoardCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventpush.EventPush/BoardCast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventPushServer).BoardCast(ctx, req.(*BoardCastReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EventPush_ServiceDesc is the grpc.ServiceDesc for EventPush service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventPush_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventpush.EventPush",
	HandlerType: (*EventPushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _EventPush_SendMsg_Handler,
		},
		{
			MethodName: "BoardCast",
			Handler:    _EventPush_BoardCast_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Join",
			Handler:       _EventPush_Join_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/event/event.proto",
}
