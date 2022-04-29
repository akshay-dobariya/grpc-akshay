package main

import (
	pb "grpc-akshay/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet invoked")
	res := "Hello, "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading the client stream %s\n", err)
			return err
		}
		log.Printf("Got req %s\n", req)
		res += req.FirstName
	}

}
