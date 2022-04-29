package main

import (
	"context"
	pb "grpc-akshay/calculator/proto"
	"io"
	"log"
	"sync"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Failed in starting stream %s", err)
	}
	reqs := []pb.CalculatorMaxRequest{
		{Num: 12},
		{Num: 2},
		{Num: 67},
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer stream.CloseSend()
		for _, req := range reqs {
			log.Printf("Sending ... %d", req.Num)
			err = stream.Send(&req)
			if err != nil {
				log.Fatalf("Failed in sending")
			}
			time.Sleep(time.Second * 5)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed in getting resp")
				break
			}
			log.Println("RESP MAX: %d", res.Result)
		}
	}()
	wg.Wait()
}
