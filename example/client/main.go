package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	hrmpb "github.com/chetanborase/gRPC-server-side-streaming/grpc/gen/go/HeartRateMonitor/v1"
	"google.golang.org/grpc"
)

func main() {
	dialer, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := hrmpb.NewHeartRateMonitorServiceClient(dialer)

	res, err := client.BeatsPerMinute(context.Background(), &hrmpb.BeatsPerMinuteRequest{Uuid: "mario"})
	if err != nil {
		log.Fatal(err)
	}

	//use waitgroup so that dont exit from main function and all goroutine will not terminate
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			resp, err := res.Recv()
			if err == io.EOF {
				log.Println("End of stream by server")
				return
			}

			if err != nil {
				log.Fatalln("Receiving", err)
				return
			}

			//lets print the response
			fmt.Println(resp)
		}
	}()

	//lets say client want to sample the heart rate for only 10 seconds
	//and after 10 seconds it tells server to stop streaming.

	timer := time.AfterFunc(time.Duration(5*time.Second), func() { log.Fatal("closing stream from client side"); res.CloseSend(); res.Context().Done() })
	defer timer.Stop()

	//lets wait here, and goroutines will do there work
	wg.Wait()
}
