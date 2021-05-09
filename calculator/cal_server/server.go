package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calpb "github.com/rprajapati0067/grpc-go-example/calculator/calpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calpb.SumRequest) (*calpb.SumResponse, error) {
	fmt.Println("Sum function was invoked with %v", req)
	firstNum := req.GetFirstNumber()
	secondNum := req.GetSecondNumber()
	resSum := firstNum + secondNum
	res := &calpb.SumResponse{
		SumResult: resSum,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World Calculator!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
