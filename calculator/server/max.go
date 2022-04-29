package main

import (
	pb "grpc-akshay/calculator/proto"
	"io"
	"log"
	"math"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max invoked")
	max := int64(math.MinInt64)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Failed getting req from client %s", err)
		}
		log.Printf("Processing ... %d", req.Num)
		if max < req.Num {
			max = req.Num
		}
		err = stream.Send(&pb.CalculatorMaxResponse{
			Result: max,
		})
		if err != nil {
			log.Fatalf("Failed sending resp")
		}
	}
}
