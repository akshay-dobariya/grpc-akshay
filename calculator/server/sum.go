package main

import (
	"context"
	"log"

	pb "grpc-akshay/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculatorSumRequest) (*pb.CalculatorSumResponse, error) {
	log.Printf("Sum has been Invoked INPUT: %+v", in)
	return &pb.CalculatorSumResponse{
		Result: in.Var1 + in.Var2,
	}, nil
}
