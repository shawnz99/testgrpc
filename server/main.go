package main

import (
	"context"
	"net"

	pb "testgrpc"

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
	chat.RegisterChatServer(s, &server{})
}
