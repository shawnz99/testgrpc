package main

import (
	"context"
	"net"

	pb "github.com/shawnz99/testgrpc/api"

	"google.golang.org/grpc"
)

type server struct {
	ln   net.Listener
	id   string
	addr string
	// messages chan *api.message
}

func (s *server) Chat(ctx context.Context, in *pb.Msg) (*pb.Reply, error) {

}

func main() {
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
}
