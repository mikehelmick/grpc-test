package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	cpb "github.com/mikehelmick/grpc-test/counter/proto"

	grpc "google.golang.org/grpc"
)

var port = flag.Int("port", 3232, "port number to listen on")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	cpb.RegisterEchoServer(grpcServer, NewServer())
	grpcServer.Serve(lis)
}
