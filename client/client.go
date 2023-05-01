package main

import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
	"williamzelesny/release-namer/data"
	pb "williamzelesny/release-namer/releasenamer"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

// getCandies gets the candy.
func getCandies(logger *slog.Logger, client pb.ReleaseNamerClient, empty *empty.Empty) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.GetCandies(ctx, empty)
	if err != nil {
		logger.Error("client.GetFeature failed: %v", err)
	}
	log.Println(feature)
}

// getCandies gets the candy.
func getRandomCandy(logger *slog.Logger, client pb.ReleaseNamerClient, empty *empty.Empty) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.GetRandomCandy(ctx, empty)
	if err != nil {
		logger.Error("client.GetFeature failed: %v", err)
	}
	log.Println(feature)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))

	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			logger.Error("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		logger.Error("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReleaseNamerClient(conn)

	// run getCandies
	getCandies(logger, client, &empty.Empty{})

	getRandomCandy(logger, client, &empty.Empty{})
}
