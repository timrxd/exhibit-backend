package handlers

import (
	"context"

	pb "github.com/timrxd/exhibit-backend/proto/players"
)

func (s *server) GetPlayer(ctx context.Context, in *pb.GetPlayerReq) (*pb.PlayerResp, error) {
	return &pb.PlayerResp{Player: playerDB[in.Name]}, nil
}
