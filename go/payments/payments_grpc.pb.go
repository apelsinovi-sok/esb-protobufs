// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: proto/payments.proto

package payments

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

// PaymentsClient is the client API for Payments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentsClient interface {
	GetCardsByUserID(ctx context.Context, in *ParamsGetCardByUserID, opts ...grpc.CallOption) (*ResponseGetCardByUserID, error)
	SaveUserCard(ctx context.Context, in *UserCard, opts ...grpc.CallOption) (*ResponseSuccess, error)
	DeleteCardByID(ctx context.Context, in *ParamsDeleteCardByID, opts ...grpc.CallOption) (*ResponseSuccess, error)
}

type paymentsClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsClient(cc grpc.ClientConnInterface) PaymentsClient {
	return &paymentsClient{cc}
}

func (c *paymentsClient) GetCardsByUserID(ctx context.Context, in *ParamsGetCardByUserID, opts ...grpc.CallOption) (*ResponseGetCardByUserID, error) {
	out := new(ResponseGetCardByUserID)
	err := c.cc.Invoke(ctx, "/payments.Payments/GetCardsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsClient) SaveUserCard(ctx context.Context, in *UserCard, opts ...grpc.CallOption) (*ResponseSuccess, error) {
	out := new(ResponseSuccess)
	err := c.cc.Invoke(ctx, "/payments.Payments/SaveUserCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsClient) DeleteCardByID(ctx context.Context, in *ParamsDeleteCardByID, opts ...grpc.CallOption) (*ResponseSuccess, error) {
	out := new(ResponseSuccess)
	err := c.cc.Invoke(ctx, "/payments.Payments/DeleteCardByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsServer is the server API for Payments service.
// All implementations should embed UnimplementedPaymentsServer
// for forward compatibility
type PaymentsServer interface {
	GetCardsByUserID(context.Context, *ParamsGetCardByUserID) (*ResponseGetCardByUserID, error)
	SaveUserCard(context.Context, *UserCard) (*ResponseSuccess, error)
	DeleteCardByID(context.Context, *ParamsDeleteCardByID) (*ResponseSuccess, error)
}

// UnimplementedPaymentsServer should be embedded to have forward compatible implementations.
type UnimplementedPaymentsServer struct {
}

func (UnimplementedPaymentsServer) GetCardsByUserID(context.Context, *ParamsGetCardByUserID) (*ResponseGetCardByUserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardsByUserID not implemented")
}
func (UnimplementedPaymentsServer) SaveUserCard(context.Context, *UserCard) (*ResponseSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserCard not implemented")
}
func (UnimplementedPaymentsServer) DeleteCardByID(context.Context, *ParamsDeleteCardByID) (*ResponseSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCardByID not implemented")
}

// UnsafePaymentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentsServer will
// result in compilation errors.
type UnsafePaymentsServer interface {
	mustEmbedUnimplementedPaymentsServer()
}

func RegisterPaymentsServer(s grpc.ServiceRegistrar, srv PaymentsServer) {
	s.RegisterService(&Payments_ServiceDesc, srv)
}

func _Payments_GetCardsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsGetCardByUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).GetCardsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payments.Payments/GetCardsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).GetCardsByUserID(ctx, req.(*ParamsGetCardByUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payments_SaveUserCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).SaveUserCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payments.Payments/SaveUserCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).SaveUserCard(ctx, req.(*UserCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payments_DeleteCardByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsDeleteCardByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).DeleteCardByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payments.Payments/DeleteCardByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).DeleteCardByID(ctx, req.(*ParamsDeleteCardByID))
	}
	return interceptor(ctx, in, info, handler)
}

// Payments_ServiceDesc is the grpc.ServiceDesc for Payments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payments.Payments",
	HandlerType: (*PaymentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCardsByUserID",
			Handler:    _Payments_GetCardsByUserID_Handler,
		},
		{
			MethodName: "SaveUserCard",
			Handler:    _Payments_SaveUserCard_Handler,
		},
		{
			MethodName: "DeleteCardByID",
			Handler:    _Payments_DeleteCardByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/payments.proto",
}