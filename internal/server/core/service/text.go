package service

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/logger"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/projectError"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
)

type TextService struct {
	log  logger.Logger
	repo port.ITextRepository
}

func NewTextService(repo port.ITextRepository, log logger.Logger) *TextService {
	return &TextService{
		log,
		repo,
	}
}

func (t *TextService) CreateByUserId(ctx context.Context, user models.User, text models.Text) (*models.Text, error) {
	newText, err := t.repo.CreateByUserId(ctx, &user, &text)
	if err != nil {
		if err == projectError.ErrConflictingData {
			return nil, projectError.ErrConflictingData
		}
		return nil, err
	}

	return newText, nil
}

func (t *TextService) CreateByUsername(ctx context.Context, user models.User, text models.Text) (*models.Text, error) {
	_, err := t.repo.CreateByUsername(ctx, &user, &text)
	if err != nil {
		if err == projectError.ErrConflictingData {
			return nil, projectError.ErrConflictingData
		}
		return nil, err
	}

	return &text, nil
}

func (t *TextService) GetTextsByUsername(ctx context.Context, user models.User) (*[]models.Text, error) {
	texts, err := t.repo.GetTextsByUsername(ctx, &user)
	if err != nil {
		if err == projectError.ErrDataNotFound {
			return nil, projectError.ErrDataNotFound
		}
		return nil, err
	}

	return texts, nil
}

func (t *TextService) GetTextsByUserID(ctx context.Context, user models.User) (*[]models.Text, error) {
	texts, err := t.repo.GetTextsByUserID(ctx, &user)
	if err != nil {
		if err == projectError.ErrDataNotFound {
			return nil, projectError.ErrDataNotFound
		}
		return nil, err
	}

	return texts, nil
}

func (t *TextService) Update(ctx context.Context, text models.Text) (*models.Text, error) {
	_, err := t.repo.GetTextsByID(ctx, &text)
	if err != nil {
		return nil, err
	}

	_, err = t.repo.Update(ctx, &text)
	if err != nil {
		return nil, err
	}

	return &text, nil
}

func (t *TextService) DeleteById(ctx context.Context, text models.Text) (*models.Text, error) {
	_, err := t.repo.GetTextsByID(ctx, &text)
	if err != nil {
		return nil, err
	}

	_, err = t.repo.DeleteById(ctx, &text)
	if err != nil {
		return nil, err
	}

	return &text, nil
}

func (t *TextService) DeleteByUsername(ctx context.Context, user models.User) error {
	_, err := t.repo.GetTextsByUsername(ctx, &user)
	if err != nil {
		return err
	}

	err = t.repo.DeleteByUsername(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (t *TextService) DeleteByUserID(ctx context.Context, user models.User) error {
	_, err := t.repo.GetTextsByUserID(ctx, &user)
	if err != nil {
		return err
	}

	err = t.repo.DeleteByUserID(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
