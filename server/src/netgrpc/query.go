package netgrpc

import (
	"context"
	"fmt"

	pb "github.com/anwesh-b/MeroDB/server/pb"
	"github.com/anwesh-b/MeroDB/server/src/parser"
)

func executeAndReturnQueryResult(query string) string {
	fmt.Println("\nExecuting Query: ", query)

	res := parser.EvaluateInput(query)
	return res
}

func (s *server) ExamineAndExecuteQuery(ctx context.Context, req *pb.QueryServiceRequest) (*pb.QueryServiceResponse, error) {
	fmt.Println("Log req: ", req)
	fmt.Println("Log Query: ", req.Query)
	fmt.Println("Log sessionid: ", req.SessionId)
	fmt.Println("Log getquery: ", req.GetQuery())
	fmt.Println("Log getsessionid: ", req.GetSessionId())

	query := req.GetQuery()
	fmt.Println("Query: ", query)
	queryOutput := executeAndReturnQueryResult(query)

	return &pb.QueryServiceResponse{
		Result: queryOutput,
	}, nil
}
