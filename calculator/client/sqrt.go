package main

import (
	"context"
	pb "grpc-akshay/calculator/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int64) {
	log.Printf("Invoked doSqrt")
	res, err := c.Sqrt(context.Background(), &pb.CalculatorSqrtRequest{
		Num: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We sent a -ve num")
				return
			}
		} else {
			log.Fatalf("rpc call failed non gRPC err %s", err)
		}
	}
	log.Printf("Sqrt: %+v", res.Result)
}
