package handlers

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/pass"
	proto "github.com/dlc-01/GophKeeper/internal/general/proto/gen"

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

func (c *AuthClient) Register(ctx context.Context, login, password, key string) (string, error) {
	client := proto.NewAuthClient(c.conn)

	hashedPass := pass.HashH512Password(pass.HashData{Data: password, SecretKey: key})

	req := &proto.RegisterUserRequest{
		Login:        login,
		PasswordHash: hashedPass,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	defer cancel()

	resp, err := client.Register(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetToken(), nil
}

func (c *AuthClient) Login(ctx context.Context, login, password, key string) (string, error) {
	client := proto.NewAuthClient(c.conn)

	hashedPass := pass.HashH512Password(pass.HashData{Data: password, SecretKey: key})

	req := &proto.LoginUserRequest{
		Login:        login,
		PasswordHash: hashedPass,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	defer cancel()

	resp, err := client.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetToken(), nil
}
