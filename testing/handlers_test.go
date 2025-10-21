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

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewPlayerServiceClient(conn)
	resp, err := client.GetPlayer(ctx, &pb.GetPlayerReq{Name: "Federer"})
	if err != nil {
		t.Fatalf("GetPlayer failed: %v", err)
	}

	if resp.Player.Name != "Federer" {
		t.Errorf("Error: Body mismatch\tExpected %s\tGot %s", "Federer", resp.Player.Name)
	}

}
