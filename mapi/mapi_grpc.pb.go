// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package math_over_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MathOverGrpcClient is the client API for MathOverGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathOverGrpcClient interface {
	//supports only one method
	DoOperation(ctx context.Context, in *OpRequest, opts ...grpc.CallOption) (*OpResponse, error)
}

type mathOverGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewMathOverGrpcClient(cc grpc.ClientConnInterface) MathOverGrpcClient {
	return &mathOverGrpcClient{cc}
}

func (c *mathOverGrpcClient) DoOperation(ctx context.Context, in *OpRequest, opts ...grpc.CallOption) (*OpResponse, error) {
	out := new(OpResponse)
	err := c.cc.Invoke(ctx, "/math_over_grpc.MathOverGrpc/DoOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathOverGrpcServer is the server API for MathOverGrpc service.
// All implementations must embed UnimplementedMathOverGrpcServer
// for forward compatibility
type MathOverGrpcServer interface {
	//supports only one method
	DoOperation(context.Context, *OpRequest) (*OpResponse, error)
	mustEmbedUnimplementedMathOverGrpcServer()
}

// UnimplementedMathOverGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedMathOverGrpcServer struct {
}

func (UnimplementedMathOverGrpcServer) DoOperation(context.Context, *OpRequest) (*OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoOperation not implemented")
}
func (UnimplementedMathOverGrpcServer) mustEmbedUnimplementedMathOverGrpcServer() {}

// UnsafeMathOverGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathOverGrpcServer will
// result in compilation errors.
type UnsafeMathOverGrpcServer interface {
	mustEmbedUnimplementedMathOverGrpcServer()
}

func RegisterMathOverGrpcServer(s grpc.ServiceRegistrar, srv MathOverGrpcServer) {
	s.RegisterService(&_MathOverGrpc_serviceDesc, srv)
}

func _MathOverGrpc_DoOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathOverGrpcServer).DoOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math_over_grpc.MathOverGrpc/DoOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathOverGrpcServer).DoOperation(ctx, req.(*OpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathOverGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "math_over_grpc.MathOverGrpc",
	HandlerType: (*MathOverGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoOperation",
			Handler:    _MathOverGrpc_DoOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mapi/mapi.proto",
}
