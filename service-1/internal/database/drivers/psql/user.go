package psql

import (
	"context"
	"database/sql"
	"fmt"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
)

var _ drivers.UserRepository = &userImpl{}

type userImpl struct {
	tableName string
	db        *sql.DB
}

func (u *userImpl) Create(ctx context.Context, req models.User) error {
	insertStmt := fmt.Sprint("INSERT INTO ", u.tableName, " (name, login) VALUES ($1, $2)")
	_, err := u.db.Exec(insertStmt, req.FullName, req.Login)
	if err != nil {
		return err
	}
	return nil

}

func NewUserRepository(db *sql.DB, tName string) drivers.UserRepository {
	return &userImpl{
		tableName: tName,
		db:        db,
	}
}
