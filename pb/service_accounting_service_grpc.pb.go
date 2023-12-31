// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: service_accounting_service.proto

package pb

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

// AccountingClient is the client API for Accounting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountingClient interface {
	CreateClientProfile(ctx context.Context, in *CreateClientProfileRequest, opts ...grpc.CallOption) (*CreateClientProfileResponse, error)
	CreateClientToken(ctx context.Context, in *CreateClientTokenRequest, opts ...grpc.CallOption) (*CreateClientTokenResponse, error)
}

type accountingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingClient(cc grpc.ClientConnInterface) AccountingClient {
	return &accountingClient{cc}
}

func (c *accountingClient) CreateClientProfile(ctx context.Context, in *CreateClientProfileRequest, opts ...grpc.CallOption) (*CreateClientProfileResponse, error) {
	out := new(CreateClientProfileResponse)
	err := c.cc.Invoke(ctx, "/pb.Accounting/CreateClientProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) CreateClientToken(ctx context.Context, in *CreateClientTokenRequest, opts ...grpc.CallOption) (*CreateClientTokenResponse, error) {
	out := new(CreateClientTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.Accounting/CreateClientToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServer is the server API for Accounting service.
// All implementations must embed UnimplementedAccountingServer
// for forward compatibility
type AccountingServer interface {
	CreateClientProfile(context.Context, *CreateClientProfileRequest) (*CreateClientProfileResponse, error)
	CreateClientToken(context.Context, *CreateClientTokenRequest) (*CreateClientTokenResponse, error)
	mustEmbedUnimplementedAccountingServer()
}

// UnimplementedAccountingServer must be embedded to have forward compatible implementations.
type UnimplementedAccountingServer struct {
}

func (UnimplementedAccountingServer) CreateClientProfile(context.Context, *CreateClientProfileRequest) (*CreateClientProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientProfile not implemented")
}
func (UnimplementedAccountingServer) CreateClientToken(context.Context, *CreateClientTokenRequest) (*CreateClientTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientToken not implemented")
}
func (UnimplementedAccountingServer) mustEmbedUnimplementedAccountingServer() {}

// UnsafeAccountingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountingServer will
// result in compilation errors.
type UnsafeAccountingServer interface {
	mustEmbedUnimplementedAccountingServer()
}

func RegisterAccountingServer(s grpc.ServiceRegistrar, srv AccountingServer) {
	s.RegisterService(&Accounting_ServiceDesc, srv)
}

func _Accounting_CreateClientProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).CreateClientProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Accounting/CreateClientProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).CreateClientProfile(ctx, req.(*CreateClientProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_CreateClientToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).CreateClientToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Accounting/CreateClientToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).CreateClientToken(ctx, req.(*CreateClientTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounting_ServiceDesc is the grpc.ServiceDesc for Accounting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Accounting",
	HandlerType: (*AccountingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientProfile",
			Handler:    _Accounting_CreateClientProfile_Handler,
		},
		{
			MethodName: "CreateClientToken",
			Handler:    _Accounting_CreateClientToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_accounting_service.proto",
}
