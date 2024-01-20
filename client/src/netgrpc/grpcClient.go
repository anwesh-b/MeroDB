package netgrpc

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/anwesh-b/MeroDB/client/pb"
	transporter "github.com/anwesh-b/MeroDB/client/src/transporter"
	strings "github.com/anwesh-b/MeroDB/lib/string"
	"google.golang.org/grpc"
)

func getSessionId() string {
	opts := grpc.WithInsecure()
	fmt.Println("nice")
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal("there is an error")

		log.Fatal(err)
	}
	defer cc.Close()
	fmt.Print("\nGetting session Id\n")

	client := pb.NewSessionServiceClient(cc)

	request := &pb.SessionServiceRequest{
		Message: "Hello from client",
	}

	resp, err := client.GetSessionId(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	return resp.GetSessionId()
}

func readerAndTransporter(sessionID string) {
	fmt.Println("For sessionId: ", sessionID)
	for {
		fmt.Print("MeroDB > ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		text = strings.TrimString(text)

		if text == ".exit" {
			break
		}

		res := transporter.TransportQuery(sessionID, text)

		fmt.Println(res)
	}

}

func InitClient() {
	sId := getSessionId()

	readerAndTransporter(sId)
}
