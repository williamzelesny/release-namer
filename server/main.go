package main

import (
	"flag"
	"fmt"
	"golang.org/x/exp/slog"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "williamzelesny/release-namer/releasenamer"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedReleaseNamerServer
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReleaseNamerServer(s, &server{})
	logger.Info("server listening at " + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
