package service

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"github.com/dlc-01/GophKeeper/internal/general/pass"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/projectError"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
)

type AuthService struct {
	log  *logger.Logger
	repo port.IUsersRepository
	ts   port.ITokenService
}

func NewAuthService(repo port.IUsersRepository, ts port.ITokenService, log *logger.Logger) *AuthService {
	return &AuthService{
		log,
		repo,
		ts,
	}
}

func (as *AuthService) Login(ctx context.Context, userRPC *models.User) (string, error) {
	user, err := as.repo.GetByUsername(ctx, userRPC)
	if err != nil {
		return "", err
	}

	check := pass.ComparePasswordHash(userRPC.PasswordHash, user.PasswordHash)
	if check != true {
		return "", projectError.ErrInvalidCredentials
	}

	accessToken, err := as.ts.CreateToken(user)
	if err != nil {
		return "", projectError.ErrTokenCreation
	}

	return accessToken, nil
}

func (as *AuthService) Register(ctx context.Context, user *models.User) (string, error) {
	var err error
	user, err = as.repo.Create(ctx, user)
	if err != nil {
		return "", err
	}

	accessToken, err := as.ts.CreateToken(user)
	if err != nil {
		return "", projectError.ErrTokenCreation
	}

	return accessToken, nil
}
