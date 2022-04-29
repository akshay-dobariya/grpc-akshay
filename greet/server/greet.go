package main

import (
	"context"
	pb "grpc-akshay/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet func invoked in server with '%v'", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
