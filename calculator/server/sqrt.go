package main

import (
	"context"
	"fmt"
	pb "grpc-akshay/calculator/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.CalculatorSqrtRequest) (*pb.CalculatorSqrtResponse, error) {
	log.Printf("Sqrt invoked with %s", in)
	number := in.Num
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a -ve number %d", number),
		)
	}
	return &pb.CalculatorSqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
