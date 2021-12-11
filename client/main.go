package main

import (
	"context"
	pb "github.com/kaantecik/grpc-demo/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	conn, err := grpc.Dial(net.JoinHostPort("localhost", "8080"), grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	client := pb.NewTestClient(conn)

	ctx := context.Background()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "kaan"})

	if err != nil {
		panic(err)
	}

	log.Println(resp.GetMessage())
}
