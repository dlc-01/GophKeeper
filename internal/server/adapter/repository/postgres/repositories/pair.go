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

type PairRepository struct {
	*sql.DB
	lgr *logger.Logger
}

func NewPairRepository(lgr *logger.Logger, client *sql.DB) (*PairRepository, error) {
	return &PairRepository{
		DB:  client,
		lgr: lgr,
	}, nil
}

func (p *PairRepository) CreateByUserId(ctx context.Context, pair *models.Pair, user *models.User) (*models.Pair, error) {
	var stored models.Pair

	tx, err := p.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.CreatePair, user.ID, pair.Username, pair.PasswordHash, pair.NonceHex, pair.Metadata).
		Scan(&stored.ID,
			&stored.Username,
			&stored.PasswordHash,
			&stored.NonceHex,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating pair: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (p *PairRepository) CreateByUsername(ctx context.Context, pair *models.Pair, user *models.User) (*models.Pair, error) {
	var stored models.Pair

	tx, err := p.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.CreatePairByUsername, user.Username, pair.Username, pair.PasswordHash, pair.NonceHex, pair.Metadata).
		Scan(&stored.ID,
			&stored.Username,
			&stored.PasswordHash,
			&stored.NonceHex,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating pair: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (p *PairRepository) GetPairsByUsername(ctx context.Context, user *models.User) (*[]models.Pair, error) {
	stored := make([]models.Pair, 0)

	tx, err := p.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	row, err := tx.QueryContext(ctx, query.GetPairsByUsername, user.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting pair: %w", err)
	}

	for row.Next() {
		pair := models.Pair{}
		if err = row.Scan(&pair.ID, &pair.Username, &pair.PasswordHash, &pair.NonceHex, &pair.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}
			return nil, fmt.Errorf("error while scaning pair: %w", err)
		}
		stored = append(stored, pair)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (p *PairRepository) GetPairByID(ctx context.Context, pair *models.Pair) (*models.Pair, error) {
	var stored models.Pair

	tx, err := p.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.GetPairsByID, pair.ID).
		Scan(&stored.ID,
			&stored.Username,
			&stored.PasswordHash,
			&stored.NonceHex,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while geting pair: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (p *PairRepository) GetPairsByUserID(ctx context.Context, user *models.User) (*[]models.Pair, error) {
	stored := make([]models.Pair, 0)

	tx, err := p.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	row, err := tx.QueryContext(ctx, query.GetPairsByUserID, *user.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting pair: %w", err)
	}

	for row.Next() {
		pair := models.Pair{}
		if err = row.Scan(&pair.ID, &pair.Username, &pair.PasswordHash, &pair.NonceHex, &pair.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}
			return nil, fmt.Errorf("error while scaning pair: %w", err)
		}
		stored = append(stored, pair)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (p *PairRepository) Update(ctx context.Context, pair *models.Pair) (*models.Pair, error) {
	tx, err := p.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.UpdatePair, pair.Username, pair.PasswordHash, pair.NonceHex, pair.Metadata, pair.ID).Err()
	if err != nil {
		return nil, fmt.Errorf("error while updating pair: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return pair, nil
}

func (p *PairRepository) DeleteById(ctx context.Context, pair *models.Pair) (*models.Pair, error) {
	tx, err := p.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeletePair, pair.ID).Err()
	if err != nil {
		return nil, fmt.Errorf("error while deleting pair: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return pair, nil
}

func (p *PairRepository) DeleteByUsername(ctx context.Context, user *models.User) error {
	tx, err := p.Begin()
	if err != nil {
		return fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeletePairByUsername, user.Username).Err()
	if err != nil {
		return fmt.Errorf("error while deleting pairs: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return nil
}

func (p *PairRepository) DeleteByUserId(ctx context.Context, user *models.User) error {
	tx, err := p.Begin()
	if err != nil {
		return fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeletePairByUserID, user.Username).Err()
	if err != nil {
		return fmt.Errorf("error while deleting pairs: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return nil
}
