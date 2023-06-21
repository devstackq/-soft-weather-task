package psql

import (
	"context"
	"database/sql"
	"fmt"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
)

type taskImpl struct {
	db        *sql.DB
	tableName string
}

func (t *taskImpl) Update(ctx context.Context, req models.Task) error {
	query := fmt.Sprint("UPDATE ", t.tableName, " SET price = ", req.Price, " WHERE id = ", req.ID)
	result, err := t.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (t *taskImpl) GetByID(ctx context.Context, id string) (models.Task, error) {
	query := fmt.Sprint("SELECT price FROM ", t.tableName, " WHERE id = ", id)
	var task models.Task
	if err := t.db.QueryRowContext(ctx, query).Scan(&task.Price); err != nil {
		return task, err
	}
	return task, nil
}

func (t *taskImpl) GetList(ctx context.Context) (result []models.Task, err error) {
	query := fmt.Sprint(` SELECT id, name, price FROM `, t.tableName)
	rows, err := t.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Task
		err = rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return
}

func NewTaskRepository(db *sql.DB, tableName string) drivers.TaskRepository {
	return &taskImpl{
		db:        db,
		tableName: tableName,
	}
}
