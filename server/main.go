package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "training_grpc/server/server"
	training_grpc "training_grpc/server/server/User"
)
var (
	port = flag.Int("port", 50051, "The Server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:= grpc.NewServer()
	pb.RegisterUserServer(s, &training_grpc.Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
