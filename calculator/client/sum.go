package main

import (
	"context"
	pb "grpc-akshay/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("Invoked doSum")
	res, err := c.Sum(context.Background(), &pb.CalculatorSumRequest{
		Var1: 273984.348975,
		Var2: 523436.456,
	})
	if err != nil {
		log.Fatalf("rpc call failed %s", err)
	}
	log.Printf("Sum: %+v", res.Result)
}
