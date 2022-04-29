package main

import (
	"context"
	pb "grpc-akshay/greet/proto"
	"io"
	"log"
	"sync"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream %s", err)
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Akshay"},
		{FirstName: "Yagnik"},
		{FirstName: "Dobariya"},
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for _, req := range reqs {
			log.Printf("Sending %s\n", req)
			stream.Send(req)
			time.Sleep(5 * time.Second)
		}
		stream.CloseSend()
		wg.Done()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while getting resp %s\n", err)
				break
			}
			log.Printf("RECV: %s\n", res.Result)
		}
		wg.Done()
	}()
	wg.Wait()
}
