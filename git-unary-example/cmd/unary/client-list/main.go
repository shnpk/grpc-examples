package main

import (
	"context"
	"flag"
	"fmt"
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

	list, e := c.EntryList(context.Background(), &pb.Void{})
	if e != nil {
		logrus.Error(e)
		return
	}
	for k, v := range list.Guest {
		logrus.Info(fmt.Sprintf("key : %s, value : %s", k, v.Name))
	}
}
