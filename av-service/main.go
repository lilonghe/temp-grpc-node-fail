package main

import (
	"av-service/src/handlers"
	av "av-service/src/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	addr := flag.String("addr", ":8090", "service addr")
	flag.Parse()

	listen, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	av.RegisterAVServiceServer(grpcServer, &handlers.AvServer{})
	fmt.Println("listen -> ", *addr)
	grpcServer.Serve(listen)
}