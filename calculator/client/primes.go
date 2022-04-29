package main

import (
	"context"
	"io"
	"log"

	pb "grpc-akshay/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("Invoked doPrimes")
	req := pb.CalculatorPrimesRequest{
		Num: 12545,
	}
	stream, err := c.Primes(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed in getting stream %s\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed in getting message from stream %s\n", err)
		}
		log.Printf("VALUES: %d\n", msg.Result)
	}
}
