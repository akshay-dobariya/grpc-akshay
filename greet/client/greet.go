package main

import (
	"context"
	pb "grpc-akshay/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked in client")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "akshay",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s", res.Result)
}
