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

type UserRepository struct {
	*sql.DB
	lgr *logger.Logger
}

func NewUserRepository(lgr *logger.Logger, client *sql.DB) (*UserRepository, error) {
	return &UserRepository{
		DB:  client,
		lgr: lgr,
	}, nil
}

func (u *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	var stored models.User
	err := u.QueryRowContext(ctx, query.CreateUser, user.Username, user.PasswordHash).
		Scan(&stored.ID,
			&stored.Username,
			&stored.PasswordHash)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}

		return nil, fmt.Errorf("error while creating user: %w", err)
	}

	return &stored, nil
}

func (u *UserRepository) GetByUserID(ctx context.Context, user *models.User) (*models.User, error) {
	var stored models.User
	err := u.QueryRowContext(ctx, query.GetUserByID, user.ID).
		Scan(&stored.ID,
			&stored.Username,
			&stored.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting user: %w", err)
	}

	return &stored, nil
}

func (u *UserRepository) GetByUsername(ctx context.Context, user *models.User) (*models.User, error) {
	err := u.QueryRowContext(ctx, query.GetUserByUsername, user.Username).
		Scan(&user.ID,
			&user.Username,
			&user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting user: %w", err)
	}

	return user, nil
}

func (u *UserRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	err := u.QueryRowContext(ctx, query.UpdateUser, user.Username, user.PasswordHash, user.ID).Err()
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while updating user: %w", err)
	}

	return user, nil
}

func (u *UserRepository) Delete(ctx context.Context, user *models.User) (*models.User, error) {
	err := u.QueryRowContext(ctx, query.DeleteUsers, user.ID).Err()
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while deleting user: %w", err)
	}

	return user, nil
}
