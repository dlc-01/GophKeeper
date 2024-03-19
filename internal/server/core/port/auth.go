package port

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/auth/jwt/manager"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
)

//go:generate mockgen -source=auth.go -destination=mock/auth.go -package=mock

type ITokenService interface {
	CreateToken(user *models.User) (string, error)

	VerifyToken(accessToken string) (*manager.UserClaims, error)
}

type IAuthService interface {
	Login(ctx context.Context, user *models.User) (string, error)
	Register(ctx context.Context, user *models.User) (string, error)
}
