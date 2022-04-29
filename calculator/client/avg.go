package main

import (
	"context"
	pb "grpc-akshay/calculator/proto"
	"log"
	"time"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("invoked doAverage")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Failed in Average call %s", err)
	}
	nums := []int64{1221, 534, 643, 786, 32, 534, 743, 568}
	for _, num := range nums {
		log.Printf("Sending %d\n", num)
		stream.Send(&pb.CalculatorAvgRequest{
			Num: num,
		})
		time.Sleep(5 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed in getting Average res %s", err)
	}
	log.Printf("RESP: %f", res.Result)
}
