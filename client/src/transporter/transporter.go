package transporter

import (
	"context"
	"fmt"
	"log"

	"github.com/anwesh-b/MeroDB/client/pb"
	"google.golang.org/grpc"
)

func TransportQuery(sessionId string, query string) string {
	fmt.Println("For sessionId: ", sessionId)
	fmt.Println("Sending query: ", query)
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	fmt.Print("\nGetting session Id\n")

	client := pb.NewQueryServiceClient(cc)

	request := &pb.QueryServiceRequest{}

	resp, err := client.ExamineAndExecuteQuery(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	return resp.GetResult()
}

func ExamineAndExecuteQuery(query string) string {
	return query
}
