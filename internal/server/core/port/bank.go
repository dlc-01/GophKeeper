package port

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
)

//go:generate mockgen -source=bank.go -destination=mock/bank.go -package=mock

type IBankRepository interface {
	CreateByUserId(ctx context.Context, bank *models.BankAccount, user *models.User) (*models.BankAccount, error)

	CreateByUsername(ctx context.Context, bank *models.BankAccount, user *models.User) (*models.BankAccount, error)

	GetBankAccountsByUsername(ctx context.Context, user *models.User) (*[]models.BankAccount, error)

	GetBankAccountByID(ctx context.Context, bank *models.BankAccount) (*models.BankAccount, error)

	GetBankAccountsByUserID(ctx context.Context, user *models.User) (*[]models.BankAccount, error)

	Update(ctx context.Context, bank *models.BankAccount) (*models.BankAccount, error)

	DeleteById(ctx context.Context, bank *models.BankAccount) error

	DeleteByUserId(ctx context.Context, user models.User) error

	DeleteByUsername(ctx context.Context, user *models.User) error
}

type IBankService interface {
	CreateByUserId(ctx context.Context, bank models.BankAccount, user models.User) (*models.BankAccount, error)

	CreateByUsername(ctx context.Context, bank models.BankAccount, user models.User) (*models.BankAccount, error)

	GetBankAccountsByUsername(ctx context.Context, user models.User) (*[]models.BankAccount, error)

	GetPairsByUserID(ctx context.Context, user models.User) (*[]models.BankAccount, error)

	Update(ctx context.Context, acc models.BankAccount) (*models.BankAccount, error)

	DeleteById(ctx context.Context, acc models.BankAccount) error

	DeleteByUserId(ctx context.Context, user models.User) error

	DeleteByUsername(ctx context.Context, user models.User) error
}
