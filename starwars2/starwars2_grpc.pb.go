// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_starwars_grpc

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

// StarWars1Client is the client API for StarWars1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StarWars1Client interface {
	CityMgmtFulcrum(ctx context.Context, in *NewCity1, opts ...grpc.CallOption) (*RespFulcrum1, error)
	CityBrokerFulcrum(ctx context.Context, in *NewCity1, opts ...grpc.CallOption) (*RespFulcrum2, error)
	RelojesBrokerFulcrum(ctx context.Context, in *Planeta, opts ...grpc.CallOption) (*RespFulcrum1, error)
	ConsistenciaEventual(ctx context.Context, in *SolMerge, opts ...grpc.CallOption) (*RespBroker3, error)
	PreguntarRelojesYRegistros(ctx context.Context, in *SolMerge, opts ...grpc.CallOption) (*RelojesYRegistros, error)
	MandarFulcrums(ctx context.Context, in *RelojesYRegistros, opts ...grpc.CallOption) (*SolMerge, error)
}

type starWars1Client struct {
	cc grpc.ClientConnInterface
}

func NewStarWars1Client(cc grpc.ClientConnInterface) StarWars1Client {
	return &starWars1Client{cc}
}

func (c *starWars1Client) CityMgmtFulcrum(ctx context.Context, in *NewCity1, opts ...grpc.CallOption) (*RespFulcrum1, error) {
	out := new(RespFulcrum1)
	err := c.cc.Invoke(ctx, "/starwars2.StarWars1/CityMgmtFulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starWars1Client) CityBrokerFulcrum(ctx context.Context, in *NewCity1, opts ...grpc.CallOption) (*RespFulcrum2, error) {
	out := new(RespFulcrum2)
	err := c.cc.Invoke(ctx, "/starwars2.StarWars1/CityBrokerFulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starWars1Client) RelojesBrokerFulcrum(ctx context.Context, in *Planeta, opts ...grpc.CallOption) (*RespFulcrum1, error) {
	out := new(RespFulcrum1)
	err := c.cc.Invoke(ctx, "/starwars2.StarWars1/RelojesBrokerFulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starWars1Client) ConsistenciaEventual(ctx context.Context, in *SolMerge, opts ...grpc.CallOption) (*RespBroker3, error) {
	out := new(RespBroker3)
	err := c.cc.Invoke(ctx, "/starwars2.StarWars1/ConsistenciaEventual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starWars1Client) PreguntarRelojesYRegistros(ctx context.Context, in *SolMerge, opts ...grpc.CallOption) (*RelojesYRegistros, error) {
	out := new(RelojesYRegistros)
	err := c.cc.Invoke(ctx, "/starwars2.StarWars1/preguntarRelojesYRegistros", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starWars1Client) MandarFulcrums(ctx context.Context, in *RelojesYRegistros, opts ...grpc.CallOption) (*SolMerge, error) {
	out := new(SolMerge)
	err := c.cc.Invoke(ctx, "/starwars2.StarWars1/mandarFulcrums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarWars1Server is the server API for StarWars1 service.
// All implementations must embed UnimplementedStarWars1Server
// for forward compatibility
type StarWars1Server interface {
	CityMgmtFulcrum(context.Context, *NewCity1) (*RespFulcrum1, error)
	CityBrokerFulcrum(context.Context, *NewCity1) (*RespFulcrum2, error)
	RelojesBrokerFulcrum(context.Context, *Planeta) (*RespFulcrum1, error)
	ConsistenciaEventual(context.Context, *SolMerge) (*RespBroker3, error)
	PreguntarRelojesYRegistros(context.Context, *SolMerge) (*RelojesYRegistros, error)
	MandarFulcrums(context.Context, *RelojesYRegistros) (*SolMerge, error)
	mustEmbedUnimplementedStarWars1Server()
}

// UnimplementedStarWars1Server must be embedded to have forward compatible implementations.
type UnimplementedStarWars1Server struct {
}

func (UnimplementedStarWars1Server) CityMgmtFulcrum(context.Context, *NewCity1) (*RespFulcrum1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CityMgmtFulcrum not implemented")
}
func (UnimplementedStarWars1Server) CityBrokerFulcrum(context.Context, *NewCity1) (*RespFulcrum2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CityBrokerFulcrum not implemented")
}
func (UnimplementedStarWars1Server) RelojesBrokerFulcrum(context.Context, *Planeta) (*RespFulcrum1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelojesBrokerFulcrum not implemented")
}
func (UnimplementedStarWars1Server) ConsistenciaEventual(context.Context, *SolMerge) (*RespBroker3, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsistenciaEventual not implemented")
}
func (UnimplementedStarWars1Server) PreguntarRelojesYRegistros(context.Context, *SolMerge) (*RelojesYRegistros, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreguntarRelojesYRegistros not implemented")
}
func (UnimplementedStarWars1Server) MandarFulcrums(context.Context, *RelojesYRegistros) (*SolMerge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarFulcrums not implemented")
}
func (UnimplementedStarWars1Server) mustEmbedUnimplementedStarWars1Server() {}

// UnsafeStarWars1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StarWars1Server will
// result in compilation errors.
type UnsafeStarWars1Server interface {
	mustEmbedUnimplementedStarWars1Server()
}

func RegisterStarWars1Server(s grpc.ServiceRegistrar, srv StarWars1Server) {
	s.RegisterService(&StarWars1_ServiceDesc, srv)
}

func _StarWars1_CityMgmtFulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCity1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarWars1Server).CityMgmtFulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starwars2.StarWars1/CityMgmtFulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarWars1Server).CityMgmtFulcrum(ctx, req.(*NewCity1))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarWars1_CityBrokerFulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCity1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarWars1Server).CityBrokerFulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starwars2.StarWars1/CityBrokerFulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarWars1Server).CityBrokerFulcrum(ctx, req.(*NewCity1))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarWars1_RelojesBrokerFulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Planeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarWars1Server).RelojesBrokerFulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starwars2.StarWars1/RelojesBrokerFulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarWars1Server).RelojesBrokerFulcrum(ctx, req.(*Planeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarWars1_ConsistenciaEventual_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolMerge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarWars1Server).ConsistenciaEventual(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starwars2.StarWars1/ConsistenciaEventual",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarWars1Server).ConsistenciaEventual(ctx, req.(*SolMerge))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarWars1_PreguntarRelojesYRegistros_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolMerge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarWars1Server).PreguntarRelojesYRegistros(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starwars2.StarWars1/preguntarRelojesYRegistros",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarWars1Server).PreguntarRelojesYRegistros(ctx, req.(*SolMerge))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarWars1_MandarFulcrums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelojesYRegistros)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarWars1Server).MandarFulcrums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starwars2.StarWars1/mandarFulcrums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarWars1Server).MandarFulcrums(ctx, req.(*RelojesYRegistros))
	}
	return interceptor(ctx, in, info, handler)
}

// StarWars1_ServiceDesc is the grpc.ServiceDesc for StarWars1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StarWars1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "starwars2.StarWars1",
	HandlerType: (*StarWars1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CityMgmtFulcrum",
			Handler:    _StarWars1_CityMgmtFulcrum_Handler,
		},
		{
			MethodName: "CityBrokerFulcrum",
			Handler:    _StarWars1_CityBrokerFulcrum_Handler,
		},
		{
			MethodName: "RelojesBrokerFulcrum",
			Handler:    _StarWars1_RelojesBrokerFulcrum_Handler,
		},
		{
			MethodName: "ConsistenciaEventual",
			Handler:    _StarWars1_ConsistenciaEventual_Handler,
		},
		{
			MethodName: "preguntarRelojesYRegistros",
			Handler:    _StarWars1_PreguntarRelojesYRegistros_Handler,
		},
		{
			MethodName: "mandarFulcrums",
			Handler:    _StarWars1_MandarFulcrums_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "starwars2/starwars2.proto",
}
