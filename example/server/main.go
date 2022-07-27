package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	hrmpb "github.com/chetanborase/gRPC-server-side-streaming/grpc/gen/go/HeartRateMonitor/v1"
)

// HeartRateMonitorService
// Lets implement HeartRateMonitorService by complying service definition given in heart-rate-monitor-service.proto.proto file
type HeartRateMonitorService struct {
	hrmpb.BeatsPerMinuteRequest
}

func main() {
	// create listener that is required for our grpc server
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create new server
	grpcServer := grpc.NewServer()

	//grpc reflection to discover endpoints from our server.
	reflection.Register(grpcServer)

	//we have listener, server and service
	//however the service is still not bounded to server
	//lets do this by Registering out implemented service to the grpc service.
	//hrm.RegisterGreetServiceServer(grpcServer, &GreetService{})

	log.Println("Starting grpc server.")
	//lets start serving
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server with error %v", err)
	}
}
