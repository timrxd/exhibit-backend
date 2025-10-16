package handlers

import (
	pb "github.com/timrxd/exhibit-backend/proto/players"
)

type server struct {
	pb.UnimplementedPlayerServiceServer
}

func NewServer() *server {
	return &server{}
}
