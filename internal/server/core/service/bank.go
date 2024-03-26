package service

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/projectError"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
)

type BankService struct {
	repo port.IBankRepository
	lgr  *logger.Logger
}

func NewBankService(repo port.IBankRepository, lgr *logger.Logger) *BankService {
	return &BankService{
		repo: repo,
		lgr:  lgr,
	}
}

func (b BankService) CreateByUserId(ctx context.Context, bank models.BankAccount, user models.User) (*models.BankAccount, error) {
	newBank, err := b.repo.CreateByUserId(ctx, &bank, &user)
	if err != nil {
		if err == projectError.ErrConflictingData {
			return nil, projectError.ErrConflictingData
		}
		return nil, err
	}

	return newBank, nil
}

func (b BankService) CreateByUsername(ctx context.Context, bank models.BankAccount, user models.User) (*models.BankAccount, error) {
	_, err := b.repo.CreateByUsername(ctx, &bank, &user)
	if err != nil {
		if err == projectError.ErrConflictingData {
			return nil, projectError.ErrConflictingData
		}
		return nil, err
	}

	return &bank, nil
}

func (b BankService) GetBankAccountsByUsername(ctx context.Context, user models.User) (*[]models.BankAccount, error) {
	pairs, err := b.repo.GetBankAccountsByUsername(ctx, &user)
	if err != nil {
		if err == projectError.ErrDataNotFound {
			return nil, projectError.ErrDataNotFound
		}
		return nil, err
	}

	return pairs, nil
}

func (b BankService) GetPairsByUserID(ctx context.Context, user models.User) (*[]models.BankAccount, error) {
	pairs, err := b.repo.GetBankAccountsByUserID(ctx, &user)
	if err != nil {
		if err == projectError.ErrDataNotFound {
			return nil, projectError.ErrDataNotFound
		}
		return nil, err
	}

	return pairs, nil
}

func (b BankService) Update(ctx context.Context, bank models.BankAccount) (*models.BankAccount, error) {
	_, err := b.repo.GetBankAccountByID(ctx, &bank)
	if err != nil {
		return nil, err
	}

	_, err = b.repo.Update(ctx, &bank)
	if err != nil {
		return nil, err
	}

	return &bank, nil
}

func (b BankService) DeleteById(ctx context.Context, bank models.BankAccount) error {
	_, err := b.repo.GetBankAccountByID(ctx, &bank)
	if err != nil {
		return err
	}

	err = b.repo.DeleteById(ctx, &bank)
	if err != nil {
		return err
	}

	return nil
}

func (b BankService) DeleteByUsername(ctx context.Context, user models.User) error {
	_, err := b.repo.GetBankAccountsByUsername(ctx, &user)
	if err != nil {
		return err
	}

	err = b.repo.DeleteByUsername(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (b *BankService) DeleteByUserId(ctx context.Context, user models.User) error {
	_, err := b.repo.GetBankAccountsByUserID(ctx, &user)
	if err != nil {
		return err
	}

	err = b.repo.DeleteByUsername(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
