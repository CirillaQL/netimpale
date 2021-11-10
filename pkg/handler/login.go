package handler

import (
	"context"
	loginservice "netimpale/proto/login"
)

type LoginHandler struct{}

func (l *LoginHandler) Login(ctx context.Context, req *loginservice.LoginRequest) (*loginservice.LoginResponse, error) {
	loginInfo := req.GetLoginInfo()

	return &loginservice.LoginResponse{
		PublicKey: "",
		Time:      "",
		Result:    false,
	}, nil
}
