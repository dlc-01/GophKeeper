package handlers

import (
	"context"
	proto "github.com/dlc-01/GophKeeper/internal/general/proto/gen"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PairServer struct {
	proto.UnimplementedPairServer
	pair port.IPairService
}

func NewPairServer(pair port.IPairService) *PairServer {
	return &PairServer{
		pair: pair,
	}
}

func (p *PairServer) CreatePair(ctx context.Context, req *proto.CreatePairRequest) (*proto.CreatePairResponse, error) {
	var resp proto.CreatePairResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	pair := models.Pair{
		Metadata:     req.Pair.GetMetadata(),
		Username:     req.Pair.GetLogin(),
		PasswordHash: req.Pair.GetPasswordHash(),
		NonceHex:     req.Pair.GetNonceHex(),
	}

	pairNew, err := p.pair.CreateByUserId(ctx, pair, user)
	if err != nil {
		return nil, err
	}

	resp.Pair = req.Pair
	resp.Pair.Id = pairNew.ID

	return &resp, nil
}

func (p *PairServer) GetPair(ctx context.Context, req *proto.GetPairRequest) (*proto.GetPairResponse, error) {
	var resp proto.GetPairResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	stored, err := p.pair.GetPairsByUserID(ctx, user)
	if err != nil {
		return nil, err
	}

	for _, pair := range *stored {
		resp.Pairs = append(resp.Pairs, &proto.PairMsg{
			Id:           pair.ID,
			Login:        pair.Username,
			PasswordHash: pair.PasswordHash,
			Metadata:     pair.Metadata,
			NonceHex:     pair.NonceHex,
		})
	}

	return &resp, nil
}

func (p *PairServer) UpdatePair(ctx context.Context, req *proto.UpdatePairRequest) (*proto.UpdatePairResponse, error) {
	var resp proto.UpdatePairResponse

	_, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	pair := models.Pair{
		ID:           req.Pair.GetId(),
		Username:     req.Pair.GetLogin(),
		PasswordHash: req.Pair.GetPasswordHash(),
		Metadata:     req.Pair.GetMetadata(),
		NonceHex:     req.Pair.GetNonceHex(),
	}

	_, err := p.pair.Update(ctx, pair)
	if err != nil {
		return nil, err
	}

	resp.Pair = req.Pair

	return &resp, nil
}

func (p *PairServer) DeletePair(ctx context.Context, req *proto.DeletePairRequest) (*proto.DeletePairResponse, error) {
	var resp proto.DeletePairResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	err := p.pair.DeleteByUserId(ctx, user)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
