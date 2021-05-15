package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	calpb "github.com/rprajapati0067/grpc-go-example/calculator/calpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (*server) SquareRoot(ctx context.Context, req *calpb.SquareRootRequest) (*calpb.SquareRootResponse, error) {
	fmt.Println("Request SquareRoot RPC")

	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &calpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
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
