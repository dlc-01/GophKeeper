package service

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/projectError"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
)

type UserService struct {
	log  *logger.Logger
	repo port.IUsersRepository
}

func NewUserService(repo port.IUsersRepository, log *logger.Logger) *UserService {
	return &UserService{
		log,
		repo,
	}
}

func (us *UserService) GetByUserID(ctx context.Context, user *models.User) (*models.User, error) {

	user, err := us.repo.GetByUserID(ctx, user)
	if err != nil {
		us.log.Info("some trubles")
		if err == projectError.ErrDataNotFound {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetByUsername(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := us.repo.GetByUsername(ctx, user)
	if err != nil {
		us.log.Info("some trubles")
		return nil, err
	}

	return user, nil
}

func (us *UserService) Update(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := us.repo.GetByUsername(ctx, user)
	if err != nil {
		return nil, err
	}

	_, err = us.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) Delete(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := us.repo.GetByUserID(ctx, user)
	if err != nil {
		return nil, err
	}
	user, err = us.repo.Delete(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
