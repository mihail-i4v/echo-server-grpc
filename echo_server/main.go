// Package main implements a server for Echo service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "go-grpc-echo/echo"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement echo.EchoService.
type server struct {
	pb.UnimplementedEchoServiceServer
}

// Echo implements EchoService.Echo
func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.String())
	return &pb.EchoResponse{Output: in.GetInput()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
