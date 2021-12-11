package main

import (
	"context"
	pb "github.com/kaantecik/grpc-demo/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myTestServer struct {
	pb.UnimplementedTestServer
}

func (s *myTestServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println(in.GetName())

	return &pb.HelloReply{
		Message: "Yey!" + in.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterTestServer(srv, &myTestServer{})

	log.Println("starting server...")
	log.Fatalln(srv.Serve(lis))
}
