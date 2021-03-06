// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: billetterie.proto

package billetterie

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

// BilletterieClient is the client API for Billetterie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BilletterieClient interface {
	Issue(ctx context.Context, in *IssueRequest, opts ...grpc.CallOption) (*IssueResponse, error)
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type billetterieClient struct {
	cc grpc.ClientConnInterface
}

func NewBilletterieClient(cc grpc.ClientConnInterface) BilletterieClient {
	return &billetterieClient{cc}
}

func (c *billetterieClient) Issue(ctx context.Context, in *IssueRequest, opts ...grpc.CallOption) (*IssueResponse, error) {
	out := new(IssueResponse)
	err := c.cc.Invoke(ctx, "/Billetterie/Issue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billetterieClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/Billetterie/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billetterieClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/Billetterie/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BilletterieServer is the server API for Billetterie service.
// All implementations must embed UnimplementedBilletterieServer
// for forward compatibility
type BilletterieServer interface {
	Issue(context.Context, *IssueRequest) (*IssueResponse, error)
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	mustEmbedUnimplementedBilletterieServer()
}

// UnimplementedBilletterieServer must be embedded to have forward compatible implementations.
type UnimplementedBilletterieServer struct {
}

func (UnimplementedBilletterieServer) Issue(context.Context, *IssueRequest) (*IssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedBilletterieServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedBilletterieServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedBilletterieServer) mustEmbedUnimplementedBilletterieServer() {}

// UnsafeBilletterieServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BilletterieServer will
// result in compilation errors.
type UnsafeBilletterieServer interface {
	mustEmbedUnimplementedBilletterieServer()
}

func RegisterBilletterieServer(s grpc.ServiceRegistrar, srv BilletterieServer) {
	s.RegisterService(&Billetterie_ServiceDesc, srv)
}

func _Billetterie_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilletterieServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Billetterie/Issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilletterieServer).Issue(ctx, req.(*IssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billetterie_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilletterieServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Billetterie/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilletterieServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billetterie_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BilletterieServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Billetterie/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BilletterieServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Billetterie_ServiceDesc is the grpc.ServiceDesc for Billetterie service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Billetterie_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Billetterie",
	HandlerType: (*BilletterieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issue",
			Handler:    _Billetterie_Issue_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Billetterie_Transfer_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Billetterie_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "billetterie.proto",
}
