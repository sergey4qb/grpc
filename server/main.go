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

//type Server struct {
//	pb.UnimplementedUserServer
//}
//
////func (s *Server) CreateUser(ctx context.Context, data *pb.UserData) (*pb.Response, error) {
////
////	return &pb.Response{
////		Name: data.Name,
////		Surname:  data.Surname,
////	}, nil
////}
////
////func (s *Server) mustEmbedUnimplementedUserServer() {
////	panic("implement me")
////}
//
////func (r *User) CreateUser(context.Context, *pb.UserData) (*pb.Response, error) {
////
////}

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
