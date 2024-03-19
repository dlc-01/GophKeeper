// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/bank.proto

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

const (
	Banks_CreateBank_FullMethodName = "/proto.Banks/CreateBank"
	Banks_GetBank_FullMethodName    = "/proto.Banks/GetBank"
	Banks_UpdateBank_FullMethodName = "/proto.Banks/UpdateBank"
	Banks_DeleteBank_FullMethodName = "/proto.Banks/DeleteBank"
)

// BanksClient is the client API for Banks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BanksClient interface {
	CreateBank(ctx context.Context, in *CreateBankRequest, opts ...grpc.CallOption) (*CreateBankResponse, error)
	GetBank(ctx context.Context, in *GetBankRequest, opts ...grpc.CallOption) (*GetBankResponse, error)
	UpdateBank(ctx context.Context, in *UpdateBankRequest, opts ...grpc.CallOption) (*UpdateBankResponse, error)
	DeleteBank(ctx context.Context, in *DeleteBankRequest, opts ...grpc.CallOption) (*DeleteBankResponse, error)
}

type banksClient struct {
	cc grpc.ClientConnInterface
}

func NewBanksClient(cc grpc.ClientConnInterface) BanksClient {
	return &banksClient{cc}
}

func (c *banksClient) CreateBank(ctx context.Context, in *CreateBankRequest, opts ...grpc.CallOption) (*CreateBankResponse, error) {
	out := new(CreateBankResponse)
	err := c.cc.Invoke(ctx, Banks_CreateBank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) GetBank(ctx context.Context, in *GetBankRequest, opts ...grpc.CallOption) (*GetBankResponse, error) {
	out := new(GetBankResponse)
	err := c.cc.Invoke(ctx, Banks_GetBank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) UpdateBank(ctx context.Context, in *UpdateBankRequest, opts ...grpc.CallOption) (*UpdateBankResponse, error) {
	out := new(UpdateBankResponse)
	err := c.cc.Invoke(ctx, Banks_UpdateBank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) DeleteBank(ctx context.Context, in *DeleteBankRequest, opts ...grpc.CallOption) (*DeleteBankResponse, error) {
	out := new(DeleteBankResponse)
	err := c.cc.Invoke(ctx, Banks_DeleteBank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BanksServer is the server API for Banks service.
// All implementations must embed UnimplementedBanksServer
// for forward compatibility
type BanksServer interface {
	CreateBank(context.Context, *CreateBankRequest) (*CreateBankResponse, error)
	GetBank(context.Context, *GetBankRequest) (*GetBankResponse, error)
	UpdateBank(context.Context, *UpdateBankRequest) (*UpdateBankResponse, error)
	DeleteBank(context.Context, *DeleteBankRequest) (*DeleteBankResponse, error)
	mustEmbedUnimplementedBanksServer()
}

// UnimplementedBanksServer must be embedded to have forward compatible implementations.
type UnimplementedBanksServer struct {
}

func (UnimplementedBanksServer) CreateBank(context.Context, *CreateBankRequest) (*CreateBankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBank not implemented")
}
func (UnimplementedBanksServer) GetBank(context.Context, *GetBankRequest) (*GetBankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBank not implemented")
}
func (UnimplementedBanksServer) UpdateBank(context.Context, *UpdateBankRequest) (*UpdateBankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBank not implemented")
}
func (UnimplementedBanksServer) DeleteBank(context.Context, *DeleteBankRequest) (*DeleteBankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBank not implemented")
}
func (UnimplementedBanksServer) mustEmbedUnimplementedBanksServer() {}

// UnsafeBanksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BanksServer will
// result in compilation errors.
type UnsafeBanksServer interface {
	mustEmbedUnimplementedBanksServer()
}

func RegisterBanksServer(s grpc.ServiceRegistrar, srv BanksServer) {
	s.RegisterService(&Banks_ServiceDesc, srv)
}

func _Banks_CreateBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).CreateBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_CreateBank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).CreateBank(ctx, req.(*CreateBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_GetBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).GetBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_GetBank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).GetBank(ctx, req.(*GetBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_UpdateBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).UpdateBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_UpdateBank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).UpdateBank(ctx, req.(*UpdateBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_DeleteBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).DeleteBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_DeleteBank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).DeleteBank(ctx, req.(*DeleteBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Banks_ServiceDesc is the grpc.ServiceDesc for Banks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Banks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Banks",
	HandlerType: (*BanksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBank",
			Handler:    _Banks_CreateBank_Handler,
		},
		{
			MethodName: "GetBank",
			Handler:    _Banks_GetBank_Handler,
		},
		{
			MethodName: "UpdateBank",
			Handler:    _Banks_UpdateBank_Handler,
		},
		{
			MethodName: "DeleteBank",
			Handler:    _Banks_DeleteBank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bank.proto",
}