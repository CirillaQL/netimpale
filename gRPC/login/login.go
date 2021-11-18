package login_proto

import (
	"context"
	"netimpale/utils/log"
)

var LOG = log.LOG

type LoginServer struct {
	UnimplementedLoginServiceServer
}

func (l *LoginServer) Login(ctx context.Context,
	req *LoginRequest) (*LoginResponse, error) {
	LOG.Infof("User: %s Login", req.Username)
	LOG.Infof("Login Infomation: %+v", req)

	s := &LoginResponse{
		Uuid:      "ssss",
		PublicKey: "aaaaa",
		Time:      "ccdcd",
	}
	return s, nil
}
