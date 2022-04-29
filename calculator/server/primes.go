package main

import (
	"log"

	pb "grpc-akshay/calculator/proto"
)

func getPrimeFactors(num uint64) *[]uint64 {
	primes := []uint64{}
	var k uint64
	k = 2
	for num > 1 {
		if num%k == 0 {
			primes = append(primes, k)
			num = num / k
		} else {
			k = k + 1
		}
	}
	return &primes
}

func (s *Server) Primes(in *pb.CalculatorPrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes invoked : %+v", in)
	primes := getPrimeFactors(in.Num)
	for _, p := range *primes {
		stream.Send(&pb.CalculatorPrimesResponse{
			Result: p,
		})
	}
	return nil
}
