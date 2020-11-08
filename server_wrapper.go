package math_over_grpc

import (
	"context"
	"fmt"
	"math/big"
	"net"

	pb "github.com/nayas360/math_over_grpc/mapi"
	"google.golang.org/grpc"
)

// A basic abstraction over the grpc server construct
type MApiServer struct {
	pb.UnimplementedMathOverGrpcServer
	listener net.Listener
	server   *grpc.Server
}

// returns a new server
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

// starts serving
func (ms *MApiServer) Serve() error {
	if e := ms.server.Serve(ms.listener); e != nil {
		return e
	}
	return nil
}

// the handler for the rpc
func (ms *MApiServer) DoOperation(ctx context.Context, in *pb.OpRequest) (*pb.OpResponse, error) {
	switch in.Op {
	/*case pb.Opcode_NOP:
	return nil, fmt.Errorf("no operation selected")*/
	case pb.Opcode_ADD:
		Unary, _ := new(big.Float).SetString(in.Unary)
		Binary, _ := new(big.Float).SetString(in.Binary)
		return &pb.OpResponse{Result: Unary.Add(Unary, Binary).String()}, nil
	case pb.Opcode_SUB:
		Unary, _ := new(big.Float).SetString(in.Unary)
		Binary, _ := new(big.Float).SetString(in.Binary)
		return &pb.OpResponse{Result: Unary.Sub(Unary, Binary).String()}, nil
	case pb.Opcode_MUL:
		Unary, _ := new(big.Float).SetString(in.Unary)
		Binary, _ := new(big.Float).SetString(in.Binary)
		return &pb.OpResponse{Result: Unary.Mul(Unary, Binary).String()}, nil
	case pb.Opcode_DIV:
		Unary, _ := new(big.Float).SetString(in.Unary)
		Binary, _ := new(big.Float).SetString(in.Binary)
		return &pb.OpResponse{Result: Unary.Quo(Unary, Binary).String()}, nil
	}

	return nil, fmt.Errorf("no operation selected")
}
