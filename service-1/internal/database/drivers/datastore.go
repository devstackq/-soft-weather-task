package drivers

import (
	"context"
	"soft-weater/internal/config"
	"soft-weater/internal/models"
)

type Datastore interface {
	Base
	UserRepository() UserRepository
	AccountRepository() AccountRepository
}

type Base interface {
	Connect(cfg config.DataStoreConfiguration) error
}

type UserRepository interface {
	Create(ctx context.Context, req models.User) error
}

type AccountRepository interface {
	IncreaseDebt(ctx context.Context, account models.Account) error
	DecreaseDebt(ctx context.Context, account models.Account) error
	Get(ctx context.Context, id string) (models.Account, error)
}
