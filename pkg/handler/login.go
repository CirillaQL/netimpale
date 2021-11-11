package handler

import (
	"context"
	"fmt"
	"netimpale/proto/login"
	"netimpale/utils/log"
	"time"
)

var LOG = log.LOG

type LoginHandler struct {
	login.UnimplementedLoginServiceServer
}

func (l *LoginHandler) Login(ctx context.Context,
	req *login.LoginRequest) (*login.LoginResponse, error) {
	LOG.Infof("Client login %+v", req.LoginInfo)
	loginInfo := req.GetLoginInfo()
	fmt.Println(loginInfo)
	return &login.LoginResponse{
		PublicKey: "aaa",
		Time:      time.Now().String(),
		Result:    true,
	}, nil
}
