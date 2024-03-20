package handlers

import (
	"context"
	proto "github.com/dlc-01/GophKeeper/internal/general/proto/gen"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
)

type AuthServer struct {
	proto.UnimplementedAuthServer
	auth port.IAuthService
}

func NewAuthServer(auth port.IAuthService) *AuthServer {
	return &AuthServer{
		auth: auth,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	var resp proto.RegisterUserResponse
	user := models.User{
		Username:     req.GetLogin(),
		PasswordHash: req.GetPasswordHash()}
	token, err := s.auth.Register(ctx, &user)
	if err != nil {
		return nil, err
	}

	resp.Token = token
	return &resp, nil
}

func (s *AuthServer) Login(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	var resp proto.LoginUserResponse
	user := models.User{
		Username:     req.GetLogin(),
		PasswordHash: req.GetPasswordHash()}
	token, err := s.auth.Login(ctx, &user)
	if err != nil {
		return nil, err
	}

	resp.Token = token
	return &resp, nil
}
