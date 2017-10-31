# math_over_grpc
A gRPC server and client that can do maths

## Try it out
To get the server, do
>$ go get github.com/nayas360/math_over_grpc/cmd/mserver

To get the the client, do
>$ go get github.com/nayas360/math_over_grpc/cmd/mclient

To get both do
>$ go get github.com/nayas360/math_over_grpc/cmd/...

Run the Server by
>$ mserver

Query the server from the client
>$ mclient -op add 10 20  
$ mclient -op sub 20 30  
$ mclient -op mul 30 40  
$ mclient -op div 40 50

The proto definition in mapi folder the generated Go bindings are provided with the source,
you are free to generate your own bindings
