package drivers

import (
	"context"
	"soft-weater/internal/config"
	"soft-weater/internal/models"
)

type Datastore interface {
	Base
	TaskRepository() TaskRepository
	HistoryRepository() HistoryRepository
}

type Base interface {
	Connect(cfg config.DataStoreConfiguration) error
}

type AccountRepository interface {
	IncreaseDebt(ctx context.Context) error
	DecreaseDebt(ctx context.Context) error
}

type TaskRepository interface {
	GetList(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, req models.Task) error
	GetByID(ctx context.Context, id string) (models.Task, error)
}

type HistoryRepository interface {
	Create(ctx context.Context, history models.History) error
}
