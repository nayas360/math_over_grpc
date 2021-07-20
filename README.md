# math_over_grpc
A gRPC server and client that can do maths

## Try it out
If you don't have _Go_ installed then you can download the latest binary
[release](https://github.com/nayas360/math_over_grpc/releases)
else,

To get both the server and client run
>$ go get github.com/nayas360/math_over_grpc/cmd/...

To test the setup run the server in a terminal
> $ mserver

In another terminal send queries to the server like so
> $ mclient -op add 10 20  
> $ mclient -op sub 20 30  
> $ mclient -op mul 30 40  
> $ mclient -op div 40 50

The default address for both the server, and the client is `localhost:9999`. The server can be started at a different
address by using the `-address` flag.

## Is the library available only for _Go_ ??

Fun fact _gRPC_ can generate bindings for a lot of languages and _Go_ is just one of them. You can use
the [mapi protobuf definition](./mapi/mapi.proto)
to generate bindings for your favorite language. The _Go_ binding have been generated with the following command from
the project root directory.
> $ protoc --go_out=. --go_opt=paths=source_relative
> --go-grpc_out=. --go-grpc_opt=paths=source_relative
> mapi/mapi.proto

Consult the [_gRPC_ Docs](https://grpc.io/docs/) to see how to generate bindings for other languages.
