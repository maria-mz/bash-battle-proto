// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: service.proto

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

// BashBattleClient is the client API for BashBattle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BashBattleClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetGameConfig(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*GameConfig, error)
	GetPlayers(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*Players, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (BashBattle_StreamClient, error)
	Logout(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*EmptyMsg, error)
}

type bashBattleClient struct {
	cc grpc.ClientConnInterface
}

func NewBashBattleClient(cc grpc.ClientConnInterface) BashBattleClient {
	return &bashBattleClient{cc}
}

func (c *bashBattleClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.BashBattle/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bashBattleClient) GetGameConfig(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*GameConfig, error) {
	out := new(GameConfig)
	err := c.cc.Invoke(ctx, "/proto.BashBattle/GetGameConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bashBattleClient) GetPlayers(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*Players, error) {
	out := new(Players)
	err := c.cc.Invoke(ctx, "/proto.BashBattle/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bashBattleClient) Stream(ctx context.Context, opts ...grpc.CallOption) (BashBattle_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &BashBattle_ServiceDesc.Streams[0], "/proto.BashBattle/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &bashBattleStreamClient{stream}
	return x, nil
}

type BashBattle_StreamClient interface {
	Send(*AwkMsg) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type bashBattleStreamClient struct {
	grpc.ClientStream
}

func (x *bashBattleStreamClient) Send(m *AwkMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bashBattleStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bashBattleClient) Logout(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*EmptyMsg, error) {
	out := new(EmptyMsg)
	err := c.cc.Invoke(ctx, "/proto.BashBattle/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BashBattleServer is the server API for BashBattle service.
// All implementations must embed UnimplementedBashBattleServer
// for forward compatibility
type BashBattleServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetGameConfig(context.Context, *EmptyMsg) (*GameConfig, error)
	GetPlayers(context.Context, *EmptyMsg) (*Players, error)
	Stream(BashBattle_StreamServer) error
	Logout(context.Context, *EmptyMsg) (*EmptyMsg, error)
	mustEmbedUnimplementedBashBattleServer()
}

// UnimplementedBashBattleServer must be embedded to have forward compatible implementations.
type UnimplementedBashBattleServer struct {
}

func (UnimplementedBashBattleServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBashBattleServer) GetGameConfig(context.Context, *EmptyMsg) (*GameConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameConfig not implemented")
}
func (UnimplementedBashBattleServer) GetPlayers(context.Context, *EmptyMsg) (*Players, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedBashBattleServer) Stream(BashBattle_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedBashBattleServer) Logout(context.Context, *EmptyMsg) (*EmptyMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedBashBattleServer) mustEmbedUnimplementedBashBattleServer() {}

// UnsafeBashBattleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BashBattleServer will
// result in compilation errors.
type UnsafeBashBattleServer interface {
	mustEmbedUnimplementedBashBattleServer()
}

func RegisterBashBattleServer(s grpc.ServiceRegistrar, srv BashBattleServer) {
	s.RegisterService(&BashBattle_ServiceDesc, srv)
}

func _BashBattle_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BashBattleServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BashBattle/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BashBattleServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BashBattle_GetGameConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BashBattleServer).GetGameConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BashBattle/GetGameConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BashBattleServer).GetGameConfig(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _BashBattle_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BashBattleServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BashBattle/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BashBattleServer).GetPlayers(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _BashBattle_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BashBattleServer).Stream(&bashBattleStreamServer{stream})
}

type BashBattle_StreamServer interface {
	Send(*Event) error
	Recv() (*AwkMsg, error)
	grpc.ServerStream
}

type bashBattleStreamServer struct {
	grpc.ServerStream
}

func (x *bashBattleStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bashBattleStreamServer) Recv() (*AwkMsg, error) {
	m := new(AwkMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BashBattle_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BashBattleServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BashBattle/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BashBattleServer).Logout(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// BashBattle_ServiceDesc is the grpc.ServiceDesc for BashBattle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BashBattle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BashBattle",
	HandlerType: (*BashBattleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _BashBattle_Login_Handler,
		},
		{
			MethodName: "GetGameConfig",
			Handler:    _BashBattle_GetGameConfig_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _BashBattle_GetPlayers_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _BashBattle_Logout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _BashBattle_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
