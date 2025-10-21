package handlers

import (
	"context"

	pb "github.com/timrxd/exhibit-backend/proto/players"
)

func (s *server) DeletePlayer(ctx context.Context, in *pb.DeletePlayerReq) (*pb.EmptyMessage, error) {
	return &pb.EmptyMessage{}, nil
}
