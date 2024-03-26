package port

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
)

//go:generate mockgen -source=text.go -destination=mock/text.go -package=mock

type ITextRepository interface {
	CreateByUserId(ctx context.Context, user *models.User, text *models.Text) (*models.Text, error)

	CreateByUsername(ctx context.Context, user *models.User, text *models.Text) (*models.Text, error)

	GetTextsByUsername(ctx context.Context, user *models.User) (*[]models.Text, error)

	GetTextsByUserID(ctx context.Context, user *models.User) (*[]models.Text, error)

	GetTextsByID(ctx context.Context, text *models.Text) (*models.Text, error)

	Update(ctx context.Context, text *models.Text) (*models.Text, error)

	DeleteById(ctx context.Context, text *models.Text) (*models.Text, error)

	DeleteByUsername(ctx context.Context, user *models.User) error

	DeleteByUserID(ctx context.Context, user *models.User) error
}

type ITextService interface {
	CreateByUserId(ctx context.Context, user models.User, text models.Text) (*models.Text, error)

	CreateByUsername(ctx context.Context, user models.User, text models.Text) (*models.Text, error)

	GetTextsByUsername(ctx context.Context, user models.User) (*[]models.Text, error)

	GetTextsByUserID(ctx context.Context, user models.User) (*[]models.Text, error)

	Update(ctx context.Context, text models.Text) (*models.Text, error)

	DeleteById(ctx context.Context, text models.Text) (*models.Text, error)

	DeleteByUsername(ctx context.Context, user models.User) error

	DeleteByUserID(ctx context.Context, user models.User) error
}
