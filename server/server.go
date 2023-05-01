package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slog"
	"net"
	"os"
	"time"
	"williamzelesny/release-namer/scrapers"

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

func (s *releaseNamerServer) GetCandies(ctx context.Context, empty *empty.Empty) (*pb.CandyResponse, error) {
	//candies := []*pb.Candy{
	//	&pb.Candy{Name: "Twizzlers"},
	//	&pb.Candy{Name: "Skittles"},
	//}
	candies, _ := scrapers.GetCandies()

	return &pb.CandyResponse{Candies: candies}, nil
}

func (s *releaseNamerServer) GetRandomCandy(ctx context.Context, empty *empty.Empty) (*pb.CandyResponse, error) {
	candies, _ := scrapers.GetCandies()
	rand.Seed(uint64(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(candies))
	return &pb.CandyResponse{Candies: []*pb.Candy{candies[randomIndex]}}, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Error("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReleaseNamerServer(s, &releaseNamerServer{})
	logger.Info("server listening at " + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve: %v", err)
	}
}
