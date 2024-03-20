package handlers

import (
	"context"
	proto "github.com/dlc-01/GophKeeper/internal/general/proto/gen"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type BankServer struct {
	proto.UnimplementedBanksServer
	bank port.IBankService
}

func NewBankServer(bank port.IBankService) *BankServer {
	return &BankServer{
		bank: bank,
	}
}

func (b *BankServer) CreateBank(ctx context.Context, req *proto.CreateBankRequest) (*proto.CreateBankResponse, error) {
	var resp proto.CreateBankResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	date, err := time.Parse("2006-01-02", req.Card.GetExpirationDate())
	if err != nil {
		return nil, status.Error(codes.Aborted, "cannot parse ExpirationDate")
	}

	card := models.BankAccount{
		Metadata:         req.Card.GetMetadata(),
		Number:           req.Card.GetNumber(),
		CardHolder:       req.Card.GetCardHolder(),
		SecurityCodeHash: req.Card.GetSecurityCodeHash(),
		NonceHex:         req.Card.GetNonceHex(),
		ExpirationDate:   date,
	}

	bankNew, err := b.bank.CreateByUserId(ctx, card, user)
	if err != nil {
		return nil, err
	}

	resp.Card = req.Card
	resp.Card.Id = bankNew.ID

	return &resp, nil
}
func (b *BankServer) GetBank(ctx context.Context, req *proto.GetBankRequest) (*proto.GetBankResponse, error) {
	var resp proto.GetBankResponse

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

		resp.Cards = append(resp.Cards, &proto.CardMsg{
			Id:               bank.ID,
			Number:           bank.Number,
			CardHolder:       bank.CardHolder,
			ExpirationDate:   bank.ExpirationDate.String(),
			Metadata:         bank.Metadata,
			SecurityCodeHash: bank.SecurityCodeHash,
			NonceHex:         bank.NonceHex,
		})
	}
	return &resp, nil
}
func (b *BankServer) UpdateBank(ctx context.Context, req *proto.UpdateBankRequest) (*proto.UpdateBankResponse, error) {
	var resp proto.UpdateBankResponse

	_, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	date, err := time.Parse("2006-01-02", req.Card.GetExpirationDate())
	if err != nil {
		return nil, status.Error(codes.Aborted, "cannot parse ExpirationDate")
	}

	card := models.BankAccount{
		Metadata:         req.Card.GetMetadata(),
		Number:           req.Card.GetNumber(),
		CardHolder:       req.Card.GetCardHolder(),
		ExpirationDate:   date,
		SecurityCodeHash: req.Card.GetSecurityCodeHash(),
		NonceHex:         req.Card.GetNonceHex(),
	}

	_, err = b.bank.Update(ctx, card)
	if err != nil {
		return nil, err
	}

	resp.Card = req.Card

	return &resp, nil
}
func (b *BankServer) DeleteBank(ctx context.Context, req *proto.DeleteBankRequest) (*proto.DeleteBankResponse, error) {
	var resp proto.DeleteBankResponse

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
