package service

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/logger"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/projectError"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
)

type PairService struct {
	log  logger.Logger
	repo port.IPairRepository
}

func NewPairService(repo port.IPairRepository, log logger.Logger) *PairService {
	return &PairService{
		log,
		repo,
	}
}

func (p *PairService) CreateByUserId(ctx context.Context, pair models.Pair, user models.User) (*models.Pair, error) {
	newPair, err := p.repo.CreateByUserId(ctx, &pair, &user)
	if err != nil {
		if err == projectError.ErrConflictingData {
			return nil, projectError.ErrConflictingData
		}
		return nil, err
	}

	return newPair, nil
}

func (p *PairService) CreateByUsername(ctx context.Context, pair models.Pair, user models.User) (*models.Pair, error) {
	_, err := p.repo.CreateByUsername(ctx, &pair, &user)
	if err != nil {
		if err == projectError.ErrConflictingData {
			return nil, projectError.ErrConflictingData
		}
		return nil, err
	}

	return &pair, nil
}

func (p *PairService) GetPairsByUsername(ctx context.Context, user models.User) (*[]models.Pair, error) {
	pairs, err := p.repo.GetPairsByUsername(ctx, &user)
	if err != nil {
		if err == projectError.ErrDataNotFound {
			return nil, projectError.ErrDataNotFound
		}
		return nil, err
	}

	return pairs, nil
}

func (p *PairService) GetPairsByUserID(ctx context.Context, user models.User) (*[]models.Pair, error) {
	pairs, err := p.repo.GetPairsByUserID(ctx, &user)
	if err != nil {
		if err == projectError.ErrDataNotFound {
			return nil, projectError.ErrDataNotFound
		}
		return nil, err
	}

	return pairs, nil
}

func (p *PairService) Update(ctx context.Context, pair models.Pair) (*models.Pair, error) {
	_, err := p.repo.GetPairByID(ctx, &pair)
	if err != nil {
		return nil, err
	}

	_, err = p.repo.Update(ctx, &pair)
	if err != nil {
		return nil, err
	}

	return &pair, nil
}

func (p *PairService) DeleteById(ctx context.Context, pair models.Pair) (*models.Pair, error) {
	_, err := p.repo.GetPairByID(ctx, &pair)
	if err != nil {
		return nil, err
	}

	_, err = p.repo.DeleteById(ctx, &pair)
	if err != nil {
		return nil, err
	}

	return &pair, nil
}

func (p PairService) DeleteByUserId(ctx context.Context, user models.User) error {
	_, err := p.repo.GetPairsByUserID(ctx, &user)
	if err != nil {
		return err
	}

	err = p.repo.DeleteByUserId(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (p *PairService) DeleteByUsername(ctx context.Context, user models.User) error {
	_, err := p.repo.GetPairsByUsername(ctx, &user)
	if err != nil {
		return err
	}

	err = p.repo.DeleteByUsername(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
