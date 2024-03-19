package handlers

import (
	"context"
	proto2 "github.com/dlc-01/GophKeeper/internal/general/proto"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PairServer struct {
	proto2.UnimplementedPairServer
	pair port.IPairService
}

func NewPairServer(pair port.IPairService) *PairServer {
	return &PairServer{
		pair: pair,
	}
}

func (p *PairServer) CreatePair(ctx context.Context, req *proto2.CreatePairRequest) (*proto2.CreatePairResponse, error) {
	var resp proto2.CreatePairResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	pair := models.Pair{
		Metadata:     req.Pair.Metadata,
		Username:     req.Pair.Login,
		PasswordHash: req.Pair.PasswordHash,
	}

	pairNew, err := p.pair.CreateByUserId(ctx, pair, user)
	if err != nil {
		return nil, err
	}

	resp.Pair = req.Pair
	resp.Pair.Id = pairNew.ID

	return &resp, nil
}

func (p *PairServer) GetPair(ctx context.Context, req *proto2.GetPairRequest) (*proto2.GetPairResponse, error) {
	var resp proto2.GetPairResponse

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
		resp.Pairs = append(resp.Pairs, &proto2.PairMsg{
			Id:           pair.ID,
			Login:        pair.Username,
			PasswordHash: pair.PasswordHash,
			Metadata:     pair.Metadata,
		})
	}

	return &resp, nil
}

func (p *PairServer) UpdatePair(ctx context.Context, req *proto2.UpdatePairRequest) (*proto2.UpdatePairResponse, error) {
	var resp proto2.UpdatePairResponse

	_, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	pair := models.Pair{
		ID:           req.Pair.Id,
		Username:     req.Pair.Login,
		PasswordHash: req.Pair.PasswordHash,
		Metadata:     req.Pair.Metadata,
	}

	_, err := p.pair.Update(ctx, pair)
	if err != nil {
		return nil, err
	}

	resp.Pair = req.Pair

	return &resp, nil
}

func (p *PairServer) DeletePair(ctx context.Context, req *proto2.DeletePairRequest) (*proto2.DeletePairResponse, error) {
	var resp proto2.DeletePairResponse

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
