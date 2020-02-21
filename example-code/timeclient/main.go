package main

import (
	"context"
	"flag"
	"log"

	"github.com/nileshsimaria/grpclb/example-code/timeclient/timep"
	"google.golang.org/grpc"
)

var (
	host = flag.String("host", "localhost:50051", "grpc server host:port")
	n    = flag.Int("count", 1, "number of rpc calls")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := timep.NewTimeServerClient(conn)

	for i := 0; i < *n; i++ {
		r, err := client.GetTime(context.Background(), &timep.TimeRequest{
			Name: "test",
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Printf("Reply: %v", r)
	}
}
