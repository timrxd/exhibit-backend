package handlers

import (
	"context"

	pb "github.com/timrxd/exhibit-backend/proto/players"
)

func (s *server) UpdatePlayer(ctx context.Context, in *pb.UpdatePlayerReq) (*pb.PlayerResp, error) {
	return &pb.PlayerResp{Player: playerDB[in.Player.Name]}, nil
}
