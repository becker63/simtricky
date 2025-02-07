// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: machine_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Tricky_Start_FullMethodName         = "/pb.tricky/Start"
	Tricky_Stop_FullMethodName          = "/pb.tricky/Stop"
	Tricky_ChangeConfig_FullMethodName  = "/pb.tricky/ChangeConfig"
	Tricky_Create_Server_FullMethodName = "/pb.tricky/Create_Server"
)

// TrickyClient is the client API for Tricky service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrickyClient interface {
	// Machine api
	Start(ctx context.Context, in *StartStopReq, opts ...grpc.CallOption) (*Sucess, error)
	Stop(ctx context.Context, in *StartStopReq, opts ...grpc.CallOption) (*Sucess, error)
	ChangeConfig(ctx context.Context, in *ConfigReq, opts ...grpc.CallOption) (*Sucess, error)
	Create_Server(ctx context.Context, in *ConfigReq, opts ...grpc.CallOption) (*Sucess, error)
}

type trickyClient struct {
	cc grpc.ClientConnInterface
}

func NewTrickyClient(cc grpc.ClientConnInterface) TrickyClient {
	return &trickyClient{cc}
}

func (c *trickyClient) Start(ctx context.Context, in *StartStopReq, opts ...grpc.CallOption) (*Sucess, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Sucess)
	err := c.cc.Invoke(ctx, Tricky_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trickyClient) Stop(ctx context.Context, in *StartStopReq, opts ...grpc.CallOption) (*Sucess, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Sucess)
	err := c.cc.Invoke(ctx, Tricky_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trickyClient) ChangeConfig(ctx context.Context, in *ConfigReq, opts ...grpc.CallOption) (*Sucess, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Sucess)
	err := c.cc.Invoke(ctx, Tricky_ChangeConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trickyClient) Create_Server(ctx context.Context, in *ConfigReq, opts ...grpc.CallOption) (*Sucess, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Sucess)
	err := c.cc.Invoke(ctx, Tricky_Create_Server_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrickyServer is the server API for Tricky service.
// All implementations must embed UnimplementedTrickyServer
// for forward compatibility.
type TrickyServer interface {
	// Machine api
	Start(context.Context, *StartStopReq) (*Sucess, error)
	Stop(context.Context, *StartStopReq) (*Sucess, error)
	ChangeConfig(context.Context, *ConfigReq) (*Sucess, error)
	Create_Server(context.Context, *ConfigReq) (*Sucess, error)
	mustEmbedUnimplementedTrickyServer()
}

// UnimplementedTrickyServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrickyServer struct{}

func (UnimplementedTrickyServer) Start(context.Context, *StartStopReq) (*Sucess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedTrickyServer) Stop(context.Context, *StartStopReq) (*Sucess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedTrickyServer) ChangeConfig(context.Context, *ConfigReq) (*Sucess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeConfig not implemented")
}
func (UnimplementedTrickyServer) Create_Server(context.Context, *ConfigReq) (*Sucess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create_Server not implemented")
}
func (UnimplementedTrickyServer) mustEmbedUnimplementedTrickyServer() {}
func (UnimplementedTrickyServer) testEmbeddedByValue()                {}

// UnsafeTrickyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrickyServer will
// result in compilation errors.
type UnsafeTrickyServer interface {
	mustEmbedUnimplementedTrickyServer()
}

func RegisterTrickyServer(s grpc.ServiceRegistrar, srv TrickyServer) {
	// If the following call pancis, it indicates UnimplementedTrickyServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Tricky_ServiceDesc, srv)
}

func _Tricky_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartStopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrickyServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tricky_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrickyServer).Start(ctx, req.(*StartStopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tricky_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartStopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrickyServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tricky_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrickyServer).Stop(ctx, req.(*StartStopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tricky_ChangeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrickyServer).ChangeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tricky_ChangeConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrickyServer).ChangeConfig(ctx, req.(*ConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tricky_Create_Server_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrickyServer).Create_Server(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tricky_Create_Server_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrickyServer).Create_Server(ctx, req.(*ConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Tricky_ServiceDesc is the grpc.ServiceDesc for Tricky service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tricky_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.tricky",
	HandlerType: (*TrickyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Tricky_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Tricky_Stop_Handler,
		},
		{
			MethodName: "ChangeConfig",
			Handler:    _Tricky_ChangeConfig_Handler,
		},
		{
			MethodName: "Create_Server",
			Handler:    _Tricky_Create_Server_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "machine_service.proto",
}
