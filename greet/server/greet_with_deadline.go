package main

import (
	"context"
	"fmt"
	pb "grpc-akshay/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("GreetWithDeadline invoked")
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the req")
			return nil, status.Errorf(
				codes.Canceled,
				fmt.Sprintf("Deadline exceeded"),
			)
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "hello ----- " + in.FirstName,
	}, nil
}
