package main

import (
	"context"
	"fmt"
	"log"

	calpb "github.com/rprajapati0067/grpc-go-example/calculator/calpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect: %v", err)
	}
	defer cc.Close()
	c := calpb.NewCalculatorServiceClient(cc)

	doSum(c)
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
