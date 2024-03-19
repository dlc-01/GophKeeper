package handlers

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/proto"
	"google.golang.org/grpc"
	"time"
)

type AuthClient struct {
	conn *grpc.ClientConn
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		conn: conn,
	}
}

// TODO: захешируй пароль
func (c *AuthClient) Register(ctx context.Context, login, password string) (string, error) {
	client := proto.NewAuthClient(c.conn)
	req := &proto.RegisterUserRequest{
		Login:        login,
		PasswordHash: password,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	defer cancel()

	resp, err := client.Register(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetToken(), nil
}

func (c *AuthClient) Login(ctx context.Context, login, password string) (string, error) {
	client := proto.NewAuthClient(c.conn)
	req := &proto.LoginUserRequest{
		Login:        login,
		PasswordHash: password,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	resp, err := client.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetToken(), nil
}
