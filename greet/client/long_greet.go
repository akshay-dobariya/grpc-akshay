package main

import (
	"context"
	pb "grpc-akshay/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Akshay"},
		{FirstName: "Yagnik"},
		{FirstName: "Dobariya"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while LongGreet %s", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req %v\n", req)
		stream.Send(req)
		time.Sleep(5 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed in recv")
	}
	log.Printf("RESPONSE: %s\n", res.Result)
}
