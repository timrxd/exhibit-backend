package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	playerspb "github.com/timrxd/exhibit-backend/proto/players"
)

type server struct {
	playerspb.UnimplementedPlayerServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetPlayer(ctx context.Context, in *playerspb.PlayerReq) (*playerspb.PlayerResp, error) {
	return &playerspb.PlayerResp{Name: in.Name}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	playerspb.RegisterPlayerServiceServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
