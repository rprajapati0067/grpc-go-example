package main

import (
	"context"
	"fmt"
	"log"

	blogpb "github.com/rprajapati0067/grpc-go-example/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect: %v", err)
	}
	defer cc.Close()
	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating the Blog")
	blog := blogpb.Blog{
		AuthorId: "Stephene",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: &blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println("Blog has been created: %v", createBlogRes)
}
