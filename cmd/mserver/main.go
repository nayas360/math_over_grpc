package main

import (
	"flag"
	"log"

	mserver "github.com/nayas360/math_over_grpc"
)

func main() {
	address := flag.String("address", "localhost:9999", "MApiServer address")
	flag.Parse()
	log.Printf("creating MApiServer to serve at address %s", *address)
	server, err := mserver.NewMApiServer(*address)
	if err != nil {
		log.Fatalf("could not create a MApiServer at address %s: %s", *address, err)
	}
	log.Printf("starting MApiServer at address %s", *address)
	if err := server.Serve(); err != nil {
		log.Fatalf("MApiServer could not serve at address %s: %s", *address, err)
	}
}
