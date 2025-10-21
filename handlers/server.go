package handlers

import (
	pb "github.com/timrxd/exhibit-backend/proto/players"
)

var playerDB map[string]*pb.Player

type server struct {
	pb.UnimplementedPlayerServiceServer
}

func NewServer() *server {
	return &server{}
}

func LoadTestData() {
	playerDB = map[string]*pb.Player{
		"Federer":  &pb.Player{Name: "Federer", Country: "SUI", Points: 10550, Age: 28},
		"Nadal":    &pb.Player{Name: "Nadal", Country: "ESP", Points: 9205, Age: 23},
		"Djokovic": &pb.Player{Name: "Djokovic", Country: "SRB", Points: 8310, Age: 22},
		"Murray":   &pb.Player{Name: "Murray", Country: "GBR", Points: 7030, Age: 22},
		"delPotro": &pb.Player{Name: "delPotro", Country: "ARG", Points: 6785, Age: 21},
	}
}
