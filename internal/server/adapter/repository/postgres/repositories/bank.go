package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/repository/postgres/query"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/projectError"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

type BankRepository struct {
	*sql.DB
	lgr *logger.Logger
}

func NewBankRepository(lgr *logger.Logger, client *sql.DB) (*BankRepository, error) {
	return &BankRepository{
		DB:  client,
		lgr: lgr,
	}, nil
}

func (b *BankRepository) CreateByUserId(ctx context.Context, bank *models.BankAccount, user *models.User) (*models.BankAccount, error) {
	var stored models.BankAccount
	err := b.QueryRowContext(ctx, query.CreateBankAcc, user.ID, bank.Number, bank.CardHolder, bank.ExpirationDate, bank.SecurityCode, bank.Metadata).
		Scan(&stored.ID,
			&stored.Number,
			&stored.CardHolder,
			&stored.ExpirationDate,
			&stored.SecurityCode,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating bankAcc: %w", err)
	}

	return &stored, nil
}

func (b *BankRepository) CreateByUsername(ctx context.Context, bank *models.BankAccount, user *models.User) (*models.BankAccount, error) {
	var stored models.BankAccount
	err := b.QueryRowContext(ctx, query.CreateBankAccByUsername, user.Username, bank.Number, bank.CardHolder, bank.ExpirationDate, bank.SecurityCode, bank.Metadata).
		Scan(&stored.ID,
			&stored.Number,
			&stored.CardHolder,
			&stored.ExpirationDate,
			&stored.SecurityCode,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating bankAcc: %w", err)
	}

	return &stored, nil
}

func (b *BankRepository) GetBankAccountsByUsername(ctx context.Context, user *models.User) (*[]models.BankAccount, error) {
	stored := make([]models.BankAccount, 0)
	row, err := b.DB.QueryContext(ctx, query.GetBankAccountsByUsername, user.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting bankAcc: %w", err)
	}
	for row.Next() {
		acc := models.BankAccount{}
		if err = row.Scan(&acc.ID, &acc.Number, &acc.CardHolder, &acc.ExpirationDate, &acc.SecurityCode, &acc.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}
			return nil, fmt.Errorf("error while scaning bankAcc: %w", err)
		}
		stored = append(stored, acc)
	}

	return &stored, nil
}

func (b *BankRepository) GetBankAccountByID(ctx context.Context, bank *models.BankAccount) (*models.BankAccount, error) {
	var stored models.BankAccount
	err := b.QueryRowContext(ctx, query.GetBankAccountByID, bank.ID).
		Scan(&stored.ID,
			&stored.Number,
			&stored.CardHolder,
			&stored.ExpirationDate,
			&stored.SecurityCode,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while geting bankAcc: %w", err)
	}
	return &stored, nil
}

func (b *BankRepository) GetBankAccountsByUserID(ctx context.Context, user *models.User) (*[]models.BankAccount, error) {
	stored := make([]models.BankAccount, 0)
	row, err := b.DB.QueryContext(ctx, query.GetBankAccountsByUserID, *user.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting bankAcc: %w", err)
	}
	for row.Next() {
		acc := models.BankAccount{}
		if err = row.Scan(&acc.ID, &acc.Number, &acc.CardHolder, &acc.ExpirationDate, &acc.SecurityCode, &acc.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}
			return nil, fmt.Errorf("error while scaning bankAcc: %w", err)
		}
		stored = append(stored, acc)
	}

	return &stored, nil
}

func (b *BankRepository) Update(ctx context.Context, bank *models.BankAccount) (*models.BankAccount, error) {
	err := b.QueryRowContext(ctx, query.UpdateBankAcc, bank.ID, bank.Number, bank.CardHolder, bank.ExpirationDate, bank.SecurityCode, bank.Metadata).Err()
	if err != nil {
		return nil, fmt.Errorf("error while updating bankAcc: %w", err)
	}

	return bank, nil
}

func (b *BankRepository) DeleteById(ctx context.Context, bank *models.BankAccount) error {
	err := b.QueryRowContext(ctx, query.DeleteBankAcc, bank.ID).Err()
	if err != nil {
		return fmt.Errorf("error while deleting bankAcc: %w", err)
	}

	return nil
}

func (b *BankRepository) DeleteByUsername(ctx context.Context, user *models.User) error {
	err := b.QueryRowContext(ctx, query.DeleteBankAccByUsername, user.Username).Err()
	if err != nil {
		return fmt.Errorf("error while deleting bankAcc: %w", err)
	}

	return nil
}

func (b *BankRepository) DeleteByUserId(ctx context.Context, user models.User) error {
	err := b.QueryRowContext(ctx, query.DeleteBankAccByUserID, user.ID).Err()
	if err != nil {
		return fmt.Errorf("error while deleting bankAcc: %w", err)
	}

	return nil
}
