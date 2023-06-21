package managers

import (
	"context"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
)

type Task interface {
	GetList(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, req models.Task) error
}

type TaskManager struct {
	taskRepo drivers.TaskRepository
}

func NewTask(ds drivers.Datastore) Task {
	return &TaskManager{
		taskRepo: ds.TaskRepository(),
	}
}

func (t *TaskManager) GetList(ctx context.Context) ([]models.Task, error) {
	return t.taskRepo.GetList(ctx)
}

func (t *TaskManager) Update(ctx context.Context, req models.Task) error {
	return t.taskRepo.Update(ctx, req)
}
