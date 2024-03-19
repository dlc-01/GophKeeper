package handlers

import (
	"context"
	proto2 "github.com/dlc-01/GophKeeper/internal/general/proto"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type BankServer struct {
	proto2.UnimplementedBanksServer
	bank port.IBankService
}

func NewBankServer(bank port.IBankService) *BankServer {
	return &BankServer{
		bank: bank,
	}
}

func (b *BankServer) CreateBank(ctx context.Context, req *proto2.CreateBankRequest) (*proto2.CreateBankResponse, error) {
	var resp proto2.CreateBankResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	date, err := time.Parse("2006-01-02", req.Card.ExpirationDate)
	if err != nil {
		return nil, status.Error(codes.Aborted, "cannot parse ExpirationDate")
	}

	card := models.BankAccount{
		Metadata:       req.Card.Metadata,
		Number:         req.Card.Number,
		CardHolder:     req.Card.CardHolder,
		ExpirationDate: date,
	}

	bankNew, err := b.bank.CreateByUserId(ctx, card, user)
	if err != nil {
		return nil, err
	}

	resp.Card = req.Card
	resp.Card.Id = bankNew.ID

	return &resp, nil
}
func (b *BankServer) GetBank(ctx context.Context, req *proto2.GetBankRequest) (*proto2.GetBankResponse, error) {
	var resp proto2.GetBankResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	stored, err := b.bank.GetPairsByUserID(ctx, user)
	if err != nil {
		return nil, err
	}

	for _, bank := range *stored {
		resp.Cards = append(resp.Cards, &proto2.CardMsg{
			Id:             bank.ID,
			Number:         bank.Number,
			CardHolder:     bank.CardHolder,
			ExpirationDate: bank.ExpirationDate.String(),
			Metadata:       bank.Metadata,
		})
	}
	return &resp, nil
}
func (b *BankServer) UpdateBank(ctx context.Context, req *proto2.UpdateBankRequest) (*proto2.UpdateBankResponse, error) {
	var resp proto2.UpdateBankResponse

	_, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	date, err := time.Parse("2006-01-02", req.Card.ExpirationDate)
	if err != nil {
		return nil, status.Error(codes.Aborted, "cannot parse ExpirationDate")
	}

	card := models.BankAccount{
		Metadata:       req.Card.Metadata,
		Number:         req.Card.Number,
		CardHolder:     req.Card.CardHolder,
		ExpirationDate: date,
	}

	_, err = b.bank.Update(ctx, card)
	if err != nil {
		return nil, err
	}

	resp.Card = req.Card

	return &resp, nil
}
func (b *BankServer) DeleteBank(ctx context.Context, req *proto2.DeleteBankRequest) (*proto2.DeleteBankResponse, error) {
	var resp proto2.DeleteBankResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	err := b.bank.DeleteByUserId(ctx, user)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
