package notification

import (
	"context"
	pb_push "github.com/onezerobinary/push-box/proto"
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

func (s *PushServiceServer) SendNotifications(ctx context.Context, info *pb_push.Info) (*pb_push.PushResponse, error) {

	statusCode, err := SendNotification(info)

	response := pb_push.PushResponse{}

	if err != nil {
		response.Code = int32(*statusCode)
		return &response, err
	}

	return &response, nil
}

func (s *PushServiceServer) StopNotifications (ctx context.Context, stop *pb_push.Stop) (*pb_push.StopResponse, error) {

	stopResponse, err := StopNotifications(stop)

	if err != nil {
		stopResponse.IsClosed = false
	}

	return stopResponse, nil
}