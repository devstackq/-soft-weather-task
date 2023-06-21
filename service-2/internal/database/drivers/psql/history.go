package psql

import (
	"context"
	"database/sql"
	"fmt"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
	"strconv"
)

type historyImpl struct {
	db        *sql.DB
	tableName string
}

func NewHistoryRepository(db *sql.DB, tableName string) drivers.HistoryRepository {
	return &historyImpl{
		db:        db,
		tableName: tableName,
	}
}

func (h *historyImpl) Create(ctx context.Context, history models.History) error {
	query := fmt.Sprint(`INSERT INTO `, h.tableName, `(user_id, task_id) VALUES ($1, $2)`)
	userIDInt, err := strconv.Atoi(history.UserID)
	if err != nil {
		return err
	}
	taskIDInt, err := strconv.Atoi(history.TaskID)
	if err != nil {
		return err
	}

	_, err = h.db.ExecContext(ctx, query, userIDInt, taskIDInt)
	if err != nil {
		return err
	}
	return nil
}
