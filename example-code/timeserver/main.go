package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"

	"github.com/nileshsimaria/grpclb/example-code/timeserver/timep"
)

var (
	port = flag.String("port", ":50051", "grpc server port number")
)

type server struct {
}

func (s *server) GetTime(ctx context.Context, in *timep.TimeRequest) (*timep.TimeReply, error) {
	log.Printf("Req from %s", in.GetName())
	hName, err := os.Hostname()
	if err != nil {
		hName = "unknown"
	}
	out := &timep.TimeReply{
		Time: fmt.Sprintf("[time=%s] [host=%s]", time.Now().String(), hName),
	}
	return out, nil
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("%s", err)
	}

	s := grpc.NewServer()
	timep.RegisterTimeServerServer(s, &server{})
	s.Serve(l)
}
