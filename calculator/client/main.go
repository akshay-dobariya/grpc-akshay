package main

import (
	pb "grpc-akshay/calculator/proto"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Filed in initialising connection %s", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	// doPrimes(c)
	// doAverage(c)
	// doMax(c)
	doSqrt(c, -2)
}
