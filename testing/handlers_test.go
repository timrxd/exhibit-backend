package testing

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/timrxd/exhibit-backend/handlers"
	pb "github.com/timrxd/exhibit-backend/proto/players"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterPlayerServiceServer(s, handlers.NewServer())
	handlers.LoadTestData()
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAPI(t *testing.T) {

	// Connect to mock grpc
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewPlayerServiceClient(conn)

	// Test GetPlayer
	resp, err := client.GetPlayer(ctx, &pb.GetPlayerReq{Name: "Federer"})
	if err != nil {
		t.Fatalf("GetPlayer failed: %v", err)
	}

	if resp.Player.Name != "Federer" {
		t.Errorf("Error: Body mismatch\tExpected %s\tGot %s", "Federer", resp.Player.Name)
	}

	// Test GetAllPlayers
	allResp, err := client.GetAllPlayers(ctx, &pb.EmptyMessage{})
	if err != nil {
		t.Fatalf("GetPlayer failed: %v", err)
	}

	if allResp.String() != `players:{name:"Federer"  country:"SUI"  age:28  points:10550}  players:{name:"Nadal"  country:"ESP"  age:23  points:9205}  players:{name:"Djokovic"  country:"SRB"  age:22  points:8310}  players:{name:"Murray"  country:"GBR"  age:22  points:7030}  players:{name:"delPotro"  country:"ARG"  age:21  points:6785}` {
		t.Errorf("Error: Body mismatch\tExpected %s\tGot %s", "Federer", allResp.String())
	}

}
