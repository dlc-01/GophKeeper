package handlers

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type UserClient struct {
	conn *grpc.ClientConn
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		conn: conn,
	}
}

func (c *UserClient) DeleteUser(ctx context.Context, token string) error {
	client := proto.NewUserClient(c.conn)
	req := &proto.DeleteUserRequest{
		Token: token,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	_, err := client.DeleteUser(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserClient) UpdateUser(ctx context.Context, token, login, password string) (string, error) {
	client := proto.NewUserClient(c.conn)
	req := &proto.UpdateUserRequest{
		Token: token,
		User:  &proto.UserRequest{Login: login, PasswordHash: password},
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	resp, err := client.UpdateUser(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}
