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

type TextRepository struct {
	*sql.DB
	lgr *logger.Logger
}

func NewTextRepository(lgr *logger.Logger, client *sql.DB) (*TextRepository, error) {
	return &TextRepository{
		DB:  client,
		lgr: lgr,
	}, nil
}

func (t *TextRepository) CreateByUserId(ctx context.Context, user *models.User, text *models.Text) (*models.Text, error) {
	var stored models.Text

	tx, err := t.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.CreateTextInf, user.ID, text.Note, text.Metadata).
		Scan(&stored.ID,
			&stored.Note,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating text: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (t *TextRepository) CreateByUsername(ctx context.Context, user *models.User, text *models.Text) (*models.Text, error) {
	var stored models.Text

	tx, err := t.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.CreateTextInfByUsername, user.Username, text.Note, text.Metadata).
		Scan(&stored.ID,
			&stored.Note,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating text: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (t *TextRepository) GetTextsByUsername(ctx context.Context, user *models.User) (*[]models.Text, error) {
	stored := make([]models.Text, 0)

	tx, err := t.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	row, err := tx.QueryContext(ctx, query.GetTextInfsByUsername, user.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting text: %w", err)
	}

	for row.Next() {
		text := models.Text{}
		if err = row.Scan(&text.ID, &text.Note, &text.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}

			return nil, fmt.Errorf("error while scaning text: %w", err)
		}
		stored = append(stored, text)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (t *TextRepository) GetTextsByUserID(ctx context.Context, user *models.User) (*[]models.Text, error) {
	stored := make([]models.Text, 0)

	tx, err := t.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	row, err := tx.QueryContext(ctx, query.GetTextInfsByUserID, user.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrDataNotFound
		}
		return nil, fmt.Errorf("error while geting text: %w", err)
	}

	for row.Next() {
		text := models.Text{}
		if err = row.Scan(&text.ID, &text.Note, &text.Metadata); err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrDataNotFound
			}
			return nil, fmt.Errorf("error while scaning text: %w", err)
		}
		stored = append(stored, text)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}
func (t *TextRepository) GetTextsByID(ctx context.Context, text *models.Text) (*models.Text, error) {
	var stored models.Text

	tx, err := t.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.GetBTextInfByID, text.ID).
		Scan(&stored.ID,
			&stored.Note,
			&stored.Metadata)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrConflictingData
		}
		return nil, fmt.Errorf("error while creating text: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return &stored, nil
}

func (t *TextRepository) Update(ctx context.Context, text *models.Text) (*models.Text, error) {
	tx, err := t.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.UpdateTextInf, text.ID, text.Note, text.Metadata).Err()
	if err != nil {
		return nil, fmt.Errorf("error while updating text: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return text, nil
}

func (t *TextRepository) DeleteById(ctx context.Context, text *models.Text) (*models.Text, error) {
	tx, err := t.Begin()
	if err != nil {
		return nil, fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeleteTextInf, text.ID).Err()
	if err != nil {
		return nil, fmt.Errorf("error while deleting text: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return text, nil
}

func (t *TextRepository) DeleteByUsername(ctx context.Context, user *models.User) error {
	tx, err := t.Begin()
	if err != nil {
		return fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeleteTextInfByUsername, user.Username).Err()
	if err != nil {
		return fmt.Errorf("error while deleting text: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error while commiting transacsion: %s", err)
	}
	return nil
}

func (t *TextRepository) DeleteByUserID(ctx context.Context, user *models.User) error {
	tx, err := t.Begin()
	if err != nil {
		return fmt.Errorf("error while creating transacsion: %s", err)
	}

	err = tx.QueryRowContext(ctx, query.DeleteTextInfByUserID, user.ID).Err()
	if err != nil {
		return fmt.Errorf("error while deleting text: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error while commiting transacsion: %s", err)
	}

	return nil
}
