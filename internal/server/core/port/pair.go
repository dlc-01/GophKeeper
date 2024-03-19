package port

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
)

//go:generate mockgen -source=pair.go -destination=mock/pair.go -package=mock

type IPairRepository interface {
	CreateByUserId(ctx context.Context, pair *models.Pair, user *models.User) (*models.Pair, error)

	CreateByUsername(ctx context.Context, pair *models.Pair, user *models.User) (*models.Pair, error)

	GetPairsByUsername(ctx context.Context, user *models.User) (*[]models.Pair, error)

	GetPairByID(ctx context.Context, pair *models.Pair) (*models.Pair, error)

	GetPairsByUserID(ctx context.Context, user *models.User) (*[]models.Pair, error)

	Update(ctx context.Context, pair *models.Pair) (*models.Pair, error)

	DeleteById(ctx context.Context, pair *models.Pair) (*models.Pair, error)

	DeleteByUserId(ctx context.Context, user *models.User) error

	DeleteByUsername(ctx context.Context, user *models.User) error
}

type IPairService interface {
	CreateByUserId(ctx context.Context, pair models.Pair, user models.User) (*models.Pair, error)

	CreateByUsername(ctx context.Context, pair models.Pair, user models.User) (*models.Pair, error)

	GetPairsByUsername(ctx context.Context, user models.User) (*[]models.Pair, error)

	GetPairsByUserID(ctx context.Context, user models.User) (*[]models.Pair, error)

	Update(ctx context.Context, pair models.Pair) (*models.Pair, error)

	DeleteById(ctx context.Context, pair models.Pair) (*models.Pair, error)

	DeleteByUserId(ctx context.Context, user models.User) error

	DeleteByUsername(ctx context.Context, user models.User) error
}
