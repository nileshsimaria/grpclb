package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/nileshsimaria/grpclb/example-code/timeclient/timep"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	host       = flag.String("host", "localhost:50051", "grpc server host:port")
	n          = flag.Int("count", 1, "number of rpc calls")
	cacert     = flag.String("cacert", "CA.crt", "CACert for server")
	serverName = flag.String("servername", "timeserver.gke.net", "CACert for server")
	insecure   = flag.Bool("insecure", false, "connect without TLS")
	api        = flag.Int("api", 1, "GetTime(1), GetTimeSOut(2), GetTimeSIn(3), GetTimeSInSOut(4)")
	sin        = flag.Int("sin", 1, "number of messages in GetTimeSIn (input stream) rpc calls")
)

func main() {
	flag.Parse()
	var conn *grpc.ClientConn
	var err error
	if *insecure {
		conn, err = grpc.Dial(*host, grpc.WithInsecure())
	} else {
		var tlsCfg tls.Config
		rootCAs := x509.NewCertPool()
		pem, err := ioutil.ReadFile(*cacert)
		if err != nil {
			log.Fatalf("failed to load root CA certificates %v", err)
		}
		if !rootCAs.AppendCertsFromPEM(pem) {
			log.Fatalf("no root CA certs parsed from file ")
		}
		tlsCfg.RootCAs = rootCAs
		tlsCfg.ServerName = *serverName

		ce := credentials.NewTLS(&tlsCfg)
		conn, err = grpc.Dial(*host, grpc.WithTransportCredentials(ce))

	}
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := timep.NewTimeServerClient(conn)

	switch *api {
	case 1:
		fmt.Printf("Calling GetTime %d times\n", *n)
		for i := 0; i < *n; i++ {
			r, err := client.GetTime(context.Background(), &timep.TimeRequest{
				Name: "test",
			})
			if err != nil {
				log.Fatalf("%v", err)
			}
			log.Printf("Reply: %v", r)
		}
	case 2:
		fmt.Printf("Calling GetTimeSOut\n")
		stream, err := client.GetTimeSOut(context.Background(), &timep.TimeRequest{
			Name: "test",
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
		for {
			r, err := stream.Recv()
			if err == io.EOF {
				log.Println("Done")
			}
			if err != nil {
				log.Fatalf("%v", err)
			}
			log.Printf("StreamOut Reply: %v", r)
		}
	case 3:
		fmt.Printf("Calling GetTimeSIn\n")
		for i := 0; i < *n; i++ {
			stream, err := client.GetTimeSIn(context.Background())
			if err != nil {
				log.Fatalf("%v", err)
			}
			for j := 0; j < *sin; j++ {
				stream.Send(&timep.TimeRequest{
					Name: "test",
				})
			}
			r, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("%v", err)
			}
			log.Printf("[#%d]StreamIn Reply: %v", i, r)
		}
	case 4:
		fmt.Printf("Calling GetTimeSInSOut\n")
		for i := 0; i < *n; i++ {
			stream, err := client.GetTimeSInSOut(context.Background())
			if err != nil {
				log.Fatalf("%v", err)
			}
			// recv go routine

			go func(i int) {
				for {
					r, err := stream.Recv()
					if err == io.EOF {
						return
					}
					if err != nil {
						log.Fatalf("%v", err)
					}
					log.Printf("[#%d]StreamInOUT Reply: %v", i, r)
				}
			}(i)

			for j := 0; j < *sin; j++ {
				stream.Send(&timep.TimeRequest{
					Name: "test",
				})
			}
			time.Sleep(2 * time.Second)
			err = stream.CloseSend()
			if err != nil {
				log.Fatalf("%v", err)
			}
		}
	}
}
