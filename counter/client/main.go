package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	cpb "github.com/mikehelmick/grpc-test/counter/proto"
)

var serverAddr = flag.String("addr", "127.0.0.1:3232", "address of server")

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := cpb.NewEchoClient(conn)

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		request := &cpb.IncrementRequest{
			Name: "counter",
		}
		result, err := client.Increment(ctx, request)
		if err != nil {
			log.Fatalf("error calling Increment: %v", err)
		}
		log.Printf("called increment: %+v", result)
	}

	conn.Close()
}
