package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rprajapati0067/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
