package handlers

import (
	"context"

	pb "github.com/timrxd/exhibit-backend/proto/players"
)

func (s *server) GetPlayer(ctx context.Context, in *pb.PlayerReq) (*pb.PlayerResp, error) {
	return &pb.PlayerResp{Player: &pb.Player{Name: in.Name}}, nil
}
