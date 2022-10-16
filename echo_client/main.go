// Package main implements a client for Echo service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "go-grpc-echo/echo"
)

const (
	defaultInput = "bar"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	input = flag.String("input", defaultInput, "input to pass")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{Input: *input})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Output: %s", r.GetOutput())
}
