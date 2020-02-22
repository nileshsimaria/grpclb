package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"

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
