package login_proto

import "context"

type LoginServer struct {
	UnimplementedLoginServiceServer
}

func (l *LoginServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	s := &LoginResponse{
		Uuid:      "ssss",
		PublicKey: "aaaaa",
		Time:      "ccdcd",
	}
	return s, nil
}
