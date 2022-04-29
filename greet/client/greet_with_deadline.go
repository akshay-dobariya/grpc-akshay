package main

import (
	"context"
	pb "grpc-akshay/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req := &pb.GreetRequest{
		FirstName: "Akshay",
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("got exp err %s", err)
				return
			} else {
				log.Printf("unexp gRPC err %s", err)
				return
			}
		} else {
			log.Fatalf("non gRPC err %s", err)
		}
	}
	log.Printf("RES: %s", res.Result)
}
