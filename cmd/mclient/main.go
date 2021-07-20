package main

import (
	"flag"
	"fmt"
	"log"

	mclient "github.com/nayas360/math_over_grpc"
	"google.golang.org/grpc"
)

func main() {
	address := flag.String("address", "localhost:9999", "MApiServer address")
	operation := flag.String("op", "add", "operation to be done [add, sub, mul, div]")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("no operands were provided")
	}
	client, err := mclient.NewMApiClient(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not create MApiClient for server at address %s: %s", *address, err)
	}
	defer func(client *mclient.MApiClient) {
		if err := client.Close(); err != nil {
			log.Fatalf("could not close the client, %s", err)
		}
	}(client)
	var resp string
	switch *operation {
	case "add":
		resp, err = client.Add(flag.Arg(0), flag.Arg(1))
	case "sub":
		resp, err = client.Sub(flag.Arg(0), flag.Arg(1))
	case "mul":
		resp, err = client.Mul(flag.Arg(0), flag.Arg(1))
	case "div":
		resp, err = client.Div(flag.Arg(0), flag.Arg(1))
	default:
		log.Fatalf("expected op to be one of [add, sub, mul, div], got '%s'", *operation)
	}
	if err != nil {
		log.Fatalf("could not do operation: %s", err)
	}
	fmt.Println(resp)
}
