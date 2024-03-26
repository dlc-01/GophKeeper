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

	tx, err := b.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.CreateBankAcc, user.ID, bank.Number, bank.CardHolder, bank.ExpirationDate, bank.SecurityCodeHash, bank.NonceHex, bank.Metadata).
		Scan(&stored.ID,
			&stored.Number,
			&stored.CardHolder,
			&stored.ExpirationDate,
			&stored.SecurityCodeHash,
			&stored.NonceHex,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating bankAcc: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}
	return &stored, nil
}

func (b *BankRepository) CreateByUsername(ctx context.Context, bank *models.BankAccount, user *models.User) (*models.BankAccount, error) {
	var stored models.BankAccount

	tx, err := b.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.CreateBankAccByUsername, user.Username, bank.Number, bank.CardHolder, bank.ExpirationDate, bank.SecurityCodeHash, bank.NonceHex, bank.Metadata).
		Scan(&stored.ID,
			&stored.Number,
			&stored.CardHolder,
			&stored.ExpirationDate,
			&stored.SecurityCodeHash,
			&stored.NonceHex,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating bankAcc: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (b *BankRepository) GetBankAccountsByUsername(ctx context.Context, user *models.User) (*[]models.BankAccount, error) {
	stored := make([]models.BankAccount, 0)

	tx, err := b.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	row, err := tx.QueryContext(ctx, query.GetBankAccountsByUsername, user.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting bankAcc: %w", err)
	}

	for row.Next() {
		acc := models.BankAccount{}
		if err = row.Scan(&acc.ID, &acc.Number, &acc.CardHolder, &acc.ExpirationDate, &acc.SecurityCodeHash, &acc.NonceHex, &acc.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}
			return nil, fmt.Errorf("error while scaning bankAcc: %w", err)
		}
		stored = append(stored, acc)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (b *BankRepository) GetBankAccountByID(ctx context.Context, bank *models.BankAccount) (*models.BankAccount, error) {
	var stored models.BankAccount

	tx, err := b.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.GetBankAccountByID, bank.ID).
		Scan(&stored.ID,
			&stored.Number,
			&stored.CardHolder,
			&stored.ExpirationDate,
			&stored.SecurityCodeHash,
			&stored.NonceHex,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while geting bankAcc: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (b *BankRepository) GetBankAccountsByUserID(ctx context.Context, user *models.User) (*[]models.BankAccount, error) {
	stored := make([]models.BankAccount, 0)

	tx, err := b.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	row, err := tx.QueryContext(ctx, query.GetBankAccountsByUserID, *user.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting bankAcc: %w", err)
	}

	for row.Next() {
		acc := models.BankAccount{}
		if err = row.Scan(&acc.ID, &acc.Number, &acc.CardHolder, &acc.ExpirationDate, &acc.SecurityCodeHash, &acc.NonceHex, &acc.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}
			return nil, fmt.Errorf("error while scaning bankAcc: %w", err)
		}
		stored = append(stored, acc)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (b *BankRepository) Update(ctx context.Context, bank *models.BankAccount) (*models.BankAccount, error) {
	tx, err := b.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.UpdateBankAcc, bank.ID, bank.Number, bank.CardHolder, bank.ExpirationDate, bank.SecurityCodeHash, bank.NonceHex, bank.Metadata).Err()
	if err != nil {
		return nil, fmt.Errorf("error while updating bankAcc: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return bank, nil
}

func (b *BankRepository) DeleteById(ctx context.Context, bank *models.BankAccount) error {
	tx, err := b.Begin()
	if err != nil {
		return fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeleteBankAcc, bank.ID).Err()
	if err != nil {
		return fmt.Errorf("error while deleting bankAcc: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return nil
}

func (b *BankRepository) DeleteByUsername(ctx context.Context, user *models.User) error {
	tx, err := b.Begin()
	if err != nil {
		return fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeleteBankAccByUsername, user.Username).Err()
	if err != nil {
		return fmt.Errorf("error while deleting bankAcc: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return nil
}

func (b *BankRepository) DeleteByUserId(ctx context.Context, user models.User) error {
	tx, err := b.Begin()
	if err != nil {
		return fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeleteBankAccByUserID, user.ID).Err()
	if err != nil {
		return fmt.Errorf("error while deleting bankAcc: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return nil
}
