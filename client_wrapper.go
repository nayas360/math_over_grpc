package math_over_grpc

import (
	pb "github.com/nayas360/math_over_grpc/mapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// basic client construct
type MApiClient struct {
	conn   *grpc.ClientConn
	client pb.MathOverGrpcClient
}

// returns a new client
func NewMApiClient(address string, opts ...grpc.DialOption) (*MApiClient, error) {
	cn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	c := pb.NewMathOverGrpcClient(cn)
	g := &MApiClient{conn: cn, client: c}
	return g, nil
}

// internal wrapper over do operation rpc
func (mc *MApiClient) doOperation(Op pb.Opcode, Unary string, Binary string) (*pb.OpResponse, error) {
	ctx := context.Background()
	return mc.client.DoOperation(ctx, &pb.OpRequest{Op: Op, Unary: Unary, Binary: Binary})
}

// helper methods wrapping the do operation rpc
func (mc *MApiClient) Add(Unary string, Binary string) (string, error) {
	res, err := mc.doOperation(pb.Opcode_ADD, Unary, Binary)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// helper methods wrapping the do operation rpc
func (mc *MApiClient) Sub(Unary string, Binary string) (string, error) {
	res, err := mc.doOperation(pb.Opcode_SUB, Unary, Binary)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// helper methods wrapping the do operation rpc
func (mc *MApiClient) Mul(Unary string, Binary string) (string, error) {
	res, err := mc.doOperation(pb.Opcode_MUL, Unary, Binary)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// helper methods wrapping the do operation rpc
func (mc *MApiClient) Div(Unary string, Binary string) (string, error) {
	res, err := mc.doOperation(pb.Opcode_DIV, Unary, Binary)
	if err != nil {
		return "", err
	}
	return res.Result, nil
}

// wrapper to close the connection
func (mc *MApiClient) Close() {
	mc.conn.Close()
}
