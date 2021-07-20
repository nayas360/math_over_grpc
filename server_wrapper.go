package math_over_grpc

import (
	"context"
	"math/big"
	"net"

	pb "github.com/nayas360/math_over_grpc/mapi"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// MApiServer A basic abstraction over the grpc server construct
type MApiServer struct {
	pb.UnimplementedMathOverGrpcServer
	listener net.Listener
	server   *grpc.Server
}

// NewMApiServer returns a new server
func NewMApiServer(address string, opts ...grpc.ServerOption) (*MApiServer, error) {
	s := grpc.NewServer(opts...)
	l, e := net.Listen("tcp", address)
	if e != nil {
		return nil, e
	}
	g := &MApiServer{listener: l, server: s}
	pb.RegisterMathOverGrpcServer(s, g)
	return g, nil
}

// Serve starts serving
func (ms *MApiServer) Serve() error {
	if e := ms.server.Serve(ms.listener); e != nil {
		return e
	}
	return nil
}

// DoOperation the handler for the rpc
func (ms *MApiServer) DoOperation(_ context.Context, in *pb.OpRequest) (*pb.OpResponse, error) {
	Unary, ok := new(big.Float).SetString(in.Unary)
	if !ok {
		Unary = Unary.SetInt64(0)
	}
	Binary, ok := new(big.Float).SetString(in.Binary)
	if !ok {
		Binary = Binary.SetInt64(0)
	}
	switch in.Op {
	case pb.Opcode_ADD:
		return &pb.OpResponse{Result: Unary.Add(Unary, Binary).String()}, nil
	case pb.Opcode_SUB:
		return &pb.OpResponse{Result: Unary.Sub(Unary, Binary).String()}, nil
	case pb.Opcode_MUL:
		return &pb.OpResponse{Result: Unary.Mul(Unary, Binary).String()}, nil
	case pb.Opcode_DIV:
		return &pb.OpResponse{Result: Unary.Quo(Unary, Binary).String()}, nil
	}

	return nil, errors.Errorf("no operation selected")
}
