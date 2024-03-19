package handlers

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/proto"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

type BankClient struct {
	conn *grpc.ClientConn
}

func NewBankClient(conn *grpc.ClientConn) *BankClient {
	return &BankClient{
		conn: conn,
	}
}

func (c *BankClient) CreateBank(ctx context.Context, token string, card models.BankAccountString) (*models.BankAccountString, error) {
	client := proto.NewBanksClient(c.conn)
	req := &proto.CreateBankRequest{
		Token: token,
		Card:  &proto.CardMsg{CardHolder: card.CardHolder, Number: card.Number, ExpirationDate: card.ExpirationDate, Metadata: card.Metadata, SecurityCode: card.SecurityCode},
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	resp, err := client.CreateBank(ctx, req)
	if err != nil {
		return nil, err
	}
	card.ID = resp.Card.GetId()
	return &card, nil
}

func (c *BankClient) GetBank(ctx context.Context, token string) ([]models.BankAccount, error) {
	client := proto.NewBanksClient(c.conn)
	req := &proto.GetBankRequest{
		Token: token,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	resp, err := client.GetBank(ctx, req)
	if err != nil {
		return nil, err
	}

	out := make([]models.BankAccount, len(resp.Cards))
	for i, card := range resp.GetCards() {
		data := card.ExpirationDate[:10]
		date, err := time.Parse("2006-01-02", data)
		if err != nil {
			return nil, status.Error(codes.Aborted, "cannot parse ExpirationDate")
		}
		out[i] = models.BankAccount{
			ID:             card.GetId(),
			CardHolder:     card.GetCardHolder(),
			Number:         card.GetNumber(),
			ExpirationDate: date,
			Metadata:       card.GetMetadata(),
			SecurityCode:   card.GetSecurityCode(),
		}
	}

	return out, nil
}

func (c *BankClient) UpdateBank(ctx context.Context, card models.BankAccount) (*models.BankAccount, error) {
	client := proto.NewBanksClient(c.conn)
	req := &proto.UpdateBankRequest{
		Card: &proto.CardMsg{CardHolder: card.CardHolder, Number: card.Number, ExpirationDate: card.ExpirationDate.String(), Metadata: card.Metadata},
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	resp, err := client.UpdateBank(ctx, req)
	if err != nil {
		return nil, err
	}
	card.ID = resp.Card.GetId()
	return &card, nil
}

func (c *BankClient) DeleteBank(ctx context.Context, token string) error {
	client := proto.NewBanksClient(c.conn)
	req := &proto.DeleteBankRequest{
		Token: token,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	_, err := client.DeleteBank(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
