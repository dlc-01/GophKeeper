package port

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
)

//go:generate mockgen -source=users.go -destination=mock/payment.go -package=mock

type IUsersRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)

	GetByUserID(ctx context.Context, user *models.User) (*models.User, error)

	GetByUsername(ctx context.Context, user *models.User) (*models.User, error)

	Update(ctx context.Context, user *models.User) (*models.User, error)

	Delete(ctx context.Context, user *models.User) (*models.User, error)
}

type IUsersService interface {
	GetByUserID(ctx context.Context, user *models.User) (*models.User, error)

	GetByUsername(ctx context.Context, user *models.User) (*models.User, error)

	Update(ctx context.Context, user *models.User) (*models.User, error)

	Delete(ctx context.Context, user *models.User) (*models.User, error)
}
