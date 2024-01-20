package netgrpc

import (
	"context"
	"fmt"

	pb "github.com/anwesh-b/MeroDB/server/pb"
	"google.golang.org/grpc/peer"
)

func generateUniqueSessionId() string {
	return "1234"
}

func (s *server) GetSessionId(ctx context.Context, req *pb.SessionServiceRequest) (*pb.SessionServiceResponse, error) {
	fmt.Println("\nGenerating new unique client Id")
	p, _ := peer.FromContext(ctx)
	fmt.Println("\nnGenerating new unique client Id for request from Client IP:", p.Addr.String())
	fmt.Println("Logging req: ", req)
	msg := req.GetMessage()

	fmt.Println("Message from client: ", msg)

	sId := generateUniqueSessionId()

	return &pb.SessionServiceResponse{
		SessionId: sId,
	}, nil
}
