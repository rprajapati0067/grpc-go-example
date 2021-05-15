package main

import (
	"context"
	"fmt"
	"log"

	calpb "github.com/rprajapati0067/grpc-go-example/calculator/calpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello I'm client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect: %v", err)
	}
	defer cc.Close()
	c := calpb.NewCalculatorServiceClient(cc)

	//doSum(c)
	doErrorUnary(c)
}

func doErrorUnary(c calpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}

func doErrorCall(c calpb.CalculatorServiceClient, n int32) {
	// correct call
	res, err := c.SquareRoot(context.Background(), &calpb.SquareRootRequest{Number: n})
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message rom server: %v", resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			} else {
				log.Fatalf("Big Error calling SquareRoot: %v", err)
				return
			}
		}
	}
	fmt.Printf("Result of SquareRoot of %v: %v\n", n, res.GetNumberRoot())
}

func doSum(c calpb.CalculatorServiceClient) {
	req := &calpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling sum RPC: %v", err)
	}

	fmt.Printf("Response from Sum RPC: %v", res.SumResult)

}
