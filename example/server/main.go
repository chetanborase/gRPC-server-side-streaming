package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	hrmpb "github.com/chetanborase/gRPC-server-side-streaming/grpc/gen/go/HeartRateMonitor/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// HeartRateMonitorService
// Lets implement HeartRateMonitorService by complying service definition given in heart-rate-monitor-service.proto.proto file
type HeartRateMonitorService struct {
	hrmpb.UnimplementedHeartRateMonitorServiceServer
}

func (hrms *HeartRateMonitorService) BeatsPerMinute(request *hrmpb.BeatsPerMinuteRequest, stream hrmpb.HeartRateMonitorService_BeatsPerMinuteServer) error {
	defer func() { /*lets assert if we are out of loop */ fmt.Println("Session closed.") }()

	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			time.Sleep(1 * time.Second)
			value := 30 + rand.Int31n(80)
			if err := stream.SendMsg(&hrmpb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			}); err != nil {
				return status.Error(codes.Canceled, "Stream has ended")
			}

			fmt.Println("Msg Sent.")
		}
	}

	return nil
}

func main() {
	// create listener that is required for our grpc server
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create new server
	grpcServer := grpc.NewServer()

	//grpc reflection to discover endpoints from our server.
	reflection.Register(grpcServer)

	//we have listener, server and service
	//however the service is still not bounded to server
	//lets do this by Registering out implemented service to the grpc service.ree
	hrmpb.RegisterHeartRateMonitorServiceServer(grpcServer, &HeartRateMonitorService{})

	log.Println("Starting grpc server.")
	//lets start serving
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server with error %v", err)
	}
}
