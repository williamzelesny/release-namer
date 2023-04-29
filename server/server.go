package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
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

// server is used to implement releasenamerserver.
type releaseNamerServer struct {
	pb.UnimplementedReleaseNamerServer
}

func (s *releaseNamerServer) GetCandies(ctx context.Context, empty *empty.Empty) (*pb.CandyReply, error) {
	return &pb.CandyReply{Name: "Twizzlers"}, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReleaseNamerServer(s, &releaseNamerServer{})
	logger.Info("server listening at " + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
