// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.9
// source: billing.proto

package api

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
	BilligService_ProcessPayment_FullMethodName = "/api.BilligService/ProcessPayment"
	BilligService_HandleWebhook_FullMethodName  = "/api.BilligService/HandleWebhook"
)

// BilligServiceClient is the client API for BilligService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BilligServiceClient interface {
	ProcessPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	HandleWebhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookReponse, error)
}

type billigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBilligServiceClient(cc grpc.ClientConnInterface) BilligServiceClient {
	return &billigServiceClient{cc}
}

func (c *billigServiceClient) ProcessPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, BilligService_ProcessPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billigServiceClient) HandleWebhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookReponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookReponse)
	err := c.cc.Invoke(ctx, BilligService_HandleWebhook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BilligServiceServer is the server API for BilligService service.
// All implementations must embed UnimplementedBilligServiceServer
// for forward compatibility.
type BilligServiceServer interface {
	ProcessPayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	HandleWebhook(context.Context, *WebhookRequest) (*WebhookReponse, error)
	mustEmbedUnimplementedBilligServiceServer()
}

// UnimplementedBilligServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBilligServiceServer struct{}

func (UnimplementedBilligServiceServer) ProcessPayment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedBilligServiceServer) HandleWebhook(context.Context, *WebhookRequest) (*WebhookReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleWebhook not implemented")
}
func (UnimplementedBilligServiceServer) mustEmbedUnimplementedBilligServiceServer() {}
func (UnimplementedBilligServiceServer) testEmbeddedByValue()                       {}

// UnsafeBilligServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BilligServiceServer will
// result in compilation errors.
type UnsafeBilligServiceServer interface {
	mustEmbedUnimplementedBilligServiceServer()
}

func RegisterBilligServiceServer(s grpc.ServiceRegistrar, srv BilligServiceServer) {
	// If the following call pancis, it indicates UnimplementedBilligServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BilligService_ServiceDesc, srv)
}

func _BilligService_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilligServiceServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BilligService_ProcessPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilligServiceServer).ProcessPayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BilligService_HandleWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilligServiceServer).HandleWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BilligService_HandleWebhook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilligServiceServer).HandleWebhook(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BilligService_ServiceDesc is the grpc.ServiceDesc for BilligService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BilligService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BilligService",
	HandlerType: (*BilligServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPayment",
			Handler:    _BilligService_ProcessPayment_Handler,
		},
		{
			MethodName: "HandleWebhook",
			Handler:    _BilligService_HandleWebhook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "billing.proto",
}
