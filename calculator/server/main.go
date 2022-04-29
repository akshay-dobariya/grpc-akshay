package main

import (
	"log"
	"net"

	pb "grpc-akshay/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listner, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed in initialising listner on '%s': %s", err)
	}
	log.Printf("Listening on %s", addr)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(listner); err != nil {
		log.Fatalf("Failed in serve %s", err)
	}
}
