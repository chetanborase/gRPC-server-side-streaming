package main

//import (
//	"context"
//	"fmt"
//	greet "github.com/chetanborase/grpc-greet-proto/grpc/gen/go/greeting/v1"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"log"
//)
//
//func main() {
//	//we dont have ssl certificate defined at and gRPC uses https communication protocol
//	//so, we need to tell grpc client that connect insecurely without certificate.
//
//	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
//
//	clientConn, err := grpc.Dial(":9000", opt)
//	if err != nil {
//		log.Fatalf("server unreachable, %v", err)
//	}
//
//	//don't forget to close otherwise it lead to serious performance issue when multiple client environment
//	//however, notice we didnt started establishing connection yet. we will do that shortly.
//	defer clientConn.Close()
//
//	greetClient := greet.NewGreetServiceClient(clientConn)
//
//	requestData := &greet.GreetRequest{Name: "Chetan"}
//
//	//now lets create request object
//	response, err := greetClient.SayHello(context.Background(), requestData)
//	if err != nil {
//		log.Fatalf("Something went wrong, %v", err)
//	}
//
//	//lets print our response
//	fmt.Printf("Response : %+v\n", response)
//}
