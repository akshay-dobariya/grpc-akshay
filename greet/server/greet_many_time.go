package main

import (
	"fmt"
	"log"
	"time"

	pb "grpc-akshay/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes func invoked in server with '%v'", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s number:%d", in.FirstName, i)
		log.Printf("Sending %s\n", res)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
		time.Sleep(5 * time.Second)
	}
	return nil
}
