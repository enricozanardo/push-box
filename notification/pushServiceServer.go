package notification

import (
	"context"
	pb_push "github.com/onezerobinary/push-box/proto"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/goinggo/tracelog"
	"net"
	"os"
	"google.golang.org/grpc"
)

const (
	GRPC_PORT = ":1972"
)

type PushServiceServer struct {

}

func StartPushService(){

	// Start the Push Service
	listen, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		tracelog.Errorf(err, "app", "main", "Failed to start the service")
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	// Add to the grpcServer the Service
	pb_push.RegisterPushServiceServer(grpcServer, &PushServiceServer{})

	tracelog.Trace("main", "main", "Grpc Server Listening on port 1972")

	grpcServer.Serve(listen)
}

func (s *PushServiceServer) SendNotifications(ctx context.Context, accountToken *pb_push.AccountToken) (*pb_push.PushResponse, error) {

	tokens := []*pb_account.Token{}

	for _, token := range accountToken.Tokens {
		tmpTopken := pb_account.Token{token}
		tokens = append(tokens, &tmpTopken)
	}

	statusCode, err := SendNotification(tokens)

	response := pb_push.PushResponse{}

	if err != nil {
		response.Code = int32(*statusCode)
		return &response, err
	}

	return &response, nil
}
