package main

import (
	pb "grpc-akshay/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("invoked Avg")
	var result float64
	total := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CalculatorAvgResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatal("Failed in receive\n")
		}
		if total == 0 {
			result = float64(req.Num)
		} else {
			result = ((result * float64(total)) + float64(req.Num)) / (float64(total + 1))
		}
		log.Printf("We received %d now Avg is %f\n", req.Num, result)
		total++
	}
}
