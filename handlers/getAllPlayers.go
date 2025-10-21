package handlers

import (
	"context"
	"sort"

	pb "github.com/timrxd/exhibit-backend/proto/players"
)

func (s *server) GetAllPlayers(ctx context.Context, in *pb.EmptyMessage) (*pb.PlayerListResp, error) {
	// Get players from db
	playerList := []*pb.Player{}
	for _, p := range playerDB {
		playerList = append(playerList, p)
	}

	// Sort players by ranking points
	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].Points > playerList[j].Points
	})

	return &pb.PlayerListResp{Players: playerList}, nil
}
