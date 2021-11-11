package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	login "netimpale/proto/login"
	"netimpale/utils/log"
	"time"
)

var LOG = log.LOG

func main() {
	fmt.Println("Hello, NetImpale!")
	conn, err := grpc.Dial("119.28.77.233:23456", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		LOG.Errorf("gRPC Start Failed: %+v", err)
	}
	defer conn.Close()

	c := login.NewLoginServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	a := login.LoginInfo{
		Uuid:        "aaaaa",
		Username:    "pp",
		Password:    "aaa",
		IfEncrypted: false,
		ClientAddr:  "ssss",
		ClientPort:  0,
		ServerAddr:  "aaaaa",
		ServerPort:  0,
	}
	r, err := c.Login(ctx, &login.LoginRequest{
		LoginInfo: &a,
		Time:      time.Now().String(),
	})
	fmt.Println(r)
}
