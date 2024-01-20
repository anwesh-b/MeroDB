package netgrpc

import (
	"fmt"

	// "grpc-golang/pb"
	"log"
	"net"

	pb "github.com/anwesh-b/MeroDB/server/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.SessionServiceServer
	pb.QueryServiceServer
}

func InitServer() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterSessionServiceServer(s, &server{})
	pb.RegisterQueryServiceServer(s, &server{})

	fmt.Print("\ninitializing server \n")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
	fmt.Print("\nserver initialized\n")
}
