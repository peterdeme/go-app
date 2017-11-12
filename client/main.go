package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/peterdeme/go-app/generatedcode"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Peter"})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Message)
}
