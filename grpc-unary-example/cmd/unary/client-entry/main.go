package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "grpc/grpc-unary-example/core/unary"
)

var (
	name string
)

func init() {
	flag.StringVar(&name, "name", "karl", "input name")
	flag.Parse()
}

func main() {
	conn, e := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if e != nil {
		logrus.Error(e)
		return
	}
	defer conn.Close()

	c := pb.NewRoomClient(conn)

	member := pb.Guest{
		Name: name,
	}

	cb, e := c.Entry(context.Background(), &member)
	if e != nil {
		logrus.Error(e)
		return
	}

	logrus.Info(cb.Message)
}
