package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "grpc/grpc-unary-example/core/unary"
	"net"
)

// room 클래스 구현
type room struct {
	list *pb.Guests
	pb.UnimplementedRoomServer
}

// Entry (방문) 기능 구현
func (t *room) Entry(ctx context.Context, guest *pb.Guest) (*pb.Message, error) {
	var welcomeMsg pb.Message

	if _, ok := t.list.Guest[guest.Name]; !ok {
		t.list.Guest[guest.Name] = guest

		welcomeMsg = pb.Message{Message: "Welcome! " + guest.Name}
	} else {
		welcomeMsg = pb.Message{Message: "Welcome Back! " + guest.Name}
	}
	return &welcomeMsg, nil
}

// EntryList (방문자 목록 조회) 기능 구현
func (t *room) EntryList(ctx context.Context, void *pb.Void) (*pb.Guests, error) {
	return t.list, nil
}

func main() {
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		logrus.Error(e)
		return
	}

	srv := grpc.NewServer()

	wsrv := &room{
		list: &pb.Guests{
			Guest: map[string]*pb.Guest{},
		},
	}
	pb.RegisterRoomServer(srv, wsrv)

	logrus.Info(fmt.Sprintf("gRPC Server (%s)", l.Addr().String()))

	if e := srv.Serve(l); e != nil {
		logrus.Error(e)
	}
}
