package handlers

import (
	"context"
	"fmt"
	"github.com/dlc-01/GophKeeper/internal/general/pass"
	proto "github.com/dlc-01/GophKeeper/internal/general/proto/gen"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type PairClient struct {
	conn *grpc.ClientConn
}

func NewPairClient(conn *grpc.ClientConn) *PairClient {
	return &PairClient{
		conn: conn,
	}
}

func (c *PairClient) CreatePair(ctx context.Context, token string, secretKey string, pair models.Pair) (*models.Pair, error) {
	client := proto.NewPairClient(c.conn)

	cipher, err := pass.Encrypt(pass.HashData{Data: pair.PasswordHash, SecretKey: secretKey})
	if err != nil {
		return nil, fmt.Errorf("cannot decrypt")
	}

	req := &proto.CreatePairRequest{
		Token: token,
		Pair:  &proto.PairMsg{Login: pair.Username, PasswordHash: cipher.Data, NonceHex: cipher.NonceHex, Metadata: pair.Metadata},
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	resp, err := client.CreatePair(ctx, req)
	if err != nil {
		return nil, err
	}
	pair.ID = resp.Pair.GetId()
	return &pair, nil
}

func (c *PairClient) GetPair(ctx context.Context, token string, secretKey string) ([]models.Pair, error) {
	client := proto.NewPairClient(c.conn)
	req := &proto.GetPairRequest{
		Token: token,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	resp, err := client.GetPair(ctx, req)
	if err != nil {
		return nil, err
	}

	out := make([]models.Pair, len(resp.Pairs))
	for i, pair := range resp.GetPairs() {

		cipher, err := pass.Decrypt(pass.HashData{Data: pair.GetPasswordHash(), SecretKey: secretKey, NonceHex: pair.GetNonceHex()})
		if err != nil {
			return nil, fmt.Errorf("cannot decrypt")
		}

		out[i] = models.Pair{
			ID:           pair.GetId(),
			Username:     pair.GetLogin(),
			PasswordHash: cipher.Data,
			Metadata:     pair.GetMetadata(),
		}
	}

	return out, nil
}

func (c *PairClient) UpdatePair(ctx context.Context, pair models.Pair) (*models.Pair, error) {
	client := proto.NewPairClient(c.conn)
	req := &proto.UpdatePairRequest{
		Pair: &proto.PairMsg{Id: pair.ID, PasswordHash: pair.PasswordHash, Login: pair.Username, Metadata: pair.Metadata},
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	resp, err := client.UpdatePair(ctx, req)
	if err != nil {
		return nil, err
	}
	pair = models.Pair{ID: resp.Pair.GetId(), Username: req.Pair.GetLogin(), Metadata: req.Pair.Metadata, PasswordHash: req.Pair.PasswordHash}
	return &pair, nil
}

func (c *PairClient) DeletePair(ctx context.Context, token string) error {
	client := proto.NewPairClient(c.conn)
	req := &proto.DeletePairRequest{
		Token: token,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	_, err := client.DeletePair(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
