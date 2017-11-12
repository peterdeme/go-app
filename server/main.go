package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/peterdeme/go-app/generatedcode"
	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":5000")

	defer lis.Close()

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &GreeterServer{})

	log.SetOutput(os.Stdout)

	log.Println("Starting the server.")

	s.Serve(lis)

	waitForStopSignal(s)
}

func waitForStopSignal(s *grpc.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, os.Interrupt)
	log.Println("Terminating server:", <-c)
	s.GracefulStop()
}

// GreeterServer represents the Greeter service
type GreeterServer struct{}

// SayHello returns HelloReply to the given HelloRequest
func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}
