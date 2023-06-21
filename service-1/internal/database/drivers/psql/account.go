package psql

import (
	"context"
	"database/sql"
	"fmt"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
)

type accountImpl struct {
	db    *sql.DB
	tName string
}

func (a *accountImpl) DecreaseDebt(ctx context.Context, req models.Account) error {
	updateStmt := fmt.Sprint("UPDATE ", a.tName, " SET debt=debt- ", req.Debt, " WHERE user_id =", req.UserID)
	_, err := a.db.ExecContext(ctx, updateStmt)
	if err != nil {
		return err
	}
	return nil
}

func (a *accountImpl) IncreaseDebt(ctx context.Context, req models.Account) error {
	updateStmt := fmt.Sprint("UPDATE ", a.tName, " SET debt=debt+ ", req.Debt, " WHERE user_id =", req.UserID)
	_, err := a.db.ExecContext(ctx, updateStmt)
	if err != nil {
		return err
	}
	return nil

}

func (a *accountImpl) Get(ctx context.Context, uid string) (models.Account, error) {
	query := fmt.Sprint("SELECT id, debt, balance FROM ", a.tName, " WHERE user_id = ", uid)
	//fmt.Printf("repo layer getAccount %+v \n", query)

	var res models.Account
	if err := a.db.QueryRowContext(ctx, query).Scan(&res.ID, &res.Debt, &res.Balance); err != nil {
		return res, err
	}
	return res, nil
}
func NewAccountRepository(db *sql.DB, tName string) drivers.AccountRepository {
	return &accountImpl{
		tName: tName,
		db:    db,
	}
}
