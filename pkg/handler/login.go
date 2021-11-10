package handler

import (
	"context"
	"fmt"
	loginservice "netimpale/proto/login"
	"netimpale/utils/log"
)

type LoginHandler struct{}

func (l *LoginHandler) Login(ctx context.Context,
	req *loginservice.LoginRequest) (*loginservice.LoginResponse, error) {
	log.LOG.Debug("")
	loginInfo := req.GetLoginInfo()
	fmt.Println(loginInfo)
	return &loginservice.LoginResponse{
		PublicKey: "",
		Time:      "",
		Result:    false,
	}, nil
}
