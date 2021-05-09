package main

import (
	"context"
	"fmt"
	"log"

	greetpb "github.com/rprajapati0067/grpc-go-example/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}
func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ravi Shankar",
			LastName:  "Prajapati",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling greet RPC: %v", err)
	}

	fmt.Printf("Response from Greet: %v", res.Result)
}
