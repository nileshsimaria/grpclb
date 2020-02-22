package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/nileshsimaria/grpclb/example-code/timeserver/timep"
)

var (
	port     = flag.String("port", ":50051", "grpc server port number")
	insecure = flag.Bool("insecure", false, "startup without TLS")
)

type server struct {
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

// grpc health check
type healthServer struct{}

func (s *healthServer) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	log.Printf("received healthServer.Check()")
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}
func (s *healthServer) Watch(in *healthpb.HealthCheckRequest, srv healthpb.Health_WatchServer) error {
	log.Printf("received healthServer.Watch()")
	return status.Error(codes.Unimplemented, "Watch is not implemented")
}

func (s *server) GetTime(ctx context.Context, in *timep.TimeRequest) (*timep.TimeReply, error) {
	log.Printf("GetTime Req from %s", in.GetName())
	hName, err := os.Hostname()
	if err != nil {
		hName = "unknown"
	}
	out := &timep.TimeReply{
		Time: fmt.Sprintf("[time=%s] [host=%s]", time.Now().String(), hName),
	}
	return out, nil
}

func (s *server) GetTimeSOut(in *timep.TimeRequest, srv timep.TimeServer_GetTimeSOutServer) error {
	log.Printf("GetTimeStreamOut Req from %s", in.GetName())
	hName, err := os.Hostname()
	if err != nil {
		hName = "unknown"
	}

	ctx := srv.Context()
	var hdr = metadata.MD{
		"k1": []string{"v1"},
	}
	if err := grpc.SendHeader(ctx, hdr); err != nil {
		log.Fatalf("%v", err)
	}
	for i := 1; i <= 10; i++ {
		out := &timep.TimeReply{
			Time: fmt.Sprintf("[#%d][time=%s] [host=%s]", i, time.Now().String(), hName),
		}
		srv.Send(out)
	}
	hdr = metadata.MD{
		"k2": []string{"v2"},
	}
	grpc.SetTrailer(ctx, hdr)

	return nil
}

func (s *server) GetTimeSIn(srv timep.TimeServer_GetTimeSInServer) error {
	hName, err := os.Hostname()
	if err != nil {
		hName = "unknown"
	}

	i := 0
	for {
		_, err := srv.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		i++
	}
	err = srv.SendAndClose(&timep.TimeReply{
		Time: fmt.Sprintf("[InStreamMsg count #%d][time=%s] [host=%s]", i, time.Now().String(), hName),
	})
	return nil
}

func (s *server) GetTimeSInSOut(srv timep.TimeServer_GetTimeSInSOutServer) error {
	hName, err := os.Hostname()
	if err != nil {
		hName = "unknown"
	}

	i := 0
	for {
		_, err := srv.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		i++
		if i%1000 == 0 {
			err = srv.Send(&timep.TimeReply{
				Time: fmt.Sprintf("[InOUTStreamMsg count #%d][time=%s] [host=%s]", i, time.Now().String(), hName),
			})
		}
	}

	return nil
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handleHTTP)
	http.HandleFunc("/_ah/health", handleHTTP)

	s := grpc.NewServer()
	timep.RegisterTimeServerServer(s, &server{})
	healthpb.RegisterHealthServer(s, &healthServer{})

	muxHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cType := r.Header.Get("Content-Type")
		if r.ProtoMajor == 2 &&
			strings.HasPrefix(cType, "application/grpc") {
			s.ServeHTTP(w, r)
			return
		}
		http.DefaultServeMux.ServeHTTP(w, r)
	})
	log.Printf("Starting gRPC server(port=%v)", *port)
	if *insecure == false {
		log.Fatal(http.ListenAndServeTLS(*port, "server.crt", "server.key", h2c.NewHandler(muxHandler, &http2.Server{})))
	}
	log.Fatal(http.ListenAndServe(*port, h2c.NewHandler(muxHandler, &http2.Server{})))
}
