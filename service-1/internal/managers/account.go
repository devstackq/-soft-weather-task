package managers

import (
	"context"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
)

type Account interface {
	Get(ctx context.Context, id string) (models.Account, error)
	IncreaseDebt(ctx context.Context, account models.Account) error
	DecreaseDebt(ctx context.Context, account models.Account) error
}

type AccountManager struct {
	accountRepo drivers.AccountRepository
}

func NewAccount(ds drivers.Datastore) Account {
	return &AccountManager{
		accountRepo: ds.AccountRepository(),
	}
}

func (a *AccountManager) IncreaseDebt(ctx context.Context, req models.Account) error {
	return a.accountRepo.IncreaseDebt(ctx, req)
}
func (a *AccountManager) DecreaseDebt(ctx context.Context, account models.Account) error {
	return a.accountRepo.DecreaseDebt(ctx, account)
}

func (a *AccountManager) Get(ctx context.Context, id string) (models.Account, error) {
	return a.accountRepo.Get(ctx, id)
}
