package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb_login "netimpale/gRPC/login"
	"netimpale/utils/log"
)

var LOG = log.LOG

func main() {
	fmt.Println("Hello, NetImpale!")
	//首先监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		LOG.Infof("listener err: %v", err)
	}
	LOG.Infof(" net.Listing...")
	// 实例化gRPC
	grpcServer := grpc.NewServer()
	pb_login.RegisterLoginServiceServer(grpcServer, &pb_login.LoginServer{})

	//启动服务器
	err = grpcServer.Serve(listener)
	if err != nil {
		LOG.Errorf("grpc server err: %v", err)
	}
}
